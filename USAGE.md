<!-- Start SDK Example Usage -->
```go
package main

import (
    "context"
    "log"
    "github.com/speakeasy-sdks/lago-go"
    "github.com/speakeasy-sdks/lago-go/pkg/models/shared"
    "github.com/speakeasy-sdks/lago-go/pkg/models/operations"
)

func main() {
    s := lago.New(
        lago.WithSecurity(shared.Security{
            BearerAuth: "Bearer YOUR_BEARER_TOKEN_HERE",
        }),
    )

    req := operations.FindInvoiceRequest{
        ID: "1a901a90-1a90-1a90-1a90-1a901a901a90",
    }

    ctx := context.Background()
    res, err := s.Invoices.Find(ctx, req)
    if err != nil {
        log.Fatal(err)
    }

    if res.Invoice != nil {
        // handle response
    }
}
```
<!-- End SDK Example Usage -->