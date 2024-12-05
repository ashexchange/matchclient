package client

import (
	"fmt"
	"log"
	"os"
)

var logger = log.New(os.Stderr, "DEBUG ", log.LstdFlags)

func debugf(format string, v ...any) {
	_ = logger.Output(2, fmt.Sprintf(format, v...))
}
