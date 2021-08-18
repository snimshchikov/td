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

// PhotosPhoto represents TL type `photos.photo#20212ca8`.
// Photo with auxiliary data.
//
// See https://core.telegram.org/constructor/photos.photo for reference.
type PhotosPhoto struct {
	// Photo
	Photo PhotoClass
	// Users
	Users []UserClass
}

// PhotosPhotoTypeID is TL type id of PhotosPhoto.
const PhotosPhotoTypeID = 0x20212ca8

func (p *PhotosPhoto) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Photo == nil) {
		return false
	}
	if !(p.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PhotosPhoto) String() string {
	if p == nil {
		return "PhotosPhoto(nil)"
	}
	type Alias PhotosPhoto
	return fmt.Sprintf("PhotosPhoto%+v", Alias(*p))
}

// FillFrom fills PhotosPhoto from given interface.
func (p *PhotosPhoto) FillFrom(from interface {
	GetPhoto() (value PhotoClass)
	GetUsers() (value []UserClass)
}) {
	p.Photo = from.GetPhoto()
	p.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PhotosPhoto) TypeID() uint32 {
	return PhotosPhotoTypeID
}

// TypeName returns name of type in TL schema.
func (*PhotosPhoto) TypeName() string {
	return "photos.photo"
}

// TypeInfo returns info about TL type.
func (p *PhotosPhoto) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "photos.photo",
		ID:   PhotosPhotoTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Photo",
			SchemaName: "photo",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PhotosPhoto) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photos.photo#20212ca8 as nil")
	}
	b.PutID(PhotosPhotoTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PhotosPhoto) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode photos.photo#20212ca8 as nil")
	}
	if p.Photo == nil {
		return fmt.Errorf("unable to encode photos.photo#20212ca8: field photo is nil")
	}
	if err := p.Photo.Encode(b); err != nil {
		return fmt.Errorf("unable to encode photos.photo#20212ca8: field photo: %w", err)
	}
	b.PutVectorHeader(len(p.Users))
	for idx, v := range p.Users {
		if v == nil {
			return fmt.Errorf("unable to encode photos.photo#20212ca8: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode photos.photo#20212ca8: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// GetPhoto returns value of Photo field.
func (p *PhotosPhoto) GetPhoto() (value PhotoClass) {
	return p.Photo
}

// GetPhotoAsNotEmpty returns mapped value of Photo field.
func (p *PhotosPhoto) GetPhotoAsNotEmpty() (*Photo, bool) {
	return p.Photo.AsNotEmpty()
}

// GetUsers returns value of Users field.
func (p *PhotosPhoto) GetUsers() (value []UserClass) {
	return p.Users
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (p *PhotosPhoto) MapUsers() (value UserClassArray) {
	return UserClassArray(p.Users)
}

// Decode implements bin.Decoder.
func (p *PhotosPhoto) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photos.photo#20212ca8 to nil")
	}
	if err := b.ConsumeID(PhotosPhotoTypeID); err != nil {
		return fmt.Errorf("unable to decode photos.photo#20212ca8: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PhotosPhoto) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode photos.photo#20212ca8 to nil")
	}
	{
		value, err := DecodePhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode photos.photo#20212ca8: field photo: %w", err)
		}
		p.Photo = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode photos.photo#20212ca8: field users: %w", err)
		}

		if headerLen > 0 {
			p.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode photos.photo#20212ca8: field users: %w", err)
			}
			p.Users = append(p.Users, value)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for PhotosPhoto.
var (
	_ bin.Encoder     = &PhotosPhoto{}
	_ bin.Decoder     = &PhotosPhoto{}
	_ bin.BareEncoder = &PhotosPhoto{}
	_ bin.BareDecoder = &PhotosPhoto{}
)
