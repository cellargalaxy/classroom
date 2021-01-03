package main

import (
	"fmt"
	"github.com/cellargalaxy/classroom/service"
)

func main() {
	fmt.Printf("%+v", string(service.GetSignMessage()))
}
