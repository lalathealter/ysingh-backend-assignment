package commlogic

import (
	"fmt"
	"strconv"
)

var Stock = BindStock(1, "%vx UNITS OF PRODUCT [%v] WERE ADDED TO THE WAREHOUSE #%v")
var Unstock = BindStock(-1, "%vx UNITS OF PRODUCT [%v] WERE SUBTRACTED FROM THE WAREHOUSE #%v")

func BindStock(multiplier int, msgTemp string) func([]string) (string, error) {
	return func(args []string) (string, error) {
		ans := ""
		SKU := args[0]
		wareInt, err := strconv.Atoi(args[1])
		QTY, err := strconv.Atoi(args[2])
		if err != nil {
			return ans, ErrArguementNotInteger
		}

		cat := GetCatalog()
		_, isInCatalog := cat[SKU]
		if !isInCatalog {
			return ans, ErrProductDoesntExist
		}

		wares := GetWarehousesColl()
		ware, isWareValid := wares[wareInt]
		if !isWareValid {
			return ans, ErrWarehouseDoesntExist
		}

		diff := ware.Store(SKU, multiplier*QTY)
		ans = fmt.Sprintf(msgTemp, diff, SKU, wareInt)
		return ans, nil
	}
}
