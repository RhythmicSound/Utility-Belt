package geo

import (
	"fmt"
	"math"
)

//LonLatDistance calculates the distance between two points of longitude and latitude
// using Haversine formula: http://en.wikipedia.org/wiki/Haversine_formula
//   Definitions:
//     South latitudes are negative, east longitudes are positive
//
//   Passed to function:
//     lat1, lon1 = Latitude and Longitude of point 1 (in decimal degrees)
//     lat2, lon2 = Latitude and Longitude of point 2 (in decimal degrees)
//     unit = the unit you desire for results
//            where: 'M' is statute miles (default)
//                   'K' is kilometers
//                   'N' is nautical miles
//
//Based on: https://www.geodatasource.com/developers/go
func LonLatDistance(lat1 float64, lng1 float64, lat2 float64, lng2 float64, unit string) (float64, error) {
	const PI float64 = 3.141592653589793

	radlat1 := float64(PI * lat1 / 180)
	radlat2 := float64(PI * lat2 / 180)

	theta := float64(lng1 - lng2)
	radtheta := float64(PI * theta / 180)

	dist := math.Sin(radlat1)*math.Sin(radlat2) + math.Cos(radlat1)*math.Cos(radlat2)*math.Cos(radtheta)

	if dist > 1 {
		dist = 1
	}

	dist = math.Acos(dist)
	dist = dist * 180 / PI
	dist = dist * 60 * 1.1515

	switch unit {
	case "K":
		dist = dist * 1.609344
	case "N":
		dist = dist * 0.8684
	case "M":
		//already in miles
	default:
		return 0, fmt.Errorf("Unit '%s' not recognised in LonLatDistance func", unit)

	}

	return dist, nil
}
