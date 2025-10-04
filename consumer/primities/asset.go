package primities

type AssetDetails struct {
	Currency uint16 `json:"currency"`

	AppID      uint64 `json:"appid"`
	ContextID  string `json:"contextid"`
	ID         string `json:"id"`
	ClassID    string `json:"classid"`
	InstanceID string `json:"instanceid"`

	Amount         string `json:"amount"`
	Status         int    `json:"status"`
	OriginalAmount string `json:"original_amount"`

	UnownedID        string `json:"unowned_id"`
	UnownedContextID string `json:"unowned_contextid"`

	BackgroupColor string         `json:"background_color"`
	IconUrl        string         `json:"icon_url"`
	Descriptions   []descriptions `json:"descriptions"`
	Tradable       int            `json:"tradable"`

	Name           string `json:"name"`
	NameColor      string `json:"name_color"`
	Type           string `json:"type"`
	MarketName     string `json:"market_name"`
	MarketHashName string `json:"market_hash_name"`

	Commodity                   int `json:"commodity"`
	MarketTradableRestriction   int `json:"market_tradable_restriction"`
	MarketMarketableRestriction int `json:"market_marketable_restriction"`

	AppIcon string `json:"app_icon"`
	Owner   int    `json:"owner"`

	IconUrlLarge  *string `json:"icon_url_large"`
	Actions       *[]marketActions
	MarketActions *[]marketActions
}

type descriptions struct {
	Name  string  `json:"name"`
	Color string  `json:"color"`
	Value string  `json:"value"`
	Type  *string `json:"type"`
}
