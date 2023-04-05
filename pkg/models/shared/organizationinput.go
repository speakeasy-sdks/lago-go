// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

type OrganizationInputOrganizationEmailSettingsEnum string

const (
	OrganizationInputOrganizationEmailSettingsEnumInvoiceFinalized  OrganizationInputOrganizationEmailSettingsEnum = "invoice.finalized"
	OrganizationInputOrganizationEmailSettingsEnumCreditNoteCreated OrganizationInputOrganizationEmailSettingsEnum = "credit_note.created"
)

func (e *OrganizationInputOrganizationEmailSettingsEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	switch s {
	case "invoice.finalized":
		fallthrough
	case "credit_note.created":
		*e = OrganizationInputOrganizationEmailSettingsEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid value for OrganizationInputOrganizationEmailSettingsEnum: %s", s)
	}
}

type OrganizationInputOrganization struct {
	AddressLine1         *string                                          `json:"address_line1,omitempty"`
	AddressLine2         *string                                          `json:"address_line2,omitempty"`
	BillingConfiguration *BillingConfigurationOrganization                `json:"billing_configuration,omitempty"`
	City                 *string                                          `json:"city,omitempty"`
	Country              *string                                          `json:"country,omitempty"`
	Email                *string                                          `json:"email,omitempty"`
	EmailSettings        []OrganizationInputOrganizationEmailSettingsEnum `json:"email_settings,omitempty"`
	LegalName            *string                                          `json:"legal_name,omitempty"`
	LegalNumber          *string                                          `json:"legal_number,omitempty"`
	State                *string                                          `json:"state,omitempty"`
	Timezone             *string                                          `json:"timezone,omitempty"`
	WebhookURL           *string                                          `json:"webhook_url,omitempty"`
	Zipcode              *string                                          `json:"zipcode,omitempty"`
}

// OrganizationInput - Update an existing organization
type OrganizationInput struct {
	Organization OrganizationInputOrganization `json:"organization"`
}
