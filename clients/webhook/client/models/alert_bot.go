// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// AlertBot alert bot
//
// swagger:model AlertBot
type AlertBot struct {

	// id
	// Example: 0x17381ae942ee1fe141d0652e9dad7d001761552f906fb1684b2812603de31049
	ID string `json:"id,omitempty"`

	// Docker image reference (Disco)
	// Example: bafybeibrigevnhic4befnkqbaagzgxqtdyv2fdgcbqwxe7ees3hw6fymme@sha256:9ca1547e130a6264bb1b4ad6b10f17cabf404957f23d457a30046b9afdf29fc8
	Image string `json:"image,omitempty"`

	// Bot reference (IPFS hash)
	// Example: QmU6L9Zo5rweF6QZLhLfwAAFUFRMF3uFdSnMiJzENXr37R
	Reference string `json:"reference,omitempty"`
}

// Validate validates this alert bot
func (m *AlertBot) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this alert bot based on context it is used
func (m *AlertBot) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *AlertBot) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AlertBot) UnmarshalBinary(b []byte) error {
	var res AlertBot
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
