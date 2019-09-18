package main

import "fmt"

const MAX_WIDTH = 16
var WORDS = [...]string{"this", "is", "an", "example", "of", "text", "justification"}

func main() {
	var begin_slice = 0
	var end_slice = 0
	var current_length = 0

	for i, word := range WORDS {
		// print out line
		if MAX_WIDTH < len(word) + 1 + current_length {

			total_word_lengths := 0
			for _, j := range WORDS[begin_slice:end_slice] {
				total_word_lengths += len(j)
			}

			//fmt.Println("Toral word lengths", total_word_lengths)

			remaining_spaces := MAX_WIDTH - total_word_lengths

			spaces_in_between := remaining_spaces / (len(WORDS[begin_slice:end_slice]) - 1)


			extra_space := false
			if (remaining_spaces % 2) != 0 {
				extra_space = true
			}

			//fmt.Println("Spaces in between", spaces_in_between)
			//fmt.Println("extra space", extra_space)
			fmt.Printf(WORDS[begin_slice])
			if extra_space {
				fmt.Print(" ")
			}

			for i = begin_slice + 1; i < end_slice; i++ {
				for j := 0; j < spaces_in_between; j++ {
					fmt.Printf(" ")
				}
				fmt.Printf(WORDS[i])
			}
			fmt.Printf("\n")

			begin_slice = i
			current_length = 0
		}

		if i == len(WORDS) - 1 {
			fmt.Println(word)
		}

		current_length += len(word) + 1
		end_slice = i + 1
	}
}
