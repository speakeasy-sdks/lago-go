# Invoices

## Overview

Everything about Invoice collection

Find out more
<https://doc.getlago.com/docs/api/invoices/invoice-object>
### Available Operations

* [Download](#download) - Download an existing invoice
* [Finalize](#finalize) - Finalize a draft invoice
* [Find](#find) - Find invoice by ID
* [FindAll](#findall) - Find all invoices
* [Retry](#retry) - Retry invoice payment
* [Update](#update) - Update an existing invoice status
* [Void](#void) - Refresh a draft invoice

## Download

Download an existing invoice

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
    res, err := s.Invoices.Download(ctx, operations.DownloadInvoiceRequest{
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

## Finalize

Finalize a draft invoice

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
    res, err := s.Invoices.Finalize(ctx, operations.FinalizeInvoiceRequest{
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

## Find

Return a single invoice

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
    res, err := s.Invoices.Find(ctx, operations.FindInvoiceRequest{
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

## FindAll

Find all invoices in certain organisation

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
    res, err := s.Invoices.FindAll(ctx, operations.FindAllInvoicesRequest{
        ExternalCustomerID: lago.String("12345"),
        IssuingDateFrom: types.MustDateFromString("2022-07-08"),
        IssuingDateTo: types.MustDateFromString("2022-08-09"),
        Page: lago.Int64(2),
        PerPage: lago.Int64(20),
        Status: operations.FindAllInvoicesStatusEnumFinalized.ToPointer(),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Invoices != nil {
        // handle response
    }
}
```

## Retry

Retry invoice payment

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
    res, err := s.Invoices.Retry(ctx, operations.RetryPaymentRequest{
        ID: "1a901a90-1a90-1a90-1a90-1a901a901a90",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

## Update

Update an existing invoice

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
    res, err := s.Invoices.Update(ctx, operations.UpdateInvoiceRequest{
        InvoiceInput: shared.InvoiceInput{
            Invoice: shared.InvoiceInputInvoice{
                Metadata: []shared.InvoiceInputInvoiceMetadata{
                    shared.InvoiceInputInvoiceMetadata{
                        ID: lago.String("1a901a90-1a90-1a90-1a90-1a901a901a90"),
                        Key: lago.String("name"),
                        Value: lago.String("John"),
                    },
                    shared.InvoiceInputInvoiceMetadata{
                        ID: lago.String("1a901a90-1a90-1a90-1a90-1a901a901a90"),
                        Key: lago.String("name"),
                        Value: lago.String("John"),
                    },
                    shared.InvoiceInputInvoiceMetadata{
                        ID: lago.String("1a901a90-1a90-1a90-1a90-1a901a901a90"),
                        Key: lago.String("name"),
                        Value: lago.String("John"),
                    },
                    shared.InvoiceInputInvoiceMetadata{
                        ID: lago.String("1a901a90-1a90-1a90-1a90-1a901a901a90"),
                        Key: lago.String("name"),
                        Value: lago.String("John"),
                    },
                },
                PaymentStatus: shared.InvoiceInputInvoicePaymentStatusEnumSucceeded,
            },
        },
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

## Void

Refresh a draft invoice

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
    res, err := s.Invoices.Void(ctx, operations.RefreshInvoiceRequest{
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
