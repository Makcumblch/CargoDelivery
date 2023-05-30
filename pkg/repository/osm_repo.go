package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/spf13/viper"
)

type OSMRepo struct{}

func NewOSMRepo() *OSMRepo {
	return &OSMRepo{}
}

func (o *OSMRepo) GetDistanceMatrix(clients []cargodelivery.Client) ([][]float32, error) {
	coords := make([]string, 0)
	for _, client := range clients {
		coords = append(coords, fmt.Sprintf("%f,%f", client.CoordX, client.CoordY))
	}

	resp, err := http.Get(fmt.Sprintf("http://%s:%s/table/v1/driving/%s?annotations=distance", viper.GetString("osmr.addres"), viper.GetString("osmr.port"), strings.Join(coords, ";")))
	if err != nil {
		return make([][]float32, 0), err
	}

	var result cargodelivery.OSMRTableResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return make([][]float32, 0), err
	}

	return result.Distances, nil
}

func (o *OSMRepo) GetRoutePoints(clients []cargodelivery.Client) ([][]float32, error) {
	coords := make([]string, 0)
	for _, client := range clients {
		coords = append(coords, fmt.Sprintf("%f,%f", client.CoordX, client.CoordY))
	}

	resp, err := http.Get(fmt.Sprintf("http://%s:%s/route/v1/driving/%s?geometries=geojson&overview=full", viper.GetString("osmr.addres"), viper.GetString("osmr.port"), strings.Join(coords, ";")))
	if err != nil {
		return make([][]float32, 0), err
	}

	var route cargodelivery.OSMRRoteResponse
	err = json.NewDecoder(resp.Body).Decode(&route)
	if err != nil {
		return make([][]float32, 0), err
	}

	if len(route.Routes) == 0 {
		return make([][]float32, 0), errors.New("не найдены точки маршрута")
	}

	var newCoords = make([][]float32, 0)
	for _, c := range route.Routes[0].Geometry.Coordinates {
		coord := make([]float32, 0)
		coord = append(coord, c[1], c[0])
		newCoords = append(newCoords, coord)
	}

	return newCoords, nil
}
