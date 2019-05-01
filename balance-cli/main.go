package main

import (
	"flag"
	"fmt"
	"github.com/lancatlin/balance"
)

func main() {
	equation := flag.String("s", "H2 + O2 = H2O", "Chemical equation")
	flag.Parse()
	fmt.Println(*equation)
	fmt.Println(balance.Balance(*equation))
}
