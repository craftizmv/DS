package loops

import "fmt"

type SpotType int

const (
	FREE SpotType = iota
	AVAILABLE
)

var spots = []SpotType{AVAILABLE, FREE}

func LoopPractice() {
	for i, spot := range spots {
		fmt.Printf("Spot %d: %v\n", i, spot)
	}
}
