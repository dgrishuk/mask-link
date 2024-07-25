package main

import (
	"bytes"
	"strings"
)

func MaskLink(message string) string {
	prefixes := []byte("http")
	sourceStr := "Here's my spammy page: http://hehefouls.netHAHAHA see you."
	arrWords := strings.Split(sourceStr, " ")
	var resultStr string

	for _, word := range arrWords {
		var maskWord string

		if len(word) > 7 {
			byteWord := []byte(word[:4])
			if bytes.Equal(prefixes, byteWord) {
				indexForMask := strings.Index(word, "/")
				maskStars := strings.Repeat("*", len(word[indexForMask+2:]))
				maskWord = word[:indexForMask+2] + maskStars
				resultStr = strings.Replace(sourceStr, word, maskWord, -1)
			}
		}
	}
	return resultStr
}
