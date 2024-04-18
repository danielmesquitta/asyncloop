package main

import (
	"fmt"

	"github.com/danielmesquitta/asyncloop"
)

func main() {
	asyncloop.LoopN(3, func(i int) {
		fmt.Println(i)
	})
}
