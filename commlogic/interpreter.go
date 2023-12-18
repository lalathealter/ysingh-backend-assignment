package commlogic

import (
	"strings"
)

type CommandFunction func([]string) (string, error)
type CommandInterpreter map[string]CommandFunction

func (ci CommandInterpreter) Execute(args []string) (string, error) {
	com := strings.ToUpper(args[0])
	comfun, ok := ci[com]
	if !ok {
		return "", ErrUnknownCommand
	}
	ans, err := comfun(args[1:])
	return ans, err
}

func ProduceCommandInterpreter() CommandInterpreter {
	return CommandInterpreter{
		StopTerm: ExitApp,
		AddTerm: BranchCommands(CommandInterpreter{
			ProductTerm:   AddProduct,
			WarehouseTerm: AddWarehouse,
		}),
		// StockTerm:   StockItem,
		// UnstockTerm: UnstockItem,
		// ListTerm:    ListItems,
	}
}

const (
	AddTerm        = "ADD"
	StockTerm      = "STOCK"
	UnstockTerm    = "UNSTOCK"
	ListTerm       = "LIST"
	StopTerm       = "STOP"
	ProductTerm    = "PRODUCT"
	WarehouseTerm  = "WAREHOUSE"
	WarehousesTerm = "WAREHOUSES"
)

func BranchCommands(ci CommandInterpreter) CommandFunction {
	return ci.Execute
}
