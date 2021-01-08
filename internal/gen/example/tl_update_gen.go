// Code generated by gotdgen, DO NOT EDIT.

package td

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

// Update represents TL type `update#b03e2ef8`.
//
// See https://localhost:80/doc/constructor/update for reference.
type Update struct {
	// Msg field of Update.
	Msg AbstractMessageClass
	// Delay field of Update.
	Delay int32
}

// UpdateTypeID is TL type id of Update.
const UpdateTypeID = 0xb03e2ef8

func (u *Update) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Msg == nil) {
		return false
	}
	if !(u.Delay == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *Update) String() string {
	if u == nil {
		return "Update(nil)"
	}
	var sb strings.Builder
	sb.WriteString("Update")
	sb.WriteString("{\n")
	sb.WriteString("\tMsg: ")
	sb.WriteString(fmt.Sprint(u.Msg))
	sb.WriteString(",\n")
	sb.WriteString("\tDelay: ")
	sb.WriteString(fmt.Sprint(u.Delay))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (u *Update) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode update#b03e2ef8 as nil")
	}
	b.PutID(UpdateTypeID)
	if u.Msg == nil {
		return fmt.Errorf("unable to encode update#b03e2ef8: field msg is nil")
	}
	if err := u.Msg.Encode(b); err != nil {
		return fmt.Errorf("unable to encode update#b03e2ef8: field msg: %w", err)
	}
	b.PutInt32(u.Delay)
	return nil
}

// Decode implements bin.Decoder.
func (u *Update) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode update#b03e2ef8 to nil")
	}
	if err := b.ConsumeID(UpdateTypeID); err != nil {
		return fmt.Errorf("unable to decode update#b03e2ef8: %w", err)
	}
	{
		value, err := DecodeAbstractMessage(b)
		if err != nil {
			return fmt.Errorf("unable to decode update#b03e2ef8: field msg: %w", err)
		}
		u.Msg = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode update#b03e2ef8: field delay: %w", err)
		}
		u.Delay = value
	}
	return nil
}

// Ensuring interfaces in compile-time for Update.
var (
	_ bin.Encoder = &Update{}
	_ bin.Decoder = &Update{}
)
