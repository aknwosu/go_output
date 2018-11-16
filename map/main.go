package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#tb564",
		"white": "#ffffff",
	}
	fmt.Println(colors)
	printMap(colors)
}

func printMap(c map[string]string) {
	for color, h := range c {
		fmt.Println("Hex code for", color, "is", h)
	}
}
