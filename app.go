package main

import (
	"fmt"
	oe "github.com/winlinvip/errors"
	"github.com/pkg/errors"
)

func main() {
	fmt.Println("Complex error is", errors.Errorf("created by github.com/pkg/errors"))
	fmt.Println(oe.Errorf("Hello, %v", "Go modules"))
}
