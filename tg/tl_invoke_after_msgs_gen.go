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

// InvokeAfterMsgsRequest represents TL type `invokeAfterMsgs#3dc4b4f0`.
// Invokes a query after a successfull completion of previous queries
//
// See https://core.telegram.org/constructor/invokeAfterMsgs for reference.
type InvokeAfterMsgsRequest struct {
	// List of messages on which a current query depends
	MsgIds []int64
	// The query itself
	Query bin.Object
}

// InvokeAfterMsgsRequestTypeID is TL type id of InvokeAfterMsgsRequest.
const InvokeAfterMsgsRequestTypeID = 0x3dc4b4f0

func (i *InvokeAfterMsgsRequest) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.MsgIds == nil) {
		return false
	}
	if !(i.Query == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InvokeAfterMsgsRequest) String() string {
	if i == nil {
		return "InvokeAfterMsgsRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InvokeAfterMsgsRequest")
	sb.WriteString("{\n")
	sb.WriteByte('[')
	for _, v := range i.MsgIds {
		sb.WriteString(fmt.Sprint(v))
	}
	sb.WriteByte(']')
	sb.WriteString("\tQuery: ")
	sb.WriteString(fmt.Sprint(i.Query))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InvokeAfterMsgsRequest) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode invokeAfterMsgs#3dc4b4f0 as nil")
	}
	b.PutID(InvokeAfterMsgsRequestTypeID)
	b.PutVectorHeader(len(i.MsgIds))
	for _, v := range i.MsgIds {
		b.PutLong(v)
	}
	if err := i.Query.Encode(b); err != nil {
		return fmt.Errorf("unable to encode invokeAfterMsgs#3dc4b4f0: field query: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InvokeAfterMsgsRequest) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode invokeAfterMsgs#3dc4b4f0 to nil")
	}
	if err := b.ConsumeID(InvokeAfterMsgsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode invokeAfterMsgs#3dc4b4f0: %w", err)
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode invokeAfterMsgs#3dc4b4f0: field msg_ids: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode invokeAfterMsgs#3dc4b4f0: field msg_ids: %w", err)
			}
			i.MsgIds = append(i.MsgIds, value)
		}
	}
	{
		if err := i.Query.Decode(b); err != nil {
			return fmt.Errorf("unable to decode invokeAfterMsgs#3dc4b4f0: field query: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for InvokeAfterMsgsRequest.
var (
	_ bin.Encoder = &InvokeAfterMsgsRequest{}
	_ bin.Decoder = &InvokeAfterMsgsRequest{}
)
