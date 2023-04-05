// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type CustomerInputCustomerMetadata struct {
	DisplayInInvoice *bool   `json:"display_in_invoice,omitempty"`
	ID               *string `json:"id,omitempty"`
	Key              *string `json:"key,omitempty"`
	Value            *string `json:"value,omitempty"`
}

type CustomerInputCustomer struct {
	AddressLine1         *string                         `json:"address_line1,omitempty"`
	AddressLine2         *string                         `json:"address_line2,omitempty"`
	BillingConfiguration map[string]interface{}          `json:"billing_configuration,omitempty"`
	City                 *string                         `json:"city,omitempty"`
	Country              *string                         `json:"country,omitempty"`
	Currency             *string                         `json:"currency,omitempty"`
	Email                *string                         `json:"email,omitempty"`
	ExternalID           string                          `json:"external_id"`
	LagoURL              *string                         `json:"lago_url,omitempty"`
	LegalName            *string                         `json:"legal_name,omitempty"`
	LegalNumber          *string                         `json:"legal_number,omitempty"`
	Metadata             []CustomerInputCustomerMetadata `json:"metadata,omitempty"`
	Name                 *string                         `json:"name,omitempty"`
	Phone                *string                         `json:"phone,omitempty"`
	State                *string                         `json:"state,omitempty"`
	Timezone             *string                         `json:"timezone,omitempty"`
	URL                  *string                         `json:"url,omitempty"`
	Zipcode              *string                         `json:"zipcode,omitempty"`
}

// CustomerInput - Customer payload
type CustomerInput struct {
	Customer CustomerInputCustomer `json:"customer"`
}
