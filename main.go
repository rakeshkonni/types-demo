package main

import (
	"github.com/rakeshkonni/types-demo/organization"
)

func main() {

	p := organization.NewPerson("Rakesh", "Konni")

	println(p.FullName())
}
