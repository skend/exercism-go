package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
        return -1, errors.New("input must be the same length")
    }

    count := 0
    for i,_ := range(a) {
        if a[i] != b[i] {
            count += 1 
        }
    }

    return count, nil
}
