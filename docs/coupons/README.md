# Coupons

## Overview

Everything about Coupon collection

Find out more
<https://doc.getlago.com/docs/api/coupons/coupon-object>
### Available Operations

* [AppliedCoupons](#appliedcoupons) - Find Applied Coupons
* [Apply](#apply) - Apply a coupon to a customer
* [Create](#create) - Create a new coupon
* [Destroy](#destroy) - Delete a coupon
* [Find](#find) - Find coupon by code
* [FindAll](#findall) - Find Coupons
* [Update](#update) - Update an existing coupon

## AppliedCoupons

Find all applied coupons

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
    res, err := s.Coupons.AppliedCoupons(ctx, operations.FindAllAppliedCouponsRequest{
        ExternalCustomerID: lago.String("12345"),
        Page: lago.Int64(2),
        PerPage: lago.Int64(20),
        Status: operations.FindAllAppliedCouponsStatusEnumTerminated.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppliedCoupons != nil {
        // handle response
    }
}
```

## Apply

Apply a coupon to a customer

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
    res, err := s.Coupons.Apply(ctx, shared.AppliedCouponInput{
        AppliedCoupon: shared.AppliedCouponInputAppliedCoupon{
            AmountCents: lago.Int64(1200),
            AmountCurrency: lago.String("EUR"),
            CouponCode: "example_code",
            ExternalCustomerID: "123456",
            Frequency: shared.AppliedCouponInputAppliedCouponFrequencyEnumRecurring.ToPointer(),
            FrequencyDuration: lago.Int64(3),
            PercentageRate: lago.Float64(25),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.AppliedCoupon != nil {
        // handle response
    }
}
```

## Create

Create a new coupon

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
    res, err := s.Coupons.Create(ctx, shared.CouponInput{
        Coupon: shared.CouponInputCoupon{
            AmountCents: lago.Int64(1200),
            AmountCurrency: lago.String("EUR"),
            AppliesTo: &shared.CouponInputCouponAppliesTo{
                PlanCodes: []string{
                    "code1",
                    "code1",
                    "code1",
                },
            },
            Code: lago.String("example_code"),
            CouponType: shared.CouponInputCouponCouponTypeEnumPercentage.ToPointer(),
            Expiration: shared.CouponInputCouponExpirationEnumNoExpiration.ToPointer(),
            ExpirationAt: types.MustTimeFromString("2022-09-14T23:59:59Z"),
            Frequency: shared.CouponInputCouponFrequencyEnumRecurring.ToPointer(),
            FrequencyDuration: lago.Int64(3),
            Name: lago.String("coupon1"),
            PercentageRate: lago.Float64(25),
            Reusable: lago.Bool(true),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Coupon != nil {
        // handle response
    }
}
```

## Destroy

Delete a coupon

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
    res, err := s.Coupons.Destroy(ctx, operations.DestroyCouponRequest{
        Code: "example_code",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Coupon != nil {
        // handle response
    }
}
```

## Find

Return a single coupon

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
    res, err := s.Coupons.Find(ctx, operations.FindCouponRequest{
        Code: "example_code",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Coupon != nil {
        // handle response
    }
}
```

## FindAll

Find all coupons in certain organisation

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
    res, err := s.Coupons.FindAll(ctx, operations.FindAllCouponsRequest{
        Page: lago.Int64(2),
        PerPage: lago.Int64(20),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Coupons != nil {
        // handle response
    }
}
```

## Update

Update an existing coupon by code

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
    res, err := s.Coupons.Update(ctx, operations.UpdateCouponRequest{
        CouponInput: shared.CouponInput{
            Coupon: shared.CouponInputCoupon{
                AmountCents: lago.Int64(1200),
                AmountCurrency: lago.String("EUR"),
                AppliesTo: &shared.CouponInputCouponAppliesTo{
                    PlanCodes: []string{
                        "code1",
                        "code1",
                        "code1",
                    },
                },
                Code: lago.String("example_code"),
                CouponType: shared.CouponInputCouponCouponTypeEnumFixedAmount.ToPointer(),
                Expiration: shared.CouponInputCouponExpirationEnumNoExpiration.ToPointer(),
                ExpirationAt: types.MustTimeFromString("2022-09-14T23:59:59Z"),
                Frequency: shared.CouponInputCouponFrequencyEnumOnce.ToPointer(),
                FrequencyDuration: lago.Int64(3),
                Name: lago.String("coupon1"),
                PercentageRate: lago.Float64(25),
                Reusable: lago.Bool(true),
            },
        },
        Code: "example_code",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Coupon != nil {
        // handle response
    }
}
```
