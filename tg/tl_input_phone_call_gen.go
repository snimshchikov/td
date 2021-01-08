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

// InputPhoneCall represents TL type `inputPhoneCall#1e36fded`.
// Phone call
//
// See https://core.telegram.org/constructor/inputPhoneCall for reference.
type InputPhoneCall struct {
	// Call ID
	ID int64
	// Access hash
	AccessHash int64
}

// InputPhoneCallTypeID is TL type id of InputPhoneCall.
const InputPhoneCallTypeID = 0x1e36fded

func (i *InputPhoneCall) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == 0) {
		return false
	}
	if !(i.AccessHash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPhoneCall) String() string {
	if i == nil {
		return "InputPhoneCall(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputPhoneCall")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(i.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tAccessHash: ")
	sb.WriteString(fmt.Sprint(i.AccessHash))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InputPhoneCall) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPhoneCall#1e36fded as nil")
	}
	b.PutID(InputPhoneCallTypeID)
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPhoneCall) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPhoneCall#1e36fded to nil")
	}
	if err := b.ConsumeID(InputPhoneCallTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPhoneCall#1e36fded: %w", err)
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhoneCall#1e36fded: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputPhoneCall#1e36fded: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for InputPhoneCall.
var (
	_ bin.Encoder = &InputPhoneCall{}
	_ bin.Decoder = &InputPhoneCall{}
)
