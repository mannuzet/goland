package main

import "fmt"

func analyzeNotes(note1, note2 float64) (float64, string) {
	media := (note1 + note2) / 2

	var text string
	if media <= 6 {
		text = "you are failed"
	} else if media >= 7 {
		text = "you are approved"
	} else {
		text = "in recovery"
	}
	return media, text
}

func main() {
	average, result := analyzeNotes(7.5, 7.5)
	fmt.Println("average", average)
	fmt.Println("result", result)
}
