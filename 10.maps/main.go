package main

import "fmt"

func main() {
	fmt.Println("Maps in golang")

	languages := make(map[string]string)

	languages["CPP"] = "c plus plus"
	languages["JS"] = "javscript"
	languages["PY"] = "python"

	fmt.Println("languages map: ", languages)
	fmt.Println("JS: ", languages["JS"])

	// delete
	delete(languages, "PY")
	fmt.Println("languages now: ", languages)

	// loops
	for key, value := range languages { // can use comma, ok syntax for key/value like _, value or key, _
		fmt.Println(key, " ", value)
	}
}
