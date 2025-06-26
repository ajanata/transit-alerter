package main

import (
	"context"
	"fmt"
	"time"

	"github.com/ajanata/transit-alerter/internal/config"
	"github.com/ajanata/transit-alerter/internal/transit"

	"github.com/alexflint/go-arg"
	"github.com/fatih/color"
	"github.com/rodaine/table"
)

type args struct {
	NearbyRoutes      *NearbyRoutesCmd   `arg:"subcommand:nearby-routes" help:"Get Nearby Routes"`
	NearbyStops       *NearbyStopsCmd    `arg:"subcommand:nearby-stops" help:"Get Nearby Stops"`
	StopDeparturesCmd *StopDeparturesCmd `arg:"subcommand:stop-departures" help:"Get Departures for a stop"`
}

type NearbyRoutesCmd struct {
	Lat         float64 `arg:"positional,required" help:"Latitude, in decimal format"`
	Lon         float64 `arg:"positional,required" help:"Longitude, in decimal format"`
	MaxDistance int     `arg:"-m" help:"Meters"`
}

type NearbyStopsCmd struct {
	Lat         float64 `arg:"positional,required" help:"Latitude, in decimal format"`
	Lon         float64 `arg:"positional,required" help:"Longitude, in decimal format"`
	MaxDistance int     `arg:"-m" help:"Meters"`
}

type StopDeparturesCmd struct {
	GlobalStopID    string    `arg:"positional,required" help:"Global stop ID"`
	RemoveCancelled bool      `arg:"-r" help:"Remove cancelled trips"`
	Time            time.Time `arg:"-t" help:"Time of departure"`
}

var c *transit.Client

var ctx = context.Background()

func main() {
	cfg, err := config.Load()
	if err != nil {
		panic(err)
	}

	a := args{}
	arg.MustParse(&a)

	c = transit.New(cfg.Transit.APIKey)

	switch {
	case a.NearbyRoutes != nil:
		nearbyRoutes(*a.NearbyRoutes)
	case a.NearbyStops != nil:
		nearbyStops(*a.NearbyStops)
	case a.StopDeparturesCmd != nil:
		stopDepartures(*a.StopDeparturesCmd)
	default:
		fmt.Println("no subcommand")
	}
}

func newTable(columnHeaders ...any) table.Table {
	headerFmt := color.New(color.FgGreen, color.Underline).SprintfFunc()
	columnFmt := color.New(color.FgYellow).SprintfFunc()

	tbl := table.New(columnHeaders...)
	tbl.WithHeaderFormatter(headerFmt).WithFirstColumnFormatter(columnFmt)

	return tbl
}

func nearbyRoutes(n NearbyRoutesCmd) {
	nearby := transit.GetNearbyRoutesRequest{
		Lat:         n.Lat,
		Lon:         n.Lon,
		MaxDistance: n.MaxDistance,
		RealTime:    true,
	}

	resp, err := c.GetNearbyRoutes(ctx, nearby)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", resp)
}

func nearbyStops(n NearbyStopsCmd) {
	nearby := transit.GetNearbyStopsRequest{
		Lat:         n.Lat,
		Lon:         n.Lon,
		MaxDistance: n.MaxDistance,
	}

	resp, err := c.GetNearbyStops(ctx, nearby)
	if err != nil {
		panic(err)
	}

	tbl := newTable("Global ID", "Parent", "Name", "Type", "Distance")

	for _, stop := range resp.Stops {
		parent := ""
		if stop.ParentStationGlobalStopId != nil {
			parent = *stop.ParentStationGlobalStopId
		}
		tbl.AddRow(stop.GlobalStopId, parent, stop.StopName, stop.RouteType, stop.Distance)
	}

	tbl.Print()
}

func stopDepartures(s StopDeparturesCmd) {
	d := transit.StopDeparturesRequest{
		GlobalStopID:    s.GlobalStopID,
		RemoveCancelled: s.RemoveCancelled,
	}
	if !s.Time.IsZero() {
		d.Time = s.Time
	}

	resp, err := c.GetStopDepartures(ctx, d)
	if err != nil {
		panic(err)
	}

	alerts := newTable("Severity", "Created", "Effect", "Severity", "Title", "Description")
	tbl := newTable("Trip ID", "Global Route ID", "Branch", "Headsign", "Scheduled", "Actual")

	for _, route := range resp.RouteDepartures {
		for _, a := range route.Alerts {
			alerts.AddRow(a.Severity, a.CreatedAt, a.Effect, a.Severity, a.Title, a.Description)
		}
		for _, it := range route.Itineraries {
			for _, si := range it.ScheduleItems {
				tbl.AddRow(si.RtTripId, route.GlobalRouteId, it.BranchCode, it.MergedHeadsign, si.ScheduledDepartureTime, si.DepartureTime)
			}
		}
	}

	alerts.Print()
	tbl.Print()
}
