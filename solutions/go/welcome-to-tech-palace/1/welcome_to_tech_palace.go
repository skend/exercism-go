package techpalace

import (
    "strings"
    "fmt"
)

// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
	return fmt.Sprintf("Welcome to the Tech Palace, %s", strings.ToUpper(customer))
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	var output string
    for i := 0; i < numStarsPerLine; i++ {
        output += "*"
    }

    output += fmt.Sprintf("\n%s\n", welcomeMsg)

    for i := 0; i < numStarsPerLine; i++ {
        output += "*"
    }

    return output
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	oldMsg = strings.ReplaceAll(oldMsg, "*", "")
    oldMsg = strings.TrimSpace(oldMsg)
    return oldMsg
}
