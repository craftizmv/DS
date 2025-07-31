// Had implementation issues.

package lld_type_questions

import (
	"errors"
	"fmt"
)

// VehicleType represents the type of vehicle trying to park.
type VehicleType int
const (
	Motorcycle VehicleType = iota
	Car
	Van
)

// SpotType represents the type of parking spot.
type SpotType int
const (
	MotorcycleSpot SpotType = iota
	CarSpot
	VanSpot
	Invalid
)

var allValidSpotTypes = []SpotType{
	MotorcycleSpot,
	CarSpot,
	VanSpot,
}

var allVehicleTypes = []VehicleType{
	Motorcycle,
	Car,
	Van,
}

type ParkingStatus int
const(
	Available ParkingStatus = iota
	NotAvailable
)

var allParkingStatus = []ParkingStatus{
	Available,
	NotAvailable,
}

// Vehicle is a simple struct that wraps a VehicleType.
type Vehicle struct {
	vType VehicleType
}

// NewVehicle creates a new Vehicle of a specific type.
func NewVehicle(vType VehicleType) Vehicle {
	return Vehicle{vType: vType}
}

// ParkingSpot holds the type of the spot and whether it's occupied.
type ParkingSpot struct {
	spotType SpotType
	occupied bool
	position int
}


type ParkingLot struct {
	// Feedback - this is extra complex ..
	// Not worth
	// Identifying car adjacency for van parking is difficult.
	spots map[SpotType]map[ParkingStatus][]*ParkingSpot
	pLCap map[SpotType]int
}

// NewParkingSpot creates a new spot of a given type.
func NewParkingSpot(spotType SpotType) ParkingSpot {
	return ParkingSpot{spotType: spotType, occupied: false}
}

// TODO : ISSUE -  THIS WAS NOT NEEDED .. you did not read the problem carefully.
//type Position struct {
//	posNum int
//	used bool
//}

// ParkingLot manages all parking spots and handles parking logic.
//type ParkingLot struct {
//	// TODO: Use ParkingSpot struct to manage spots in the parking lot
//	spotStatus map[SpotType]int
//	// track position as well.
//
//	// max limit for PL
//	maxCapPL map[SpotType]int
//
//	// remaining posn
//	remainingSpotPostions map[SpotType][]Position
//	// TODO: Ideally this should have been map[SpotType]map[int]*Position
//}


// NewParkingLot initializes a parking lot with the specified number of each spot type.
func NewParkingLot(motorcycle, car, van int) *ParkingLot {
	p := &ParkingLot{
		pLCap: make(map[SpotType]int),
		spots: make(map[SpotType]map[ParkingStatus][]*ParkingSpot),
	}

	// initialise the spots.
	for _, st := range allValidSpotTypes {
			p.spots[st] = make(map[ParkingStatus][]*ParkingSpot)
		for _, pStatus := range allParkingStatus {
			p.spots[st][pStatus] = make([]*ParkingSpot, 0)
		}
	}

	p.pLCap = make(map[SpotType]int)
	p.pLCap[MotorcycleSpot] = motorcycle

	// set all the spot as available / free spot
	for i:=0;i<motorcycle;i++ {
		// lets assume .. if the motorcyle pos starts from 1000
		pS := &ParkingSpot{
			occupied: false,
			position: 1000 + i,
			spotType: MotorcycleSpot,
		}
		p.spots[MotorcycleSpot][Available] = append(p.spots[MotorcycleSpot][Available], pS)
	}

	// init car
	p.pLCap[CarSpot] = car
	for i:=0;i<car;i++ {
		// lets assume .. if the motorcyle pos starts from 1000
		pS := &ParkingSpot{
			occupied: false,
			position: 2000 + i,
			spotType: CarSpot,
		}
		p.spots[CarSpot][Available] = append(p.spots[CarSpot][Available], pS)
	}

	p.pLCap[VanSpot] = van
	for i:=0;i<van;i++ {
		// lets assume .. if the motorcyle pos starts from 1000
		pS := &ParkingSpot{
			occupied: false,
			position: 2000 + i,
			spotType: VanSpot,
		}
		p.spots[VanSpot][Available] = append(p.spots[VanSpot][Available], pS)
	}


	return p
}

// Park tries to park the vehicle using appropriate spot selection rules.
func (pl *ParkingLot) Park(vehicle Vehicle) bool {
	return pl.ParkVehicle(vehicle)
}


//TODO : Make this generic impl
func (pl *ParkingLot) ParkVehicle(vehicle Vehicle) bool {
	spotType, err := getSpotTypeForVehicle(vehicle.vType)
	if err != nil{
		return false
	}

	if pl.spotStatus[spotType] < pl.maxCapPL[spotType] {
		// try to park
		pl.spotStatus[spotType]  = pl.spotStatus[spotType] + 1

		// Need to delete for remaininig pos
		remainingPos := pl.remainingSpotPostions[MotorcycleSpot]
		// remove from the last.
		remainingPos = remainingPos[:len(remainingPos)-1]
		// again setting it in case of re-slicing.
		pl.remainingSpotPostions[MotorcycleSpot]  = remainingPos
		return true

	} else {
		// see peers positions in hierarchy.
		spotType, position, err := getEligibleVehiclePeers(SpotType)


		return false
	}
}


// implement a chain of resposibility
func  (pl *ParkingLot) getEligibleVehiclePeers(sType SpotType) (SpotType, *[]Position, error) {
	switch sType {
	case MotorcycleSpot:
		// check if car positions are remaining and then van pos.
		if pl.spotStatus[CarSpot] > 0 {
			// we can assign car loc.
			// obtain position and assign
			positions := pl.remainingSpotPostions[CarSpot]
			result := []Position{positions[len(positions)-1]}

			// re-slicing.
			pl.remainingSpotPostions[CarSpot] = pl.remainingSpotPostions[CarSpot][:len(positions)-1]

			return CarSpot, &result , nil
		}

	case CarSpot:
		// check Van first
		if pl.spotStatus[VanSpot] > 0 {
			// we can assign car loc.
			// obtain position and assign
			positions := pl.remainingSpotPostions[VanSpot]

			result := []Position{positions[len(positions)-1]}

			// re-slicing.
			pl.remainingSpotPostions[VanSpot] = pl.remainingSpotPostions[VanSpot][:len(positions)-1]

			// need to check if position point to mutated slice.
			return VanSpot, &result, nil
		}

		// TODO : check motorcycle at least 3 - Not in scope.

		return Invalid, nil, errors.New("could not find the eligible location")

	case VanSpot:
		// check car
		if pl.spotStatus[CarSpot] > 2 {
			// we can assign car loc.
			// obtain position and assign

			positions := pl.remainingSpotPostions[VanSpot]

			// need to check if the available positions are consecutive in nature.
			consecutivePosMap := make(map[int]struct{}, len(pl.spotStatus))
			for i:= range positions {
				// create a positionMap and see if the prev and next positions are found.
				consecutivePosMap[positions[i].posNum] = struct{}{}
			}

			// pass 2. check consucutive.
			for i:= range positions {
				currentPos := positions[i].posNum
				// check if next 2 entries exist.
				if v1, ok := consecutivePosMap[currentPos+1];ok {
					if v2, ok := consecutivePosMap[currentPos+2];ok {
						result := []Position{positions[i], v1, v2}

						// delete position for the remaining positions and reduce available.
						pl.spotStatus[CarSpot] = pl.spotStatus[CarSpot]-3

						posToRemove := []int{currentPos, currentPos+1, currentPos+2}
						removeFromAvailablePos(posToRemove)

						return CarSpot, &result, nil
					}
				}

				// re-slicing.
				pl.remainingSpotPostions[VanSpot] = pl.remainingSpotPostions[VanSpot][:len(positions)-1]

				// need to check if position point to mutated slice.
				return VanSpot, &result, nil
			}

			// TODO : check motorcycle at least 3 - Not in scope.

			return Invalid, nil, errors.New("could not find the eligible location")

		}
	}


	func (pl *ParkingLot) removeFromAvailablePos(rPosns []int, sType SpotType) bool {

		spotList := ppl.remainingSpotPostions[ssType]
		newList := make([]Position, 0,len(spotList))

		for i := range spotList {
		spotPos := spotList[i]
		for pos := range rPosns {
		if spotPos != pos {
		// build a toRemove list.
		newList = append(newList, spotPos)
	}
	}
	}
	}


	func getSpotTypeForVehicle(vType VehicleType) (SpotType, error) {

		switch vType {
		case Motorcycle:
			return MotorcycleSpot, nil
		case Car:
			return CarSpot, nil
		case Van:
			return VanSpot, nil
		default:
			return Invalid, errors.New("invalid error")
		}
	}


	// SpotsRemaining returns the number of available spots.
	func (pl *ParkingLot) TotalSpots() int {
		return pl.spotStatus[CarSpot] + pl.spotStatus[MotorcycleSpot] + pl.spotStatus[VanSpot]
	}

	// SpotsRemaining returns the number of available spots.
	func (pl *ParkingLot) SpotsRemaining() int {
		return len(pl.remainingSpotPostions[CarSpot]) + len(pl.remainingSpotPostions[MotorcycleSpot]) + len(pl.remainingSpotPostions[VanSpot])
	}

	// IsFull returns true if all spots are taken.
	func (pl *ParkingLot) IsFull() bool {
		return false
	}

	// IsEmpty returns true if all spots are free.
	func (pl *ParkingLot) IsEmpty() bool {
		if (pl.spotStatus[CarSpot] == pl.maxCapPL[CarSpot]) && (pl.spotStatus[MotorcycleSpot] == pl.maxCapPL[MotorcycleSpot]) && (pl.spotStatus[VanSpot] == pl.maxCapPL[VanSpot]) {
		return true
	}

		return false
	}

	// IsSpotTypeFull returns true if all spots of a given type are occupied.
	func (pl *ParkingLot) IsSpotTypeFull(spotType SpotType) bool {
		if pl.spotStatus[spotType] == 0 {
		return true
	}

		return false
	}

	// SpotsUsedByVans returns how many spots are currently used by vans.
	func (pl *ParkingLot) SpotsUsedByVans() int {
		return pl.spotStatus[VanSpot]
	}

	// Demo usage of the parking lot
	func driveFunc() {
		pl := NewParkingLot(2, 4, 3) // 2 motorcycle, 4 car, 3 van spots

		fmt.Println("Total spots:", pl.TotalSpots())
		fmt.Println("Remaining spots:", pl.SpotsRemaining())

		//Park some vehicles
		pl.Park(NewVehicle(Motorcycle)) // Should park anywhere
		pl.Park(NewVehicle(Car))        // Should take a car spot
		pl.Park(NewVehicle(Van))        // Should take a van spot
		pl.Park(NewVehicle(Van))        // Should take 3 car spots if available

		fmt.Println("Remaining spots:", pl.SpotsRemaining())
		fmt.Println("Is parking lot full?", pl.IsFull())
		fmt.Println("Is parking lot empty?", pl.IsEmpty())
		fmt.Println("Are motorcycle spots full?", pl.IsSpotTypeFull(MotorcycleSpot))
		fmt.Println("Spots used by vans:", pl.SpotsUsedByVans())
	}
