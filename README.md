# Reverse Words in a String

This Go program takes an input string and reverses the order of the words in the string. The program processes the string from the end to the beginning, reversing each word and maintaining the original spaces between words.

## Code Explanation

```go
package main

import "fmt"

func main() {
    input := "Kenapa Internet Kok Lambat"

    // Convert the input string to a slice of runes
    perHuruf := []rune(input)

    // Initialize an empty slice to hold the reversed words
    penampung := []rune{}

    // Initialize an empty slice to temporarily hold characters of a word
    tampung := []rune{}

    // Iterate over the runes from the end to the beginning
    for i := len(perHuruf) - 1; i >= 0; i-- {
        // If a space is encountered, reverse the collected characters and add to penampung
        if string(perHuruf[i]) == " " {
            // Reverse the characters in tampung and add to penampung
            for j := len(tampung) - 1; j >= 0; j-- {
                penampung = append(penampung, tampung[j])
            }
            // Add a space after the word
            penampung = append(penampung, ' ')
            // Reset tampung for the next word
            tampung = []rune{}
        } else {
            // Collect characters of the current word
            tampung = append(tampung, perHuruf[i])
        }
    }

    // Add the last word to penampung
    for j := len(tampung) - 1; j >= 0; j-- {
        penampung = append(penampung, tampung[j])
    }

    // Print the reversed string
    fmt.Println(string(penampung))
}
```

### Detailed Steps

1. **Convert Input to Runes**: The input string is converted to a slice of runes (`perHuruf`). This is necessary to handle multi-byte characters correctly.

2. **Initialize Slices**: Two slices are initialized:
   - `penampung`: To hold the final reversed words.
   - `tampung`: To temporarily hold characters of the current word being processed.

3. **Iterate from End to Beginning**: The program iterates over the runes from the end to the beginning of the input string.

4. **Handle Spaces**: When a space is encountered:
   - The characters collected in `tampung` are reversed and added to `penampung`.
   - A space is added to `penampung` to separate words.
   - `tampung` is reset to collect the next word.

5. **Handle Last Word**: After the loop, the last word collected in `tampung` is reversed and added to `penampung`.

6. **Print Result**: The final reversed string is printed.

### Example

For the input string `"Kenapa Internet Kok Lambat"`, the output will be:

```
Lambat Kok Internet Kenapa
```

This program effectively reverses the order of words in the input string while maintaining the original spaces between words.
