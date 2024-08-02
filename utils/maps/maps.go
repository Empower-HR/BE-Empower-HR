package maps

import (
	"be-empower-hr/app/config"
	"fmt"
	"math"

	geo "github.com/martinlindhe/google-geolocate"
)

type MapsUtilityInterface interface {
	GeoCode(address string) (float64, float64, error)
	Geolocate() (float64, float64, error)
	Haversine(lat1, lon1, lat2, lon2 float64) float64
}

type MapsUtility struct{}

func NewMapsUtility() MapsUtilityInterface {
	return &MapsUtility{}
}

func (mu *MapsUtility) GeoCode(address string) (float64, float64, error) {
	keyApi := config.API_KEYS
	fmt.Println("api key", keyApi)
	client := geo.NewGoogleGeo(keyApi)
	geoCodeRes, _ := client.Geocode(address)
	fmt.Println("GEOCODENYA BROOO", geoCodeRes)
	
	return geoCodeRes.Lat,geoCodeRes.Lng, nil
}

func (mu *MapsUtility) Geolocate() (float64, float64, error) {
	keyApi := config.API_KEYS
	client := geo.NewGoogleGeo(keyApi)
	geolocateRes, _ := client.Geolocate()
	return geolocateRes.Lat,geolocateRes.Lng, nil
}

func (mu *MapsUtility) Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const earthRadius = 6371000
	// Mengubah derajat ke radian
	lat1Rad := lat1 * math.Pi / 180
	lon1Rad := lon1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	lon2Rad := lon2 * math.Pi / 180

	// Perhitungan Haversine
	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c
}