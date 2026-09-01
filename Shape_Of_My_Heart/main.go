// Shape of My Heart - Sting
package main

import "fmt"

func main() {
	// Lirik lagu Shape of My Heart
	lyrics := []string{
		"He deals the cards as a meditation",
		"And those words are like a sign",
		"And if you read between the lines",
		"You'll know that this is the shape of my heart",
		"",
		"Shape of my heart...",
	}

	// Print title
	fmt.Println("Shape of My Heart - Sting")
	fmt.Println("------------------------------")
	fmt.Println()

	// Print lyrics
	for _, line := range lyrics {
		fmt.Println(line)
	}
}
