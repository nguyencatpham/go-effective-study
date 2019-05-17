package helper

import (
	// "bytes"
	"encoding/json"
	// "fmt"
	"regexp"
	// "time"
	"unicode"
	"unicode/utf8"
)

// Regexp definitions
var keyMatchRegex = regexp.MustCompile(`\"(\w+)\":`)
var wordBarrierRegex = regexp.MustCompile(`(\w)([A-Z])`)

type ConventionalMarshaller struct {
	Value interface{}
}

func (c *ConventionalMarshaller) MarshalJSON() ([]byte, error) {
	marshalled, err := json.Marshal(c.Value)

	converted := keyMatchRegex.ReplaceAllFunc(
		marshalled,
		func(match []byte) []byte {
			// Empty keys are valid JSON, only lowercase if we do not have an
			// empty key.
			if string(match) == `"ID":` {
				return []byte(`"id":`)
			}
			if len(match) > 2 {
				// fmt.Println(string(match))
				// Decode first rune after the double quotes
				r, width := utf8.DecodeRune(match[1:])
				r = unicode.ToLower(r)
				utf8.EncodeRune(match[1:width+1], r)
			}
			// fmt.Println(string(match))
			if string(match) == `"ID":` {
				match = []byte(`"id":`)
			}
			return match
		},
	)

	return converted, err
}

// func (c *ConventionalMarshaller) MarshalJSON() ([]byte, error) {
// 	marshalled, err := json.MarshalIndent(c.Value, "", "  ")

// 	converted := keyMatchRegex.ReplaceAllFunc(
// 		marshalled,
// 		func(match []byte) []byte {
// 			return bytes.ToLower(wordBarrierRegex.ReplaceAll(
// 				match,
// 				[]byte(`${1}_${2}`),
// 			))
// 		},
// 	)

// 	return converted, err
// }
