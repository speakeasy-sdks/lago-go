# CreditNotes

## Overview

Everything about Credit notes collection

### Available Operations

* [Create](#create) - Create a new Credit note
* [Download](#download) - Download an existing credit note
* [Find](#find) - Find credit note
* [FindAll](#findall) - Find Credit notes
* [Update](#update) - Update an existing credit note
* [Void](#void) - Void existing credit note

## Create

Create a new credit note

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
    res, err := s.CreditNotes.Create(ctx, shared.CreditNoteInput{
        CreditNote: shared.CreditNoteInputCreditNote{
            CreditAmountCents: 20,
            Description: lago.String("description"),
            InvoiceID: "1a901a90-1a90-1a90-1a90-1a901a901a90",
            Items: []shared.CreditNoteInputCreditNoteItems{
                shared.CreditNoteInputCreditNoteItems{
                    AmountCents: 20,
                    FeeID: "1a901a90-1a90-1a90-1a90-1a901a901a90",
                },
                shared.CreditNoteInputCreditNoteItems{
                    AmountCents: 20,
                    FeeID: "1a901a90-1a90-1a90-1a90-1a901a901a90",
                },
                shared.CreditNoteInputCreditNoteItems{
                    AmountCents: 20,
                    FeeID: "1a901a90-1a90-1a90-1a90-1a901a901a90",
                },
                shared.CreditNoteInputCreditNoteItems{
                    AmountCents: 20,
                    FeeID: "1a901a90-1a90-1a90-1a90-1a901a901a90",
                },
            },
            Reason: shared.CreditNoteInputCreditNoteReasonEnumDuplicatedCharge,
            RefundAmountCents: 20,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreditNote != nil {
        // handle response
    }
}
```

## Download

Download an existing credit note

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
    res, err := s.CreditNotes.Download(ctx, operations.DownloadCreditNoteRequest{
        ID: "1a901a90-1a90-1a90-1a90-1a901a901a90",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreditNote != nil {
        // handle response
    }
}
```

## Find

Return a single credit note

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
    res, err := s.CreditNotes.Find(ctx, operations.FindCreditNoteRequest{
        ID: "12345",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreditNote != nil {
        // handle response
    }
}
```

## FindAll

Find all credit notes in certain organisation

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
    res, err := s.CreditNotes.FindAll(ctx, operations.FindAllCreditNotesRequest{
        ExternalCustomerID: lago.String("12345"),
        Page: lago.Int64(2),
        PerPage: lago.Int64(20),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreditNotes != nil {
        // handle response
    }
}
```

## Update

Update an existing credit note

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
    res, err := s.CreditNotes.Update(ctx, operations.UpdateCreditNoteRequest{
        CreditNoteUpdateInput: shared.CreditNoteUpdateInput{
            CreditNote: shared.CreditNoteUpdateInputCreditNote{
                RefundStatus: shared.CreditNoteUpdateInputCreditNoteRefundStatusEnumFailed,
            },
        },
        ID: "12345",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreditNote != nil {
        // handle response
    }
}
```

## Void

Void an existing credit note

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
    res, err := s.CreditNotes.Void(ctx, operations.VoidCreditNoteRequest{
        ID: "1a901a90-1a90-1a90-1a90-1a901a901a90",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreditNote != nil {
        // handle response
    }
}
```
