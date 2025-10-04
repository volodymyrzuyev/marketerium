package consumer

import "errors"

type Marketerium interface {
	// Fetch listings for market hash name
	// steamcommunity.com/market/listings/{appID}/{marketHashName}/render?currency={?}&country={?}&start={start}&count={count}&language={?}
	GetListings(params GetListingsRequest) (GetListingsResponse, error)
}

var (
	ErrNotSuccessfulReply = errors.New("steam identified the response as not successful")
	ErrSteamRateLimit     = errors.New("steam rate limited the request")
	ErrUnexpectedFormat   = errors.New("unexpected response format")
)
