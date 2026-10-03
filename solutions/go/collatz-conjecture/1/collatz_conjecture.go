package collatzconjecture

import "errors"

func CollatzConjecture(n int) (int, error) {
    count := 0

    for n > 1 {
        if n % 2 == 0 {
            n /= 2
        } else {
            n *= 3
            n += 1
        }
        count += 1
    }

    if n == 1 {
        return count, nil
    } else {
        return -1, errors.New("number is not 1")
    }
}
