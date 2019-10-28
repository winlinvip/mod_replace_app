package main

import (
	"fmt"
	"github.com/winterserver/errors"
)

func main() {
	fmt.Println(fmt.Sprintf("Complex error is %+v", errors.Errorf("created by github.com/pkg/errors")))
}
