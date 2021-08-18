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

// AccountWebAuthorizations represents TL type `account.webAuthorizations#ed56c9fc`.
// Web authorizations
//
// See https://core.telegram.org/constructor/account.webAuthorizations for reference.
type AccountWebAuthorizations struct {
	// Web authorization list
	Authorizations []WebAuthorization
	// Users
	Users []UserClass
}

// AccountWebAuthorizationsTypeID is TL type id of AccountWebAuthorizations.
const AccountWebAuthorizationsTypeID = 0xed56c9fc

func (w *AccountWebAuthorizations) Zero() bool {
	if w == nil {
		return true
	}
	if !(w.Authorizations == nil) {
		return false
	}
	if !(w.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (w *AccountWebAuthorizations) String() string {
	if w == nil {
		return "AccountWebAuthorizations(nil)"
	}
	type Alias AccountWebAuthorizations
	return fmt.Sprintf("AccountWebAuthorizations%+v", Alias(*w))
}

// FillFrom fills AccountWebAuthorizations from given interface.
func (w *AccountWebAuthorizations) FillFrom(from interface {
	GetAuthorizations() (value []WebAuthorization)
	GetUsers() (value []UserClass)
}) {
	w.Authorizations = from.GetAuthorizations()
	w.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountWebAuthorizations) TypeID() uint32 {
	return AccountWebAuthorizationsTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountWebAuthorizations) TypeName() string {
	return "account.webAuthorizations"
}

// TypeInfo returns info about TL type.
func (w *AccountWebAuthorizations) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.webAuthorizations",
		ID:   AccountWebAuthorizationsTypeID,
	}
	if w == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Authorizations",
			SchemaName: "authorizations",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (w *AccountWebAuthorizations) Encode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode account.webAuthorizations#ed56c9fc as nil")
	}
	b.PutID(AccountWebAuthorizationsTypeID)
	return w.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (w *AccountWebAuthorizations) EncodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't encode account.webAuthorizations#ed56c9fc as nil")
	}
	b.PutVectorHeader(len(w.Authorizations))
	for idx, v := range w.Authorizations {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.webAuthorizations#ed56c9fc: field authorizations element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(w.Users))
	for idx, v := range w.Users {
		if v == nil {
			return fmt.Errorf("unable to encode account.webAuthorizations#ed56c9fc: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.webAuthorizations#ed56c9fc: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetAuthorizations returns value of Authorizations field.
func (w *AccountWebAuthorizations) GetAuthorizations() (value []WebAuthorization) {
	return w.Authorizations
}

// GetUsers returns value of Users field.
func (w *AccountWebAuthorizations) GetUsers() (value []UserClass) {
	return w.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (w *AccountWebAuthorizations) MapUsers() (value UserClassArray) {
	return UserClassArray(w.Users)
}

// Decode implements bin.Decoder.
func (w *AccountWebAuthorizations) Decode(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode account.webAuthorizations#ed56c9fc to nil")
	}
	if err := b.ConsumeID(AccountWebAuthorizationsTypeID); err != nil {
		return fmt.Errorf("unable to decode account.webAuthorizations#ed56c9fc: %w", err)
	}
	return w.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (w *AccountWebAuthorizations) DecodeBare(b *bin.Buffer) error {
	if w == nil {
		return fmt.Errorf("can't decode account.webAuthorizations#ed56c9fc to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.webAuthorizations#ed56c9fc: field authorizations: %w", err)
		}

		if headerLen > 0 {
			w.Authorizations = make([]WebAuthorization, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value WebAuthorization
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode account.webAuthorizations#ed56c9fc: field authorizations: %w", err)
			}
			w.Authorizations = append(w.Authorizations, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.webAuthorizations#ed56c9fc: field users: %w", err)
		}

		if headerLen > 0 {
			w.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode account.webAuthorizations#ed56c9fc: field users: %w", err)
			}
			w.Users = append(w.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountWebAuthorizations.
var (
	_ bin.Encoder     = &AccountWebAuthorizations{}
	_ bin.Decoder     = &AccountWebAuthorizations{}
	_ bin.BareEncoder = &AccountWebAuthorizations{}
	_ bin.BareDecoder = &AccountWebAuthorizations{}
)
