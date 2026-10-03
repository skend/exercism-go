package luhn

import "strings"

func Valid(id string) bool {
    id = strings.ReplaceAll(id, " ", "")
    
	if len(id) <= 1 {
        return false
    }
    
	sum := 0
    for i := 0; i < len(id); i += 1 {
		if id[i] < '0' || id[i] > '9' {
            return false
        }
        
        digit := int(id[i] - '0')
        offsetFromEnd := len(id) - i
        if offsetFromEnd % 2 == 1 {
            sum = sum + digit
        } else if digit < 5 {
            sum = sum + 2 * digit
        } else {
            sum = sum + 2 * digit - 9
        }
    }
    return (sum % 10) == 0
}
