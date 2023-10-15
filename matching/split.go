package matching

import "strings"

func Escape(str, delimiter string) string {
	return strings.ReplaceAll(str, delimiter, "\\"+delimiter)
}

// SplitIgnoringEscaped splits the input string str by the delimiter, ignoring escaped delimiters.
func SplitIgnoringEscaped(str, delimiter string) []string {
	var result []string

	// boolean variable to keep track whether the current character is escaped or not
	isEscaped := false
	// slice to accumulate characters for the current chunk
	var currentChunk strings.Builder

	for _, char := range str {
		// check if the current character is a backslash and toggle the isEscaped flag
		if string(char) == "\\" && !isEscaped {
			isEscaped = true
			continue // skip to the next character
		}
		// check if the current character is the delimiter and it is not escaped
		if string(char) == delimiter && !isEscaped {
			// if we have a delimiter, we append the current chunk to result and reset currentChunk
			result = append(result, currentChunk.String())
			currentChunk.Reset()
		} else {
			// handle escaped delimiter by adding a delimiter after a backslash
			if isEscaped {
				currentChunk.WriteRune('\\')
			}
			// otherwise, append the current character to the current chunk and reset the isEscaped flag
			currentChunk.WriteRune(char)
			isEscaped = false
		}
	}

	// add the last chunk to the result after the loop ends
	result = append(result, currentChunk.String())
	return result
}
