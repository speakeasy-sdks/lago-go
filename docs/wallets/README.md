# Wallets

## Overview

Everything about Wallet collection

Find out more
<https://doc.getlago.com/docs/api/wallets/wallet-object>
### Available Operations

* [Create](#create) - Create a new wallet
* [CreateTransaction](#createtransaction) - Create a new wallet transaction
* [Destroy](#destroy) - Delete a wallet
* [Find](#find) - Find wallet
* [FindAll](#findall) - Find wallets
* [FindAllTransactions](#findalltransactions) - Find wallet transactions
* [Update](#update) - Update an existing wallet

## Create

Create a new wallet

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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Wallets.Create(ctx, shared.WalletInput{
        Wallet: &shared.WalletInputWallet{
            Currency: "EUR",
            ExpirationAt: types.MustTimeFromString("2022-09-14T23:59:59Z"),
            ExternalCustomerID: "12345",
            GrantedCredits: lago.Float64(10),
            Name: lago.String("Wallet name"),
            PaidCredits: lago.Float64(500),
            RateAmount: 2,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Wallet != nil {
        // handle response
    }
}
```

## CreateTransaction

Create a new wallet transaction

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
    res, err := s.Wallets.CreateTransaction(ctx, shared.WalletTransactionInput{
        WalletTransaction: shared.WalletTransactionInputWalletTransaction{
            GrantedCredits: lago.Float64(10),
            PaidCredits: lago.Float64(100),
            WalletID: "1a901a90-1a90-1a90-1a90-1a901a901a90",
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WalletTransactions != nil {
        // handle response
    }
}
```

## Destroy

Delete a wallet

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
    res, err := s.Wallets.Destroy(ctx, operations.DestroyWalletRequest{
        ID: "1a901a90-1a90-1a90-1a90-1a901a901a90",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Wallet != nil {
        // handle response
    }
}
```

## Find

Return a wallet

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
    res, err := s.Wallets.Find(ctx, operations.FindWalletRequest{
        ID: "1a901a90-1a90-1a90-1a90-1a901a901a90",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Wallet != nil {
        // handle response
    }
}
```

## FindAll

Find all wallets for certain customer

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
    res, err := s.Wallets.FindAll(ctx, operations.FindAllWalletsRequest{
        ExternalCustomerID: "12345",
        Page: lago.Int64(2),
        PerPage: lago.Int64(20),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Wallets != nil {
        // handle response
    }
}
```

## FindAllTransactions

Find all wallet transactions for certain wallet

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
    res, err := s.Wallets.FindAllTransactions(ctx, operations.FindAllWalletTransactionsRequest{
        ID: "1a901a90-1a90-1a90-1a90-1a901a901a90",
        Page: lago.Int64(2),
        PerPage: lago.Int64(20),
        Status: lago.String("pending"),
        TransactionType: lago.String("inbound"),
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.WalletTransactions != nil {
        // handle response
    }
}
```

## Update

Update an existing wallet

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
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Wallets.Update(ctx, operations.UpdateWalletRequest{
        WalletUpdateInput: shared.WalletUpdateInput{
            Wallet: shared.WalletUpdateInputWallet{
                ExpirationAt: types.MustTimeFromString("2022-09-14T23:59:59Z"),
                Name: lago.String("Wallet name"),
            },
        },
        ID: "1a901a90-1a90-1a90-1a90-1a901a901a90",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Wallet != nil {
        // handle response
    }
}
```
