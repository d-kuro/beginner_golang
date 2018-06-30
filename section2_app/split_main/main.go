package main

import (
	"fmt"

	"./language"
)

func main() {
	fmt.Println(AppName())

	fmt.Println(language.English())
	fmt.Println(language.Japanese())
}
