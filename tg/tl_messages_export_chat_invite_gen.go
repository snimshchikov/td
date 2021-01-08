// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}

// MessagesExportChatInviteRequest represents TL type `messages.exportChatInvite#df7534c`.
// Export an invite link for a chat
//
// See https://core.telegram.org/method/messages.exportChatInvite for reference.
type MessagesExportChatInviteRequest struct {
	// Chat
	Peer InputPeerClass
}

// MessagesExportChatInviteRequestTypeID is TL type id of MessagesExportChatInviteRequest.
const MessagesExportChatInviteRequestTypeID = 0xdf7534c

func (e *MessagesExportChatInviteRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Peer == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *MessagesExportChatInviteRequest) String() string {
	if e == nil {
		return "MessagesExportChatInviteRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("MessagesExportChatInviteRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(fmt.Sprint(e.Peer))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (e *MessagesExportChatInviteRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.exportChatInvite#df7534c as nil")
	}
	b.PutID(MessagesExportChatInviteRequestTypeID)
	if e.Peer == nil {
		return fmt.Errorf("unable to encode messages.exportChatInvite#df7534c: field peer is nil")
	}
	if err := e.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.exportChatInvite#df7534c: field peer: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (e *MessagesExportChatInviteRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.exportChatInvite#df7534c to nil")
	}
	if err := b.ConsumeID(MessagesExportChatInviteRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.exportChatInvite#df7534c: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode messages.exportChatInvite#df7534c: field peer: %w", err)
		}
		e.Peer = value
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesExportChatInviteRequest.
var (
	_ bin.Encoder = &MessagesExportChatInviteRequest{}
	_ bin.Decoder = &MessagesExportChatInviteRequest{}
)

// MessagesExportChatInvite invokes method messages.exportChatInvite#df7534c returning error if any.
// Export an invite link for a chat
//
// Possible errors:
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  400 CHAT_ID_INVALID: The provided chat id is invalid
//  403 CHAT_WRITE_FORBIDDEN: You can't write in this chat
//  400 PEER_ID_INVALID: The provided peer id is invalid
//
// See https://core.telegram.org/method/messages.exportChatInvite for reference.
// Can be used by bots.
func (c *Client) MessagesExportChatInvite(ctx context.Context, peer InputPeerClass) (ExportedChatInviteClass, error) {
	var result ExportedChatInviteBox

	request := &MessagesExportChatInviteRequest{
		Peer: peer,
	}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.ExportedChatInvite, nil
}
