# Events

## Overview

Everything about Event collection

Find out more
<https://doc.getlago.com/docs/api/events/event-object>
### Available Operations

* [EstimateFees](#estimatefees) - Estimate fees for an instant charge
* [BatchCreate](#batchcreate) - Create batch events
* [Create](#create) - Create a new event
* [Find](#find) - Find event by transaction ID

## EstimateFees

Estimate the fees that would be created after reception of an event for a billable metric attached to one or multiple instant charges

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
    res, err := s.Events.EstimateFees(ctx, shared.EventEstimateFeesInput{
        Event: shared.EventEstimateFeesInputEvent{
            Code: "code",
            ExternalCustomerID: lago.String("654321"),
            ExternalSubscriptionID: lago.String("123456"),
            Properties: map[string]interface{}{
                "excepturi": "nisi",
                "recusandae": "temporibus",
            },
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Fees != nil {
        // handle response
    }
}
```

## BatchCreate

Create batch events

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
    res, err := s.Events.BatchCreate(ctx, shared.BatchEventInput{
        Event: shared.BatchEventInputEvent{
            Code: "code",
            ExternalCustomerID: lago.String("654321"),
            ExternalSubscriptionIds: []string{
                "quis",
            },
            Properties: map[string]interface{}{
                "deserunt": "perferendis",
            },
            Timestamp: lago.Int64(1669823754),
            TransactionID: "123456",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## Create

Create a new event

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
    res, err := s.Events.Create(ctx, shared.EventInput{
        Event: shared.EventInputEvent{
            Code: "code",
            ExternalCustomerID: lago.String("654321"),
            ExternalSubscriptionID: lago.String("123456"),
            Properties: map[string]interface{}{
                "repellendus": "sapiente",
                "quo": "odit",
            },
            Timestamp: lago.Int64(1669823754),
            TransactionID: "123456",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## Find

Return a single event

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
    res, err := s.Events.Find(ctx, operations.FindEventRequest{
        ID: "12345",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Event != nil {
        // handle response
    }
}
```
