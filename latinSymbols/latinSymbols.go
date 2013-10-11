// latinSymbols.go
package main

import (
	"fmt"
	"unicode"
)

func hasOnlyLatinSymbols(word string) bool {
	for _, currentChar := range word {
		if !unicode.IsLetter(currentChar) {
			//fmt.Println(currentChar)
			continue
		}
		if !unicode.Is(unicode.Latin, currentChar) {
			return false
		}
	}
	return true
}

func main() {
	fmt.Printf("%t", hasOnlyLatinSymbols(`
    package main

    import "fmt"

    func main() {
        баба := 1
    }
`))
}
