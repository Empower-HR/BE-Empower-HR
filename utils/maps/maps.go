package maps

import (
	"fmt"
	"math"
	"os"

	"github.com/joho/godotenv"
	geo "github.com/martinlindhe/google-geolocate"
)

type mapsUtilityInterface interface {
	GeoCode(address string) (*geo.Point, error)
	GeoLocate() (float64, float64, error)
	haversine(lat1, lon1, lat2, lon2 float64) (float64, error)
}

type mapsUtility struct {
	client *geo.GoogleGeo
}

func NewMapsUtility() (mapsUtilityInterface, error) {
	if err := godotenv.Load(".env"); err != nil {
		return nil, fmt.Errorf("error loading .env file: %v", err)
	}

	keyApi := os.Getenv("API_KEYS")
	if keyApi == "" {
		return nil, fmt.Errorf("API_KEYS environment variable is not set")
	}

	client := geo.NewGoogleGeo(keyApi)
	return &mapsUtility{client: client}, nil
}

func (mu *mapsUtility) GeoCode(address string) (*geo.Point, error) {
	geoCodeRes, err := mu.client.Geocode(address)
	if err != nil {
		return nil, fmt.Errorf("error geocoding address: %v", err)
	}

	fmt.Println("geoCode", geoCodeRes)
	return geoCodeRes, nil
}

func (mu *mapsUtility) GeoLocate() (float64, float64, error) {
	geolocateRes, err := mu.client.Geolocate()
	if err != nil {
		return 0, 0, fmt.Errorf("error locating geolocation: %v", err)
	}

	fmt.Println("geoLocate", geolocateRes)
	return geolocateRes.Lat, geolocateRes.Lng, nil
}

func (mu *mapsUtility) haversine(lat1, lon1, lat2, lon2 float64) (float64, error) {
	const earthRadius = 6371000
	lat1Rad := lat1 * math.Pi / 180
	lon1Rad := lon1 * math.Pi / 180
	lat2Rad := lat2 * math.Pi / 180
	lon2Rad := lon2 * math.Pi / 180

	dLat := lat2Rad - lat1Rad
	dLon := lon2Rad - lon1Rad
	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1Rad)*math.Cos(lat2Rad)*
			math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return earthRadius * c, nil
}
