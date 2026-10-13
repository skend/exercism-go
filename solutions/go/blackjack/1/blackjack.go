package blackjack


// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch (card) {
    case "ace":
        return 11
    case "king","queen","jack","ten":
        return 10
    case "nine":
        return 9
    case "eight":
        return 8
    case "seven":
        return 7
    case "six":
        return 6
    case "five":
        return 5
    case "four":
        return 4
    case "three":
        return 3
    case "two":
        return 2
    default:
    	return 0
    }

    return 0
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    parsed1 := ParseCard(card1)
    parsed2 := ParseCard(card2)
    parsed3 := ParseCard(dealerCard)
    
	if card1 == card2 && card1 == "ace" {
        return "P"
    }

    sum := parsed1 + parsed2

    if sum == 21 {
        if parsed3 < 10 {
            return "W"
        } else {
            return "S"
        }
    } else if sum >= 17 && sum <= 20 {
        return "S"
    } else if sum >= 12 && sum <= 16 {
        if parsed3 >= 7 {
            return "H"
        } else {
            return "S"
        }
    } else if sum <= 11 {
        return "H"
    } else {
        return "H"
    }
}
