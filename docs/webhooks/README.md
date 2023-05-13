# Webhooks

## Overview

Everything about Webhooks

Find out more
<https://doc.getlago.com/docs/api/webhooks/message-signature#1-retrieve-the-public-key>
### Available Operations

* [FetchPublicKey](#fetchpublickey) - Fetch webhook public key

## FetchPublicKey

Webhook public key

### Example Usage

```go
package main

import(
	"context"
	"log"
	"github.com/speakeasy-sdks/lago-go"
)

func main() {
    s := lago.New(
        lago.WithSecurity(shared.Security{
            BearerAuth: "YOUR_BEARER_TOKEN_HERE",
        }),
    )

    ctx := context.Background()
    res, err := s.Webhooks.FetchPublicKey(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.FetchPublicKey200TextPlainString != nil {
        // handle response
    }
}
```
