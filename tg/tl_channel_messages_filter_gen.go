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

// ChannelMessagesFilterEmpty represents TL type `channelMessagesFilterEmpty#94d42ee7`.
// No filter
//
// See https://core.telegram.org/constructor/channelMessagesFilterEmpty for reference.
type ChannelMessagesFilterEmpty struct {
}

// ChannelMessagesFilterEmptyTypeID is TL type id of ChannelMessagesFilterEmpty.
const ChannelMessagesFilterEmptyTypeID = 0x94d42ee7

func (c *ChannelMessagesFilterEmpty) Zero() bool {
	if c == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelMessagesFilterEmpty) String() string {
	if c == nil {
		return "ChannelMessagesFilterEmpty(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelMessagesFilterEmpty")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *ChannelMessagesFilterEmpty) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelMessagesFilterEmpty#94d42ee7 as nil")
	}
	b.PutID(ChannelMessagesFilterEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChannelMessagesFilterEmpty) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelMessagesFilterEmpty#94d42ee7 to nil")
	}
	if err := b.ConsumeID(ChannelMessagesFilterEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode channelMessagesFilterEmpty#94d42ee7: %w", err)
	}
	return nil
}

// construct implements constructor of ChannelMessagesFilterClass.
func (c ChannelMessagesFilterEmpty) construct() ChannelMessagesFilterClass { return &c }

// Ensuring interfaces in compile-time for ChannelMessagesFilterEmpty.
var (
	_ bin.Encoder = &ChannelMessagesFilterEmpty{}
	_ bin.Decoder = &ChannelMessagesFilterEmpty{}

	_ ChannelMessagesFilterClass = &ChannelMessagesFilterEmpty{}
)

// ChannelMessagesFilter represents TL type `channelMessagesFilter#cd77d957`.
// Filter for getting only certain types of channel messages
//
// See https://core.telegram.org/constructor/channelMessagesFilter for reference.
type ChannelMessagesFilter struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether to exclude new messages from the search
	ExcludeNewMessages bool
	// A range of messages to fetch
	Ranges []MessageRange
}

// ChannelMessagesFilterTypeID is TL type id of ChannelMessagesFilter.
const ChannelMessagesFilterTypeID = 0xcd77d957

func (c *ChannelMessagesFilter) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Flags.Zero()) {
		return false
	}
	if !(c.ExcludeNewMessages == false) {
		return false
	}
	if !(c.Ranges == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelMessagesFilter) String() string {
	if c == nil {
		return "ChannelMessagesFilter(nil)"
	}
	var sb strings.Builder
	sb.WriteString("ChannelMessagesFilter")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(c.Flags))
	sb.WriteString(",\n")
	sb.WriteByte('[')
	for _, v := range c.Ranges {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (c *ChannelMessagesFilter) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channelMessagesFilter#cd77d957 as nil")
	}
	b.PutID(ChannelMessagesFilterTypeID)
	if !(c.ExcludeNewMessages == false) {
		c.Flags.Set(1)
	}
	if err := c.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channelMessagesFilter#cd77d957: field flags: %w", err)
	}
	b.PutVectorHeader(len(c.Ranges))
	for idx, v := range c.Ranges {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channelMessagesFilter#cd77d957: field ranges element with index %d: %w", idx, err)
		}
	}
	return nil
}

// SetExcludeNewMessages sets value of ExcludeNewMessages conditional field.
func (c *ChannelMessagesFilter) SetExcludeNewMessages(value bool) {
	if value {
		c.Flags.Set(1)
		c.ExcludeNewMessages = true
	} else {
		c.Flags.Unset(1)
		c.ExcludeNewMessages = false
	}
}

// Decode implements bin.Decoder.
func (c *ChannelMessagesFilter) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channelMessagesFilter#cd77d957 to nil")
	}
	if err := b.ConsumeID(ChannelMessagesFilterTypeID); err != nil {
		return fmt.Errorf("unable to decode channelMessagesFilter#cd77d957: %w", err)
	}
	{
		if err := c.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channelMessagesFilter#cd77d957: field flags: %w", err)
		}
	}
	c.ExcludeNewMessages = c.Flags.Has(1)
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channelMessagesFilter#cd77d957: field ranges: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value MessageRange
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode channelMessagesFilter#cd77d957: field ranges: %w", err)
			}
			c.Ranges = append(c.Ranges, value)
		}
	}
	return nil
}

// construct implements constructor of ChannelMessagesFilterClass.
func (c ChannelMessagesFilter) construct() ChannelMessagesFilterClass { return &c }

// Ensuring interfaces in compile-time for ChannelMessagesFilter.
var (
	_ bin.Encoder = &ChannelMessagesFilter{}
	_ bin.Decoder = &ChannelMessagesFilter{}

	_ ChannelMessagesFilterClass = &ChannelMessagesFilter{}
)

// ChannelMessagesFilterClass represents ChannelMessagesFilter generic type.
//
// See https://core.telegram.org/type/ChannelMessagesFilter for reference.
//
// Example:
//  g, err := DecodeChannelMessagesFilter(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *ChannelMessagesFilterEmpty: // channelMessagesFilterEmpty#94d42ee7
//  case *ChannelMessagesFilter: // channelMessagesFilter#cd77d957
//  default: panic(v)
//  }
type ChannelMessagesFilterClass interface {
	bin.Encoder
	bin.Decoder
	construct() ChannelMessagesFilterClass

	fmt.Stringer
	Zero() bool
}

// DecodeChannelMessagesFilter implements binary de-serialization for ChannelMessagesFilterClass.
func DecodeChannelMessagesFilter(buf *bin.Buffer) (ChannelMessagesFilterClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case ChannelMessagesFilterEmptyTypeID:
		// Decoding channelMessagesFilterEmpty#94d42ee7.
		v := ChannelMessagesFilterEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChannelMessagesFilterClass: %w", err)
		}
		return &v, nil
	case ChannelMessagesFilterTypeID:
		// Decoding channelMessagesFilter#cd77d957.
		v := ChannelMessagesFilter{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ChannelMessagesFilterClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ChannelMessagesFilterClass: %w", bin.NewUnexpectedID(id))
	}
}

// ChannelMessagesFilter boxes the ChannelMessagesFilterClass providing a helper.
type ChannelMessagesFilterBox struct {
	ChannelMessagesFilter ChannelMessagesFilterClass
}

// Decode implements bin.Decoder for ChannelMessagesFilterBox.
func (b *ChannelMessagesFilterBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ChannelMessagesFilterBox to nil")
	}
	v, err := DecodeChannelMessagesFilter(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ChannelMessagesFilter = v
	return nil
}

// Encode implements bin.Encode for ChannelMessagesFilterBox.
func (b *ChannelMessagesFilterBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ChannelMessagesFilter == nil {
		return fmt.Errorf("unable to encode ChannelMessagesFilterClass as nil")
	}
	return b.ChannelMessagesFilter.Encode(buf)
}
