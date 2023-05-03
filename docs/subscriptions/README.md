# Subscriptions

## Overview

Everything about Subscription collection

Find out more
<https://doc.getlago.com/docs/api/subscriptions/subscription-object>
### Available Operations

* [Create](#create) - Assign a plan to a customer
* [Destroy](#destroy) - Terminate a subscription
* [FindAll](#findall) - Find subscriptions
* [Update](#update) - Update an existing subscription

## Create

Assign a plan to a customer

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/lago-go"
	"github.com/speakeasy-sdks/lago-go/pkg/models/shared"
	"github.com/speakeasy-sdks/lago-go/pkg/types"
)

func main() {
    s := lago.New(
        lago.WithSecurity(shared.Security{
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Subscriptions.Create(ctx, shared.SubscriptionCreateInput{
        Subscription: shared.SubscriptionCreateInputSubscription{
            BillingTime: shared.SubscriptionCreateInputSubscriptionBillingTimeEnumAnniversary.ToPointer(),
            ExternalCustomerID: "12345",
            ExternalID: "54321",
            Name: lago.String("Test name"),
            PlanCode: "example_code",
            SubscriptionAt: types.MustTimeFromString("2022-08-08T00:00:00Z"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Subscription != nil {
        // handle response
    }
}
```

## Destroy

Terminate a subscription

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
    res, err := s.Subscriptions.Destroy(ctx, operations.DestroySubscriptionRequest{
        ExternalID: "example_id",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Subscription != nil {
        // handle response
    }
}
```

## FindAll

Find all suscriptions for certain customer

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
    res, err := s.Subscriptions.FindAll(ctx, operations.FindAllSubscriptionsRequest{
        ExternalCustomerID: "12345",
        Page: lago.Int64(2),
        PerPage: lago.Int64(20),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Subscriptions != nil {
        // handle response
    }
}
```

## Update

Update an existing subscription by external ID

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/lago-go"
	"github.com/speakeasy-sdks/lago-go/pkg/models/operations"
	"github.com/speakeasy-sdks/lago-go/pkg/models/shared"
	"github.com/speakeasy-sdks/lago-go/pkg/types"
)

func main() {
    s := lago.New(
        lago.WithSecurity(shared.Security{
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Subscriptions.Update(ctx, operations.UpdateSubscriptionRequest{
        SubscriptionUpdateInput: shared.SubscriptionUpdateInput{
            Subscription: shared.SubscriptionUpdateInputSubscription{
                Name: lago.String("New name"),
                SubscriptionAt: types.MustTimeFromString("2022-08-08T00:00:00Z"),
            },
        },
        ExternalID: "example_id",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Subscription != nil {
        // handle response
    }
}
```
