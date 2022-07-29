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
	"github.com/gotd/td/tdjson"
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
	_ = tdjson.Encoder{}
)

// PremiumGiftOption represents TL type `premiumGiftOption#74c34319`.
//
// See https://core.telegram.org/constructor/premiumGiftOption for reference.
type PremiumGiftOption struct {
	// Flags field of PremiumGiftOption.
	Flags bin.Fields
	// Months field of PremiumGiftOption.
	Months int
	// Currency field of PremiumGiftOption.
	Currency string
	// Amount field of PremiumGiftOption.
	Amount int64
	// BotURL field of PremiumGiftOption.
	BotURL string
	// StoreProduct field of PremiumGiftOption.
	//
	// Use SetStoreProduct and GetStoreProduct helpers.
	StoreProduct string
}

// PremiumGiftOptionTypeID is TL type id of PremiumGiftOption.
const PremiumGiftOptionTypeID = 0x74c34319

// Ensuring interfaces in compile-time for PremiumGiftOption.
var (
	_ bin.Encoder     = &PremiumGiftOption{}
	_ bin.Decoder     = &PremiumGiftOption{}
	_ bin.BareEncoder = &PremiumGiftOption{}
	_ bin.BareDecoder = &PremiumGiftOption{}
)

func (p *PremiumGiftOption) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Flags.Zero()) {
		return false
	}
	if !(p.Months == 0) {
		return false
	}
	if !(p.Currency == "") {
		return false
	}
	if !(p.Amount == 0) {
		return false
	}
	if !(p.BotURL == "") {
		return false
	}
	if !(p.StoreProduct == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PremiumGiftOption) String() string {
	if p == nil {
		return "PremiumGiftOption(nil)"
	}
	type Alias PremiumGiftOption
	return fmt.Sprintf("PremiumGiftOption%+v", Alias(*p))
}

// FillFrom fills PremiumGiftOption from given interface.
func (p *PremiumGiftOption) FillFrom(from interface {
	GetMonths() (value int)
	GetCurrency() (value string)
	GetAmount() (value int64)
	GetBotURL() (value string)
	GetStoreProduct() (value string, ok bool)
}) {
	p.Months = from.GetMonths()
	p.Currency = from.GetCurrency()
	p.Amount = from.GetAmount()
	p.BotURL = from.GetBotURL()
	if val, ok := from.GetStoreProduct(); ok {
		p.StoreProduct = val
	}

}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PremiumGiftOption) TypeID() uint32 {
	return PremiumGiftOptionTypeID
}

// TypeName returns name of type in TL schema.
func (*PremiumGiftOption) TypeName() string {
	return "premiumGiftOption"
}

// TypeInfo returns info about TL type.
func (p *PremiumGiftOption) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "premiumGiftOption",
		ID:   PremiumGiftOptionTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Months",
			SchemaName: "months",
		},
		{
			Name:       "Currency",
			SchemaName: "currency",
		},
		{
			Name:       "Amount",
			SchemaName: "amount",
		},
		{
			Name:       "BotURL",
			SchemaName: "bot_url",
		},
		{
			Name:       "StoreProduct",
			SchemaName: "store_product",
			Null:       !p.Flags.Has(0),
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (p *PremiumGiftOption) SetFlags() {
	if !(p.StoreProduct == "") {
		p.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (p *PremiumGiftOption) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumGiftOption#74c34319 as nil")
	}
	b.PutID(PremiumGiftOptionTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PremiumGiftOption) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode premiumGiftOption#74c34319 as nil")
	}
	p.SetFlags()
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode premiumGiftOption#74c34319: field flags: %w", err)
	}
	b.PutInt(p.Months)
	b.PutString(p.Currency)
	b.PutLong(p.Amount)
	b.PutString(p.BotURL)
	if p.Flags.Has(0) {
		b.PutString(p.StoreProduct)
	}
	return nil
}

// Decode implements bin.Decoder.
func (p *PremiumGiftOption) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumGiftOption#74c34319 to nil")
	}
	if err := b.ConsumeID(PremiumGiftOptionTypeID); err != nil {
		return fmt.Errorf("unable to decode premiumGiftOption#74c34319: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PremiumGiftOption) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode premiumGiftOption#74c34319 to nil")
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode premiumGiftOption#74c34319: field flags: %w", err)
		}
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode premiumGiftOption#74c34319: field months: %w", err)
		}
		p.Months = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode premiumGiftOption#74c34319: field currency: %w", err)
		}
		p.Currency = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode premiumGiftOption#74c34319: field amount: %w", err)
		}
		p.Amount = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode premiumGiftOption#74c34319: field bot_url: %w", err)
		}
		p.BotURL = value
	}
	if p.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode premiumGiftOption#74c34319: field store_product: %w", err)
		}
		p.StoreProduct = value
	}
	return nil
}

// GetMonths returns value of Months field.
func (p *PremiumGiftOption) GetMonths() (value int) {
	if p == nil {
		return
	}
	return p.Months
}

// GetCurrency returns value of Currency field.
func (p *PremiumGiftOption) GetCurrency() (value string) {
	if p == nil {
		return
	}
	return p.Currency
}

// GetAmount returns value of Amount field.
func (p *PremiumGiftOption) GetAmount() (value int64) {
	if p == nil {
		return
	}
	return p.Amount
}

// GetBotURL returns value of BotURL field.
func (p *PremiumGiftOption) GetBotURL() (value string) {
	if p == nil {
		return
	}
	return p.BotURL
}

// SetStoreProduct sets value of StoreProduct conditional field.
func (p *PremiumGiftOption) SetStoreProduct(value string) {
	p.Flags.Set(0)
	p.StoreProduct = value
}

// GetStoreProduct returns value of StoreProduct conditional field and
// boolean which is true if field was set.
func (p *PremiumGiftOption) GetStoreProduct() (value string, ok bool) {
	if p == nil {
		return
	}
	if !p.Flags.Has(0) {
		return value, false
	}
	return p.StoreProduct, true
}
