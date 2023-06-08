package packing

import (
	"errors"
	"math"
	"math/rand"
	"time"

	cargodelivery "github.com/Makcumblch/CargoDelivery"
	"github.com/Makcumblch/CargoDelivery/pkg/service/solutionDeliveryTask/utils"
)

func GetPackingCost(solution cargodelivery.RouteSolution) float32 {

	var packingCost float32

	for _, car := range solution.CarsRouteSolution {

		var Xcg float32
		var Ycg float32
		var Zcg float32
		var m float32

		for _, client := range car.Items {

			for _, item := range client {

				Xcg += item.Position.X * *item.Cargo.Weight
				Ycg += item.Position.Y * *item.Cargo.Weight
				Zcg += item.Position.Z * *item.Cargo.Weight
				m += *item.Cargo.Weight
			}

		}

		if m == 0 {
			continue
		}
		Xcg = Xcg / m
		Ycg = Ycg / m
		Zcg = Zcg / m

		packingCost += float32(math.Pow(float64(Xcg-*car.Width/2.0), 2)) + float32(math.Pow(float64(Ycg-*car.Length/2.0), 2)) + float32(math.Pow(float64(Zcg-*car.Height/2.0), 2))
	}

	return packingCost

}

func CheckUnderItem(packedItems []cargodelivery.Item, position cargodelivery.Position, item cargodelivery.Item) bool {

	for _, packedItem := range packedItems {

		if (packedItem.Position.Z+*packedItem.Cargo.Height == position.Z) &&
			(packedItem.Position.X < position.X+*item.Cargo.Width/2.0 && packedItem.Position.X+*packedItem.Cargo.Width > position.X+*item.Cargo.Width/2.0) &&
			(packedItem.Position.Y < position.Y+*item.Cargo.Length/2.0 && packedItem.Position.Y+*packedItem.Cargo.Length > position.Y+*item.Cargo.Length/2.0) {

			return true

		}

	}

	return false
}

// Удаление недопустимых позиций
func DeleteInvalidPositions(car cargodelivery.Car, item cargodelivery.Item, positionList []cargodelivery.Position, packedItems []cargodelivery.Item) []cargodelivery.Position {

	for _, packedItem := range packedItems {

		for posInd := len(positionList) - 1; posInd >= 0; posInd-- {
			position := positionList[posInd]
			if (packedItem.Position.X >= position.X+*item.Cargo.Width ||
				position.X >= packedItem.Position.X+*packedItem.Cargo.Width ||
				packedItem.Position.Y >= position.Y+*item.Cargo.Length ||
				position.Y >= packedItem.Position.Y+*packedItem.Cargo.Length ||
				packedItem.Position.Z >= position.Z+*item.Cargo.Height ||
				position.Z >= packedItem.Position.Z+*packedItem.Cargo.Height) &&
				(position.X >= 0 && position.Y >= 0 && position.Z >= 0 && (position.X+*item.Cargo.Width) <= *car.Width && (position.Y+*item.Cargo.Length <= *car.Length) && (position.Z+*item.Cargo.Height) <= *car.Height) {

				if position.Z != 0 && CheckUnderItem(packedItems, position, item) {

					continue

				}

				if packedItem.Position.Z+*packedItem.Cargo.Height != position.Z {

					if packedItem.Position.X < position.X+*item.Cargo.Width/2.0 &&
						packedItem.Position.X+*packedItem.Cargo.Width > position.X+*item.Cargo.Width/2.0 &&
						packedItem.Position.Y < position.Y+*item.Cargo.Length/2.0 &&
						packedItem.Position.Y+*packedItem.Cargo.Length > position.Y+*item.Cargo.Length/2.0 {

						positionList = append(positionList[:posInd], positionList[posInd+1:]...)
						continue

					}

				} else {

					positionList = append(positionList[:posInd], positionList[posInd+1:]...)
					continue

				}

			} else {

				positionList = append(positionList[:posInd], positionList[posInd+1:]...)
				continue

			}

		}
	}

	return positionList
}

func SetCoordinatesAtPosition(X float32, Y float32, Z float32) cargodelivery.Position {

	var position cargodelivery.Position

	position.X = X
	position.Y = Y
	position.Z = Z

	return position
}

// Добавление угловых позиций контейнера
func AddCornerPositions(car cargodelivery.Car, item cargodelivery.Item, positionList []cargodelivery.Position) []cargodelivery.Position {

	var newPosition cargodelivery.Position

	newPosition = SetCoordinatesAtPosition(0, 0, 0)
	positionList = append(positionList, newPosition)

	newPosition = SetCoordinatesAtPosition((*car.Width - *item.Cargo.Width), 0, 0)
	positionList = append(positionList, newPosition)

	newPosition = SetCoordinatesAtPosition(0, (*car.Length - *item.Cargo.Length), 0)
	positionList = append(positionList, newPosition)

	newPosition = SetCoordinatesAtPosition((*car.Width - *item.Cargo.Width), (*car.Length - *item.Cargo.Length), 0)
	positionList = append(positionList, newPosition)

	return positionList

}

// Получение возможных позиций
func GetPositions(car cargodelivery.Car, item cargodelivery.Item, packedItems []cargodelivery.Item) []cargodelivery.Position {

	positionList := make([]cargodelivery.Position, 0)

	positionList = AddCornerPositions(car, item, positionList)

	for _, packedItem := range packedItems {

		newPosition := SetCoordinatesAtPosition(packedItem.Position.X, (packedItem.Position.Y - *item.Cargo.Length), packedItem.Position.Z)
		positionList = append(positionList, newPosition)

		newPosition = SetCoordinatesAtPosition((packedItem.Position.X + *packedItem.Cargo.Width), packedItem.Position.Y, packedItem.Position.Z)
		positionList = append(positionList, newPosition)

		newPosition = SetCoordinatesAtPosition(packedItem.Position.X, (packedItem.Position.Y + *packedItem.Cargo.Length), packedItem.Position.Z)
		positionList = append(positionList, newPosition)

		newPosition = SetCoordinatesAtPosition((packedItem.Position.X - *item.Cargo.Width), packedItem.Position.Y, packedItem.Position.Z)
		positionList = append(positionList, newPosition)

		newPosition = SetCoordinatesAtPosition(packedItem.Position.X, packedItem.Position.Y, (packedItem.Position.Z + *packedItem.Cargo.Height))
		positionList = append(positionList, newPosition)

		positionList = DeleteInvalidPositions(car, item, positionList, packedItems)

	}

	return positionList

}

// Выбор лучшей позиции
func BestPosition(car cargodelivery.Car, item cargodelivery.Item, PackedItems []cargodelivery.Item, settingsRoute cargodelivery.RouteSettings) (cargodelivery.Position, error) {

	var bestPosition cargodelivery.Position

	positionList := GetPositions(car, item, PackedItems)
	var min = cargodelivery.Position{Z: math.MaxFloat32}
	if *settingsRoute.PackingType == false {

		bestZ := make([]cargodelivery.Position, 0)

		for _, position := range positionList {

			if position.Z < min.Z {
				min = position
			}

		}

		for _, position := range positionList {

			if position.Z == min.Z {

				bestZ = append(bestZ, position)

			}

		}

		min.Y = math.MaxFloat32
		bestY := make([]cargodelivery.Position, 0)

		for _, position := range bestZ {

			if position.Y < min.Y {
				min = position
			}

		}

		for _, position := range bestZ {

			if position.Y == min.Y {

				bestY = append(bestY, position)

			}

		}

		min.X = math.MaxFloat32
		bestX := make([]cargodelivery.Position, 0)

		for _, position := range bestY {

			if position.X < min.X {
				min = position
			}

		}

		for _, position := range bestY {

			if position.X == min.X {

				bestX = append(bestX, position)

			}

		}

		if len(bestX) == 0 {
			return cargodelivery.Position{}, errors.New("не удалось найти позицию для размещения")
		}

		bestPosition = bestX[0]

	} else {

		// var min = cargodelivery.Position{Y: math.MaxFloat32}

		bestY := make([]cargodelivery.Position, 0)
		min.Y = math.MaxFloat32
		for _, position := range positionList {

			if position.Y < min.Y {
				min = position
			}

		}

		for _, position := range positionList {

			if position.Y == min.Y {

				bestY = append(bestY, position)

			}

		}

		min.X = math.MaxFloat32
		bestX := make([]cargodelivery.Position, 0)

		for _, position := range bestY {

			if position.X < min.X {
				min = position
			}

		}

		for _, position := range bestY {

			if position.X == min.X {

				bestX = append(bestX, position)

			}

		}

		min.Z = math.MaxFloat32
		bestZ := make([]cargodelivery.Position, 0)

		for _, position := range bestX {

			if position.Z < min.Z {
				min = position
			}

		}

		for _, position := range bestX {

			if position.Z == min.Z {

				bestZ = append(bestZ, position)

			}

		}

		if len(bestZ) == 0 {
			return cargodelivery.Position{}, errors.New("не удалось найти позицию для размещения")
		}

		bestPosition = bestZ[0]

	}

	return bestPosition, nil
}

func SetCoordinates(item *cargodelivery.Item, X float32, Y float32, Z float32) *cargodelivery.Item {

	item.Position.X = X
	item.Position.Y = Y
	item.Position.Z = Z

	return item

}

func Decoder(solution cargodelivery.RouteSolution, settingsRoute cargodelivery.RouteSettings) (cargodelivery.RouteSolution, error) {

	for carId := 0; carId < len(solution.CarsRouteSolution); carId++ {
		car := &solution.CarsRouteSolution[carId]
		packedItems := make([]cargodelivery.Item, 0)

		for clientId := 0; clientId < len(car.Items); clientId++ {
			client := &car.Items[clientId]
			for itemId := 0; itemId < len(*client); itemId++ {
				item := &(*client)[itemId]
				if len(packedItems) == 0 {

					item = SetCoordinates(item, 0, 0, 0)
					packedItems = append(packedItems, *item)
					continue

				}

				bestPosition, err := BestPosition(car.Car, *item, packedItems, settingsRoute) //error!!

				if err != nil {
					return solution, err
				}

				item = SetCoordinates(item, bestPosition.X, bestPosition.Y, bestPosition.Z)
				packedItems = append(packedItems, *item)
			}

		}

	}

	solution.PackingCost = GetPackingCost(solution)

	return solution, nil
}

func EvMutation(Items [][]cargodelivery.Item) [][]cargodelivery.Item {
	if len(Items) == 0 {
		return Items
	}
	clientInd := utils.GetRandInt(len(Items)-1, 0)

	if len(Items[clientInd]) == 0 {
		return Items
	}
	minItemInd := utils.GetRandInt(len(Items[clientInd])-1, 0)
	maxItemInd := utils.GetRandInt(len(Items[clientInd])-1, 0)

	Items[clientInd][minItemInd], Items[clientInd][maxItemInd] = Items[clientInd][maxItemInd], Items[clientInd][minItemInd]

	return Items
}

func RotateItem(Items [][]cargodelivery.Item) [][]cargodelivery.Item {
	if len(Items) == 0 {
		return Items
	}
	clientInd := utils.GetRandInt(len(Items)-1, 0)

	if len(Items[clientInd]) == 0 {
		return Items
	}
	ItemInd := utils.GetRandInt(len(Items[clientInd])-1, 0)

	Items[clientInd][ItemInd].Cargo.Length, Items[clientInd][ItemInd].Cargo.Width = Items[clientInd][ItemInd].Cargo.Width, Items[clientInd][ItemInd].Cargo.Length

	return Items
}

func PackingProcedure(settingsRoute cargodelivery.RouteSettings, solution cargodelivery.RouteSolution) (cargodelivery.RouteSolution, error) {
	rand.Seed(time.Now().UnixNano())

	bestSolution := solution

	flag := true

	for i := 0; i < int(settingsRoute.EvCount); i++ {

		newSolution := utils.CloneSolution(bestSolution)

		for valueId := 0; valueId < len(newSolution.CarsRouteSolution); valueId++ {
			value := &newSolution.CarsRouteSolution[valueId]
			probability := rand.Float32()

			if probability < 0.5 {

				value.Items = EvMutation(value.Items)

			} else {

				value.Items = RotateItem(value.Items)

			}

		}

		newSolution, err := Decoder(newSolution, settingsRoute)

		if err != nil {

			continue
		}

		flag = false

		if newSolution.PackingCost < bestSolution.PackingCost {

			bestSolution = newSolution

		}

	}

	if flag == true {
		return cargodelivery.RouteSolution{}, errors.New("не удалось найти позицию для размещения")
	}

	return bestSolution, nil
}
