// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"time"
)

type OrganizationObject struct {
	AddressLine1         *string                           `json:"address_line1,omitempty"`
	AddressLine2         *string                           `json:"address_line2,omitempty"`
	BillingConfiguration *BillingConfigurationOrganization `json:"billing_configuration,omitempty"`
	City                 *string                           `json:"city,omitempty"`
	Country              *string                           `json:"country,omitempty"`
	CreatedAt            time.Time                         `json:"created_at"`
	Email                *string                           `json:"email,omitempty"`
	LagoID               *string                           `json:"lago_id,omitempty"`
	LegalName            *string                           `json:"legal_name,omitempty"`
	LegalNumber          *string                           `json:"legal_number,omitempty"`
	Name                 string                            `json:"name"`
	State                *string                           `json:"state,omitempty"`
	Timezone             *string                           `json:"timezone,omitempty"`
	WebhookURL           *string                           `json:"webhook_url,omitempty"`
	Zipcode              *string                           `json:"zipcode,omitempty"`
}
