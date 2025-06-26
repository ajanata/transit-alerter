package transit

type RouteType int

const (
	TramRoute RouteType = iota
	SubwayRoute
	RailRoute
	BusRoute
	FerryRoute
	CableCarRoute
	GondolaRoute
	FunicularRoute
)
const (
	TrolleybusRoute = iota + 11
	MonorailRoute
)

type Intinerary struct {
	BranchCode        string          `json:"branch_code"`
	ClosestStop       Stop            `json:"closest_stop"`
	DirectionHeadsign string          `json:"direction_headsign"`
	DirectionId       int             `json:"direction_id"`
	Headsign          string          `json:"headsign"`
	MergedHeadsign    string          `json:"merged_headsign"`
	ScheduleItems     []ScheduleItems `json:"schedule_items"`
}

type Stop struct {
	GlobalStopId              string    `json:"global_stop_id"`
	LocationType              int       `json:"location_type"`
	ParentStationId           *Stop     `json:"parent_station_id"`
	ParentStationGlobalStopId *string   `json:"parent_station_global_stop_id"`
	RouteType                 RouteType `json:"route_type"`
	RtStopId                  string    `json:"rt_stop_id"`
	StopCode                  string    `json:"stop_code"`
	StopLat                   float64   `json:"stop_lat"`
	StopLon                   float64   `json:"stop_lon"`
	StopName                  string    `json:"stop_name"`
	WheelchairBoarding        int       `json:"wheelchair_boarding"`
}

type ScheduleItems struct {
	DepartureTime          UnixTime `json:"departure_time"`
	IsCancelled            bool     `json:"is_cancelled"`
	IsRealTime             bool     `json:"is_real_time,omitempty"`
	RtTripId               string   `json:"rt_trip_id"`
	ScheduledDepartureTime UnixTime `json:"scheduled_departure_time,omitempty"`
	TripSearchKey          string   `json:"trip_search_key"`
	WheelchairAccessible   int      `json:"wheelchair_accessible"`
}

func (rt RouteType) String() string {
	switch rt {
	case TramRoute:
		return "Tram"
	case SubwayRoute:
		return "Subway"
	case RailRoute:
		return "Rail"
	case BusRoute:
		return "Bus"
	case FerryRoute:
		return "Ferry"
	case CableCarRoute:
		return "CableCar"
	case GondolaRoute:
		return "Gondola"
	case FunicularRoute:
		return "Funicular"
	case TrolleybusRoute:
		return "Trolleybus"
	case MonorailRoute:
		return "Monorail"
	default:
		return "Unknown"
	}
}
