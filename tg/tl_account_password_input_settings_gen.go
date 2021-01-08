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

// AccountPasswordInputSettings represents TL type `account.passwordInputSettings#c23727c9`.
// Settings for setting up a new password
//
// See https://core.telegram.org/constructor/account.passwordInputSettings for reference.
type AccountPasswordInputSettings struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// The SRP algorithm¹ to use
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	//
	// Use SetNewAlgo and GetNewAlgo helpers.
	NewAlgo PasswordKdfAlgoClass
	// The computed password hash¹
	//
	// Links:
	//  1) https://core.telegram.org/api/srp
	//
	// Use SetNewPasswordHash and GetNewPasswordHash helpers.
	NewPasswordHash []byte
	// Text hint for the password
	//
	// Use SetHint and GetHint helpers.
	Hint string
	// Password recovery email
	//
	// Use SetEmail and GetEmail helpers.
	Email string
	// Telegram passport¹ settings
	//
	// Links:
	//  1) https://core.telegram.org/passport
	//
	// Use SetNewSecureSettings and GetNewSecureSettings helpers.
	NewSecureSettings SecureSecretSettings
}

// AccountPasswordInputSettingsTypeID is TL type id of AccountPasswordInputSettings.
const AccountPasswordInputSettingsTypeID = 0xc23727c9

func (p *AccountPasswordInputSettings) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.Flags.Zero()) {
		return false
	}
	if !(p.NewAlgo == nil) {
		return false
	}
	if !(p.NewPasswordHash == nil) {
		return false
	}
	if !(p.Hint == "") {
		return false
	}
	if !(p.Email == "") {
		return false
	}
	if !(p.NewSecureSettings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *AccountPasswordInputSettings) String() string {
	if p == nil {
		return "AccountPasswordInputSettings(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountPasswordInputSettings")
	sb.WriteString("{\n")
	sb.WriteString("\tFlags: ")
	sb.WriteString(fmt.Sprint(p.Flags))
	sb.WriteString(",\n")
	if p.Flags.Has(0) {
		sb.WriteString("\tNewAlgo: ")
		sb.WriteString(fmt.Sprint(p.NewAlgo))
		sb.WriteString(",\n")
	}
	if p.Flags.Has(0) {
		sb.WriteString("\tNewPasswordHash: ")
		sb.WriteString(fmt.Sprint(p.NewPasswordHash))
		sb.WriteString(",\n")
	}
	if p.Flags.Has(0) {
		sb.WriteString("\tHint: ")
		sb.WriteString(fmt.Sprint(p.Hint))
		sb.WriteString(",\n")
	}
	if p.Flags.Has(1) {
		sb.WriteString("\tEmail: ")
		sb.WriteString(fmt.Sprint(p.Email))
		sb.WriteString(",\n")
	}
	if p.Flags.Has(2) {
		sb.WriteString("\tNewSecureSettings: ")
		sb.WriteString(fmt.Sprint(p.NewSecureSettings))
		sb.WriteString(",\n")
	}
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (p *AccountPasswordInputSettings) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode account.passwordInputSettings#c23727c9 as nil")
	}
	b.PutID(AccountPasswordInputSettingsTypeID)
	if !(p.NewAlgo == nil) {
		p.Flags.Set(0)
	}
	if !(p.NewPasswordHash == nil) {
		p.Flags.Set(0)
	}
	if !(p.Hint == "") {
		p.Flags.Set(0)
	}
	if !(p.Email == "") {
		p.Flags.Set(1)
	}
	if !(p.NewSecureSettings.Zero()) {
		p.Flags.Set(2)
	}
	if err := p.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.passwordInputSettings#c23727c9: field flags: %w", err)
	}
	if p.Flags.Has(0) {
		if p.NewAlgo == nil {
			return fmt.Errorf("unable to encode account.passwordInputSettings#c23727c9: field new_algo is nil")
		}
		if err := p.NewAlgo.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.passwordInputSettings#c23727c9: field new_algo: %w", err)
		}
	}
	if p.Flags.Has(0) {
		b.PutBytes(p.NewPasswordHash)
	}
	if p.Flags.Has(0) {
		b.PutString(p.Hint)
	}
	if p.Flags.Has(1) {
		b.PutString(p.Email)
	}
	if p.Flags.Has(2) {
		if err := p.NewSecureSettings.Encode(b); err != nil {
			return fmt.Errorf("unable to encode account.passwordInputSettings#c23727c9: field new_secure_settings: %w", err)
		}
	}
	return nil
}

// SetNewAlgo sets value of NewAlgo conditional field.
func (p *AccountPasswordInputSettings) SetNewAlgo(value PasswordKdfAlgoClass) {
	p.Flags.Set(0)
	p.NewAlgo = value
}

// GetNewAlgo returns value of NewAlgo conditional field and
// boolean which is true if field was set.
func (p *AccountPasswordInputSettings) GetNewAlgo() (value PasswordKdfAlgoClass, ok bool) {
	if !p.Flags.Has(0) {
		return value, false
	}
	return p.NewAlgo, true
}

// SetNewPasswordHash sets value of NewPasswordHash conditional field.
func (p *AccountPasswordInputSettings) SetNewPasswordHash(value []byte) {
	p.Flags.Set(0)
	p.NewPasswordHash = value
}

// GetNewPasswordHash returns value of NewPasswordHash conditional field and
// boolean which is true if field was set.
func (p *AccountPasswordInputSettings) GetNewPasswordHash() (value []byte, ok bool) {
	if !p.Flags.Has(0) {
		return value, false
	}
	return p.NewPasswordHash, true
}

// SetHint sets value of Hint conditional field.
func (p *AccountPasswordInputSettings) SetHint(value string) {
	p.Flags.Set(0)
	p.Hint = value
}

// GetHint returns value of Hint conditional field and
// boolean which is true if field was set.
func (p *AccountPasswordInputSettings) GetHint() (value string, ok bool) {
	if !p.Flags.Has(0) {
		return value, false
	}
	return p.Hint, true
}

// SetEmail sets value of Email conditional field.
func (p *AccountPasswordInputSettings) SetEmail(value string) {
	p.Flags.Set(1)
	p.Email = value
}

// GetEmail returns value of Email conditional field and
// boolean which is true if field was set.
func (p *AccountPasswordInputSettings) GetEmail() (value string, ok bool) {
	if !p.Flags.Has(1) {
		return value, false
	}
	return p.Email, true
}

// SetNewSecureSettings sets value of NewSecureSettings conditional field.
func (p *AccountPasswordInputSettings) SetNewSecureSettings(value SecureSecretSettings) {
	p.Flags.Set(2)
	p.NewSecureSettings = value
}

// GetNewSecureSettings returns value of NewSecureSettings conditional field and
// boolean which is true if field was set.
func (p *AccountPasswordInputSettings) GetNewSecureSettings() (value SecureSecretSettings, ok bool) {
	if !p.Flags.Has(2) {
		return value, false
	}
	return p.NewSecureSettings, true
}

// Decode implements bin.Decoder.
func (p *AccountPasswordInputSettings) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode account.passwordInputSettings#c23727c9 to nil")
	}
	if err := b.ConsumeID(AccountPasswordInputSettingsTypeID); err != nil {
		return fmt.Errorf("unable to decode account.passwordInputSettings#c23727c9: %w", err)
	}
	{
		if err := p.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.passwordInputSettings#c23727c9: field flags: %w", err)
		}
	}
	if p.Flags.Has(0) {
		value, err := DecodePasswordKdfAlgo(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.passwordInputSettings#c23727c9: field new_algo: %w", err)
		}
		p.NewAlgo = value
	}
	if p.Flags.Has(0) {
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode account.passwordInputSettings#c23727c9: field new_password_hash: %w", err)
		}
		p.NewPasswordHash = value
	}
	if p.Flags.Has(0) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.passwordInputSettings#c23727c9: field hint: %w", err)
		}
		p.Hint = value
	}
	if p.Flags.Has(1) {
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.passwordInputSettings#c23727c9: field email: %w", err)
		}
		p.Email = value
	}
	if p.Flags.Has(2) {
		if err := p.NewSecureSettings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.passwordInputSettings#c23727c9: field new_secure_settings: %w", err)
		}
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountPasswordInputSettings.
var (
	_ bin.Encoder = &AccountPasswordInputSettings{}
	_ bin.Decoder = &AccountPasswordInputSettings{}
)
