package main

import (
	"fmt"
	"strings"
)

func lewisText() string {
	return `As far as the eye could reach he saw nothing but the stems of the great 
		plants about him receding in the violet shade, and far overhead the 
		multiple transparency of huge leaves filtering the sunshine to the solemn
		splendour of twilight in which he walked. Whenever he felt able he ran 
		again; the ground continued soft and springy, covered with the same 
		resiliant weed which was the first his hands had touched in Malacandra.
		Once or twice a small red creature scuttled across his path, but 
		otherwise there seemed to be no life stirring in the wood; nothing to
		fear - except the fact of wandering unprovisioned and alone in a forest
		of unknown vegetation thousands or millions of miles beyond the reach
		or knowledge of man.`
}

func main() {
	mappedText := make(map[string]int)
	lewisString := lewisText()
	removeables := []string{".", ",", "-", ";"}
	frequentWords := make(map[string]int)
	fc := 1

	for i, _ := range removeables {
		lewisString = strings.ReplaceAll(lewisString, removeables[i], "")
	}

	for _, w := range strings.Fields(lewisString) {
		mappedText[w]++
		if mappedText[w] > fc {
			frequentWords[w] = mappedText[w]
		}
	}
	
	fmt.Printf("Words that occur more than once:\n%v", frequentWords)
}