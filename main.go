package main

import (
	"fmt"
	"os"
	"strings"
	"github.com/eoBattisti/proj-stack-machine/src"
)


func main () {
  text := os.Args[1]
  text = strings.TrimSpace(text)

  var calc src.Calculator = src.Calculator{}

  value := calc.Parse(text)

  fmt.Printf("%.2f \n", value)
}
