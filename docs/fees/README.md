# Fees

## Overview

Everything about Fees

Find out more
<https://doc.getlago.com/docs/api/invoices/invoice-object#fee-object>
### Available Operations

* [Find](#find) - Find fee by ID
* [FindAll](#findall) - Find all fees
* [Update](#update) - Update an existing fee

## Find

Return a single fee

### Example Usage

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
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Fees.Find(ctx, operations.FindFeeRequest{
        ID: "1a901a90-1a90-1a90-1a90-1a901a901a90",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FeeObject != nil {
        // handle response
    }
}
```

## FindAll

Find all fees of an organization and filter them

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/lago-go"
	"github.com/speakeasy-sdks/lago-go/pkg/models/operations"
	"github.com/speakeasy-sdks/lago-go/pkg/types"
)

func main() {
    s := lago.New(
        lago.WithSecurity(shared.Security{
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Fees.FindAll(ctx, operations.FindAllFeesRequest{
        BillableMetricCode: lago.String("bm_code"),
        CreatedAtFrom: types.MustTimeFromString("2023-03-28T12:21:51Z"),
        CreatedAtTo: types.MustTimeFromString("2023-03-28T12:21:51Z"),
        Currency: lago.String("EUR"),
        ExternalCustomerID: lago.String("12345"),
        ExternalSubscriptionID: lago.String("12345"),
        FailedAtFrom: types.MustTimeFromString("2023-03-28T12:21:51Z"),
        FailedAtTo: types.MustTimeFromString("2023-03-28T12:21:51Z"),
        FeeType: operations.FindAllFeesFeeTypeEnumCharge.ToPointer(),
        Page: lago.Int64(2),
        PaymentStatus: operations.FindAllFeesPaymentStatusEnumSucceeded.ToPointer(),
        PerPage: lago.Int64(20),
        RefundedAtFrom: types.MustTimeFromString("2023-03-28T12:21:51Z"),
        RefundedAtTo: types.MustTimeFromString("2023-03-28T12:21:51Z"),
        SucceededAtFrom: types.MustTimeFromString("2023-03-28T12:21:51Z"),
        SucceededAtTo: types.MustTimeFromString("2023-03-28T12:21:51Z"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FeesPaginated != nil {
        // handle response
    }
}
```

## Update

Update an existing fee

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/lago-go"
	"github.com/speakeasy-sdks/lago-go/pkg/models/operations"
	"github.com/speakeasy-sdks/lago-go/pkg/models/shared"
)

func main() {
    s := lago.New(
        lago.WithSecurity(shared.Security{
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Fees.Update(ctx, operations.UpdateFeeRequest{
        FeeUpdateInput: &shared.FeeUpdateInput{
            Invoice: shared.FeeUpdateInputInvoice{
                PaymentStatus: shared.FeeUpdateInputInvoicePaymentStatusEnumRefunded,
            },
        },
        ID: "1a901a90-1a90-1a90-1a90-1a901a901a90",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.FeeObject != nil {
        // handle response
    }
}
```
