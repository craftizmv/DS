package mc

import (
	"errors"
)

type OrderShipmentStatus int

const (
	PLACED OrderShipmentStatus = iota
	PACKED
	PICKED_UP
	SHIPPED
	IN_TRANSIT
	DELIVERED
	CANCELLED
	RETURNED_FROM_CUSTOMER
	RETURN_IN_TRANSIT
	RETURNED_TO_SOURCE
	REPLACE_REQ_FROM_CUSTOMER
)

type ShipmentStrategyType int

const (
	SIMPLE_STRATEGY ShipmentStrategyType = iota
	SILVER_CUSTOMER_STRATEGY
)

type Order struct {
	ShipmentStatus OrderShipmentStatus
	eventCh        chan OrderShipmentStatus
}

func (o *Order) updateShipmentStatus(newState OrderShipmentStatus) {
	// check the state m/c
	o.ShipmentStatus = newState
}

func (o *Order) processEvent(newState OrderShipmentStatus) {

}

type ShipmentStrategy interface {
	isStateTransitionAllowed(current OrderShipmentStatus, next OrderShipmentStatus) (bool, error)
}

type SimpleShipmentStrategy struct {
	stateTransitionMachine map[OrderShipmentStatus]map[OrderShipmentStatus]struct{}
}

func NewSimpleShipmentStrategy() *SimpleShipmentStrategy {
	shipmentStrategy := &SimpleShipmentStrategy{
		stateTransitionMachine: make(map[OrderShipmentStatus]map[OrderShipmentStatus]struct{}),
	}

	shipmentStrategy.stateTransitionMachine[PLACED] = map[OrderShipmentStatus]struct{}{
		PACKED:    {},
		CANCELLED: {},
	}
	shipmentStrategy.stateTransitionMachine[PACKED] = map[OrderShipmentStatus]struct{}{
		PICKED_UP: {},
		CANCELLED: {},
	}
	shipmentStrategy.stateTransitionMachine[PICKED_UP] = map[OrderShipmentStatus]struct{}{
		SHIPPED:   {},
		CANCELLED: {},
	}
	shipmentStrategy.stateTransitionMachine[SHIPPED] = map[OrderShipmentStatus]struct{}{
		IN_TRANSIT: {},
		CANCELLED:  {},
	}
	shipmentStrategy.stateTransitionMachine[SHIPPED] = map[OrderShipmentStatus]struct{}{
		IN_TRANSIT: {},
		CANCELLED:  {},
	}
	shipmentStrategy.stateTransitionMachine[IN_TRANSIT] = map[OrderShipmentStatus]struct{}{
		DELIVERED: {},
		CANCELLED: {},
	}
	shipmentStrategy.stateTransitionMachine[DELIVERED] = map[OrderShipmentStatus]struct{}{
		RETURNED_FROM_CUSTOMER:    {},
		REPLACE_REQ_FROM_CUSTOMER: {},
	}

	return shipmentStrategy
}

func (s *SimpleShipmentStrategy) isStateTransitionAllowed(current OrderShipmentStatus, next OrderShipmentStatus) (bool, error) {
	if s.stateTransitionMachine[current] == nil {
		return false, errors.New("invalid current state")
	}

	if s.stateTransitionMachine[next] == nil {
		return false, errors.New("invalid next state")
	}

	if _, ok := s.stateTransitionMachine[current][next]; !ok {
		// valid state transition
		return false, errors.New("invalid transition attempt")
	}

	return true, nil
}

type ShipmentStrategyProvider struct {
	currentStrategy ShipmentStrategy
}

func NewShipmentStrategyProvider() *ShipmentStrategyProvider {
	return &ShipmentStrategyProvider{
		currentStrategy: NewSimpleShipmentStrategy(),
	}
}

func (p *ShipmentStrategyProvider) provideShipmentStrategy(strategyType ShipmentStrategyType) ShipmentStrategy {
	switch strategyType {
	case SIMPLE_STRATEGY:
		strategy := NewSimpleShipmentStrategy()
		return strategy
	case SILVER_CUSTOMER_STRATEGY:
		// TODO- provide custom impl
		return NewSimpleShipmentStrategy()
	default:
		return NewSimpleShipmentStrategy()
	}
}

func (p *ShipmentStrategyProvider) getCurrentStrategy() ShipmentStrategy {
	return p.currentStrategy
}

type ShipmentManagementService struct {
	strategyProvider *ShipmentStrategyProvider
}

func (sms *ShipmentManagementService) updateShipmentState(current OrderShipmentStatus, next OrderShipmentStatus) (bool, error) {
	allowed, err := sms.strategyProvider.getCurrentStrategy().isStateTransitionAllowed(current, next)
	if err != nil {
		return false, err
	}

	return allowed, nil
}

func InitSMS(strategyType ShipmentStrategyType) *ShipmentManagementService {
	provider := NewShipmentStrategyProvider()
	provider.provideShipmentStrategy(strategyType)
	return &ShipmentManagementService{
		strategyProvider: provider,
	}
}

type OMS struct {
	OrderIDMap                map[string]*Order
	shipmentManagementService *ShipmentManagementService
}

func NewOMS() *OMS {
	return &OMS{
		shipmentManagementService: InitSMS(SIMPLE_STRATEGY),
	}
}

func (oms *OMS) UpdateOrderShipment(orderID string, orderState OrderShipmentStatus) (*Order, error) {
	if order, ok := oms.OrderIDMap[orderID]; !ok {
		return nil, errors.New("order not found")
	} else {
		// take order and update shipment
		_, err := oms.shipmentManagementService.updateShipmentState(order.ShipmentStatus, orderState)
		if err != nil {
			return nil, err
		} else {
			// update the order state
			order.updateShipmentStatus(orderState)
			return order, nil
		}
	}

	//return nil, errors.New("error updating order shipment status")
}
