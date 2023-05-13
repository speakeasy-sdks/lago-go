<div align="center">
    <img src="https://user-images.githubusercontent.com/6267663/230070609-43e6bc4c-e839-49ac-82b8-04ebc5ff3a89.svg" width="250">
    <h1>Go SDK</h1>
   <p>Open-source metering and usage-based billing</p>
   <a href="https://doc.getlago.com/docs/api/intro"><img src="https://img.shields.io/static/v1?label=Docs&message=API Ref&color=000&style=for-the-badge" /></a>
   <a href="https://github.com/speakeasy-sdks/lago-go/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/lago-go/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
  <a href="https://opensource.org/licenses/MIT"><img src="https://img.shields.io/badge/License-MIT-blue.svg?style=for-the-badge" /></a>
  <a href="https://github.com/speakeasy-sdks/lago-go/releases"><img src="https://img.shields.io/github/v/release/speakeasy-sdks/lago-go?sort=semver&style=for-the-badge" /></a>
</div>

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

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/lago-go"
	"github.com/speakeasy-sdks/lago-go/pkg/models/operations"
)

func main() {
    s := lago.New(
        lago.WithSecurity(shared.Security{
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Invoices.Void(ctx, operations.FindInvoiceRequest{
        ID: "1a901a90-1a90-1a90-1a90-1a901a901a90",
    })
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


### [AddOns](docs/addons/README.md)

* [Apply](docs/addons/README.md#apply) - Apply an add-on to a customer
* [Create](docs/addons/README.md#create) - Create a new add-on
* [Destroy](docs/addons/README.md#destroy) - Delete an add-on
* [Find](docs/addons/README.md#find) - Find add-on by code
* [FindAll](docs/addons/README.md#findall) - Find add-ons
* [Update](docs/addons/README.md#update) - Update an existing add-on

### [BillableMetrics](docs/billablemetrics/README.md)

* [Create](docs/billablemetrics/README.md#create) - Create a new billable metric
* [Destroy](docs/billablemetrics/README.md#destroy) - Delete a billable metric
* [Find](docs/billablemetrics/README.md#find) - Find billable metric by code
* [FindAll](docs/billablemetrics/README.md#findall) - Find Billable metrics
* [FindAllGroups](docs/billablemetrics/README.md#findallgroups) - Find Billable metric groups
* [Update](docs/billablemetrics/README.md#update) - Update an existing billable metric

### [Coupons](docs/coupons/README.md)

* [AppliedCoupons](docs/coupons/README.md#appliedcoupons) - Find Applied Coupons
* [Apply](docs/coupons/README.md#apply) - Apply a coupon to a customer
* [Create](docs/coupons/README.md#create) - Create a new coupon
* [Destroy](docs/coupons/README.md#destroy) - Delete a coupon
* [Find](docs/coupons/README.md#find) - Find coupon by code
* [FindAll](docs/coupons/README.md#findall) - Find Coupons
* [Update](docs/coupons/README.md#update) - Update an existing coupon

### [CreditNotes](docs/creditnotes/README.md)

* [Create](docs/creditnotes/README.md#create) - Create a new Credit note
* [Download](docs/creditnotes/README.md#download) - Download an existing credit note
* [Find](docs/creditnotes/README.md#find) - Find credit note
* [FindAll](docs/creditnotes/README.md#findall) - Find Credit notes
* [Update](docs/creditnotes/README.md#update) - Update an existing credit note
* [Void](docs/creditnotes/README.md#void) - Void existing credit note

### [Customers](docs/customers/README.md)

* [Create](docs/customers/README.md#create) - Create a customer
* [CurrentUsage](docs/customers/README.md#currentusage) - Find customer current usage
* [DeleteAppliedCoupon](docs/customers/README.md#deleteappliedcoupon) - Delete customer's appplied coupon
* [Destroy](docs/customers/README.md#destroy) - Delete a customer
* [Find](docs/customers/README.md#find) - Find customer by external ID
* [FindAll](docs/customers/README.md#findall) - Find customers
* [PortalURL](docs/customers/README.md#portalurl) - Get customer portal URL

### [Events](docs/events/README.md)

* [EstimateFees](docs/events/README.md#estimatefees) - Estimate fees for an instant charge
* [BatchCreate](docs/events/README.md#batchcreate) - Create batch events
* [Create](docs/events/README.md#create) - Create a new event
* [Find](docs/events/README.md#find) - Find event by transaction ID

### [Fees](docs/fees/README.md)

* [Find](docs/fees/README.md#find) - Find fee by ID
* [FindAll](docs/fees/README.md#findall) - Find all fees
* [Update](docs/fees/README.md#update) - Update an existing fee

### [Invoices](docs/invoices/README.md)

* [Refresh](docs/invoices/README.md#refresh) - Refresh a draft invoice
* [Retry](docs/invoices/README.md#retry) - Retry invoice payment
* [Void](docs/invoices/README.md#void) - Finalize a draft invoice
* [Void](docs/invoices/README.md#void) - Download an existing invoice
* [Void](docs/invoices/README.md#void) - Update an existing invoice status
* [Void](docs/invoices/README.md#void) - Find invoice by ID
* [Void](docs/invoices/README.md#void) - Find all invoices

### [Organizations](docs/organizations/README.md)

* [Update](docs/organizations/README.md#update) - Update an existing Organization

### [Plans](docs/plans/README.md)

* [Create](docs/plans/README.md#create) - Create a new plan
* [Destroy](docs/plans/README.md#destroy) - Delete a plan
* [Find](docs/plans/README.md#find) - Fin plan by code
* [FindAll](docs/plans/README.md#findall) - Find plans
* [Update](docs/plans/README.md#update) - Update an existing plan

### [Subscriptions](docs/subscriptions/README.md)

* [Create](docs/subscriptions/README.md#create) - Assign a plan to a customer
* [Destroy](docs/subscriptions/README.md#destroy) - Terminate a subscription
* [FindAll](docs/subscriptions/README.md#findall) - Find subscriptions
* [Update](docs/subscriptions/README.md#update) - Update an existing subscription

### [Wallets](docs/wallets/README.md)

* [Create](docs/wallets/README.md#create) - Create a new wallet
* [CreateTransaction](docs/wallets/README.md#createtransaction) - Create a new wallet transaction
* [Destroy](docs/wallets/README.md#destroy) - Delete a wallet
* [Find](docs/wallets/README.md#find) - Find wallet
* [FindAll](docs/wallets/README.md#findall) - Find wallets
* [FindAllTransactions](docs/wallets/README.md#findalltransactions) - Find wallet transactions
* [Update](docs/wallets/README.md#update) - Update an existing wallet

### [Webhooks](docs/webhooks/README.md)

* [FetchPublicKey](docs/webhooks/README.md#fetchpublickey) - Fetch webhook public key
<!-- End SDK Available Operations -->

### Maturity

This SDK is in beta and therefore, we recommend pinning usage to a specific package version.
This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

### Contributions

While we value open-source contributions to this SDK, this library is generated and maintained programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release !

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
