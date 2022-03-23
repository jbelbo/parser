package parser

type bracket struct {
	opening rune
	closing rune
}

var validBrackets = []bracket{
	{'{', '}'},
	{'[', ']'},
	{'(', ')'},
}

func IsValid(text string) bool {
	var stack []rune
	for _, char := range text {
		if !isBracket(char) {
			continue
		}

		if isOpeningChar(char) {
			stack = append(stack, char)
			continue
		}

		if len(stack) == 0 {
			return false
		}

		var lastBracketInStack rune
		lastBracketInStack, stack = popBracket(stack)
		for _, bracket := range validBrackets {
			if bracket.closing == char && bracket.opening != lastBracketInStack {
				return false
			}
		}
	}

	return len(stack) == 0
}

func popBracket(stack []rune) (rune, []rune) {
	lastBracketInStack := stack[len(stack)-1]
	stack = stack[:len(stack)-1]

	return lastBracketInStack, stack
}

func isOpeningChar(r rune) bool {
	for _, bracket := range validBrackets {
		if r == bracket.opening {
			return true
		}
	}

	return false
}

func isBracket(r rune) bool {
	for _, bracket := range validBrackets {
		if r == bracket.opening || r == bracket.closing {
			return true
		}
	}

	return false
}
