package commlogic

import (
	"errors"
)

var ErrUnknownCommand = errors.New("UNKNOWN COMMAND")
var ErrUnknownCommandTerm = errors.New("UNKNOWN COMMAND TERM")
var ErrFewArguements = errors.New("NOT ENOUGH ARGUEMENTS WERE PASSED")
var ErrItemAlreadyExists = errors.New("ITEM WITH THIS SKU (UNIQUE IDENTIFIER) ALREADY EXISTS")
var ErrArguementNotInteger = errors.New("FAILED TO PARSE ONE OF THE ARGUEMENTS AS AN INTEGER")
