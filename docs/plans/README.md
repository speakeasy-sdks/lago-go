# Plans

## Overview

Everything about Plan collection

Find out more
<https://doc.getlago.com/docs/api/plans/plan-object>
### Available Operations

* [Create](#create) - Create a new plan
* [Destroy](#destroy) - Delete a plan
* [Find](#find) - Fin plan by code
* [FindAll](#findall) - Find plans
* [Update](#update) - Update an existing plan

## Create

Create a new plan

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
    req := shared.PlanInput{
        Plan: shared.PlanInputPlan{
            AmountCents: lago.Int64(1200),
            AmountCurrency: lago.String("EUR"),
            BillChargesMonthly: lago.Bool(false),
            Charges: []shared.PlanInputPlanCharges{
                shared.PlanInputPlanCharges{
                    BillableMetricID: lago.String("1a901a90-1a90-1a90-1a90-1a901a901a90"),
                    ChargeModel: shared.PlanInputPlanChargesChargeModelEnumStandard.ToPointer(),
                    GroupProperties: []shared.PlanInputPlanChargesGroupProperties{
                        shared.PlanInputPlanChargesGroupProperties{
                            GroupID: "123456",
                            Values: map[string]interface{}{
                                "occaecati": "fugit",
                                "deleniti": "hic",
                                "optio": "totam",
                            },
                        },
                        shared.PlanInputPlanChargesGroupProperties{
                            GroupID: "123456",
                            Values: map[string]interface{}{
                                "commodi": "molestiae",
                            },
                        },
                        shared.PlanInputPlanChargesGroupProperties{
                            GroupID: "123456",
                            Values: map[string]interface{}{
                                "qui": "impedit",
                                "cum": "esse",
                            },
                        },
                    },
                    ID: lago.String("1a901a90-1a90-1a90-1a90-1a901a901a90"),
                    Instant: lago.Bool(false),
                    Properties: map[string]interface{}{
                        "excepturi": "aspernatur",
                    },
                },
                shared.PlanInputPlanCharges{
                    BillableMetricID: lago.String("1a901a90-1a90-1a90-1a90-1a901a901a90"),
                    ChargeModel: shared.PlanInputPlanChargesChargeModelEnumStandard.ToPointer(),
                    GroupProperties: []shared.PlanInputPlanChargesGroupProperties{
                        shared.PlanInputPlanChargesGroupProperties{
                            GroupID: "123456",
                            Values: map[string]interface{}{
                                "sed": "iste",
                                "dolor": "natus",
                                "laboriosam": "hic",
                            },
                        },
                        shared.PlanInputPlanChargesGroupProperties{
                            GroupID: "123456",
                            Values: map[string]interface{}{
                                "fuga": "in",
                                "corporis": "iste",
                                "iure": "saepe",
                                "quidem": "architecto",
                            },
                        },
                    },
                    ID: lago.String("1a901a90-1a90-1a90-1a90-1a901a901a90"),
                    Instant: lago.Bool(false),
                    Properties: map[string]interface{}{
                        "reiciendis": "est",
                    },
                },
                shared.PlanInputPlanCharges{
                    BillableMetricID: lago.String("1a901a90-1a90-1a90-1a90-1a901a901a90"),
                    ChargeModel: shared.PlanInputPlanChargesChargeModelEnumPercentage.ToPointer(),
                    GroupProperties: []shared.PlanInputPlanChargesGroupProperties{
                        shared.PlanInputPlanChargesGroupProperties{
                            GroupID: "123456",
                            Values: map[string]interface{}{
                                "dolorem": "corporis",
                            },
                        },
                        shared.PlanInputPlanChargesGroupProperties{
                            GroupID: "123456",
                            Values: map[string]interface{}{
                                "nobis": "enim",
                            },
                        },
                        shared.PlanInputPlanChargesGroupProperties{
                            GroupID: "123456",
                            Values: map[string]interface{}{
                                "nemo": "minima",
                                "excepturi": "accusantium",
                                "iure": "culpa",
                            },
                        },
                    },
                    ID: lago.String("1a901a90-1a90-1a90-1a90-1a901a901a90"),
                    Instant: lago.Bool(false),
                    Properties: map[string]interface{}{
                        "sapiente": "architecto",
                        "mollitia": "dolorem",
                        "culpa": "consequuntur",
                        "repellat": "mollitia",
                    },
                },
            },
            Code: lago.String("example_code"),
            Description: lago.String("description"),
            Interval: shared.PlanInputPlanIntervalEnumMonthly.ToPointer(),
            Name: lago.String("example name"),
            PayInAdvance: lago.Bool(true),
            TrialPeriod: lago.Float64(2),
        },
    }

    res, err := s.Plans.Create(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Plan != nil {
        // handle response
    }
}
```

## Destroy

Delete a plan

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
    req := operations.DestroyPlanRequest{
        Code: "example_code",
    }

    res, err := s.Plans.Destroy(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Plan != nil {
        // handle response
    }
}
```

## Find

Return a single plan

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
    req := operations.FindPlanRequest{
        Code: "example_code",
    }

    res, err := s.Plans.Find(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Plan != nil {
        // handle response
    }
}
```

## FindAll

Find all plans in certain organisation

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
    req := operations.FindAllPlansRequest{
        Page: lago.Int64(2),
        PerPage: lago.Int64(20),
    }

    res, err := s.Plans.FindAll(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Plans != nil {
        // handle response
    }
}
```

## Update

Update an existing plan by code

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
    req := operations.UpdatePlanRequest{
        PlanInput: shared.PlanInput{
            Plan: shared.PlanInputPlan{
                AmountCents: lago.Int64(1200),
                AmountCurrency: lago.String("EUR"),
                BillChargesMonthly: lago.Bool(false),
                Charges: []shared.PlanInputPlanCharges{
                    shared.PlanInputPlanCharges{
                        BillableMetricID: lago.String("1a901a90-1a90-1a90-1a90-1a901a901a90"),
                        ChargeModel: shared.PlanInputPlanChargesChargeModelEnumPackage.ToPointer(),
                        GroupProperties: []shared.PlanInputPlanChargesGroupProperties{
                            shared.PlanInputPlanChargesGroupProperties{
                                GroupID: "123456",
                                Values: map[string]interface{}{
                                    "velit": "error",
                                    "quia": "quis",
                                },
                            },
                            shared.PlanInputPlanChargesGroupProperties{
                                GroupID: "123456",
                                Values: map[string]interface{}{
                                    "laborum": "animi",
                                },
                            },
                        },
                        ID: lago.String("1a901a90-1a90-1a90-1a90-1a901a901a90"),
                        Instant: lago.Bool(false),
                        Properties: map[string]interface{}{
                            "odit": "quo",
                            "sequi": "tenetur",
                        },
                    },
                    shared.PlanInputPlanCharges{
                        BillableMetricID: lago.String("1a901a90-1a90-1a90-1a90-1a901a901a90"),
                        ChargeModel: shared.PlanInputPlanChargesChargeModelEnumGraduated.ToPointer(),
                        GroupProperties: []shared.PlanInputPlanChargesGroupProperties{
                            shared.PlanInputPlanChargesGroupProperties{
                                GroupID: "123456",
                                Values: map[string]interface{}{
                                    "aut": "quasi",
                                    "error": "temporibus",
                                    "laborum": "quasi",
                                    "reiciendis": "voluptatibus",
                                },
                            },
                            shared.PlanInputPlanChargesGroupProperties{
                                GroupID: "123456",
                                Values: map[string]interface{}{
                                    "nihil": "praesentium",
                                    "voluptatibus": "ipsa",
                                    "omnis": "voluptate",
                                    "cum": "perferendis",
                                },
                            },
                            shared.PlanInputPlanChargesGroupProperties{
                                GroupID: "123456",
                                Values: map[string]interface{}{
                                    "reprehenderit": "ut",
                                },
                            },
                        },
                        ID: lago.String("1a901a90-1a90-1a90-1a90-1a901a901a90"),
                        Instant: lago.Bool(false),
                        Properties: map[string]interface{}{
                            "dicta": "corporis",
                            "dolore": "iusto",
                            "dicta": "harum",
                            "enim": "accusamus",
                        },
                    },
                },
                Code: lago.String("example_code"),
                Description: lago.String("description"),
                Interval: shared.PlanInputPlanIntervalEnumMonthly.ToPointer(),
                Name: lago.String("example name"),
                PayInAdvance: lago.Bool(true),
                TrialPeriod: lago.Float64(2),
            },
        },
        Code: "example_code",
    }

    res, err := s.Plans.Update(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Plan != nil {
        // handle response
    }
}
```
