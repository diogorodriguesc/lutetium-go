package main

import (
	"fmt"
	"sort"

	"diogocorreia.com/golang-application/logging"
)

func main() {
	logging.LogInfo("Init Execution of Application")

	names := []string{"Portugal", "England", "Brazil", "Argentina", "Uruguay", "Germany", "Italy", "Indonesia", "India", "France", "Spain"}

	sort.Strings(names)
	for _, element := range names {
		fmt.Println(transform(element))
	}

	logging.LogInfo("End Execution of Application")
}

func transform(name string) string {
	return "Hello " + name
}
