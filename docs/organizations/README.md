# Organizations

## Overview

Everything about Organization collection

Find out more
<https://doc.getlago.com/docs/api/organizations/organization-object>
### Available Operations

* [Update](#update) - Update an existing Organization

## Update

Update an existing organization

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/lago-go"
	"github.com/speakeasy-sdks/lago-go/pkg/models/shared"
)

func main() {
    s := lago.New(
        lago.WithSecurity(shared.Security{
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()    
    req := shared.OrganizationInput{
        Organization: shared.OrganizationInputOrganization{
            AddressLine1: lago.String("address1"),
            AddressLine2: lago.String("address2"),
            BillingConfiguration: &shared.BillingConfigurationOrganization{
                DocumentLocale: lago.String("fr"),
                InvoiceFooter: lago.String("text"),
                InvoiceGracePeriod: lago.Int64(5),
                VatRate: lago.Float64(25),
            },
            City: lago.String("City"),
            Country: lago.String("CZ"),
            Email: lago.String("example@example.com"),
            EmailSettings: []shared.OrganizationInputOrganizationEmailSettingsEnum{
                shared.OrganizationInputOrganizationEmailSettingsEnumCreditNoteCreated,
                shared.OrganizationInputOrganizationEmailSettingsEnumInvoiceFinalized,
                shared.OrganizationInputOrganizationEmailSettingsEnumCreditNoteCreated,
                shared.OrganizationInputOrganizationEmailSettingsEnumCreditNoteCreated,
            },
            LegalName: lago.String("name1"),
            LegalNumber: lago.String("10000"),
            State: lago.String("state1"),
            Timezone: lago.String("Europe/Paris"),
            WebhookURL: lago.String("https://example.com"),
            Zipcode: lago.String("10000"),
        },
    }

    res, err := s.Organizations.Update(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Organization != nil {
        // handle response
    }
}
```
