package main

import "fmt"

func main() {
	fmt.Println("Maps in Go Lang")

	languages := make(map[string]string)

	languages["JS"] = "Javascript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages) // Output =>> map[JS:Javascript PY:Python RB:Ruby] note that output is not comma separated values
	fmt.Println("JS short for: ", languages["JS"])

	delete(languages, "RB")
	fmt.Println(languages) // Output =>> map[JS:Javascript PY:Python]

	// Loops are interesting in go lang

	for key, value := range languages {
		fmt.Printf("For key %v, value is %v \n", key, value)
	}

}
