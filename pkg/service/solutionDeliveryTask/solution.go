package solutiondeliverytask

import (
	cargodelivery "github.com/Makcumblch/CargoDelivery"
	routing "github.com/Makcumblch/CargoDelivery/pkg/service/solutionDeliveryTask/routing"
)

func getVolumeCar(car cargodelivery.Car) float32 {
	return *car.Height * *car.Length * *car.Width
}

func getVolumeCargo(car cargodelivery.Cargo) float32 {
	return *car.Height * *car.Length * *car.Width
}

func getInitSolution(taskData cargodelivery.DeliveryTaskData) (cargodelivery.RouteSolution, error) {
	indexCar := 0
	lenCars := len(taskData.Cars)
	if lenCars == 0 {
		return cargodelivery.RouteSolution{}, cargodelivery.ErrCreateRouteCars
	}

	idxClient := 0
	lenClient := len(taskData.Clients)
	if lenClient == 1 {
		return cargodelivery.RouteSolution{}, cargodelivery.ErrCreateRouteClient
	}

	routeSolution := cargodelivery.RouteSolution{Distance: 0, Fuel: 0, PackingCost: 0}

	carsRouteSolution := make([]cargodelivery.CarRoute, lenCars)

	for indexCar < lenCars {
		carsRouteSolution[indexCar] = cargodelivery.CarRoute{
			Car: taskData.Cars[indexCar],
			Route: cargodelivery.Route{
				Clients:  make([]cargodelivery.ClientRoute, 0),
				Polyline: make([][]float32, 0),
			},
			Items:            make([][]cargodelivery.Item, 0),
			FreeVolume:       getVolumeCar(taskData.Cars[indexCar]),
			FreeLoadCapacity: *taskData.Cars[indexCar].LoadCapacity,
		}
		carsRouteSolution[indexCar].Route.Clients = append(carsRouteSolution[indexCar].Route.Clients, cargodelivery.ClientRoute{
			Client: taskData.Clients[0].Client,
			Index:  0,
		})
		carsRouteSolution[indexCar].Items = append(carsRouteSolution[indexCar].Items, make([]cargodelivery.Item, 0))
		flagTS := false
		for idxClient < lenClient {
			clientOrders := taskData.Clients[idxClient]

			if len(clientOrders.Orders) == 0 {
				idxClient++
				continue
			}

			carsRouteSolution[indexCar].Route.Clients = append(carsRouteSolution[indexCar].Route.Clients, cargodelivery.ClientRoute{
				Client: clientOrders.Client,
				Index:  idxClient,
			})
			carsRouteSolution[indexCar].Items = append(carsRouteSolution[indexCar].Items, make([]cargodelivery.Item, 0))
			idxItems := len(carsRouteSolution[indexCar].Items) - 1

			idxOrder := 0
			for idxOrder < len(clientOrders.Orders) {
				order := clientOrders.Orders[idxOrder]

				idxCargo := 0
				for idxCargo < int(order.Count) {
					item := cargodelivery.Item{
						Client: clientOrders.Client,
						Cargo:  order.Cargo,
					}
					newLoad := carsRouteSolution[indexCar].FreeLoadCapacity - *item.Cargo.Weight
					newVolume := carsRouteSolution[indexCar].FreeVolume - getVolumeCargo(item.Cargo)
					if newLoad >= 0 && newVolume >= 0 {
						carsRouteSolution[indexCar].Items[idxItems] = append(carsRouteSolution[indexCar].Items[idxItems], item)
						carsRouteSolution[indexCar].FreeLoadCapacity = newLoad
						carsRouteSolution[indexCar].FreeVolume = newVolume
					} else {
						if indexCar+1 == lenCars {
							return cargodelivery.RouteSolution{}, cargodelivery.ErrCreateRouteFewCars
						}
						flagTS = true
					}
					if flagTS {
						break
					}
					idxCargo++
				}
				if flagTS {
					break
				}
				idxOrder++
			}
			if flagTS {
				break
			}
			idxClient++
		}
		carsRouteSolution[indexCar].Route.Clients = append(carsRouteSolution[indexCar].Route.Clients, cargodelivery.ClientRoute{
			Client: taskData.Clients[0].Client,
			Index:  0,
		})
		carsRouteSolution[indexCar].Items = append(carsRouteSolution[indexCar].Items, make([]cargodelivery.Item, 0))
		indexCar++
	}

	routeSolution.CarsRouteSolution = carsRouteSolution

	return routeSolution, nil
}

func GetDeliverySolution(taskData cargodelivery.DeliveryTaskData, distanceMatrix [][]float32, settingsRoute cargodelivery.RouteSettings) (cargodelivery.RouteSolution, error) {

	bestState, err := getInitSolution(taskData)
	if err != nil {
		return cargodelivery.RouteSolution{}, err
	}

	bestState = routing.RoutingProcedure(settingsRoute, distanceMatrix, bestState)

	return bestState, nil
}
