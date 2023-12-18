package commlogic

import (
	"fmt"
	"strconv"
)

func ListProducts([]string) (string, error) {
	ans := "Contents of the product catalog:\n"
	cat := GetCatalog()
	for SKU, name := range cat {
		ans += fmt.Sprintf("[%v] %v\n", name, SKU)
	}
	ans += fmt.Sprintf("%v of items counted", len(cat))
	return ans, nil
}

func ListWarehouses([]string) (string, error) {
	ans := "Warehouses by number and storage capacity:\n"
	wares := GetWarehousesColl()
	for key, ware := range wares {
		ans += fmt.Sprintf("#%v : %v\n", key, ware.Limit)
	}
	ans += fmt.Sprintf("%v of warehouses present", len(wares))

	return ans, nil
}

func ListWarehouseContents(args []string) (string, error) {
	if len(args) < 1 {
		return "", ErrFewArguements
	}
	id, err := strconv.Atoi(args[0])
	if err != nil {
		return "", nil
	}

	wares := GetWarehousesColl()
	ware, ok := wares[id]
	if !ok {
		return "", ErrWarehouseDoesntExist
	}

	ans := fmt.Sprintf("Warehouse's #%v maximum capacity is: %v\n", id, ware.Limit)
	ans += fmt.Sprintf("Contents of the warehouse #%v:\n", id)
	for prod, count := range ware.Storage {
		ans += fmt.Sprintf("[%v]: %v units\n", prod, count)
	}
	ans += fmt.Sprintf("%v types of products present", len(ware.Storage))
	return ans, nil
}
