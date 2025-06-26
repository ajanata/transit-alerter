package transit

import (
	"context"
)

type GetNearbyStopsRequest struct {
	Lat         float64 `url:"lat"`
	Lon         float64 `url:"lon"`
	MaxDistance int     `url:"max_distance,omitempty"`

	// Defaults to RoutableStopFilter.
	StopFilter          StopFilter          `url:"stop_filter,omitempty"`
	PickupDropoffFilter PickupDropoffFilter `url:"pickup_dropoff_filter,omitempty"`
}

type StopFilter string

const (
	RoutableStopFilter                         StopFilter = "Routable"
	EntrancesAndStopsOutsideStationsStopFilter StopFilter = "EntrancesAndStopsOutsideStations"
	EntrancesStopFilter                        StopFilter = "Entrances"
	AnyStopFilter                              StopFilter = "Any"
)

type PickupDropoffFilter string

const (
	PickupAllowedOnlyFilter  PickupDropoffFilter = "PickupAllowedOnly"
	DropoffAllowedOnlyFilter PickupDropoffFilter = "DropoffAllowedOnly"
	EverythingFilter         PickupDropoffFilter = "Everything"
)

type GetNearbyStopsResponse struct {
	Stops []struct {
		Distance int `json:"distance"`
		Stop
	} `json:"stops"`
}

func (c *Client) GetNearbyStops(ctx context.Context, req GetNearbyStopsRequest) (GetNearbyStopsResponse, error) {
	var resp GetNearbyStopsResponse

	err := c.get(ctx, "/public/nearby_stops", req, &resp)
	return resp, err
}
