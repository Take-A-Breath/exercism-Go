package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
	switch {
	case card == "one":
		return 1
	case card == "two":
		return 2
	case card == "three":
		return 3
	case card == "four":
		return 4
	case card == "five":
		return 5
	case card == "six":
		return 6
	case card == "seven":
		return 7
	case card == "eight":
		return 8
	case card == "nine":
		return 9
	case card == "ten":
		return 10
	case card == "jack":
		return 10
	case card == "queen":
		return 10
	case card == "king":
		return 10
	case card == "ace":
		return 11
	default:
		return 0
	}
}

// IsBlackjack returns true if the player has a blackjack, false otherwise.
func IsBlackjack(card1, card2 string) bool {
	var t bool
	if ParseCard(card1)+ParseCard(card2) == 21 {
		t = true
	}
	return t
}

// LargeHand implements the decision tree for hand scores larger than 20 points.
func LargeHand(isBlackjack bool, dealerScore int) string {
	//  1. If you have a pair of aces you must always split them.
	//  2. If you have a Blackjack (two cards that sum up to a value of 21), and the dealer does not have an ace, a figure, or a ten
	// 		then you automatically win.
	//  3. If the dealer does have any of those cards then you'll have to stand and wait for the reveal of the other card.
	var choice string
	switch {
	case isBlackjack && dealerScore < 10:
		choice = "W"
	case !isBlackjack:
		choice = "P"
	default:
		choice = "S"
	}
	return choice
}

// SmallHand implements the decision tree for hand scores with less than 21 points.
func SmallHand(handScore, dealerScore int) string {
	// If your cards sum up to 17 or higher you should always stand.
	// If your cards sum up to 11 or lower you should always hit.
	// If your cards sum up to a value within the range [12, 16] you should always stand unless the dealer has a 7 or higher, in which case you should always hit.
	var choice string
	switch {
	case handScore >= 17:
		choice = "S"
	case handScore <= 11:
		choice = "H"
	default:
		if dealerScore >= 7 {
			choice = "H"
		} else {
			choice = "S"
		}
	}
	return choice
}
