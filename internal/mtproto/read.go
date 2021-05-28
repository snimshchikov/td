package mtproto

import (
	"context"
	"errors"
	"net"
	"time"

	"go.uber.org/zap"
	"golang.org/x/xerrors"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/internal/crypto"
	"github.com/gotd/td/internal/proto"
	"github.com/gotd/td/internal/proto/codec"
)

// https://core.telegram.org/mtproto/description#message-identifier-msg-id
// A message is rejected over 300 seconds after it is created or 30 seconds
// before it is created (this is needed to protect from replay attacks).
const (
	maxPast   = time.Second * 300
	maxFuture = time.Second * 30
)

// errRejected is returned on invalid message that should not be processed.
var errRejected = errors.New("message rejected")

func checkMessageID(now time.Time, rawID int64) error {
	id := proto.MessageID(rawID)

	// Check that message is from server.
	switch id.Type() {
	case proto.MessageFromServer, proto.MessageServerResponse:
		// Valid.
	default:
		return xerrors.Errorf("unexpected type %s: %w", id.Type(), errRejected)
	}

	created := id.Time()
	if created.Before(now) && now.Sub(created) > maxPast {
		return xerrors.Errorf("created too far in past: %w", errRejected)
	}
	if created.Sub(now) > maxFuture {
		return xerrors.Errorf("created too far in future: %w", errRejected)
	}

	return nil
}

func (c *Conn) decryptMessage(b *bin.Buffer) (*crypto.EncryptedMessageData, error) {
	session := c.session()
	msg, err := c.cipher.DecryptFromBuffer(session.Key, b)
	if err != nil {
		return nil, xerrors.Errorf("decrypt: %w", err)
	}

	// Validating message. This protects from replay attacks.
	if msg.SessionID != session.ID {
		return nil, xerrors.Errorf("invalid session (got %d, expected %d): %w", msg.SessionID, session.ID, errRejected)
	}
	if err := checkMessageID(c.clock.Now(), msg.MessageID); err != nil {
		return nil, xerrors.Errorf("bad message id %d: %w", msg.MessageID, err)
	}
	if !c.messageIDBuf.Consume(msg.MessageID) {
		return nil, xerrors.Errorf("duplicate or too low message id %d: %w", msg.MessageID, errRejected)
	}

	return msg, nil
}

func (c *Conn) consumeMessage(ctx context.Context, buf *bin.Buffer) error {
	msg, err := c.decryptMessage(buf)
	if errors.Is(err, errRejected) {
		c.log.Warn("Ignoring rejected message", zap.Error(err))
		return nil
	}
	if err != nil {
		return xerrors.Errorf("consume message: %w", err)
	}

	if err := c.handleMessage(msg.MessageID, &bin.Buffer{Buf: msg.Data()}); err != nil {
		// Probably we can return here, but this will shutdown whole
		// connection which can be unexpected.
		c.log.Warn("Error while handling message", zap.Error(err))
		// Sending acknowledge even on error. Client should restore
		// from missing updates via explicit pts check and getDiff call.
	}

	needAck := (msg.SeqNo & 0x01) != 0
	if needAck {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case c.ackSendChan <- msg.MessageID:
		}
	}

	return nil
}

func (c *Conn) noUpdates(err error) bool {
	// Checking for read timeout.
	var syscall *net.OpError
	if errors.As(err, &syscall) && syscall.Timeout() {
		// We call SetReadDeadline so such error is expected.
		c.log.Debug("No updates")
		return true
	}
	return false
}

func (c *Conn) handleAuthKeyNotFound(ctx context.Context) error {
	if c.session().ID == 0 {
		// The 404 error can also be caused by zero session id.
		// See https://github.com/gotd/td/issues/107
		//
		// We should recover from this in createAuthKey, but in general
		// this code branch should be unreachable.
		c.log.Warn("BUG: zero session id found")
	}
	c.log.Warn("Re-generating keys (server not found key that we provided)")
	if err := c.createAuthKey(ctx); err != nil {
		return xerrors.Errorf("unable to create auth key: %w", err)
	}
	c.log.Info("Re-created auth keys")
	// Request will be retried by ack loop.
	// Probably we can speed-up this.
	return nil
}

func (c *Conn) readLoop(ctx context.Context) (err error) {
	b := new(bin.Buffer)

	log := c.log.Named("read")
	log.Debug("Read loop started")
	defer func() {
		l := log
		if err != nil {
			l = log.With(zap.NamedError("reason", err))
		}
		l.Debug("Read loop done")
	}()

	for {
		b.Reset()

		err = c.conn.Recv(ctx, b)
		if err == nil {
			if err := c.consumeMessage(ctx, b); err != nil {
				return err
			}

			continue
		}

		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
			if c.noUpdates(err) {
				continue
			}
		}

		var protoErr *codec.ProtocolErr
		if errors.As(err, &protoErr) && protoErr.Code == codec.CodeAuthKeyNotFound {
			if err := c.handleAuthKeyNotFound(ctx); err != nil {
				return xerrors.Errorf("auth key not found: %w", err)
			}
			continue
		}

		select {
		case <-ctx.Done():
			return xerrors.Errorf("read loop: %w", ctx.Err())
		default:
			return xerrors.Errorf("read: %w", err)
		}
	}
}
