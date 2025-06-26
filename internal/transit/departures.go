package transit

import (
	"context"
	"time"
)

type StopDeparturesRequest struct {
	GlobalStopID         string    `url:"global_stop_id"`
	Time                 time.Time `url:"time,unix,omitempty"`
	RemoveCancelled      bool      `url:"remove_cancelled,omitempty"`
	ShouldUpdateRealtime bool      `url:"should_update_realtime,omitempty"`
}

type StopDeparturesResponse struct {
	RouteDepartures []struct {
		Alerts                  []Alert `json:"alerts"`
		CompactDisplayShortName struct {
			BoxedText           string   `json:"boxed_text"`
			Elements            []string `json:"elements"`
			RouteNameRedundancy bool     `json:"route_name_redundancy"`
		} `json:"compact_display_short_name"`
		Fares                 []Fare       `json:"fares"`
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
		RouteImage       string `json:"route_image"`
		RouteLongName    string `json:"route_long_name"`
		RouteNetworkId   string `json:"route_network_id"`
		RouteNetworkName string `json:"route_network_name"`
		RouteShortName   string `json:"route_short_name"`
		RouteTextColor   string `json:"route_text_color"`
		RouteTimezone    string `json:"route_timezone"`
		RouteType        int    `json:"route_type"`
		SortingKey       string `json:"sorting_key"`
		TtsLongName      string `json:"tts_long_name"`
		TtsShortName     string `json:"tts_short_name"`
		Vehicle          struct {
			Image          string `json:"image"`
			Name           string `json:"name"`
			NameInflection string `json:"name_inflection"`
		} `json:"vehicle"`
	} `json:"route_departures"`
}

func (c *Client) GetStopDepartures(ctx context.Context, req StopDeparturesRequest) (StopDeparturesResponse, error) {
	var resp StopDeparturesResponse

	err := c.get(ctx, "/public/stop_departures", req, &resp)
	return resp, err
}
