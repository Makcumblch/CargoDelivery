package repository

import (
	"encoding/json"
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
	json.NewDecoder(resp.Body).Decode(&result)

	return result.Distances, nil
}

func (o *OSMRepo) GetRoutePoints(clients []cargodelivery.Client) [][]float32 {

	return make([][]float32, 0)
}
