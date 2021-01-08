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

// HelpDeepLinkInfoEmpty represents TL type `help.deepLinkInfoEmpty#66afa166`.
// Deep link info empty
//
// See https://core.telegram.org/constructor/help.deepLinkInfoEmpty for reference.
type HelpDeepLinkInfoEmpty struct {
}

// HelpDeepLinkInfoEmptyTypeID is TL type id of HelpDeepLinkInfoEmpty.
const HelpDeepLinkInfoEmptyTypeID = 0x66afa166

func (d *HelpDeepLinkInfoEmpty) Zero() bool {
	if d == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (d *HelpDeepLinkInfoEmpty) String() string {
	if d == nil {
		return "HelpDeepLinkInfoEmpty(nil)"
	}
	var sb strings.Builder
	sb.WriteString("HelpDeepLinkInfoEmpty")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (d *HelpDeepLinkInfoEmpty) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode help.deepLinkInfoEmpty#66afa166 as nil")
	}
	b.PutID(HelpDeepLinkInfoEmptyTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (d *HelpDeepLinkInfoEmpty) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode help.deepLinkInfoEmpty#66afa166 to nil")
	}
	if err := b.ConsumeID(HelpDeepLinkInfoEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode help.deepLinkInfoEmpty#66afa166: %w", err)
	}
	return nil
}

// construct implements constructor of HelpDeepLinkInfoClass.
func (d HelpDeepLinkInfoEmpty) construct() HelpDeepLinkInfoClass { return &d }

// Ensuring interfaces in compile-time for HelpDeepLinkInfoEmpty.
var (
	_ bin.Encoder = &HelpDeepLinkInfoEmpty{}
	_ bin.Decoder = &HelpDeepLinkInfoEmpty{}

	_ HelpDeepLinkInfoClass = &HelpDeepLinkInfoEmpty{}
)

// HelpDeepLinkInfo represents TL type `help.deepLinkInfo#6a4ee832`.
// Deep linking info
//
// See https://core.telegram.org/constructor/help.deepLinkInfo for reference.
type HelpDeepLinkInfo struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// An update of the app is required to parse this link
	UpdateApp bool
	// Message to show to the user
	Message string
	// Message entities for styled text¹
	//
	// Links:
	//  1) https://core.telegram.org/api/entities
	//
	// Use SetEntities and GetEntities helpers.
	Entities []MessageEntityClass
}

// HelpDeepLinkInfoTypeID is TL type id of HelpDeepLinkInfo.
const HelpDeepLinkInfoTypeID = 0x6a4ee832

func (d *HelpDeepLinkInfo) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Flags.Zero()) {
		return false
	}
	if !(d.UpdateApp == false) {
		return false
	}
	if !(d.Message == "") {
		return false
	}
	if !(d.Entities == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *HelpDeepLinkInfo) String() string {
	if d == nil {
		return "HelpDeepLinkInfo(nil)"
	}
	var sb strings.Builder
	sb.WriteString("HelpDeepLinkInfo")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(d.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tMessage: ")
	sb.WriteString(fmt.Sprint(d.Message))
	sb.WriteString(",\n")
	if d.Flags.Has(1) {
		sb.WriteByte('[')
		for _, v := range d.Entities {
			sb.WriteString(fmt.Sprint(v))
		}
		sb.WriteByte(']')
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (d *HelpDeepLinkInfo) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode help.deepLinkInfo#6a4ee832 as nil")
	}
	b.PutID(HelpDeepLinkInfoTypeID)
	if !(d.UpdateApp == false) {
		d.Flags.Set(0)
	}
	if !(d.Entities == nil) {
		d.Flags.Set(1)
	}
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode help.deepLinkInfo#6a4ee832: field flags: %w", err)
	}
	b.PutString(d.Message)
	if d.Flags.Has(1) {
		b.PutVectorHeader(len(d.Entities))
		for idx, v := range d.Entities {
			if v == nil {
				return fmt.Errorf("unable to encode help.deepLinkInfo#6a4ee832: field entities element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode help.deepLinkInfo#6a4ee832: field entities element with index %d: %w", idx, err)
			}
		}
	}
	return nil
}

// SetUpdateApp sets value of UpdateApp conditional field.
func (d *HelpDeepLinkInfo) SetUpdateApp(value bool) {
	if value {
		d.Flags.Set(0)
		d.UpdateApp = true
	} else {
		d.Flags.Unset(0)
		d.UpdateApp = false
	}
}

// SetEntities sets value of Entities conditional field.
func (d *HelpDeepLinkInfo) SetEntities(value []MessageEntityClass) {
	d.Flags.Set(1)
	d.Entities = value
}

// GetEntities returns value of Entities conditional field and
// boolean which is true if field was set.
func (d *HelpDeepLinkInfo) GetEntities() (value []MessageEntityClass, ok bool) {
	if !d.Flags.Has(1) {
		return value, false
	}
	return d.Entities, true
}

// Decode implements bin.Decoder.
func (d *HelpDeepLinkInfo) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode help.deepLinkInfo#6a4ee832 to nil")
	}
	if err := b.ConsumeID(HelpDeepLinkInfoTypeID); err != nil {
		return fmt.Errorf("unable to decode help.deepLinkInfo#6a4ee832: %w", err)
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode help.deepLinkInfo#6a4ee832: field flags: %w", err)
		}
	}
	d.UpdateApp = d.Flags.Has(0)
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode help.deepLinkInfo#6a4ee832: field message: %w", err)
		}
		d.Message = value
	}
	if d.Flags.Has(1) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode help.deepLinkInfo#6a4ee832: field entities: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeMessageEntity(b)
			if err != nil {
				return fmt.Errorf("unable to decode help.deepLinkInfo#6a4ee832: field entities: %w", err)
			}
			d.Entities = append(d.Entities, value)
		}
	}
	return nil
}

// construct implements constructor of HelpDeepLinkInfoClass.
func (d HelpDeepLinkInfo) construct() HelpDeepLinkInfoClass { return &d }

// Ensuring interfaces in compile-time for HelpDeepLinkInfo.
var (
	_ bin.Encoder = &HelpDeepLinkInfo{}
	_ bin.Decoder = &HelpDeepLinkInfo{}

	_ HelpDeepLinkInfoClass = &HelpDeepLinkInfo{}
)

// HelpDeepLinkInfoClass represents help.DeepLinkInfo generic type.
//
// See https://core.telegram.org/type/help.DeepLinkInfo for reference.
//
// Example:
//  g, err := DecodeHelpDeepLinkInfo(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *HelpDeepLinkInfoEmpty: // help.deepLinkInfoEmpty#66afa166
//  case *HelpDeepLinkInfo: // help.deepLinkInfo#6a4ee832
//  default: panic(v)
//  }
type HelpDeepLinkInfoClass interface {
	bin.Encoder
	bin.Decoder
	construct() HelpDeepLinkInfoClass

	fmt.Stringer
	Zero() bool
}

// DecodeHelpDeepLinkInfo implements binary de-serialization for HelpDeepLinkInfoClass.
func DecodeHelpDeepLinkInfo(buf *bin.Buffer) (HelpDeepLinkInfoClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case HelpDeepLinkInfoEmptyTypeID:
		// Decoding help.deepLinkInfoEmpty#66afa166.
		v := HelpDeepLinkInfoEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpDeepLinkInfoClass: %w", err)
		}
		return &v, nil
	case HelpDeepLinkInfoTypeID:
		// Decoding help.deepLinkInfo#6a4ee832.
		v := HelpDeepLinkInfo{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode HelpDeepLinkInfoClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode HelpDeepLinkInfoClass: %w", bin.NewUnexpectedID(id))
	}
}

// HelpDeepLinkInfo boxes the HelpDeepLinkInfoClass providing a helper.
type HelpDeepLinkInfoBox struct {
	DeepLinkInfo HelpDeepLinkInfoClass
}

// Decode implements bin.Decoder for HelpDeepLinkInfoBox.
func (b *HelpDeepLinkInfoBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode HelpDeepLinkInfoBox to nil")
	}
	v, err := DecodeHelpDeepLinkInfo(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.DeepLinkInfo = v
	return nil
}

// Encode implements bin.Encode for HelpDeepLinkInfoBox.
func (b *HelpDeepLinkInfoBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.DeepLinkInfo == nil {
		return fmt.Errorf("unable to encode HelpDeepLinkInfoClass as nil")
	}
	return b.DeepLinkInfo.Encode(buf)
}
