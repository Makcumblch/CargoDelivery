package routing

import (
	"errors"
	"math"
	"math/rand"
	"time"

	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/Makcumblch/CargoDelivery/pkg/service/solutionDeliveryTask/packing"
	"github.com/Makcumblch/CargoDelivery/pkg/service/solutionDeliveryTask/utils"
)

func getSolutionCost(distanceMatrix [][]float32, solution *cargodelivery.RouteSolution) float32 {
	var distance float32 = 0
	var fuel float32 = 0
	for _, car := range solution.CarsRouteSolution {
		var dist float32 = 0
		for r := 0; r < len(car.Route.Clients)-1; r++ {
			dist += distanceMatrix[car.Route.Clients[r].Index][car.Route.Clients[r+1].Index] / 1000
		}
		distance += dist
		fuel += dist / 100 * *car.FuelConsumption
	}
	solution.Distance = distance
	solution.Fuel = fuel
	return fuel + solution.PackingCost
}

func getVolumeCargo(cargo cargodelivery.Cargo) float32 {
	return *cargo.Height * *cargo.Length * *cargo.Height
}

func getIndexClient(array []cargodelivery.ClientRoute, id int) int {
	for i, v := range array {
		if v.Index == id {
			return i
		}
	}
	return -1
}

func transfer(car1Index, indexClient1, car2Index, indexClient2 int, flagExistence bool, solution *cargodelivery.RouteSolution) {
	car1 := &solution.CarsRouteSolution[car1Index]
	car2 := &solution.CarsRouteSolution[car2Index]

	flag := true

	for i := len(car1.Items[indexClient1]) - 1; i >= 0; i-- {
		car1.FreeLoadCapacity += *car1.Items[indexClient1][i].Cargo.Weight
		car1.FreeVolume += getVolumeCargo(car1.Items[indexClient1][i].Cargo)
		newLoad := car2.FreeLoadCapacity - *car1.Items[indexClient1][i].Cargo.Weight
		newVolume := car2.FreeVolume - getVolumeCargo(car1.Items[indexClient1][i].Cargo)

		if newLoad >= 0 && newVolume >= 0 {
			if flag && !flagExistence {
				car2.Route.Clients = append(car2.Route.Clients[:indexClient2+1], car2.Route.Clients[indexClient2:]...)
				car2.Route.Clients[indexClient2] = car1.Route.Clients[indexClient1]

				car2.Items = append(car2.Items[:indexClient2+1], car2.Items[indexClient2:]...)
				car2.Items[indexClient2] = make([]cargodelivery.Item, 0)
			}
			car2.Items[indexClient2] = append(car2.Items[indexClient2], car1.Items[indexClient1][i])
			car2.FreeLoadCapacity = newLoad
			car2.FreeVolume = newVolume
			car1.Items[indexClient1] = append(car1.Items[indexClient1][:i], car1.Items[indexClient1][i+1:]...)
			flag = false
		} else {
			car1.FreeLoadCapacity -= *car1.Items[indexClient1][i].Cargo.Weight
			car1.FreeVolume -= getVolumeCargo(car1.Items[indexClient1][i].Cargo)
			break
		}
	}

	if len(car1.Items[indexClient1]) == 0 {
		car1.Items = append(car1.Items[:indexClient1], car1.Items[indexClient1+1:]...)
		car1.Route.Clients = append(car1.Route.Clients[:indexClient1], car1.Route.Clients[indexClient1+1:]...)
	}

}

func customerExchange(solution cargodelivery.RouteSolution) cargodelivery.RouteSolution {
	lenCars := len(solution.CarsRouteSolution)

	car1Index := rand.Intn(lenCars)
	car1 := &solution.CarsRouteSolution[car1Index]
	if len(car1.Route.Clients) < 3 {
		return solution
	}

	car2Index := rand.Intn(lenCars)
	car2 := &solution.CarsRouteSolution[car2Index]
	if len(car2.Route.Clients) < 3 {
		return solution
	}

	lenClients1 := len(car1.Route.Clients)
	indexClient1 := utils.GetRandInt(lenClients1-1, 1)

	lenClients2 := len(car2.Route.Clients)
	indexClient2 := utils.GetRandInt(lenClients2-1, 1)

	if car1.Route.Clients[indexClient1].Index == car2.Route.Clients[indexClient2].Index {
		return solution
	}

	flagExistenceCl2 := false
	idxClient2 := getIndexClient(car2.Route.Clients, car1.Route.Clients[indexClient1].Index)
	if idxClient2 == -1 {
		idxClient2 = indexClient2
	} else {
		flagExistenceCl2 = true
	}

	transfer(car1Index, indexClient1, car2Index, idxClient2, flagExistenceCl2, &solution)

	flagExistenceCl1 := false
	idxClient1 := getIndexClient(car1.Route.Clients, car2.Route.Clients[indexClient2].Index)
	if idxClient1 == -1 {
		idxClient1 = indexClient1
	} else {
		flagExistenceCl1 = true
	}

	transfer(car2Index, indexClient2, car1Index, idxClient1, flagExistenceCl1, &solution)

	return solution
}

func customerTransfer(solution cargodelivery.RouteSolution) cargodelivery.RouteSolution {
	lenCars := len(solution.CarsRouteSolution)

	car1Index := rand.Intn(lenCars)
	car1 := &solution.CarsRouteSolution[car1Index]
	for len(car1.Route.Clients) < 3 {
		car1Index = rand.Intn(lenCars)
		car1 = &solution.CarsRouteSolution[car1Index]
	}

	car2Index := rand.Intn(lenCars)
	car2 := &solution.CarsRouteSolution[car2Index]

	lenClients1 := len(car1.Route.Clients)
	indexClient1 := utils.GetRandInt(lenClients1-1, 1)

	flagExistence := false
	indexClient2 := getIndexClient(car2.Route.Clients, car1.Route.Clients[indexClient1].Index)
	if indexClient2 == -1 {
		lenClients2 := len(car2.Route.Clients)
		indexClient2 = utils.GetRandInt(lenClients2-1, 1)
	} else {
		flagExistence = true
	}

	transfer(car1Index, indexClient1, car2Index, indexClient2, flagExistence, &solution)

	return solution
}

func getNewSolution(sol cargodelivery.RouteSolution, probability float32) cargodelivery.RouteSolution {

	newSolution := utils.CloneSolution(sol)

	if probability >= rand.Float32() {
		return customerExchange(newSolution)
	}
	return customerTransfer(newSolution)
}

func getTransitionProbability(delta float32, temperature float32) float64 {
	if delta <= 0 {
		return 1
	} else {
		return math.Exp(-float64(delta / temperature))
	}
}

func getTemperatureCauchy(TMax float32, i int) float32 {
	return TMax / float32(1+i)
}

func RoutingProcedure(settingsRoute cargodelivery.RouteSettings, distanceMatrix [][]float32, solution cargodelivery.RouteSolution) (cargodelivery.RouteSolution, error) {

	rand.Seed(time.Now().UnixNano())

	var TMax = settingsRoute.TMax
	var TMin = settingsRoute.TMin

	var temperature = TMax

	flag := true

	bestSolution := solution
	bestSolutionCost := getSolutionCost(distanceMatrix, &bestSolution)

	var i = 0
	for temperature > TMin {
		newSolution := getNewSolution(bestSolution, 0.5)

		newSolution, err := packing.PackingProcedure(settingsRoute, newSolution)
		if err == nil {

			newSolutionCost := getSolutionCost(distanceMatrix, &newSolution)

			p := getTransitionProbability(newSolutionCost-bestSolutionCost, temperature)
			if p >= rand.Float64() {
				bestSolution = newSolution
				bestSolutionCost = newSolutionCost
				flag = false
			}
		}

		temperature = getTemperatureCauchy(TMax, i)
		i++
	}

	if flag == true {
		return solution, errors.New("не удалось найти позицию для размещения")
	}

	return bestSolution, nil
}
