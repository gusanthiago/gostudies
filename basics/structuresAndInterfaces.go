package main

import "fmt"

type RailroadWideChecker interface {
	CheckRailsWidth() int
}

type Railroad struct {
	Width int
}

func (r *Railroad) IsCorrectSizeTrain(p RailroadWideChecker) bool {
	return p.CheckRailsWidth() != r.Width
}

type Train struct {
	TrainWidth int
}

func (p *Train) CheckRailsWidth() int {
	return p.TrainWidth
}

func main() {

	railroad := Railroad{Width: 10}
	passengerTrain := Train{TrainWidth: 10}
	cargoTrain := Train{TrainWidth: 15}
	fmt.Println("Can passenger train pass?",  railroad.IsCorrectSizeTrain(&passengerTrain))
	fmt.Println("Can cargo train pass?", railroad.IsCorrectSizeTrain(&cargoTrain))

}
