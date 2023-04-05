# github.com/speakeasy-sdks/lago-go

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/lago-go
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/speakeasy-sdks/lago-go"
    "github.com/speakeasy-sdks/lago-go/pkg/models/shared"
    "github.com/speakeasy-sdks/lago-go/pkg/models/operations"
)

func main() {
    s := sdk.New(
        sdk.WithSecurity(shared.Security{
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    req := operations.FindInvoiceRequest{
        ID: "1a901a90-1a90-1a90-1a90-1a901a901a90",
    }

    ctx := context.Background()
    res, err := s.Invoices.Find(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Invoice != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### AddOns

* `Apply` - Apply an add-on to a customer
* `Create` - Create a new add-on
* `Destroy` - Delete an add-on
* `Find` - Find add-on by code
* `FindAll` - Find add-ons
* `Update` - Update an existing add-on

### BillableMetrics

* `Create` - Create a new billable metric
* `Destroy` - Delete a billable metric
* `Find` - Find billable metric by code
* `FindAll` - Find Billable metric groups
* `FindAll` - Find Billable metrics
* `Update` - Update an existing billable metric

### Coupons

* `AppliedCoupons` - Find Applied Coupons
* `Apply` - Apply a coupon to a customer
* `Create` - Create a new coupon
* `Destroy` - Delete a coupon
* `Find` - Find coupon by code
* `FindAll` - Find Coupons
* `Update` - Update an existing coupon

### CreditNotes

* `Create` - Create a new Credit note
* `Download` - Download an existing credit note
* `Find` - Find credit note
* `FindAll` - Find Credit notes
* `Update` - Update an existing credit note
* `Void` - Void existing credit note

### Customers

* `Create` - Create a customer
* `CurrentUsage` - Find customer current usage
* `DeleteAppliedCoupon` - Delete customer's appplied coupon
* `Destroy` - Delete a customer
* `Find` - Find customer by external ID
* `FindAll` - Find customers
* `PortalURL` - Get customer portal URL

### Events

* `EstimateFees` - Estimate fees for an instant charge
* `BatchCreate` - Create batch events
* `Create` - Create a new event
* `Find` - Find event by transaction ID

### Fees

* `Find` - Find fee by ID
* `FindAll` - Find all fees
* `Update` - Update an existing fee

### Invoices

* `Download` - Download an existing invoice
* `Find` - Find invoice by ID
* `FindAll` - Find all invoices
* `Retry` - Retry invoice payment
* `Update` - Update an existing invoice status
* `Void` - Refresh a draft invoice
* `Void` - Finalize a draft invoice

### Organizations

* `Update` - Update an existing Organization

### Plans

* `Create` - Create a new plan
* `Destroy` - Delete a plan
* `Find` - Fin plan by code
* `FindAll` - Find plans
* `Update` - Update an existing plan

### Subscriptions

* `Create` - Assign a plan to a customer
* `Destroy` - Terminate a subscription
* `FindAll` - Find subscriptions
* `Update` - Update an existing subscription

### Wallets

* `Create` - Create a new wallet
* `CreateTransaction` - Create a new wallet transaction
* `Destroy` - Delete a wallet
* `Find` - Find wallet
* `FindAll` - Find wallets
* `FindAllTransactions` - Find wallet transactions
* `Update` - Update an existing wallet

### Webhooks

* `FetchPublicKey` - Fetch webhook public key
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
