# BillableMetrics

## Overview

Everything about Billable metric collection

Find out more
<https://doc.getlago.com/docs/api/billable_metrics/billable-metric-object>
### Available Operations

* [Create](#create) - Create a new billable metric
* [Destroy](#destroy) - Delete a billable metric
* [Find](#find) - Find billable metric by code
* [FindAll](#findall) - Find Billable metrics
* [FindAllGroups](#findallgroups) - Find Billable metric groups
* [Update](#update) - Update an existing billable metric

## Create

Create a new billable metric

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
    req := shared.BillableMetricInput{
        BillableMetric: shared.BillableMetricInputBillableMetric{
            AggregationType: shared.BillableMetricInputBillableMetricAggregationTypeEnumMaxAgg.ToPointer(),
            Code: lago.String("example_code"),
            Description: lago.String("description"),
            FieldName: lago.String("amount"),
            Group: &shared.BillableMetricGroup{
                Key: "region",
                Values: []shared.BillableMetricGroupValues{
                    shared.BillableMetricGroupValues{},
                    shared.BillableMetricGroupValues{},
                    shared.BillableMetricGroupValues{},
                },
            },
            Name: lago.String("bm1"),
        },
    }

    res, err := s.BillableMetrics.Create(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BillableMetric != nil {
        // handle response
    }
}
```

## Destroy

Delete a billable metric

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
    req := operations.DestroyBillableMetricRequest{
        Code: "example_code",
    }

    res, err := s.BillableMetrics.Destroy(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BillableMetric != nil {
        // handle response
    }
}
```

## Find

Return a single billable metric

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
    req := operations.FindBillableMetricRequest{
        Code: "example_code",
    }

    res, err := s.BillableMetrics.Find(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BillableMetric != nil {
        // handle response
    }
}
```

## FindAll

Find all billable metrics in certain organisation

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
    req := operations.FindAllBillableMetricsRequest{
        Page: lago.Int64(2),
        PerPage: lago.Int64(20),
    }

    res, err := s.BillableMetrics.FindAll(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BillableMetrics != nil {
        // handle response
    }
}
```

## FindAllGroups

Find all billable metric groups in certain organisation

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
    req := operations.FindAllBillableMetricGroupsRequest{
        Code: "example_code",
        Page: lago.Int64(2),
        PerPage: lago.Int64(20),
    }

    res, err := s.BillableMetrics.FindAllGroups(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Groups != nil {
        // handle response
    }
}
```

## Update

Update an existing billable metric by code

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
    req := operations.UpdateBillableMetricRequest{
        BillableMetricInput: shared.BillableMetricInput{
            BillableMetric: shared.BillableMetricInputBillableMetric{
                AggregationType: shared.BillableMetricInputBillableMetricAggregationTypeEnumUniqueCountAgg.ToPointer(),
                Code: lago.String("example_code"),
                Description: lago.String("description"),
                FieldName: lago.String("amount"),
                Group: &shared.BillableMetricGroup{
                    Key: "region",
                    Values: []shared.BillableMetricGroupValues{
                        shared.BillableMetricGroupValues{},
                        shared.BillableMetricGroupValues{},
                        shared.BillableMetricGroupValues{},
                        shared.BillableMetricGroupValues{},
                    },
                },
                Name: lago.String("bm1"),
            },
        },
        Code: "example_code",
    }

    res, err := s.BillableMetrics.Update(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.BillableMetric != nil {
        // handle response
    }
}
```
