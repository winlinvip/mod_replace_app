package main

import (
	"fmt"
	"github.com/winlinvip/mod_replace_protocol"
	"github.com/pkg/errors"
)

func main() {
	fmt.Println("Complex error is", errors.Errorf("created by github.com/pkg/errors"))
	fmt.Println(mod_replace_protocol.Errorf("Hello, %v", "Go modules"))
}
