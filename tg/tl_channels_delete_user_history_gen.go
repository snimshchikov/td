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

// ChannelsDeleteUserHistoryRequest represents TL type `channels.deleteUserHistory#d10dd71b`.
// Delete all messages sent by a certain user in a supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.deleteUserHistory for reference.
type ChannelsDeleteUserHistoryRequest struct {
	// Supergroup¹
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Channel InputChannelClass
	// User whose messages should be deleted
	UserID InputUserClass
}

// ChannelsDeleteUserHistoryRequestTypeID is TL type id of ChannelsDeleteUserHistoryRequest.
const ChannelsDeleteUserHistoryRequestTypeID = 0xd10dd71b

func (d *ChannelsDeleteUserHistoryRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Channel == nil) {
		return false
	}
	if !(d.UserID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *ChannelsDeleteUserHistoryRequest) String() string {
	if d == nil {
		return "ChannelsDeleteUserHistoryRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelsDeleteUserHistoryRequest")
	sb.WriteString("{\n")
	sb.WriteString("\tChannel: ")
	sb.WriteString(fmt.Sprint(d.Channel))
	sb.WriteString(",\n")
	sb.WriteString("\tUserID: ")
	sb.WriteString(fmt.Sprint(d.UserID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (d *ChannelsDeleteUserHistoryRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode channels.deleteUserHistory#d10dd71b as nil")
	}
	b.PutID(ChannelsDeleteUserHistoryRequestTypeID)
	if d.Channel == nil {
		return fmt.Errorf("unable to encode channels.deleteUserHistory#d10dd71b: field channel is nil")
	}
	if err := d.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.deleteUserHistory#d10dd71b: field channel: %w", err)
	}
	if d.UserID == nil {
		return fmt.Errorf("unable to encode channels.deleteUserHistory#d10dd71b: field user_id is nil")
	}
	if err := d.UserID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.deleteUserHistory#d10dd71b: field user_id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (d *ChannelsDeleteUserHistoryRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode channels.deleteUserHistory#d10dd71b to nil")
	}
	if err := b.ConsumeID(ChannelsDeleteUserHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.deleteUserHistory#d10dd71b: %w", err)
	}
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.deleteUserHistory#d10dd71b: field channel: %w", err)
		}
		d.Channel = value
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.deleteUserHistory#d10dd71b: field user_id: %w", err)
		}
		d.UserID = value
	}
	return nil
}

// Ensuring interfaces in compile-time for ChannelsDeleteUserHistoryRequest.
var (
	_ bin.Encoder = &ChannelsDeleteUserHistoryRequest{}
	_ bin.Decoder = &ChannelsDeleteUserHistoryRequest{}
)

// ChannelsDeleteUserHistory invokes method channels.deleteUserHistory#d10dd71b returning error if any.
// Delete all messages sent by a certain user in a supergroup¹
//
// Links:
//  1) https://core.telegram.org/api/channel
//
// Possible errors:
//  400 CHANNEL_INVALID: The provided channel is invalid
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this
//  403 CHAT_WRITE_FORBIDDEN: You can't write in this chat
//  400 MSG_ID_INVALID: Invalid message ID provided
//  400 USER_ID_INVALID: The provided user ID is invalid
//
// See https://core.telegram.org/method/channels.deleteUserHistory for reference.
func (c *Client) ChannelsDeleteUserHistory(ctx context.Context, request *ChannelsDeleteUserHistoryRequest) (*MessagesAffectedHistory, error) {
	var result MessagesAffectedHistory

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
