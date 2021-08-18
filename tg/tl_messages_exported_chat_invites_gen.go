// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

// MessagesExportedChatInvites represents TL type `messages.exportedChatInvites#bdc62dcc`.
//
// See https://core.telegram.org/constructor/messages.exportedChatInvites for reference.
type MessagesExportedChatInvites struct {
	// Count field of MessagesExportedChatInvites.
	Count int
	// Invites field of MessagesExportedChatInvites.
	Invites []ChatInviteExported
	// Users field of MessagesExportedChatInvites.
	Users []UserClass
}

// MessagesExportedChatInvitesTypeID is TL type id of MessagesExportedChatInvites.
const MessagesExportedChatInvitesTypeID = 0xbdc62dcc

func (e *MessagesExportedChatInvites) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.Count == 0) {
		return false
	}
	if !(e.Invites == nil) {
		return false
	}
	if !(e.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *MessagesExportedChatInvites) String() string {
	if e == nil {
		return "MessagesExportedChatInvites(nil)"
	}
	type Alias MessagesExportedChatInvites
	return fmt.Sprintf("MessagesExportedChatInvites%+v", Alias(*e))
}

// FillFrom fills MessagesExportedChatInvites from given interface.
func (e *MessagesExportedChatInvites) FillFrom(from interface {
	GetCount() (value int)
	GetInvites() (value []ChatInviteExported)
	GetUsers() (value []UserClass)
}) {
	e.Count = from.GetCount()
	e.Invites = from.GetInvites()
	e.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesExportedChatInvites) TypeID() uint32 {
	return MessagesExportedChatInvitesTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesExportedChatInvites) TypeName() string {
	return "messages.exportedChatInvites"
}

// TypeInfo returns info about TL type.
func (e *MessagesExportedChatInvites) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.exportedChatInvites",
		ID:   MessagesExportedChatInvitesTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "Invites",
			SchemaName: "invites",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *MessagesExportedChatInvites) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.exportedChatInvites#bdc62dcc as nil")
	}
	b.PutID(MessagesExportedChatInvitesTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *MessagesExportedChatInvites) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode messages.exportedChatInvites#bdc62dcc as nil")
	}
	b.PutInt(e.Count)
	b.PutVectorHeader(len(e.Invites))
	for idx, v := range e.Invites {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.exportedChatInvites#bdc62dcc: field invites element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(e.Users))
	for idx, v := range e.Users {
		if v == nil {
			return fmt.Errorf("unable to encode messages.exportedChatInvites#bdc62dcc: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.exportedChatInvites#bdc62dcc: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetCount returns value of Count field.
func (e *MessagesExportedChatInvites) GetCount() (value int) {
	return e.Count
}

// GetInvites returns value of Invites field.
func (e *MessagesExportedChatInvites) GetInvites() (value []ChatInviteExported) {
	return e.Invites
}

// GetUsers returns value of Users field.
func (e *MessagesExportedChatInvites) GetUsers() (value []UserClass) {
	return e.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (e *MessagesExportedChatInvites) MapUsers() (value UserClassArray) {
	return UserClassArray(e.Users)
}

// Decode implements bin.Decoder.
func (e *MessagesExportedChatInvites) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.exportedChatInvites#bdc62dcc to nil")
	}
	if err := b.ConsumeID(MessagesExportedChatInvitesTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.exportedChatInvites#bdc62dcc: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *MessagesExportedChatInvites) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode messages.exportedChatInvites#bdc62dcc to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.exportedChatInvites#bdc62dcc: field count: %w", err)
		}
		e.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.exportedChatInvites#bdc62dcc: field invites: %w", err)
		}

		if headerLen > 0 {
			e.Invites = make([]ChatInviteExported, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ChatInviteExported
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.exportedChatInvites#bdc62dcc: field invites: %w", err)
			}
			e.Invites = append(e.Invites, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.exportedChatInvites#bdc62dcc: field users: %w", err)
		}

		if headerLen > 0 {
			e.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode messages.exportedChatInvites#bdc62dcc: field users: %w", err)
			}
			e.Users = append(e.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for MessagesExportedChatInvites.
var (
	_ bin.Encoder     = &MessagesExportedChatInvites{}
	_ bin.Decoder     = &MessagesExportedChatInvites{}
	_ bin.BareEncoder = &MessagesExportedChatInvites{}
	_ bin.BareDecoder = &MessagesExportedChatInvites{}
)
