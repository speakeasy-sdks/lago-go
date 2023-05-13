# AddOns

## Overview

Everything about Add-on collection

Find out more
<https://doc.getlago.com/docs/api/add_ons/add-on-object>
### Available Operations

* [Apply](#apply) - Apply an add-on to a customer
* [Create](#create) - Create a new add-on
* [Destroy](#destroy) - Delete an add-on
* [Find](#find) - Find add-on by code
* [FindAll](#findall) - Find add-ons
* [Update](#update) - Update an existing add-on

## Apply

Apply an add-on to a customer

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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.AddOns.Apply(ctx, shared.AppliedAddOnInput{
        AppliedAddOn: shared.AppliedAddOnInputAppliedAddOn{
            AddOnCode: "code",
            AmountCents: lago.Int64(1200),
            AmountCurrency: lago.String("EUR"),
            ExternalCustomerID: "1234567",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppliedAddOn != nil {
        // handle response
    }
}
```

## Create

Create a new add-on

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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.AddOns.Create(ctx, shared.AddOnInput{
        AddOn: shared.AddOnInputAddOn{
            AmountCents: lago.Int64(1200),
            AmountCurrency: lago.String("EUR"),
            Code: lago.String("example_code"),
            Description: lago.String("description"),
            Name: lago.String("example name"),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AddOn != nil {
        // handle response
    }
}
```

## Destroy

Delete an add-on

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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.AddOns.Destroy(ctx, operations.DestroyAddOnRequest{
        Code: "example_code",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AddOn != nil {
        // handle response
    }
}
```

## Find

Return a single add-on

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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.AddOns.Find(ctx, operations.FindAddOnRequest{
        Code: "example_code",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AddOn != nil {
        // handle response
    }
}
```

## FindAll

Find all add-ons in certain organisation

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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.AddOns.FindAll(ctx, operations.FindAllAddOnsRequest{
        Page: lago.Int64(2),
        PerPage: lago.Int64(20),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AddOns != nil {
        // handle response
    }
}
```

## Update

Update an existing add-on by code

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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.AddOns.Update(ctx, operations.UpdateAddOnRequest{
        AddOnInput: shared.AddOnInput{
            AddOn: shared.AddOnInputAddOn{
                AmountCents: lago.Int64(1200),
                AmountCurrency: lago.String("EUR"),
                Code: lago.String("example_code"),
                Description: lago.String("description"),
                Name: lago.String("example name"),
            },
        },
        Code: "example_code",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AddOn != nil {
        // handle response
    }
}
```
