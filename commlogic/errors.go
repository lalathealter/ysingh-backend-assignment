package commlogic

import (
	"errors"
)

var ErrUnknownCommand = errors.New("UNKNOWN COMMAND")
var ErrUnknownCommandTerm = errors.New("UNKNOWN COMMAND TERM")
var ErrFewArguements = errors.New("NOT ENOUGH ARGUEMENTS WERE PASSED")
var ErrItemAlreadyExists = errors.New("ITEM WITH THIS SKU (UNIQUE IDENTIFIER) ALREADY EXISTS")
var ErrArguementNotInteger = errors.New("FAILED TO PARSE ONE OF THE ARGUEMENTS AS AN INTEGER")
var ErrProductDoesntExist = errors.New("THE PRODUCT WITH THIS SKU (UNIQUE IDENTIFIER) DOESN'T EXIST IN THE CATALOG")
var ErrWarehouseDoesntExist = errors.New("THE WAREHOUSE WITH THE PROVIDED ID DOESN'T EXIST")
