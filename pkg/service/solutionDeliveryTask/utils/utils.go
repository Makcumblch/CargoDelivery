package utils

import cargodelivery "github.com/Makcumblch/CargoDelivery"

func CloneSolution(solution cargodelivery.RouteSolution) cargodelivery.RouteSolution {
	newSolution := cargodelivery.RouteSolution{RoutingCost: solution.RoutingCost, PackingCost: solution.PackingCost}
	newCarsRouteSolution := make([]cargodelivery.CarRoute, 0)
	for _, route := range solution.CarsRouteSolution {
		newCarRoute := cargodelivery.CarRoute{Car: route.Car, FreeVolume: route.FreeVolume, FreeLoadCapacity: route.FreeLoadCapacity}

		newRoute := cargodelivery.Route{Clients: make([]cargodelivery.ClientRoute, 0), Waypoints: make([][]float32, 0)}
		for _, client := range route.Route.Clients {
			newClientRoute := cargodelivery.ClientRoute{Client: client.Client, Index: client.Index}
			newRoute.Clients = append(newRoute.Clients, newClientRoute)
		}
		newCarRoute.Route = newRoute

		newItems := make([][]cargodelivery.Item, 0)
		for _, clientItems := range route.Items {
			newItemsClient := make([]cargodelivery.Item, 0)
			for _, items := range clientItems {
				newItemsClient = append(newItemsClient, items)
			}
			newItems = append(newItems, newItemsClient)
		}
		newCarRoute.Items = newItems

		newCarsRouteSolution = append(newCarsRouteSolution, newCarRoute)
	}

	newSolution.CarsRouteSolution = newCarsRouteSolution

	return newSolution
}
