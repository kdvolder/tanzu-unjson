package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"

	"gitlab.com/c0b/go-ordered-json"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	for scanner.Scan() {
		text := scanner.Text()
		entry := ordered.NewOrderedMap()
		err := json.Unmarshal([]byte(text), &entry)
		if err != nil {
			//not valid json
			fmt.Println(text)
		} else {
			iter := entry.EntriesIter()
			for {
				pair, ok := iter()
				if !ok {
					break
				}
				s := fmt.Sprintf("%v", pair.Value)
				if strings.Contains(s, "\n") {
					// multiline string
					fmt.Println()
					fmt.Println(strings.TrimRightFunc(s, unicode.IsSpace))
				} else {
					// singleline value
					fmt.Print(s, " ")
				}
			}
			fmt.Println()
		}
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}
