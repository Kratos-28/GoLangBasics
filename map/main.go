package main

import "fmt"

func main() {
	//first way to declare map
	colors := map[string]string{
		"red":   "#ff000",
		"green": "#4bf00",
	}
	delete(colors, "green")
	fmt.Println(colors)

	//second way to declare map
	// var laptop map[string]string

	//third way to declare map (recommended)
	laptop := make(map[string]string)
	laptop["dell"] = "premium"
	laptop["mac"] = "premimum"
	laptop["lenovo"] = "budget"

	printMap(laptop)

}
func printMap(l map[string]string) {
	for laptop, category := range l {
		fmt.Println("category for", laptop, "is", category)
	}
}
