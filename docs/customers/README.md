# Customers

## Overview

Everything about Customer collection

Find out more
<https://doc.getlago.com/docs/api/customers/customer-object>
### Available Operations

* [Create](#create) - Create a customer
* [CurrentUsage](#currentusage) - Find customer current usage
* [DeleteAppliedCoupon](#deleteappliedcoupon) - Delete customer's appplied coupon
* [Destroy](#destroy) - Delete a customer
* [Find](#find) - Find customer by external ID
* [FindAll](#findall) - Find customers
* [PortalURL](#portalurl) - Get customer portal URL

## Create

Create a new customer

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
    res, err := s.Customers.Create(ctx, shared.CustomerInput{
        Customer: shared.CustomerInputCustomer{
            AddressLine1: lago.String("address1"),
            AddressLine2: lago.String("address2"),
            BillingConfiguration: map[string]interface{}{
                "suscipit": "molestiae",
                "minus": "placeat",
            },
            City: lago.String("City"),
            Country: lago.String("CZ"),
            Currency: lago.String("EUR"),
            Email: lago.String("example@example.com"),
            ExternalID: "12345",
            LagoURL: lago.String("https://lago.url"),
            LegalName: lago.String("name1"),
            LegalNumber: lago.String("10000"),
            Metadata: []shared.CustomerInputCustomerMetadata{
                shared.CustomerInputCustomerMetadata{
                    DisplayInInvoice: lago.Bool(false),
                    ID: lago.String("1a901a90-1a90-1a90-1a90-1a901a901a90"),
                    Key: lago.String("name"),
                    Value: lago.String("John"),
                },
                shared.CustomerInputCustomerMetadata{
                    DisplayInInvoice: lago.Bool(false),
                    ID: lago.String("1a901a90-1a90-1a90-1a90-1a901a901a90"),
                    Key: lago.String("name"),
                    Value: lago.String("John"),
                },
                shared.CustomerInputCustomerMetadata{
                    DisplayInInvoice: lago.Bool(false),
                    ID: lago.String("1a901a90-1a90-1a90-1a90-1a901a901a90"),
                    Key: lago.String("name"),
                    Value: lago.String("John"),
                },
            },
            Name: lago.String("John Doe"),
            Phone: lago.String("+3551234567"),
            State: lago.String("state1"),
            Timezone: lago.String("Europe/Paris"),
            URL: lago.String("https://example.com"),
            Zipcode: lago.String("10000"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Customer != nil {
        // handle response
    }
}
```

## CurrentUsage

Return a customer current usage

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
    res, err := s.Customers.CurrentUsage(ctx, operations.FindCustomerCurrentUsageRequest{
        CustomerExternalID: "12345",
        ExternalSubscriptionID: "54321",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CustomerUsage != nil {
        // handle response
    }
}
```

## DeleteAppliedCoupon

Delete customer's appplied coupon

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
    res, err := s.Customers.DeleteAppliedCoupon(ctx, operations.DeleteAppliedCouponRequest{
        AppliedCouponID: "54321",
        CustomerExternalID: "12345",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppliedCoupon != nil {
        // handle response
    }
}
```

## Destroy

Return the deleted customer

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
    res, err := s.Customers.Destroy(ctx, operations.DeleteCustomerRequest{
        ExternalID: "12345",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Customer != nil {
        // handle response
    }
}
```

## Find

Return a single customer

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
    res, err := s.Customers.Find(ctx, operations.FindCustomerRequest{
        ExternalID: "12345",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Customer != nil {
        // handle response
    }
}
```

## FindAll

Find all customers in certain organisation

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
    res, err := s.Customers.FindAll(ctx, operations.FindAllCustomersRequest{
        Page: lago.Int64(2),
        PerPage: lago.Int64(20),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Customers != nil {
        // handle response
    }
}
```

## PortalURL

Get customer portal URL

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
    res, err := s.Customers.PortalURL(ctx, operations.GetCustomerPortalURLRequest{
        CustomerExternalID: "12345",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GetCustomerPortalURL200ApplicationJSONObject != nil {
        // handle response
    }
}
```
