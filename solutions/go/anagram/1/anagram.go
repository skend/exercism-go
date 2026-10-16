package anagram

import (
    "maps"
    "strings"
)

func Detect(subject string, candidates []string) []string {
    var res []string
    lwrsub := strings.ToLower(subject)
    m := countCharacters(lwrsub)

    for _, c := range(candidates) {
		lwrc := strings.ToLower(c)
        
		if lwrsub == lwrc {
            continue
        }
        
        cm := countCharacters(lwrc)

        if maps.Equal(m, cm) {
            res = append(res, c)
        }
    }

    return res
}

func countCharacters(s string) map[rune]int {
    m := make(map[rune]int)
	for _, ch := range(s) {
        m[ch] += 1
    }

    return m
}
