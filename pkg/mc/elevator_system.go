package mc

import (
	"fmt"
	"time"
)

// Direction represents the movement direction of the elevator.
type Direction int

const (
	Idle Direction = iota
	Up
	Down
)

func (d Direction) String() string {
	switch d {
	case Up:
		return "up"
	case Down:
		return "down"
	default:
		return "idle"
	}
}

// ElevatorEventType defines types of events.
type ElevatorEventType int

const (
	Request ElevatorEventType = iota // A floor request event.
	Tick                             // A periodic tick event to move the elevator.
)

// ElevatorEvent represents an event in the system.
type ElevatorEvent struct {
	EventType ElevatorEventType
	Floor     int       // Floor number for Request events.
	Dir       Direction // Direction for the request (optional for further extension).
}

// Elevator models an individual elevator.
type Elevator struct {
	id           int
	currentFloor int
	destinations []int
	direction    Direction
	eventCh      chan ElevatorEvent
}

// NewElevator initializes a new Elevator.
func NewElevator(id int) *Elevator {
	return &Elevator{
		id:           id,
		currentFloor: 0,
		destinations: []int{},
		direction:    Idle,
		eventCh:      make(chan ElevatorEvent, 10), // buffered channel .. why?
	}
}

// addDestination adds a new floor to the destination list if not already present.
func (e *Elevator) addDestination(floor int) {
	for _, f := range e.destinations {
		if f == floor {
			return
		}
	}
	e.destinations = append(e.destinations, floor)
}

// processEvent handles incoming events.
func (e *Elevator) processEvent(event ElevatorEvent) {
	switch event.EventType {
	case Request:
		fmt.Printf("Elevator %d received request for floor %d\n", e.id, event.Floor)
		e.addDestination(event.Floor)
	case Tick:
		e.move()
	}
}

// move advances the elevator toward its next destination.
func (e *Elevator) move() {
	if len(e.destinations) == 0 {
		e.direction = Idle
		fmt.Printf("Elevator %d is idle at floor %d\n", e.id, e.currentFloor)
		return
	}
	target := e.destinations[0]
	if e.currentFloor < target {
		e.direction = Up
		e.currentFloor++
	} else if e.currentFloor > target {
		e.direction = Down
		e.currentFloor--
	}
	fmt.Printf("Elevator %d moving %s to floor %d\n", e.id, e.direction, e.currentFloor)
	if e.currentFloor == target {
		fmt.Printf("Elevator %d arrived at floor %d\n", e.id, e.currentFloor)
		// Remove the reached destination.
		e.destinations = e.destinations[1:]
	}
}

// Run starts the elevator's event loop.
func (e *Elevator) Run() {
	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case event := <-e.eventCh:
			e.processEvent(event)
		case <-ticker.C:
			// Trigger a tick event to move the elevator.
			e.processEvent(ElevatorEvent{EventType: Tick})
		}
	}
}

// ElevatorSystem manages multiple elevators.
type ElevatorSystem struct {
	elevators []*Elevator
}

// NewElevatorSystem creates an elevator system with a given number of elevators.
func NewElevatorSystem(numElevators int) *ElevatorSystem {
	system := &ElevatorSystem{}
	for i := 0; i < numElevators; i++ {
		elevator := NewElevator(i + 1)
		system.elevators = append(system.elevators, elevator)
		// Start each elevator's event loop in its own goroutine.
		go elevator.Run()
	}
	return system
}

// Dispatch assigns a floor request event to one of the elevators.
// For simplicity, this implementation sends the event to the first elevator.
func (es *ElevatorSystem) Dispatch(event ElevatorEvent) {
	if len(es.elevators) > 0 {
		// In a more advanced system, choose the best elevator based on current status.
		es.elevators[0].eventCh <- event
	}
}

func Driver() {
	// Create an elevator system with 2 elevators.
	system := NewElevatorSystem(2)

	// Simulate external floor requests.
	go func() {
		time.Sleep(2 * time.Second)
		system.Dispatch(ElevatorEvent{EventType: Request, Floor: 5})
		time.Sleep(1 * time.Second)
		system.Dispatch(ElevatorEvent{EventType: Request, Floor: 2})
		time.Sleep(3 * time.Second)
		system.Dispatch(ElevatorEvent{EventType: Request, Floor: 8})
	}()

	// Prevent the main function from exiting.
	select {}
}
