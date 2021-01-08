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

// AccountResetWebAuthorizationsRequest represents TL type `account.resetWebAuthorizations#682d2594`.
// Reset all active web telegram login¹ sessions
//
// Links:
//  1) https://core.telegram.org/widgets/login
//
// See https://core.telegram.org/method/account.resetWebAuthorizations for reference.
type AccountResetWebAuthorizationsRequest struct {
}

// AccountResetWebAuthorizationsRequestTypeID is TL type id of AccountResetWebAuthorizationsRequest.
const AccountResetWebAuthorizationsRequestTypeID = 0x682d2594

func (r *AccountResetWebAuthorizationsRequest) Zero() bool {
	if r == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (r *AccountResetWebAuthorizationsRequest) String() string {
	if r == nil {
		return "AccountResetWebAuthorizationsRequest(nil)"
	}
	var sb strings.Builder
	sb.WriteString("AccountResetWebAuthorizationsRequest")
	sb.WriteString("{\n")
	sb.WriteString("}")
	return sb.String()
}

// Encode implements bin.Encoder.
func (r *AccountResetWebAuthorizationsRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.resetWebAuthorizations#682d2594 as nil")
	}
	b.PutID(AccountResetWebAuthorizationsRequestTypeID)
	return nil
}

// Decode implements bin.Decoder.
func (r *AccountResetWebAuthorizationsRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.resetWebAuthorizations#682d2594 to nil")
	}
	if err := b.ConsumeID(AccountResetWebAuthorizationsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.resetWebAuthorizations#682d2594: %w", err)
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountResetWebAuthorizationsRequest.
var (
	_ bin.Encoder = &AccountResetWebAuthorizationsRequest{}
	_ bin.Decoder = &AccountResetWebAuthorizationsRequest{}
)

// AccountResetWebAuthorizations invokes method account.resetWebAuthorizations#682d2594 returning error if any.
// Reset all active web telegram login¹ sessions
//
// Links:
//  1) https://core.telegram.org/widgets/login
//
// See https://core.telegram.org/method/account.resetWebAuthorizations for reference.
func (c *Client) AccountResetWebAuthorizations(ctx context.Context) (bool, error) {
	var result BoolBox

	request := &AccountResetWebAuthorizationsRequest{}
	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
