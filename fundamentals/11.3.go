package main

import "fmt"

func main() {
	capitalCities := map[string]string{}
	capitalCities["philippines"] = "Manila"
	capitalCities["malaysia"] = "Kuala Lumpur"
	capitalCities["usa"] = "Washington"

	japaneseCapital, doesExist := capitalCities["japan"]
	if doesExist {
		fmt.Println(japaneseCapital)
	}

	fmt.Println(capitalCities)
}
