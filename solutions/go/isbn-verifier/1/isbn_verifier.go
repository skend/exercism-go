package isbnverifier

import "strings"

func IsValidISBN(isbn string) bool {
	isbn = strings.ReplaceAll(isbn, "-", "")

    if len(isbn) != 10 {
        return false
    }

    total := 0
    start := 10
    for i, ch := range(isbn) {
        num := int(ch - '0')
        if ch == 'X' || ch == 'x' {
            if i != 9 {
                return false
            }
            num = 10
        }
        total += (num * start)
        start -= 1
    }

    return total % 11 == 0
}
