package transit

import (
	"context"
)

type GetNearbyRoutesRequest struct {
	Lat         float64 `url:"lat"`
	Lon         float64 `url:"lon"`
	MaxDistance int     `url:"max_distance,omitempty"`
	RealTime    bool    `url:"should_update_realtime,omitempty"`
}

type GetNearbyRoutesResponse struct {
	Routes []struct {
		CompactDisplayShortName struct {
			BoxedText           string   `json:"boxed_text"`
			Elements            []string `json:"elements"`
			RouteNameRedundancy bool     `json:"route_name_redundancy"`
		} `json:"compact_display_short_name"`
		Fares []struct {
			FareMediaType int `json:"fare_media_type"`
			PriceMin      struct {
				CurrencyCode string  `json:"currency_code"`
				Symbol       string  `json:"symbol"`
				Text         string  `json:"text"`
				Value        float64 `json:"value"`
			} `json:"price_min"`
			PriceMax struct {
				CurrencyCode string  `json:"currency_code"`
				Symbol       string  `json:"symbol"`
				Text         string  `json:"text"`
				Value        float64 `json:"value"`
			} `json:"price_max"`
		} `json:"fares"`
		GlobalRouteId         string       `json:"global_route_id"`
		Itineraries           []Intinerary `json:"itineraries"`
		ModeName              string       `json:"mode_name"`
		RealTimeRouteId       string       `json:"real_time_route_id"`
		RouteColor            string       `json:"route_color"`
		RouteDisplayShortName struct {
			BoxedText           string   `json:"boxed_text"`
			Elements            []string `json:"elements"`
			RouteNameRedundancy bool     `json:"route_name_redundancy"`
		} `json:"route_display_short_name"`
		RouteLongName  string `json:"route_long_name"`
		RouteShortName string `json:"route_short_name"`
		RouteTextColor string `json:"route_text_color"`
		RouteType      int    `json:"route_type"`
		SortingKey     string `json:"sorting_key"`
		TtsLongName    string `json:"tts_long_name"`
		TtsShortName   string `json:"tts_short_name"`
	} `json:"routes"`
}

func (c *Client) GetNearbyRoutes(ctx context.Context, req GetNearbyRoutesRequest) (GetNearbyRoutesResponse, error) {
	var resp GetNearbyRoutesResponse

	err := c.get(ctx, "/public/nearby_routes", req, &resp)
	return resp, err
}
