package transit

import (
	"encoding/json"
	"time"
)

type Alert struct {
	CreatedAt   UnixTime `json:"created_at"`
	Description string   `json:"description"`
	Effect      string   `json:"effect"`
	Severity    string   `json:"severity"`
	Title       string   `json:"title"`
}

type Price struct {
	CurrencyCode string  `json:"currency_code"`
	Symbol       string  `json:"symbol"`
	Text         string  `json:"text"`
	Value        float64 `json:"value"`
}

type Fare struct {
	FareMediaType int   `json:"fare_media_type"`
	PriceMax      Price `json:"price_max"`
	PriceMin      Price `json:"price_min"`
}

type UnixTime struct {
	time.Time
}

func (u *UnixTime) UnmarshalJSON(b []byte) error {
	var timestamp int64
	err := json.Unmarshal(b, &timestamp)
	if err != nil {
		return err
	}
	u.Time = time.Unix(timestamp, 0)
	return nil
}
