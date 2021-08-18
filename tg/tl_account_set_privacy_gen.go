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

// AccountSetPrivacyRequest represents TL type `account.setPrivacy#c9f81ce8`.
// Change privacy settings of current account
//
// See https://core.telegram.org/method/account.setPrivacy for reference.
type AccountSetPrivacyRequest struct {
	// Peers to which the privacy rules apply
	Key InputPrivacyKeyClass
	// New privacy rules
	Rules []InputPrivacyRuleClass
}

// AccountSetPrivacyRequestTypeID is TL type id of AccountSetPrivacyRequest.
const AccountSetPrivacyRequestTypeID = 0xc9f81ce8

func (s *AccountSetPrivacyRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Key == nil) {
		return false
	}
	if !(s.Rules == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSetPrivacyRequest) String() string {
	if s == nil {
		return "AccountSetPrivacyRequest(nil)"
	}
	type Alias AccountSetPrivacyRequest
	return fmt.Sprintf("AccountSetPrivacyRequest%+v", Alias(*s))
}

// FillFrom fills AccountSetPrivacyRequest from given interface.
func (s *AccountSetPrivacyRequest) FillFrom(from interface {
	GetKey() (value InputPrivacyKeyClass)
	GetRules() (value []InputPrivacyRuleClass)
}) {
	s.Key = from.GetKey()
	s.Rules = from.GetRules()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountSetPrivacyRequest) TypeID() uint32 {
	return AccountSetPrivacyRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountSetPrivacyRequest) TypeName() string {
	return "account.setPrivacy"
}

// TypeInfo returns info about TL type.
func (s *AccountSetPrivacyRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.setPrivacy",
		ID:   AccountSetPrivacyRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Key",
			SchemaName: "key",
		},
		{
			Name:       "Rules",
			SchemaName: "rules",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *AccountSetPrivacyRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.setPrivacy#c9f81ce8 as nil")
	}
	b.PutID(AccountSetPrivacyRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *AccountSetPrivacyRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.setPrivacy#c9f81ce8 as nil")
	}
	if s.Key == nil {
		return fmt.Errorf("unable to encode account.setPrivacy#c9f81ce8: field key is nil")
	}
	if err := s.Key.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.setPrivacy#c9f81ce8: field key: %w", err)
	}
	b.PutVectorHeader(len(s.Rules))
	for idx, v := range s.Rules {
		if v == nil {
			return fmt.Errorf("unable to encode account.setPrivacy#c9f81ce8: field rules element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.setPrivacy#c9f81ce8: field rules element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetKey returns value of Key field.
func (s *AccountSetPrivacyRequest) GetKey() (value InputPrivacyKeyClass) {
	return s.Key
}

// GetRules returns value of Rules field.
func (s *AccountSetPrivacyRequest) GetRules() (value []InputPrivacyRuleClass) {
	return s.Rules
}

// MapRules returns field Rules wrapped in InputPrivacyRuleClassArray helper.
func (s *AccountSetPrivacyRequest) MapRules() (value InputPrivacyRuleClassArray) {
	return InputPrivacyRuleClassArray(s.Rules)
}

// Decode implements bin.Decoder.
func (s *AccountSetPrivacyRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.setPrivacy#c9f81ce8 to nil")
	}
	if err := b.ConsumeID(AccountSetPrivacyRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.setPrivacy#c9f81ce8: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *AccountSetPrivacyRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.setPrivacy#c9f81ce8 to nil")
	}
	{
		value, err := DecodeInputPrivacyKey(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.setPrivacy#c9f81ce8: field key: %w", err)
		}
		s.Key = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode account.setPrivacy#c9f81ce8: field rules: %w", err)
		}

		if headerLen > 0 {
			s.Rules = make([]InputPrivacyRuleClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeInputPrivacyRule(b)
			if err != nil {
				return fmt.Errorf("unable to decode account.setPrivacy#c9f81ce8: field rules: %w", err)
			}
			s.Rules = append(s.Rules, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountSetPrivacyRequest.
var (
	_ bin.Encoder     = &AccountSetPrivacyRequest{}
	_ bin.Decoder     = &AccountSetPrivacyRequest{}
	_ bin.BareEncoder = &AccountSetPrivacyRequest{}
	_ bin.BareDecoder = &AccountSetPrivacyRequest{}
)

// AccountSetPrivacy invokes method account.setPrivacy#c9f81ce8 returning error if any.
// Change privacy settings of current account
//
// Possible errors:
//  400 PRIVACY_KEY_INVALID: The privacy key is invalid
//  400 PRIVACY_VALUE_INVALID: The specified privacy rule combination is invalid
//
// See https://core.telegram.org/method/account.setPrivacy for reference.
func (c *Client) AccountSetPrivacy(ctx context.Context, request *AccountSetPrivacyRequest) (*AccountPrivacyRules, error) {
	var result AccountPrivacyRules

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
