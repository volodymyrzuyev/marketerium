package consumer

import (
	"time"

	"github.com/volodymyrzuyev/marketerium/consumer/primities"
)

type GetListingsRequest struct {
	MarketHashname string
	AppID          uint
	// first index to get
	Start uint
	// number of listings to get
	Count uint
	// if not current data is needed, set the time offset that would be acceptable, based on current time
	// if it's nil, new data will be forced to be fetched from steam
	TimeOffset *time.Duration
}

type GetListingsResponse struct {
	Success     bool
	Start       uint64
	Pagesize    uint64
	TotalCount  uint64
	ResultHTML  string
	ListingInfo map[string]primities.ListingDetails
	Assets      map[string]primities.AssetDetails
	Currency    any
	Hovers      any
	AppData     any
}
