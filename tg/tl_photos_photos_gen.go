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

// PhotosPhotos represents TL type `photos.photos#8dca6aa5`.
// Full list of photos with auxiliary data.
//
// See https://core.telegram.org/constructor/photos.photos for reference.
type PhotosPhotos struct {
	// List of photos
	Photos []PhotoClass
	// List of mentioned users
	Users []UserClass
}

// PhotosPhotosTypeID is TL type id of PhotosPhotos.
const PhotosPhotosTypeID = 0x8dca6aa5

func (p *PhotosPhotos) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Photos == nil) {
		return false
	}
	if !(p.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhotosPhotos) String() string {
	if p == nil {
		return "PhotosPhotos(nil)"
	}
	type Alias PhotosPhotos
	return fmt.Sprintf("PhotosPhotos%+v", Alias(*p))
}

// FillFrom fills PhotosPhotos from given interface.
func (p *PhotosPhotos) FillFrom(from interface {
	GetPhotos() (value []PhotoClass)
	GetUsers() (value []UserClass)
}) {
	p.Photos = from.GetPhotos()
	p.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhotosPhotos) TypeID() uint32 {
	return PhotosPhotosTypeID
}

// TypeName returns name of type in TL schema.
func (*PhotosPhotos) TypeName() string {
	return "photos.photos"
}

// TypeInfo returns info about TL type.
func (p *PhotosPhotos) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "photos.photos",
		ID:   PhotosPhotosTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Photos",
			SchemaName: "photos",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PhotosPhotos) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photos.photos#8dca6aa5 as nil")
	}
	b.PutID(PhotosPhotosTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PhotosPhotos) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photos.photos#8dca6aa5 as nil")
	}
	b.PutVectorHeader(len(p.Photos))
	for idx, v := range p.Photos {
		if v == nil {
			return fmt.Errorf("unable to encode photos.photos#8dca6aa5: field photos element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode photos.photos#8dca6aa5: field photos element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(p.Users))
	for idx, v := range p.Users {
		if v == nil {
			return fmt.Errorf("unable to encode photos.photos#8dca6aa5: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode photos.photos#8dca6aa5: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetPhotos returns value of Photos field.
func (p *PhotosPhotos) GetPhotos() (value []PhotoClass) {
	return p.Photos
}

// MapPhotos returns field Photos wrapped in PhotoClassArray helper.
func (p *PhotosPhotos) MapPhotos() (value PhotoClassArray) {
	return PhotoClassArray(p.Photos)
}

// GetUsers returns value of Users field.
func (p *PhotosPhotos) GetUsers() (value []UserClass) {
	return p.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (p *PhotosPhotos) MapUsers() (value UserClassArray) {
	return UserClassArray(p.Users)
}

// Decode implements bin.Decoder.
func (p *PhotosPhotos) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photos.photos#8dca6aa5 to nil")
	}
	if err := b.ConsumeID(PhotosPhotosTypeID); err != nil {
		return fmt.Errorf("unable to decode photos.photos#8dca6aa5: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PhotosPhotos) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photos.photos#8dca6aa5 to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode photos.photos#8dca6aa5: field photos: %w", err)
		}

		if headerLen > 0 {
			p.Photos = make([]PhotoClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePhoto(b)
			if err != nil {
				return fmt.Errorf("unable to decode photos.photos#8dca6aa5: field photos: %w", err)
			}
			p.Photos = append(p.Photos, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode photos.photos#8dca6aa5: field users: %w", err)
		}

		if headerLen > 0 {
			p.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode photos.photos#8dca6aa5: field users: %w", err)
			}
			p.Users = append(p.Users, value)
		}
	}
	return nil
}

// construct implements constructor of PhotosPhotosClass.
func (p PhotosPhotos) construct() PhotosPhotosClass { return &p }

// Ensuring interfaces in compile-time for PhotosPhotos.
var (
	_ bin.Encoder     = &PhotosPhotos{}
	_ bin.Decoder     = &PhotosPhotos{}
	_ bin.BareEncoder = &PhotosPhotos{}
	_ bin.BareDecoder = &PhotosPhotos{}

	_ PhotosPhotosClass = &PhotosPhotos{}
)

// PhotosPhotosSlice represents TL type `photos.photosSlice#15051f54`.
// Incomplete list of photos with auxiliary data.
//
// See https://core.telegram.org/constructor/photos.photosSlice for reference.
type PhotosPhotosSlice struct {
	// Total number of photos
	Count int
	// List of photos
	Photos []PhotoClass
	// List of mentioned users
	Users []UserClass
}

// PhotosPhotosSliceTypeID is TL type id of PhotosPhotosSlice.
const PhotosPhotosSliceTypeID = 0x15051f54

func (p *PhotosPhotosSlice) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Count == 0) {
		return false
	}
	if !(p.Photos == nil) {
		return false
	}
	if !(p.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhotosPhotosSlice) String() string {
	if p == nil {
		return "PhotosPhotosSlice(nil)"
	}
	type Alias PhotosPhotosSlice
	return fmt.Sprintf("PhotosPhotosSlice%+v", Alias(*p))
}

// FillFrom fills PhotosPhotosSlice from given interface.
func (p *PhotosPhotosSlice) FillFrom(from interface {
	GetCount() (value int)
	GetPhotos() (value []PhotoClass)
	GetUsers() (value []UserClass)
}) {
	p.Count = from.GetCount()
	p.Photos = from.GetPhotos()
	p.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhotosPhotosSlice) TypeID() uint32 {
	return PhotosPhotosSliceTypeID
}

// TypeName returns name of type in TL schema.
func (*PhotosPhotosSlice) TypeName() string {
	return "photos.photosSlice"
}

// TypeInfo returns info about TL type.
func (p *PhotosPhotosSlice) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "photos.photosSlice",
		ID:   PhotosPhotosSliceTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Count",
			SchemaName: "count",
		},
		{
			Name:       "Photos",
			SchemaName: "photos",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PhotosPhotosSlice) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photos.photosSlice#15051f54 as nil")
	}
	b.PutID(PhotosPhotosSliceTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PhotosPhotosSlice) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photos.photosSlice#15051f54 as nil")
	}
	b.PutInt(p.Count)
	b.PutVectorHeader(len(p.Photos))
	for idx, v := range p.Photos {
		if v == nil {
			return fmt.Errorf("unable to encode photos.photosSlice#15051f54: field photos element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode photos.photosSlice#15051f54: field photos element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(p.Users))
	for idx, v := range p.Users {
		if v == nil {
			return fmt.Errorf("unable to encode photos.photosSlice#15051f54: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode photos.photosSlice#15051f54: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetCount returns value of Count field.
func (p *PhotosPhotosSlice) GetCount() (value int) {
	return p.Count
}

// GetPhotos returns value of Photos field.
func (p *PhotosPhotosSlice) GetPhotos() (value []PhotoClass) {
	return p.Photos
}

// MapPhotos returns field Photos wrapped in PhotoClassArray helper.
func (p *PhotosPhotosSlice) MapPhotos() (value PhotoClassArray) {
	return PhotoClassArray(p.Photos)
}

// GetUsers returns value of Users field.
func (p *PhotosPhotosSlice) GetUsers() (value []UserClass) {
	return p.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (p *PhotosPhotosSlice) MapUsers() (value UserClassArray) {
	return UserClassArray(p.Users)
}

// Decode implements bin.Decoder.
func (p *PhotosPhotosSlice) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photos.photosSlice#15051f54 to nil")
	}
	if err := b.ConsumeID(PhotosPhotosSliceTypeID); err != nil {
		return fmt.Errorf("unable to decode photos.photosSlice#15051f54: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PhotosPhotosSlice) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photos.photosSlice#15051f54 to nil")
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode photos.photosSlice#15051f54: field count: %w", err)
		}
		p.Count = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode photos.photosSlice#15051f54: field photos: %w", err)
		}

		if headerLen > 0 {
			p.Photos = make([]PhotoClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodePhoto(b)
			if err != nil {
				return fmt.Errorf("unable to decode photos.photosSlice#15051f54: field photos: %w", err)
			}
			p.Photos = append(p.Photos, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode photos.photosSlice#15051f54: field users: %w", err)
		}

		if headerLen > 0 {
			p.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode photos.photosSlice#15051f54: field users: %w", err)
			}
			p.Users = append(p.Users, value)
		}
	}
	return nil
}

// construct implements constructor of PhotosPhotosClass.
func (p PhotosPhotosSlice) construct() PhotosPhotosClass { return &p }

// Ensuring interfaces in compile-time for PhotosPhotosSlice.
var (
	_ bin.Encoder     = &PhotosPhotosSlice{}
	_ bin.Decoder     = &PhotosPhotosSlice{}
	_ bin.BareEncoder = &PhotosPhotosSlice{}
	_ bin.BareDecoder = &PhotosPhotosSlice{}

	_ PhotosPhotosClass = &PhotosPhotosSlice{}
)

// PhotosPhotosClass represents photos.Photos generic type.
//
// See https://core.telegram.org/type/photos.Photos for reference.
//
// Example:
//  g, err := tg.DecodePhotosPhotos(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.PhotosPhotos: // photos.photos#8dca6aa5
//  case *tg.PhotosPhotosSlice: // photos.photosSlice#15051f54
//  default: panic(v)
//  }
type PhotosPhotosClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() PhotosPhotosClass

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

	// List of photos
	GetPhotos() (value []PhotoClass)
	// List of photos
	MapPhotos() (value PhotoClassArray)
	// List of mentioned users
	GetUsers() (value []UserClass)
	// List of mentioned users
	MapUsers() (value UserClassArray)
}

// DecodePhotosPhotos implements binary de-serialization for PhotosPhotosClass.
func DecodePhotosPhotos(buf *bin.Buffer) (PhotosPhotosClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case PhotosPhotosTypeID:
		// Decoding photos.photos#8dca6aa5.
		v := PhotosPhotos{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhotosPhotosClass: %w", err)
		}
		return &v, nil
	case PhotosPhotosSliceTypeID:
		// Decoding photos.photosSlice#15051f54.
		v := PhotosPhotosSlice{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode PhotosPhotosClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode PhotosPhotosClass: %w", bin.NewUnexpectedID(id))
	}
}

// PhotosPhotos boxes the PhotosPhotosClass providing a helper.
type PhotosPhotosBox struct {
	Photos PhotosPhotosClass
}

// Decode implements bin.Decoder for PhotosPhotosBox.
func (b *PhotosPhotosBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode PhotosPhotosBox to nil")
	}
	v, err := DecodePhotosPhotos(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Photos = v
	return nil
}

// Encode implements bin.Encode for PhotosPhotosBox.
func (b *PhotosPhotosBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Photos == nil {
		return fmt.Errorf("unable to encode PhotosPhotosClass as nil")
	}
	return b.Photos.Encode(buf)
}

// PhotosPhotosClassArray is adapter for slice of PhotosPhotosClass.
type PhotosPhotosClassArray []PhotosPhotosClass

// Sort sorts slice of PhotosPhotosClass.
func (s PhotosPhotosClassArray) Sort(less func(a, b PhotosPhotosClass) bool) PhotosPhotosClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhotosPhotosClass.
func (s PhotosPhotosClassArray) SortStable(less func(a, b PhotosPhotosClass) bool) PhotosPhotosClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhotosPhotosClass.
func (s PhotosPhotosClassArray) Retain(keep func(x PhotosPhotosClass) bool) PhotosPhotosClassArray {
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
func (s PhotosPhotosClassArray) First() (v PhotosPhotosClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhotosPhotosClassArray) Last() (v PhotosPhotosClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhotosPhotosClassArray) PopFirst() (v PhotosPhotosClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhotosPhotosClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhotosPhotosClassArray) Pop() (v PhotosPhotosClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsPhotosPhotos returns copy with only PhotosPhotos constructors.
func (s PhotosPhotosClassArray) AsPhotosPhotos() (to PhotosPhotosArray) {
	for _, elem := range s {
		value, ok := elem.(*PhotosPhotos)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsPhotosPhotosSlice returns copy with only PhotosPhotosSlice constructors.
func (s PhotosPhotosClassArray) AsPhotosPhotosSlice() (to PhotosPhotosSliceArray) {
	for _, elem := range s {
		value, ok := elem.(*PhotosPhotosSlice)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// PhotosPhotosArray is adapter for slice of PhotosPhotos.
type PhotosPhotosArray []PhotosPhotos

// Sort sorts slice of PhotosPhotos.
func (s PhotosPhotosArray) Sort(less func(a, b PhotosPhotos) bool) PhotosPhotosArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhotosPhotos.
func (s PhotosPhotosArray) SortStable(less func(a, b PhotosPhotos) bool) PhotosPhotosArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhotosPhotos.
func (s PhotosPhotosArray) Retain(keep func(x PhotosPhotos) bool) PhotosPhotosArray {
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
func (s PhotosPhotosArray) First() (v PhotosPhotos, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhotosPhotosArray) Last() (v PhotosPhotos, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhotosPhotosArray) PopFirst() (v PhotosPhotos, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhotosPhotos
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhotosPhotosArray) Pop() (v PhotosPhotos, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// PhotosPhotosSliceArray is adapter for slice of PhotosPhotosSlice.
type PhotosPhotosSliceArray []PhotosPhotosSlice

// Sort sorts slice of PhotosPhotosSlice.
func (s PhotosPhotosSliceArray) Sort(less func(a, b PhotosPhotosSlice) bool) PhotosPhotosSliceArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of PhotosPhotosSlice.
func (s PhotosPhotosSliceArray) SortStable(less func(a, b PhotosPhotosSlice) bool) PhotosPhotosSliceArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of PhotosPhotosSlice.
func (s PhotosPhotosSliceArray) Retain(keep func(x PhotosPhotosSlice) bool) PhotosPhotosSliceArray {
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
func (s PhotosPhotosSliceArray) First() (v PhotosPhotosSlice, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s PhotosPhotosSliceArray) Last() (v PhotosPhotosSlice, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *PhotosPhotosSliceArray) PopFirst() (v PhotosPhotosSlice, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero PhotosPhotosSlice
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *PhotosPhotosSliceArray) Pop() (v PhotosPhotosSlice, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
