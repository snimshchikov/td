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

// StatsURL represents TL type `statsURL#47a971e0`.
// URL with chat statistics
//
// See https://core.telegram.org/constructor/statsURL for reference.
type StatsURL struct {
	// Chat statistics
	URL string
}

// StatsURLTypeID is TL type id of StatsURL.
const StatsURLTypeID = 0x47a971e0

func (s *StatsURL) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.URL == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *StatsURL) String() string {
	if s == nil {
		return "StatsURL(nil)"
	}
	var sb strings.Builder
	sb.WriteString("StatsURL")
	sb.WriteString("{\n")
	sb.WriteString("\tURL: ")
	sb.WriteString(fmt.Sprint(s.URL))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (s *StatsURL) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode statsURL#47a971e0 as nil")
	}
	b.PutID(StatsURLTypeID)
	b.PutString(s.URL)
	return nil
}

// Decode implements bin.Decoder.
func (s *StatsURL) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode statsURL#47a971e0 to nil")
	}
	if err := b.ConsumeID(StatsURLTypeID); err != nil {
		return fmt.Errorf("unable to decode statsURL#47a971e0: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode statsURL#47a971e0: field url: %w", err)
		}
		s.URL = value
	}
	return nil
}

// Ensuring interfaces in compile-time for StatsURL.
var (
	_ bin.Encoder = &StatsURL{}
	_ bin.Decoder = &StatsURL{}
)
