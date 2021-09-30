// © Copyright IBM Corporation 2020. All rights reserved.
// SPDX-License-Identifier: Apache2.0
// 
// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Obligation obligation
//
// Obligation
// swagger:model Obligation
type Obligation struct {

	// balance
	Balance *AssetBalance `json:"balance,omitempty"`

	// Identifier of the Participant who owes the balance.
	// Max Length: 32
	// Min Length: 5
	// Pattern: ^[a-zA-Z0-9-]{5,32}$
	ParticipantID string `json:"participant_id,omitempty"`
}

// Validate validates this obligation
func (m *Obligation) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBalance(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParticipantID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Obligation) validateBalance(formats strfmt.Registry) error {

	if swag.IsZero(m.Balance) { // not required
		return nil
	}

	if m.Balance != nil {
		if err := m.Balance.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("balance")
			}
			return err
		}
	}

	return nil
}

func (m *Obligation) validateParticipantID(formats strfmt.Registry) error {

	if swag.IsZero(m.ParticipantID) { // not required
		return nil
	}

	if err := validate.MinLength("participant_id", "body", string(m.ParticipantID), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("participant_id", "body", string(m.ParticipantID), 32); err != nil {
		return err
	}

	if err := validate.Pattern("participant_id", "body", string(m.ParticipantID), `^[a-zA-Z0-9-]{5,32}$`); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Obligation) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Obligation) UnmarshalBinary(b []byte) error {
	var res Obligation
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}