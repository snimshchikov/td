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

// MessagesAllStickersNotModified represents TL type `messages.allStickersNotModified#e86602c3`.
// Info about all installed stickers hasn't changed
//
// See https://core.telegram.org/constructor/messages.allStickersNotModified for reference.
type MessagesAllStickersNotModified struct {
}

// MessagesAllStickersNotModifiedTypeID is TL type id of MessagesAllStickersNotModified.
const MessagesAllStickersNotModifiedTypeID = 0xe86602c3

func (a *MessagesAllStickersNotModified) Zero() bool {
	if a == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (a *MessagesAllStickersNotModified) String() string {
	if a == nil {
		return "MessagesAllStickersNotModified(nil)"
	}
	type Alias MessagesAllStickersNotModified
	return fmt.Sprintf("MessagesAllStickersNotModified%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesAllStickersNotModified) TypeID() uint32 {
	return MessagesAllStickersNotModifiedTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesAllStickersNotModified) TypeName() string {
	return "messages.allStickersNotModified"
}

// TypeInfo returns info about TL type.
func (a *MessagesAllStickersNotModified) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.allStickersNotModified",
		ID:   MessagesAllStickersNotModifiedTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (a *MessagesAllStickersNotModified) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode messages.allStickersNotModified#e86602c3 as nil")
	}
	b.PutID(MessagesAllStickersNotModifiedTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *MessagesAllStickersNotModified) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode messages.allStickersNotModified#e86602c3 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *MessagesAllStickersNotModified) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode messages.allStickersNotModified#e86602c3 to nil")
	}
	if err := b.ConsumeID(MessagesAllStickersNotModifiedTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.allStickersNotModified#e86602c3: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *MessagesAllStickersNotModified) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode messages.allStickersNotModified#e86602c3 to nil")
	}
	return nil
}

// construct implements constructor of MessagesAllStickersClass.
func (a MessagesAllStickersNotModified) construct() MessagesAllStickersClass { return &a }

// Ensuring interfaces in compile-time for MessagesAllStickersNotModified.
var (
	_ bin.Encoder     = &MessagesAllStickersNotModified{}
	_ bin.Decoder     = &MessagesAllStickersNotModified{}
	_ bin.BareEncoder = &MessagesAllStickersNotModified{}
	_ bin.BareDecoder = &MessagesAllStickersNotModified{}

	_ MessagesAllStickersClass = &MessagesAllStickersNotModified{}
)

// MessagesAllStickers represents TL type `messages.allStickers#edfd405f`.
// Info about all installed stickers
//
// See https://core.telegram.org/constructor/messages.allStickers for reference.
type MessagesAllStickers struct {
	// Hash for pagination, for more info click here¹
	//
	// Links:
	//  1) https://core.telegram.org/api/offsets#hash-generation
	Hash int
	// All stickersets
	Sets []StickerSet
}

// MessagesAllStickersTypeID is TL type id of MessagesAllStickers.
const MessagesAllStickersTypeID = 0xedfd405f

func (a *MessagesAllStickers) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Hash == 0) {
		return false
	}
	if !(a.Sets == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *MessagesAllStickers) String() string {
	if a == nil {
		return "MessagesAllStickers(nil)"
	}
	type Alias MessagesAllStickers
	return fmt.Sprintf("MessagesAllStickers%+v", Alias(*a))
}

// FillFrom fills MessagesAllStickers from given interface.
func (a *MessagesAllStickers) FillFrom(from interface {
	GetHash() (value int)
	GetSets() (value []StickerSet)
}) {
	a.Hash = from.GetHash()
	a.Sets = from.GetSets()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesAllStickers) TypeID() uint32 {
	return MessagesAllStickersTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesAllStickers) TypeName() string {
	return "messages.allStickers"
}

// TypeInfo returns info about TL type.
func (a *MessagesAllStickers) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.allStickers",
		ID:   MessagesAllStickersTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
		{
			Name:       "Sets",
			SchemaName: "sets",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *MessagesAllStickers) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode messages.allStickers#edfd405f as nil")
	}
	b.PutID(MessagesAllStickersTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *MessagesAllStickers) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode messages.allStickers#edfd405f as nil")
	}
	b.PutInt(a.Hash)
	b.PutVectorHeader(len(a.Sets))
	for idx, v := range a.Sets {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode messages.allStickers#edfd405f: field sets element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetHash returns value of Hash field.
func (a *MessagesAllStickers) GetHash() (value int) {
	return a.Hash
}

// GetSets returns value of Sets field.
func (a *MessagesAllStickers) GetSets() (value []StickerSet) {
	return a.Sets
}

// Decode implements bin.Decoder.
func (a *MessagesAllStickers) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode messages.allStickers#edfd405f to nil")
	}
	if err := b.ConsumeID(MessagesAllStickersTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.allStickers#edfd405f: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *MessagesAllStickers) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode messages.allStickers#edfd405f to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode messages.allStickers#edfd405f: field hash: %w", err)
		}
		a.Hash = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode messages.allStickers#edfd405f: field sets: %w", err)
		}

		if headerLen > 0 {
			a.Sets = make([]StickerSet, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value StickerSet
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode messages.allStickers#edfd405f: field sets: %w", err)
			}
			a.Sets = append(a.Sets, value)
		}
	}
	return nil
}

// construct implements constructor of MessagesAllStickersClass.
func (a MessagesAllStickers) construct() MessagesAllStickersClass { return &a }

// Ensuring interfaces in compile-time for MessagesAllStickers.
var (
	_ bin.Encoder     = &MessagesAllStickers{}
	_ bin.Decoder     = &MessagesAllStickers{}
	_ bin.BareEncoder = &MessagesAllStickers{}
	_ bin.BareDecoder = &MessagesAllStickers{}

	_ MessagesAllStickersClass = &MessagesAllStickers{}
)

// MessagesAllStickersClass represents messages.AllStickers generic type.
//
// See https://core.telegram.org/type/messages.AllStickers for reference.
//
// Example:
//  g, err := tg.DecodeMessagesAllStickers(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.MessagesAllStickersNotModified: // messages.allStickersNotModified#e86602c3
//  case *tg.MessagesAllStickers: // messages.allStickers#edfd405f
//  default: panic(v)
//  }
type MessagesAllStickersClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() MessagesAllStickersClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool

	// AsModified tries to map MessagesAllStickersClass to MessagesAllStickers.
	AsModified() (*MessagesAllStickers, bool)
}

// AsModified tries to map MessagesAllStickersNotModified to MessagesAllStickers.
func (a *MessagesAllStickersNotModified) AsModified() (*MessagesAllStickers, bool) {
	return nil, false
}

// AsModified tries to map MessagesAllStickers to MessagesAllStickers.
func (a *MessagesAllStickers) AsModified() (*MessagesAllStickers, bool) {
	return a, true
}

// DecodeMessagesAllStickers implements binary de-serialization for MessagesAllStickersClass.
func DecodeMessagesAllStickers(buf *bin.Buffer) (MessagesAllStickersClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case MessagesAllStickersNotModifiedTypeID:
		// Decoding messages.allStickersNotModified#e86602c3.
		v := MessagesAllStickersNotModified{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesAllStickersClass: %w", err)
		}
		return &v, nil
	case MessagesAllStickersTypeID:
		// Decoding messages.allStickers#edfd405f.
		v := MessagesAllStickers{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode MessagesAllStickersClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode MessagesAllStickersClass: %w", bin.NewUnexpectedID(id))
	}
}

// MessagesAllStickers boxes the MessagesAllStickersClass providing a helper.
type MessagesAllStickersBox struct {
	AllStickers MessagesAllStickersClass
}

// Decode implements bin.Decoder for MessagesAllStickersBox.
func (b *MessagesAllStickersBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode MessagesAllStickersBox to nil")
	}
	v, err := DecodeMessagesAllStickers(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.AllStickers = v
	return nil
}

// Encode implements bin.Encode for MessagesAllStickersBox.
func (b *MessagesAllStickersBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.AllStickers == nil {
		return fmt.Errorf("unable to encode MessagesAllStickersClass as nil")
	}
	return b.AllStickers.Encode(buf)
}

// MessagesAllStickersClassArray is adapter for slice of MessagesAllStickersClass.
type MessagesAllStickersClassArray []MessagesAllStickersClass

// Sort sorts slice of MessagesAllStickersClass.
func (s MessagesAllStickersClassArray) Sort(less func(a, b MessagesAllStickersClass) bool) MessagesAllStickersClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesAllStickersClass.
func (s MessagesAllStickersClassArray) SortStable(less func(a, b MessagesAllStickersClass) bool) MessagesAllStickersClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesAllStickersClass.
func (s MessagesAllStickersClassArray) Retain(keep func(x MessagesAllStickersClass) bool) MessagesAllStickersClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessagesAllStickersClassArray) First() (v MessagesAllStickersClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesAllStickersClassArray) Last() (v MessagesAllStickersClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesAllStickersClassArray) PopFirst() (v MessagesAllStickersClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesAllStickersClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesAllStickersClassArray) Pop() (v MessagesAllStickersClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsMessagesAllStickers returns copy with only MessagesAllStickers constructors.
func (s MessagesAllStickersClassArray) AsMessagesAllStickers() (to MessagesAllStickersArray) {
	for _, elem := range s {
		value, ok := elem.(*MessagesAllStickers)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AppendOnlyModified appends only Modified constructors to
// given slice.
func (s MessagesAllStickersClassArray) AppendOnlyModified(to []*MessagesAllStickers) []*MessagesAllStickers {
	for _, elem := range s {
		value, ok := elem.AsModified()
		if !ok {
			continue
		}
		to = append(to, value)
	}

	return to
}

// AsModified returns copy with only Modified constructors.
func (s MessagesAllStickersClassArray) AsModified() (to []*MessagesAllStickers) {
	return s.AppendOnlyModified(to)
}

// FirstAsModified returns first element of slice (if exists).
func (s MessagesAllStickersClassArray) FirstAsModified() (v *MessagesAllStickers, ok bool) {
	value, ok := s.First()
	if !ok {
		return
	}
	return value.AsModified()
}

// LastAsModified returns last element of slice (if exists).
func (s MessagesAllStickersClassArray) LastAsModified() (v *MessagesAllStickers, ok bool) {
	value, ok := s.Last()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopFirstAsModified returns element of slice (if exists).
func (s *MessagesAllStickersClassArray) PopFirstAsModified() (v *MessagesAllStickers, ok bool) {
	value, ok := s.PopFirst()
	if !ok {
		return
	}
	return value.AsModified()
}

// PopAsModified returns element of slice (if exists).
func (s *MessagesAllStickersClassArray) PopAsModified() (v *MessagesAllStickers, ok bool) {
	value, ok := s.Pop()
	if !ok {
		return
	}
	return value.AsModified()
}

// MessagesAllStickersArray is adapter for slice of MessagesAllStickers.
type MessagesAllStickersArray []MessagesAllStickers

// Sort sorts slice of MessagesAllStickers.
func (s MessagesAllStickersArray) Sort(less func(a, b MessagesAllStickers) bool) MessagesAllStickersArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesAllStickers.
func (s MessagesAllStickersArray) SortStable(less func(a, b MessagesAllStickers) bool) MessagesAllStickersArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesAllStickers.
func (s MessagesAllStickersArray) Retain(keep func(x MessagesAllStickers) bool) MessagesAllStickersArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s MessagesAllStickersArray) First() (v MessagesAllStickers, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesAllStickersArray) Last() (v MessagesAllStickers, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesAllStickersArray) PopFirst() (v MessagesAllStickers, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesAllStickers
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesAllStickersArray) Pop() (v MessagesAllStickers, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
