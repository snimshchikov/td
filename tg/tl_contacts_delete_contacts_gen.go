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

// ContactsDeleteContactsRequest represents TL type `contacts.deleteContacts#96a0e00`.
// Deletes several contacts from the list.
//
// See https://core.telegram.org/method/contacts.deleteContacts for reference.
type ContactsDeleteContactsRequest struct {
	// User ID list
	ID []InputUserClass
}

// ContactsDeleteContactsRequestTypeID is TL type id of ContactsDeleteContactsRequest.
const ContactsDeleteContactsRequestTypeID = 0x96a0e00

func (d *ContactsDeleteContactsRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *ContactsDeleteContactsRequest) String() string {
	if d == nil {
		return "ContactsDeleteContactsRequest(nil)"
	}
	type Alias ContactsDeleteContactsRequest
	return fmt.Sprintf("ContactsDeleteContactsRequest%+v", Alias(*d))
}

// FillFrom fills ContactsDeleteContactsRequest from given interface.
func (d *ContactsDeleteContactsRequest) FillFrom(from interface {
	GetID() (value []InputUserClass)
}) {
	d.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ContactsDeleteContactsRequest) TypeID() uint32 {
	return ContactsDeleteContactsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ContactsDeleteContactsRequest) TypeName() string {
	return "contacts.deleteContacts"
}

// TypeInfo returns info about TL type.
func (d *ContactsDeleteContactsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "contacts.deleteContacts",
		ID:   ContactsDeleteContactsRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (d *ContactsDeleteContactsRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode contacts.deleteContacts#96a0e00 as nil")
	}
	b.PutID(ContactsDeleteContactsRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *ContactsDeleteContactsRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode contacts.deleteContacts#96a0e00 as nil")
	}
	b.PutVectorHeader(len(d.ID))
	for idx, v := range d.ID {
		if v == nil {
			return fmt.Errorf("unable to encode contacts.deleteContacts#96a0e00: field id element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode contacts.deleteContacts#96a0e00: field id element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetID returns value of ID field.
func (d *ContactsDeleteContactsRequest) GetID() (value []InputUserClass) {
	return d.ID
}

// MapID returns field ID wrapped in InputUserClassArray helper.
func (d *ContactsDeleteContactsRequest) MapID() (value InputUserClassArray) {
	return InputUserClassArray(d.ID)
}

// Decode implements bin.Decoder.
func (d *ContactsDeleteContactsRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode contacts.deleteContacts#96a0e00 to nil")
	}
	if err := b.ConsumeID(ContactsDeleteContactsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode contacts.deleteContacts#96a0e00: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *ContactsDeleteContactsRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode contacts.deleteContacts#96a0e00 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode contacts.deleteContacts#96a0e00: field id: %w", err)
		}

		if headerLen > 0 {
			d.ID = make([]InputUserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode contacts.deleteContacts#96a0e00: field id: %w", err)
			}
			d.ID = append(d.ID, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for ContactsDeleteContactsRequest.
var (
	_ bin.Encoder     = &ContactsDeleteContactsRequest{}
	_ bin.Decoder     = &ContactsDeleteContactsRequest{}
	_ bin.BareEncoder = &ContactsDeleteContactsRequest{}
	_ bin.BareDecoder = &ContactsDeleteContactsRequest{}
)

// ContactsDeleteContacts invokes method contacts.deleteContacts#96a0e00 returning error if any.
// Deletes several contacts from the list.
//
// See https://core.telegram.org/method/contacts.deleteContacts for reference.
func (c *Client) ContactsDeleteContacts(ctx context.Context, id []InputUserClass) (UpdatesClass, error) {
	var result UpdatesBox

	request := &ContactsDeleteContactsRequest{
		ID: id,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
