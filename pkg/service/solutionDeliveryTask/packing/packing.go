package packing

import (
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

		for posInd, position := range positionList {

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

						positionList = append(positionList[:posInd-1], positionList[posInd+1:]...)
						continue

					}

				} else {

					positionList = append(positionList[:posInd-1], positionList[posInd+1:]...)
					continue

				}

			} else {

				positionList = append(positionList[:posInd-1], positionList[posInd+1:]...)
				continue

			}

		}
	}

	return positionList
}

func SetCoordinatesAtPosition(position *cargodelivery.Position, X float32, Y float32, Z float32) {

	position.X = X
	position.Y = Y
	position.Z = Z

}

// Добавление угловых позиций контейнера
func AddCornerPositions(car cargodelivery.Car, item cargodelivery.Item, positionList []cargodelivery.Position) []cargodelivery.Position {

	var position cargodelivery.Position

	positionList = append(positionList, position)

	SetCoordinatesAtPosition(&position, 0, (*car.Length - *item.Cargo.Length), 0)
	positionList = append(positionList, position)

	SetCoordinatesAtPosition(&position, (*car.Width - *item.Cargo.Width), (*car.Length - *item.Cargo.Length), 0)
	positionList = append(positionList, position)

	return positionList

}

// Получение возможных позиций
func GetPositions(car cargodelivery.Car, item cargodelivery.Item, packedItems []cargodelivery.Item) []cargodelivery.Position {

	positionList := make([]cargodelivery.Position, 0)

	positionList = AddCornerPositions(car, item, positionList)

	for _, packedItem := range packedItems {

		var position cargodelivery.Position

		SetCoordinatesAtPosition(&position, packedItem.Position.X, (packedItem.Position.Y - *item.Cargo.Length), packedItem.Position.Z)
		positionList = append(positionList, position)

		SetCoordinatesAtPosition(&position, (packedItem.Position.X + *packedItem.Cargo.Width), packedItem.Position.Y, packedItem.Position.Z)
		positionList = append(positionList, position)

		SetCoordinatesAtPosition(&position, packedItem.Position.X, (packedItem.Position.Y + *packedItem.Cargo.Length), packedItem.Position.Z)
		positionList = append(positionList, position)

		SetCoordinatesAtPosition(&position, (packedItem.Position.X - *item.Cargo.Width), packedItem.Position.Y, packedItem.Position.Z)
		positionList = append(positionList, position)

		SetCoordinatesAtPosition(&position, packedItem.Position.X, packedItem.Position.Y, (packedItem.Position.Z + *packedItem.Cargo.Height))
		positionList = append(positionList, position)

	}

	return DeleteInvalidPositions(car, item, positionList, packedItems)

}

// Выбор лучшей позиции
func BestPosition(car cargodelivery.Car, item cargodelivery.Item, PackedItems []cargodelivery.Item) cargodelivery.Position {

	var bestPosition cargodelivery.Position

	positionList := GetPositions(car, item, PackedItems)

	var min = cargodelivery.Position{Z: math.MaxFloat32}

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

	bestPosition = bestX[0]

	return bestPosition
}

func SetCoordinates(item *cargodelivery.Item, X float32, Y float32, Z float32) {

	item.Position.X = X
	item.Position.Y = Y
	item.Position.Z = Z

}

func Decoder(solution cargodelivery.RouteSolution) (cargodelivery.RouteSolution, error) {

	for _, car := range solution.CarsRouteSolution {

		packedItems := make([]cargodelivery.Item, 0)

		for _, client := range car.Items {

			for _, item := range client {

				if len(packedItems) == 0 {

					SetCoordinates(&item, 0, 0, 0)
					packedItems = append(packedItems, item)
					continue

				}

				bestPosition := BestPosition(car.Car, item, packedItems) //error!!

				SetCoordinates(&item, bestPosition.X, bestPosition.Y, bestPosition.Z)
				packedItems = append(packedItems, item)
			}

		}

	}

	solution.PackingCost = GetPackingCost(solution)

	return solution, nil
}

func EvMutation(Items [][]cargodelivery.Item) [][]cargodelivery.Item {
	clientInd := utils.GetRandInt(len(Items)-1, 0)

	minItemInd := utils.GetRandInt(len(Items[clientInd])-1, 0)
	maxItemInd := utils.GetRandInt(len(Items[clientInd])-1, 0)

	Items[clientInd][minItemInd], Items[clientInd][maxItemInd] = Items[clientInd][maxItemInd], Items[clientInd][minItemInd]

	return Items
}

func RotateItem(Items [][]cargodelivery.Item) [][]cargodelivery.Item {
	clientInd := utils.GetRandInt(len(Items)-1, 0)
	ItemInd := utils.GetRandInt(len(Items[clientInd])-1, 0)

	Items[clientInd][ItemInd].Cargo.Length, Items[clientInd][ItemInd].Cargo.Width = Items[clientInd][ItemInd].Cargo.Width, Items[clientInd][ItemInd].Cargo.Length

	return Items
}

func PackingProcedure(settingsRoute cargodelivery.RouteSettings, solution cargodelivery.RouteSolution) (cargodelivery.RouteSolution, error) {
	rand.Seed(time.Now().UnixNano())

	bestSolution := solution

	for i := 0; i < int(settingsRoute.EvCount); i++ {

		newSolution := utils.CloneSolution(bestSolution)

		for _, value := range newSolution.CarsRouteSolution {

			probability := rand.Float32()

			if probability < 0.5 {

				value.Items = EvMutation(value.Items)

			} else {

				value.Items = RotateItem(value.Items)

			}

		}

		newSolution, _ = Decoder(newSolution)

		if newSolution.PackingCost < bestSolution.PackingCost {

			bestSolution = newSolution

		}

	}

	return bestSolution, nil
}
