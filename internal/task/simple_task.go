package task

import (
	"fmt"
)

var simpleTaskSpec = "*/10 * * * * *"

func SimpleTask() {
	fmt.Println("this is a test")
}
