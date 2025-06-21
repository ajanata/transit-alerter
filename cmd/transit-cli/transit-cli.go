package main

import (
	"context"
	"fmt"

	"github.com/ajanata/transit-alerter/internal/config"
	"github.com/ajanata/transit-alerter/internal/transit"

	"github.com/alexflint/go-arg"
)

type args struct {
	NearbyRoutes *NearbyRoutesCmd `arg:"subcommand:nearby-routes"`
}

type NearbyRoutesCmd struct {
	Lat         float64 `arg:"positional,required"`
	Lon         float64 `arg:"positional,required"`
	MaxDistance int     `arg:"-m"`
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
		nearby(*a.NearbyRoutes)
	default:
		fmt.Println("no subcommand")
	}
}

func nearby(n NearbyRoutesCmd) {
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
