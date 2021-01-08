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

// PhoneAcceptCallRequest represents TL type `phone.acceptCall#3bd2b4a0`.
// Accept incoming call
//
// See https://core.telegram.org/method/phone.acceptCall for reference.
type PhoneAcceptCallRequest struct {
	// The call to accept
	Peer InputPhoneCall
	// Parameter for E2E encryption key exchange »¹
	//
	// Links:
	//  1) https://core.telegram.org/api/end-to-end/voice-calls
	GB []byte
	// Phone call settings
	Protocol PhoneCallProtocol
}

// PhoneAcceptCallRequestTypeID is TL type id of PhoneAcceptCallRequest.
const PhoneAcceptCallRequestTypeID = 0x3bd2b4a0

func (a *PhoneAcceptCallRequest) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Peer.Zero()) {
		return false
	}
	if !(a.GB == nil) {
		return false
	}
	if !(a.Protocol.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *PhoneAcceptCallRequest) String() string {
	if a == nil {
		return "PhoneAcceptCallRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PhoneAcceptCallRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tPeer: ")
	sb.WriteString(fmt.Sprint(a.Peer))
	sb.WriteString(",\n")
	sb.WriteString("\tGB: ")
	sb.WriteString(fmt.Sprint(a.GB))
	sb.WriteString(",\n")
	sb.WriteString("\tProtocol: ")
	sb.WriteString(fmt.Sprint(a.Protocol))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (a *PhoneAcceptCallRequest) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode phone.acceptCall#3bd2b4a0 as nil")
	}
	b.PutID(PhoneAcceptCallRequestTypeID)
	if err := a.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.acceptCall#3bd2b4a0: field peer: %w", err)
	}
	b.PutBytes(a.GB)
	if err := a.Protocol.Encode(b); err != nil {
		return fmt.Errorf("unable to encode phone.acceptCall#3bd2b4a0: field protocol: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *PhoneAcceptCallRequest) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode phone.acceptCall#3bd2b4a0 to nil")
	}
	if err := b.ConsumeID(PhoneAcceptCallRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode phone.acceptCall#3bd2b4a0: %w", err)
	}
	{
		if err := a.Peer.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.acceptCall#3bd2b4a0: field peer: %w", err)
		}
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode phone.acceptCall#3bd2b4a0: field g_b: %w", err)
		}
		a.GB = value
	}
	{
		if err := a.Protocol.Decode(b); err != nil {
			return fmt.Errorf("unable to decode phone.acceptCall#3bd2b4a0: field protocol: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PhoneAcceptCallRequest.
var (
	_ bin.Encoder = &PhoneAcceptCallRequest{}
	_ bin.Decoder = &PhoneAcceptCallRequest{}
)

// PhoneAcceptCall invokes method phone.acceptCall#3bd2b4a0 returning error if any.
// Accept incoming call
//
// Possible errors:
//  400 CALL_ALREADY_ACCEPTED: The call was already accepted
//  400 CALL_ALREADY_DECLINED: The call was already declined
//  400 CALL_PEER_INVALID: The provided call peer object is invalid
//  400 CALL_PROTOCOL_FLAGS_INVALID: Call protocol flags invalid
//
// See https://core.telegram.org/method/phone.acceptCall for reference.
func (c *Client) PhoneAcceptCall(ctx context.Context, request *PhoneAcceptCallRequest) (*PhonePhoneCall, error) {
	var result PhonePhoneCall

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
