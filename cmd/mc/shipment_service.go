package main

import (
	"fmt"
	"github.com/craftizmv/DS/pkg/mc"
)

func main() {
	oms := mc.NewOMS()
	oms.OrderIDMap = map[string]*mc.Order{
		"orderID1": &mc.Order{
			ShipmentStatus: mc.PLACED,
		},
	}
	shipment, err := oms.UpdateOrderShipment("orderID1", mc.PACKED)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		// updated order
		fmt.Println("new shipment status", shipment.ShipmentStatus)
	}
}
