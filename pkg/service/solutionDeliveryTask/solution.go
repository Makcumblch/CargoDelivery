package solutiondeliverytask

import (
	cargodelivery "github.com/Makcumblch/CargoDelivery"
	packing "github.com/Makcumblch/CargoDelivery/pkg/service/solutionDeliveryTask/packing"
	routing "github.com/Makcumblch/CargoDelivery/pkg/service/solutionDeliveryTask/routing"
)

func getInitSolution(taskData cargodelivery.DeliveryTaskData) (cargodelivery.RouteSolution, error) {
	return cargodelivery.RouteSolution{}, nil
}

func getSolutionCost(solution cargodelivery.RouteSolution) float32 {
	return 0
}

func GetDeliverySolution(taskData cargodelivery.DeliveryTaskData, distanceMatrix [][]float32, settingsRoute cargodelivery.RouteSettings) (cargodelivery.RouteSolution, error) {

	bestState, err := getInitSolution(taskData)
	if err != nil {
		return cargodelivery.RouteSolution{}, err
	}

	bestStateCost := getSolutionCost(bestState)

	for t := uint(0); t < settingsRoute.Count; t++ {
		stateNew := routing.RoutingProcedure(settingsRoute.TMax, settingsRoute.TMin, bestState)
		stateNew = packing.PackingProcedure(settingsRoute.EvCount, stateNew)
		stateNewCost := getSolutionCost(stateNew)
		if stateNewCost < bestStateCost {
			bestState = stateNew
			bestStateCost = stateNewCost
		}
	}

	return bestState, nil
}
