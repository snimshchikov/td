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

// PageTableCell represents TL type `pageTableCell#34566b6a`.
// Table cell
//
// See https://core.telegram.org/constructor/pageTableCell for reference.
type PageTableCell struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Is this element part of the column header
	Header bool
	// Horizontally centered block
	AlignCenter bool
	// Right-aligned block
	AlignRight bool
	// Vertically centered block
	ValignMiddle bool
	// Block vertically-alligned to the bottom
	ValignBottom bool
	// Content
	//
	// Use SetText and GetText helpers.
	Text RichTextClass
	// For how many columns should this cell extend
	//
	// Use SetColspan and GetColspan helpers.
	Colspan int
	// For how many rows should this cell extend
	//
	// Use SetRowspan and GetRowspan helpers.
	Rowspan int
}

// PageTableCellTypeID is TL type id of PageTableCell.
const PageTableCellTypeID = 0x34566b6a

func (p *PageTableCell) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Flags.Zero()) {
		return false
	}
	if !(p.Header == false) {
		return false
	}
	if !(p.AlignCenter == false) {
		return false
	}
	if !(p.AlignRight == false) {
		return false
	}
	if !(p.ValignMiddle == false) {
		return false
	}
	if !(p.ValignBottom == false) {
		return false
	}
	if !(p.Text == nil) {
		return false
	}
	if !(p.Colspan == 0) {
		return false
	}
	if !(p.Rowspan == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PageTableCell) String() string {
	if p == nil {
		return "PageTableCell(nil)"
	}
	var sb strings.Builder
	sb.WriteString("PageTableCell")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(p.Flags))
	sb.WriteString(",\n")
	if p.Flags.Has(7) {
		sb.WriteString("\tText: ")
		sb.WriteString(fmt.Sprint(p.Text))
		sb.WriteString(",\n")
	}
	if p.Flags.Has(1) {
		sb.WriteString("\tColspan: ")
		sb.WriteString(fmt.Sprint(p.Colspan))
		sb.WriteString(",\n")
	}
	if p.Flags.Has(2) {
		sb.WriteString("\tRowspan: ")
		sb.WriteString(fmt.Sprint(p.Rowspan))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (p *PageTableCell) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode pageTableCell#34566b6a as nil")
	}
	b.PutID(PageTableCellTypeID)
	if !(p.Header == false) {
		p.Flags.Set(0)
	}
	if !(p.AlignCenter == false) {
		p.Flags.Set(3)
	}
	if !(p.AlignRight == false) {
		p.Flags.Set(4)
	}
	if !(p.ValignMiddle == false) {
		p.Flags.Set(5)
	}
	if !(p.ValignBottom == false) {
		p.Flags.Set(6)
	}
	if !(p.Text == nil) {
		p.Flags.Set(7)
	}
	if !(p.Colspan == 0) {
		p.Flags.Set(1)
	}
	if !(p.Rowspan == 0) {
		p.Flags.Set(2)
	}
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode pageTableCell#34566b6a: field flags: %w", err)
	}
	if p.Flags.Has(7) {
		if p.Text == nil {
			return fmt.Errorf("unable to encode pageTableCell#34566b6a: field text is nil")
		}
		if err := p.Text.Encode(b); err != nil {
			return fmt.Errorf("unable to encode pageTableCell#34566b6a: field text: %w", err)
		}
	}
	if p.Flags.Has(1) {
		b.PutInt(p.Colspan)
	}
	if p.Flags.Has(2) {
		b.PutInt(p.Rowspan)
	}
	return nil
}

// SetHeader sets value of Header conditional field.
func (p *PageTableCell) SetHeader(value bool) {
	if value {
		p.Flags.Set(0)
		p.Header = true
	} else {
		p.Flags.Unset(0)
		p.Header = false
	}
}

// SetAlignCenter sets value of AlignCenter conditional field.
func (p *PageTableCell) SetAlignCenter(value bool) {
	if value {
		p.Flags.Set(3)
		p.AlignCenter = true
	} else {
		p.Flags.Unset(3)
		p.AlignCenter = false
	}
}

// SetAlignRight sets value of AlignRight conditional field.
func (p *PageTableCell) SetAlignRight(value bool) {
	if value {
		p.Flags.Set(4)
		p.AlignRight = true
	} else {
		p.Flags.Unset(4)
		p.AlignRight = false
	}
}

// SetValignMiddle sets value of ValignMiddle conditional field.
func (p *PageTableCell) SetValignMiddle(value bool) {
	if value {
		p.Flags.Set(5)
		p.ValignMiddle = true
	} else {
		p.Flags.Unset(5)
		p.ValignMiddle = false
	}
}

// SetValignBottom sets value of ValignBottom conditional field.
func (p *PageTableCell) SetValignBottom(value bool) {
	if value {
		p.Flags.Set(6)
		p.ValignBottom = true
	} else {
		p.Flags.Unset(6)
		p.ValignBottom = false
	}
}

// SetText sets value of Text conditional field.
func (p *PageTableCell) SetText(value RichTextClass) {
	p.Flags.Set(7)
	p.Text = value
}

// GetText returns value of Text conditional field and
// boolean which is true if field was set.
func (p *PageTableCell) GetText() (value RichTextClass, ok bool) {
	if !p.Flags.Has(7) {
		return value, false
	}
	return p.Text, true
}

// SetColspan sets value of Colspan conditional field.
func (p *PageTableCell) SetColspan(value int) {
	p.Flags.Set(1)
	p.Colspan = value
}

// GetColspan returns value of Colspan conditional field and
// boolean which is true if field was set.
func (p *PageTableCell) GetColspan() (value int, ok bool) {
	if !p.Flags.Has(1) {
		return value, false
	}
	return p.Colspan, true
}

// SetRowspan sets value of Rowspan conditional field.
func (p *PageTableCell) SetRowspan(value int) {
	p.Flags.Set(2)
	p.Rowspan = value
}

// GetRowspan returns value of Rowspan conditional field and
// boolean which is true if field was set.
func (p *PageTableCell) GetRowspan() (value int, ok bool) {
	if !p.Flags.Has(2) {
		return value, false
	}
	return p.Rowspan, true
}

// Decode implements bin.Decoder.
func (p *PageTableCell) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode pageTableCell#34566b6a to nil")
	}
	if err := b.ConsumeID(PageTableCellTypeID); err != nil {
		return fmt.Errorf("unable to decode pageTableCell#34566b6a: %w", err)
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode pageTableCell#34566b6a: field flags: %w", err)
		}
	}
	p.Header = p.Flags.Has(0)
	p.AlignCenter = p.Flags.Has(3)
	p.AlignRight = p.Flags.Has(4)
	p.ValignMiddle = p.Flags.Has(5)
	p.ValignBottom = p.Flags.Has(6)
	if p.Flags.Has(7) {
		value, err := DecodeRichText(b)
		if err != nil {
			return fmt.Errorf("unable to decode pageTableCell#34566b6a: field text: %w", err)
		}
		p.Text = value
	}
	if p.Flags.Has(1) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode pageTableCell#34566b6a: field colspan: %w", err)
		}
		p.Colspan = value
	}
	if p.Flags.Has(2) {
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode pageTableCell#34566b6a: field rowspan: %w", err)
		}
		p.Rowspan = value
	}
	return nil
}

// Ensuring interfaces in compile-time for PageTableCell.
var (
	_ bin.Encoder = &PageTableCell{}
	_ bin.Decoder = &PageTableCell{}
)
