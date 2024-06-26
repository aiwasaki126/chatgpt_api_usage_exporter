// Code generated by go-swagger; DO NOT EDIT.

package response

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Data data
//
// swagger:model data
type Data struct {

	// Aggregation timestamp
	// Example: 1716811620
	// Required: true
	AggregationTimestamp *int64 `json:"aggregation_timestamp"`

	// API key ID
	APIKeyID *string `json:"api_key_id,omitempty"`

	// API key name
	APIKeyName *string `json:"api_key_name,omitempty"`

	// API key redacted
	APIKeyRedacted *string `json:"api_key_redacted,omitempty"`

	// API key type
	APIKeyType *string `json:"api_key_type,omitempty"`

	// Email
	Email *string `json:"email,omitempty"`

	// Number of context tokens
	// Example: 1000
	// Required: true
	NContextTokensTotal *int64 `json:"n_context_tokens_total"`

	// Number of generated tokens
	// Example: 1000
	// Required: true
	NGeneratedTokensTotal *int64 `json:"n_generated_tokens_total"`

	// Number of requests
	// Example: 100
	// Required: true
	NRequests *int64 `json:"n_requests"`

	// Operation
	// Example: completion
	// Required: true
	Operation *string `json:"operation"`

	// Organization ID
	// Example: org-1234
	// Required: true
	OrganizationID *string `json:"organization_id"`

	// Organization name
	// Example: OpenAI
	// Required: true
	OrganizationName *string `json:"organization_name"`

	// Project ID
	ProjectID *string `json:"project_id,omitempty"`

	// Project name
	ProjectName *string `json:"project_name,omitempty"`

	// Model snapshot ID
	// Example: gpt-3.5-turbo-16k-0613
	// Required: true
	SnapshotID *string `json:"snapshot_id"`
}

// Validate validates this data
func (m *Data) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAggregationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNContextTokensTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNGeneratedTokensTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNRequests(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrganizationName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Data) validateAggregationTimestamp(formats strfmt.Registry) error {

	if err := validate.Required("aggregation_timestamp", "body", m.AggregationTimestamp); err != nil {
		return err
	}

	return nil
}

func (m *Data) validateNContextTokensTotal(formats strfmt.Registry) error {

	if err := validate.Required("n_context_tokens_total", "body", m.NContextTokensTotal); err != nil {
		return err
	}

	return nil
}

func (m *Data) validateNGeneratedTokensTotal(formats strfmt.Registry) error {

	if err := validate.Required("n_generated_tokens_total", "body", m.NGeneratedTokensTotal); err != nil {
		return err
	}

	return nil
}

func (m *Data) validateNRequests(formats strfmt.Registry) error {

	if err := validate.Required("n_requests", "body", m.NRequests); err != nil {
		return err
	}

	return nil
}

func (m *Data) validateOperation(formats strfmt.Registry) error {

	if err := validate.Required("operation", "body", m.Operation); err != nil {
		return err
	}

	return nil
}

func (m *Data) validateOrganizationID(formats strfmt.Registry) error {

	if err := validate.Required("organization_id", "body", m.OrganizationID); err != nil {
		return err
	}

	return nil
}

func (m *Data) validateOrganizationName(formats strfmt.Registry) error {

	if err := validate.Required("organization_name", "body", m.OrganizationName); err != nil {
		return err
	}

	return nil
}

func (m *Data) validateSnapshotID(formats strfmt.Registry) error {

	if err := validate.Required("snapshot_id", "body", m.SnapshotID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this data based on context it is used
func (m *Data) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Data) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Data) UnmarshalBinary(b []byte) error {
	var res Data
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
