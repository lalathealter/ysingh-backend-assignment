package commlogic

import (
	"io"
	"os"
)

func ExitApp(_ []string) (string, error) {
	go os.Exit(0)
	return "", io.EOF
}
