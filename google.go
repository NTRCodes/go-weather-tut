package main

import "fmt"

type latLng struct {
	Lat float64
}

func getLatLngForPlace(place string) (latLng latLng, err error) {
	u := fmt.Sprintf("https://maps.googleapis.com/maps/api/geocode/json?key=")
}