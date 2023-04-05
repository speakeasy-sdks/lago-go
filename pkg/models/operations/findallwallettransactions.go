// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/lago-go/pkg/models/shared"
	"net/http"
)

type FindAllWalletTransactionsRequest struct {
	// Lago ID of the existing wallet
	ID string `pathParam:"style=simple,explode=false,name=id"`
	// Number of page
	Page *int64 `queryParam:"style=form,explode=true,name=page"`
	// Number of records per page
	PerPage *int64 `queryParam:"style=form,explode=true,name=per_page"`
	// Status (pending or settled)
	Status *string `queryParam:"style=form,explode=true,name=status"`
	// Transaction Type (inbound or outbound)
	TransactionType *string `queryParam:"style=form,explode=true,name=transaction_type"`
}

type FindAllWalletTransactionsResponse struct {
	// Not Found error
	APIResponseNotFound *shared.APIResponseNotFound
	// Unauthorized error
	APIResponseUnauthorized *shared.APIResponseUnauthorized
	ContentType             string
	StatusCode              int
	RawResponse             *http.Response
	// Successful response
	WalletTransactions *shared.WalletTransactions
}