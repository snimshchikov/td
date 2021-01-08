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

// ChatParticipant represents TL type `chatParticipant#c8d7493e`.
// Group member.
//
// See https://core.telegram.org/constructor/chatParticipant for reference.
type ChatParticipant struct {
	// Member user ID
	UserID int
	// ID of the user that added the member to the group
	InviterID int
	// Date added to the group
	Date int
}

// ChatParticipantTypeID is TL type id of ChatParticipant.
const ChatParticipantTypeID = 0xc8d7493e

func (c *ChatParticipant) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UserID == 0) {
		return false
	}
	if !(c.InviterID == 0) {
		return false
	}
	if !(c.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatParticipant) String() string {
	if c == nil {
		return "ChatParticipant(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChatParticipant")
	sb.WriteString("{\n")
	sb.WriteString("\tUserID: ")
	sb.WriteString(fmt.Sprint(c.UserID))
	sb.WriteString(",\n")
	sb.WriteString("\tInviterID: ")
	sb.WriteString(fmt.Sprint(c.InviterID))
	sb.WriteString(",\n")
	sb.WriteString("\tDate: ")
	sb.WriteString(fmt.Sprint(c.Date))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *ChatParticipant) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatParticipant#c8d7493e as nil")
	}
	b.PutID(ChatParticipantTypeID)
	b.PutInt(c.UserID)
	b.PutInt(c.InviterID)
	b.PutInt(c.Date)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatParticipant) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatParticipant#c8d7493e to nil")
	}
	if err := b.ConsumeID(ChatParticipantTypeID); err != nil {
		return fmt.Errorf("unable to decode chatParticipant#c8d7493e: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipant#c8d7493e: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipant#c8d7493e: field inviter_id: %w", err)
		}
		c.InviterID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipant#c8d7493e: field date: %w", err)
		}
		c.Date = value
	}
	return nil
}

// construct implements constructor of ChatParticipantClass.
func (c ChatParticipant) construct() ChatParticipantClass { return &c }

// Ensuring interfaces in compile-time for ChatParticipant.
var (
	_ bin.Encoder = &ChatParticipant{}
	_ bin.Decoder = &ChatParticipant{}

	_ ChatParticipantClass = &ChatParticipant{}
)

// ChatParticipantCreator represents TL type `chatParticipantCreator#da13538a`.
// Represents the creator of the group
//
// See https://core.telegram.org/constructor/chatParticipantCreator for reference.
type ChatParticipantCreator struct {
	// ID of the user that created the group
	UserID int
}

// ChatParticipantCreatorTypeID is TL type id of ChatParticipantCreator.
const ChatParticipantCreatorTypeID = 0xda13538a

func (c *ChatParticipantCreator) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UserID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatParticipantCreator) String() string {
	if c == nil {
		return "ChatParticipantCreator(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChatParticipantCreator")
	sb.WriteString("{\n")
	sb.WriteString("\tUserID: ")
	sb.WriteString(fmt.Sprint(c.UserID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *ChatParticipantCreator) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatParticipantCreator#da13538a as nil")
	}
	b.PutID(ChatParticipantCreatorTypeID)
	b.PutInt(c.UserID)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatParticipantCreator) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatParticipantCreator#da13538a to nil")
	}
	if err := b.ConsumeID(ChatParticipantCreatorTypeID); err != nil {
		return fmt.Errorf("unable to decode chatParticipantCreator#da13538a: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipantCreator#da13538a: field user_id: %w", err)
		}
		c.UserID = value
	}
	return nil
}

// construct implements constructor of ChatParticipantClass.
func (c ChatParticipantCreator) construct() ChatParticipantClass { return &c }

// Ensuring interfaces in compile-time for ChatParticipantCreator.
var (
	_ bin.Encoder = &ChatParticipantCreator{}
	_ bin.Decoder = &ChatParticipantCreator{}

	_ ChatParticipantClass = &ChatParticipantCreator{}
)

// ChatParticipantAdmin represents TL type `chatParticipantAdmin#e2d6e436`.
// Chat admin
//
// See https://core.telegram.org/constructor/chatParticipantAdmin for reference.
type ChatParticipantAdmin struct {
	// ID of a group member that is admin
	UserID int
	// ID of the user that added the member to the group
	InviterID int
	// Date when the user was added
	Date int
}

// ChatParticipantAdminTypeID is TL type id of ChatParticipantAdmin.
const ChatParticipantAdminTypeID = 0xe2d6e436

func (c *ChatParticipantAdmin) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.UserID == 0) {
		return false
	}
	if !(c.InviterID == 0) {
		return false
	}
	if !(c.Date == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatParticipantAdmin) String() string {
	if c == nil {
		return "ChatParticipantAdmin(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChatParticipantAdmin")
	sb.WriteString("{\n")
	sb.WriteString("\tUserID: ")
	sb.WriteString(fmt.Sprint(c.UserID))
	sb.WriteString(",\n")
	sb.WriteString("\tInviterID: ")
	sb.WriteString(fmt.Sprint(c.InviterID))
	sb.WriteString(",\n")
	sb.WriteString("\tDate: ")
	sb.WriteString(fmt.Sprint(c.Date))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *ChatParticipantAdmin) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatParticipantAdmin#e2d6e436 as nil")
	}
	b.PutID(ChatParticipantAdminTypeID)
	b.PutInt(c.UserID)
	b.PutInt(c.InviterID)
	b.PutInt(c.Date)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatParticipantAdmin) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatParticipantAdmin#e2d6e436 to nil")
	}
	if err := b.ConsumeID(ChatParticipantAdminTypeID); err != nil {
		return fmt.Errorf("unable to decode chatParticipantAdmin#e2d6e436: %w", err)
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipantAdmin#e2d6e436: field user_id: %w", err)
		}
		c.UserID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipantAdmin#e2d6e436: field inviter_id: %w", err)
		}
		c.InviterID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode chatParticipantAdmin#e2d6e436: field date: %w", err)
		}
		c.Date = value
	}
	return nil
}

// construct implements constructor of ChatParticipantClass.
func (c ChatParticipantAdmin) construct() ChatParticipantClass { return &c }

// Ensuring interfaces in compile-time for ChatParticipantAdmin.
var (
	_ bin.Encoder = &ChatParticipantAdmin{}
	_ bin.Decoder = &ChatParticipantAdmin{}

	_ ChatParticipantClass = &ChatParticipantAdmin{}
)

// ChatParticipantClass represents ChatParticipant generic type.
//
// See https://core.telegram.org/type/ChatParticipant for reference.
//
// Example:
//  g, err := DecodeChatParticipant(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *ChatParticipant: // chatParticipant#c8d7493e
//  case *ChatParticipantCreator: // chatParticipantCreator#da13538a
//  case *ChatParticipantAdmin: // chatParticipantAdmin#e2d6e436
//  default: panic(v)
//  }
type ChatParticipantClass interface {
	bin.Encoder
	bin.Decoder
	construct() ChatParticipantClass

	fmt.Stringer
	Zero() bool
}

// DecodeChatParticipant implements binary de-serialization for ChatParticipantClass.
func DecodeChatParticipant(buf *bin.Buffer) (ChatParticipantClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChatParticipantTypeID:
		// Decoding chatParticipant#c8d7493e.
		v := ChatParticipant{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatParticipantClass: %w", err)
		}
		return &v, nil
	case ChatParticipantCreatorTypeID:
		// Decoding chatParticipantCreator#da13538a.
		v := ChatParticipantCreator{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatParticipantClass: %w", err)
		}
		return &v, nil
	case ChatParticipantAdminTypeID:
		// Decoding chatParticipantAdmin#e2d6e436.
		v := ChatParticipantAdmin{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChatParticipantClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChatParticipantClass: %w", bin.NewUnexpectedID(id))
	}
}

// ChatParticipant boxes the ChatParticipantClass providing a helper.
type ChatParticipantBox struct {
	ChatParticipant ChatParticipantClass
}

// Decode implements bin.Decoder for ChatParticipantBox.
func (b *ChatParticipantBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChatParticipantBox to nil")
	}
	v, err := DecodeChatParticipant(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChatParticipant = v
	return nil
}

// Encode implements bin.Encode for ChatParticipantBox.
func (b *ChatParticipantBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChatParticipant == nil {
		return fmt.Errorf("unable to encode ChatParticipantClass as nil")
	}
	return b.ChatParticipant.Encode(buf)
}
