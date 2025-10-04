package primities

type ListingDetails struct {
	ListingID  string `json:"listingid"`
	CurrencyID uint16 `json:"currencyid"`

	Price uint64 `json:"price"`
	Fee   uint64 `json:"fee"`

	PublisherFeeApp     uint64 `json:"publisher_fee_app"`
	PublisherFeePercent string `json:"publisher_fee_percent"`

	Asset assetProto `json:"asset"`

	PublisherFee *uint64 `json:"publisher_fee"`
	SteamFee     *uint64 `json:"steam_fee"`

	ConvertedCurrencyId *uint16 `json:"converted_currencyid"`

	ConvertedPrice        *uint64 `json:"converted_price"`
	ConvertedFee          *uint64 `json:"converted_fee"`
	ConvertedSteamFee     *uint64 `json:"converted_steam_fee"`
	ConvertedPublisherFee *uint64 `json:"converted_publisher_fee"`

	ConvertedPricePerUnit        *uint64 `json:"converted_price_per_unit"`
	ConvertedFeePerUnit          *uint64 `json:"converted_fee_per_unit"`
	ConvertedSteamFeePerUnit     *uint64 `json:"converted_steam_fee_per_unit"`
	ConvertedPublisherFeePerUint *uint64 `json:"converted_publisher_fee_per_unit"`
}

func (l ListingDetails) TotalPrice() uint64 {
	totalPrice := l.Price + l.Fee

	if l.ConvertedPrice != nil && l.ConvertedFee != nil {
		totalPrice = *l.ConvertedPrice + *l.ConvertedFee
	}
	if l.ConvertedPricePerUnit != nil && l.ConvertedFeePerUnit != nil {
		totalPrice = *l.ConvertedPricePerUnit + *l.ConvertedFeePerUnit
	}
	return totalPrice
}

type assetProto struct {
	AppID         uint64          `json:"appid"`
	ContextID     string          `json:"contextid"`
	ID            string          `json:"id"`
	Amount        string          `json:"amount"`
	MarketActions []marketActions `json:"market_actions"`
}

type marketActions struct {
	Name string `json:"name"`
	Link string `json:"link"`
}
