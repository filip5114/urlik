package main

import (
	"fmt"
	"os"

	"github.com/filip5114/urlik/webserver"
)

func main() {
	if err := webserver.Run(3000); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
	}
}
