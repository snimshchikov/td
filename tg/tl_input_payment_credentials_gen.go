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

// InputPaymentCredentialsSaved represents TL type `inputPaymentCredentialsSaved#c10eb2cf`.
// Saved payment credentials
//
// See https://core.telegram.org/constructor/inputPaymentCredentialsSaved for reference.
type InputPaymentCredentialsSaved struct {
	// Credential ID
	ID string
	// Temporary password
	TmpPassword []byte
}

// InputPaymentCredentialsSavedTypeID is TL type id of InputPaymentCredentialsSaved.
const InputPaymentCredentialsSavedTypeID = 0xc10eb2cf

func (i *InputPaymentCredentialsSaved) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == "") {
		return false
	}
	if !(i.TmpPassword == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPaymentCredentialsSaved) String() string {
	if i == nil {
		return "InputPaymentCredentialsSaved(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputPaymentCredentialsSaved")
	sb.WriteString("{\n")
	sb.WriteString("\tID: ")
	sb.WriteString(fmt.Sprint(i.ID))
	sb.WriteString(",\n")
	sb.WriteString("\tTmpPassword: ")
	sb.WriteString(fmt.Sprint(i.TmpPassword))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InputPaymentCredentialsSaved) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaymentCredentialsSaved#c10eb2cf as nil")
	}
	b.PutID(InputPaymentCredentialsSavedTypeID)
	b.PutString(i.ID)
	b.PutBytes(i.TmpPassword)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPaymentCredentialsSaved) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaymentCredentialsSaved#c10eb2cf to nil")
	}
	if err := b.ConsumeID(InputPaymentCredentialsSavedTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPaymentCredentialsSaved#c10eb2cf: %w", err)
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentialsSaved#c10eb2cf: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentialsSaved#c10eb2cf: field tmp_password: %w", err)
		}
		i.TmpPassword = value
	}
	return nil
}

// construct implements constructor of InputPaymentCredentialsClass.
func (i InputPaymentCredentialsSaved) construct() InputPaymentCredentialsClass { return &i }

// Ensuring interfaces in compile-time for InputPaymentCredentialsSaved.
var (
	_ bin.Encoder = &InputPaymentCredentialsSaved{}
	_ bin.Decoder = &InputPaymentCredentialsSaved{}

	_ InputPaymentCredentialsClass = &InputPaymentCredentialsSaved{}
)

// InputPaymentCredentials represents TL type `inputPaymentCredentials#3417d728`.
// Payment credentials
//
// See https://core.telegram.org/constructor/inputPaymentCredentials for reference.
type InputPaymentCredentials struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Save payment credential for future use
	Save bool
	// Payment credentials
	Data DataJSON
}

// InputPaymentCredentialsTypeID is TL type id of InputPaymentCredentials.
const InputPaymentCredentialsTypeID = 0x3417d728

func (i *InputPaymentCredentials) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.Flags.Zero()) {
		return false
	}
	if !(i.Save == false) {
		return false
	}
	if !(i.Data.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPaymentCredentials) String() string {
	if i == nil {
		return "InputPaymentCredentials(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputPaymentCredentials")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(i.Flags))
	sb.WriteString(",\n")
	sb.WriteString("\tData: ")
	sb.WriteString(fmt.Sprint(i.Data))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InputPaymentCredentials) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaymentCredentials#3417d728 as nil")
	}
	b.PutID(InputPaymentCredentialsTypeID)
	if !(i.Save == false) {
		i.Flags.Set(0)
	}
	if err := i.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputPaymentCredentials#3417d728: field flags: %w", err)
	}
	if err := i.Data.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputPaymentCredentials#3417d728: field data: %w", err)
	}
	return nil
}

// SetSave sets value of Save conditional field.
func (i *InputPaymentCredentials) SetSave(value bool) {
	if value {
		i.Flags.Set(0)
		i.Save = true
	} else {
		i.Flags.Unset(0)
		i.Save = false
	}
}

// Decode implements bin.Decoder.
func (i *InputPaymentCredentials) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaymentCredentials#3417d728 to nil")
	}
	if err := b.ConsumeID(InputPaymentCredentialsTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPaymentCredentials#3417d728: %w", err)
	}
	{
		if err := i.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentials#3417d728: field flags: %w", err)
		}
	}
	i.Save = i.Flags.Has(0)
	{
		if err := i.Data.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentials#3417d728: field data: %w", err)
		}
	}
	return nil
}

// construct implements constructor of InputPaymentCredentialsClass.
func (i InputPaymentCredentials) construct() InputPaymentCredentialsClass { return &i }

// Ensuring interfaces in compile-time for InputPaymentCredentials.
var (
	_ bin.Encoder = &InputPaymentCredentials{}
	_ bin.Decoder = &InputPaymentCredentials{}

	_ InputPaymentCredentialsClass = &InputPaymentCredentials{}
)

// InputPaymentCredentialsApplePay represents TL type `inputPaymentCredentialsApplePay#aa1c39f`.
// Apple pay payment credentials
//
// See https://core.telegram.org/constructor/inputPaymentCredentialsApplePay for reference.
type InputPaymentCredentialsApplePay struct {
	// Payment data
	PaymentData DataJSON
}

// InputPaymentCredentialsApplePayTypeID is TL type id of InputPaymentCredentialsApplePay.
const InputPaymentCredentialsApplePayTypeID = 0xaa1c39f

func (i *InputPaymentCredentialsApplePay) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.PaymentData.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPaymentCredentialsApplePay) String() string {
	if i == nil {
		return "InputPaymentCredentialsApplePay(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputPaymentCredentialsApplePay")
	sb.WriteString("{\n")
	sb.WriteString("\tPaymentData: ")
	sb.WriteString(fmt.Sprint(i.PaymentData))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InputPaymentCredentialsApplePay) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaymentCredentialsApplePay#aa1c39f as nil")
	}
	b.PutID(InputPaymentCredentialsApplePayTypeID)
	if err := i.PaymentData.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputPaymentCredentialsApplePay#aa1c39f: field payment_data: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPaymentCredentialsApplePay) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaymentCredentialsApplePay#aa1c39f to nil")
	}
	if err := b.ConsumeID(InputPaymentCredentialsApplePayTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPaymentCredentialsApplePay#aa1c39f: %w", err)
	}
	{
		if err := i.PaymentData.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentialsApplePay#aa1c39f: field payment_data: %w", err)
		}
	}
	return nil
}

// construct implements constructor of InputPaymentCredentialsClass.
func (i InputPaymentCredentialsApplePay) construct() InputPaymentCredentialsClass { return &i }

// Ensuring interfaces in compile-time for InputPaymentCredentialsApplePay.
var (
	_ bin.Encoder = &InputPaymentCredentialsApplePay{}
	_ bin.Decoder = &InputPaymentCredentialsApplePay{}

	_ InputPaymentCredentialsClass = &InputPaymentCredentialsApplePay{}
)

// InputPaymentCredentialsAndroidPay represents TL type `inputPaymentCredentialsAndroidPay#ca05d50e`.
// Android pay payment credentials
//
// See https://core.telegram.org/constructor/inputPaymentCredentialsAndroidPay for reference.
type InputPaymentCredentialsAndroidPay struct {
	// Android pay payment token
	PaymentToken DataJSON
	// Google transaction ID
	GoogleTransactionID string
}

// InputPaymentCredentialsAndroidPayTypeID is TL type id of InputPaymentCredentialsAndroidPay.
const InputPaymentCredentialsAndroidPayTypeID = 0xca05d50e

func (i *InputPaymentCredentialsAndroidPay) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.PaymentToken.Zero()) {
		return false
	}
	if !(i.GoogleTransactionID == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputPaymentCredentialsAndroidPay) String() string {
	if i == nil {
		return "InputPaymentCredentialsAndroidPay(nil)"
	}
	var sb strings.Builder
	sb.WriteString("InputPaymentCredentialsAndroidPay")
	sb.WriteString("{\n")
	sb.WriteString("\tPaymentToken: ")
	sb.WriteString(fmt.Sprint(i.PaymentToken))
	sb.WriteString(",\n")
	sb.WriteString("\tGoogleTransactionID: ")
	sb.WriteString(fmt.Sprint(i.GoogleTransactionID))
	sb.WriteString(",\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (i *InputPaymentCredentialsAndroidPay) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputPaymentCredentialsAndroidPay#ca05d50e as nil")
	}
	b.PutID(InputPaymentCredentialsAndroidPayTypeID)
	if err := i.PaymentToken.Encode(b); err != nil {
		return fmt.Errorf("unable to encode inputPaymentCredentialsAndroidPay#ca05d50e: field payment_token: %w", err)
	}
	b.PutString(i.GoogleTransactionID)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputPaymentCredentialsAndroidPay) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputPaymentCredentialsAndroidPay#ca05d50e to nil")
	}
	if err := b.ConsumeID(InputPaymentCredentialsAndroidPayTypeID); err != nil {
		return fmt.Errorf("unable to decode inputPaymentCredentialsAndroidPay#ca05d50e: %w", err)
	}
	{
		if err := i.PaymentToken.Decode(b); err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentialsAndroidPay#ca05d50e: field payment_token: %w", err)
		}
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputPaymentCredentialsAndroidPay#ca05d50e: field google_transaction_id: %w", err)
		}
		i.GoogleTransactionID = value
	}
	return nil
}

// construct implements constructor of InputPaymentCredentialsClass.
func (i InputPaymentCredentialsAndroidPay) construct() InputPaymentCredentialsClass { return &i }

// Ensuring interfaces in compile-time for InputPaymentCredentialsAndroidPay.
var (
	_ bin.Encoder = &InputPaymentCredentialsAndroidPay{}
	_ bin.Decoder = &InputPaymentCredentialsAndroidPay{}

	_ InputPaymentCredentialsClass = &InputPaymentCredentialsAndroidPay{}
)

// InputPaymentCredentialsClass represents InputPaymentCredentials generic type.
//
// See https://core.telegram.org/type/InputPaymentCredentials for reference.
//
// Example:
//  g, err := DecodeInputPaymentCredentials(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *InputPaymentCredentialsSaved: // inputPaymentCredentialsSaved#c10eb2cf
//  case *InputPaymentCredentials: // inputPaymentCredentials#3417d728
//  case *InputPaymentCredentialsApplePay: // inputPaymentCredentialsApplePay#aa1c39f
//  case *InputPaymentCredentialsAndroidPay: // inputPaymentCredentialsAndroidPay#ca05d50e
//  default: panic(v)
//  }
type InputPaymentCredentialsClass interface {
	bin.Encoder
	bin.Decoder
	construct() InputPaymentCredentialsClass

	fmt.Stringer
	Zero() bool
}

// DecodeInputPaymentCredentials implements binary de-serialization for InputPaymentCredentialsClass.
func DecodeInputPaymentCredentials(buf *bin.Buffer) (InputPaymentCredentialsClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputPaymentCredentialsSavedTypeID:
		// Decoding inputPaymentCredentialsSaved#c10eb2cf.
		v := InputPaymentCredentialsSaved{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPaymentCredentialsClass: %w", err)
		}
		return &v, nil
	case InputPaymentCredentialsTypeID:
		// Decoding inputPaymentCredentials#3417d728.
		v := InputPaymentCredentials{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPaymentCredentialsClass: %w", err)
		}
		return &v, nil
	case InputPaymentCredentialsApplePayTypeID:
		// Decoding inputPaymentCredentialsApplePay#aa1c39f.
		v := InputPaymentCredentialsApplePay{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPaymentCredentialsClass: %w", err)
		}
		return &v, nil
	case InputPaymentCredentialsAndroidPayTypeID:
		// Decoding inputPaymentCredentialsAndroidPay#ca05d50e.
		v := InputPaymentCredentialsAndroidPay{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputPaymentCredentialsClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputPaymentCredentialsClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputPaymentCredentials boxes the InputPaymentCredentialsClass providing a helper.
type InputPaymentCredentialsBox struct {
	InputPaymentCredentials InputPaymentCredentialsClass
}

// Decode implements bin.Decoder for InputPaymentCredentialsBox.
func (b *InputPaymentCredentialsBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputPaymentCredentialsBox to nil")
	}
	v, err := DecodeInputPaymentCredentials(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputPaymentCredentials = v
	return nil
}

// Encode implements bin.Encode for InputPaymentCredentialsBox.
func (b *InputPaymentCredentialsBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputPaymentCredentials == nil {
		return fmt.Errorf("unable to encode InputPaymentCredentialsClass as nil")
	}
	return b.InputPaymentCredentials.Encode(buf)
}
