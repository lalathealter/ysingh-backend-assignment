package commlogic

import (
	"fmt"
	"strconv"
	"strings"
)

func AddWarehouse(args []string) (string, error) {
	ans := ""
	if len(args) < 1 {
		return ans, ErrFewArguements
	}

	wareInt, err := strconv.Atoi(args[0])
	if err != nil {
		return ans, ErrArguementNotInteger
	}

	var stockLimit int
	if len(args) > 1 {
		stockLimit, err = strconv.Atoi(args[1])
		if err != nil {
			return ans, ErrArguementNotInteger
		}
		if stockLimit < 0 {
			stockLimit = 0
		}
	}

	wares := GetWarehousesColl()
	_, alreadyThere := wares[wareInt]
	if alreadyThere {
		return ans, ErrItemAlreadyExists
	}

	wares[wareInt] = stockLimit
	ans = fmt.Sprintf("Warehouse #%v was added successfully", wareInt)
	return ans, nil
}

func AddProduct(args []string) (string, error) {
	ans := ""
	if len(args) < 2 {
		return ans, ErrFewArguements
	}

	prodName, args := StickTogetherQuotedArg(args)
	if len(args) < 1 {
		return ans, ErrFewArguements
	}

	SKU := args[len(args)-1]
	if SKU == "" || prodName == "" {
		return ans, ErrFewArguements
	}

	cat := GetCatalog()
	_, alreadyThere := cat[SKU]
	if alreadyThere {
		return ans, ErrItemAlreadyExists
	}

	cat[SKU] = prodName
	ans = fmt.Sprintf("[%v] was added successfully", prodName)
	return ans, nil
}

func StickTogetherQuotedArg(args []string) (string, []string) {
	result := ""
	indexToCutTo := 1
	if args[0][0] != '"' {
		result = args[0]
	} else {
		args[0] = args[0][1:]
		for i, v := range args {
			if qI := strings.LastIndex(v, "\""); qI > -1 {
				result += v[0:qI]
				result = strings.TrimSpace(result)
				indexToCutTo = i + 1
				break
			}
			result += v + " "
		}
	}

	args = args[indexToCutTo:]
	return result, args
}
