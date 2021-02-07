package segment

import "unicode"

// SimpleWordBreaker is a UnicodeBreaker which breaks at whitespace.
// Whitespace is determined by unicode.IsSpace(r) for any rune.
type SimpleWordBreaker struct {
	//penaltiesSlice []int
	penalties []int
	matchLen  [3]int // one for spaces, one for non-spaces (word)
	matchType int
}

// SimpleWordBreaker will assign the following penalties:
//
// (1) Before a sequence of whitespace runes, the penalty will be 100.
//
// (2) After a sequence of whitespace runes, the penalty will be -100 (a merit).
//
const (
	PenaltyBeforeWhitespace int = 100
	PenaltyAfterWhitespace  int = -100
)

// NewSimpleWordBreaker creates
// a new SimpleWordBreaker. Does nothing special – usually it is
// sufficient to use an empty SimpleWordBreaker struct.
func NewSimpleWordBreaker() *SimpleWordBreaker {
	swb := &SimpleWordBreaker{}
	return swb
}

const (
	spaceType int = iota
	wordType
	eofType
)

// CodePointClassFor is a very simple implementation to return code point
// classes, which will return 1 for whitespace and 0 otherwise.
// (Interface UnicodeBreaker)
func (swb *SimpleWordBreaker) CodePointClassFor(r rune) int {
	if r == 0 {
		return eofType
	}
	if unicode.IsSpace(r) {
		return spaceType
	}
	return wordType
}

// StartRulesFor is part of interface UnicodeBreaker
func (swb *SimpleWordBreaker) StartRulesFor(rune, int) {
	if swb.penalties == nil {
		swb.penalties = make([]int, 0, 64)
	}
}

// var immediateBreakBefore = []int{0, -2000}
// var inhibitBreakBefore = []int{1000, 0}

// ProceedWithRune is part of interface UnicodeBreaker
func (swb *SimpleWordBreaker) ProceedWithRune(r rune, cpClass int) {
	// if r == 0 {
	// 	if swb.matchLen > 0 { // previous rune(s) is/were whitespace
	// 		// close a match of length matchLen (= count of whitespace runes)
	// 		swb.penalties = swb.penaltiesSlice[:0]       // re-set penalties
	// 		swb.penalties = append(swb.penalties, 1000)  // inhibit break before WS span
	// 		swb.penalties = append(swb.penalties, -2000) // break point
	// 		for i := 2; i <= swb.matchLen; i++ {         // inhibit break between WS runes
	// 			swb.penalties = append(swb.penalties, 1000)
	// 		}
	// 		swb.penalties = append(swb.penalties, -900) // break before end of text
	// 		swb.matchLen = 0
	// 	} else {
	// 		swb.penalties = immediateBreakBefore // break before end of text
	// 	}
	// 	return
	// }
	if r != 0 && cpClass == swb.matchType {
		swb.penalties = swb.penalties[:0]
		swb.matchLen[swb.matchType]++
		//swb.penalties = nil
		//swb.penalties = inhibitBreakBefore // do not break between non-whitespace
		return
	}
	if cpClass != swb.matchType && swb.matchLen[swb.matchType] > 0 {
		// close a match of length matchLen (= count of runes)
		swb.penalties = swb.penalties[:0]
		if swb.matchType == spaceType {
			CT().Debugf("word breaker: closing spaces run")
			swb.penalties = append(swb.penalties, 0)
			swb.penalties = append(swb.penalties, PenaltyAfterWhitespace)
		} else {
			CT().Debugf("word breaker: closing word run")
			swb.penalties = append(swb.penalties, 0)
			swb.penalties = append(swb.penalties, PenaltyBeforeWhitespace)
		}
		// swb.penalties = append(swb.penalties, 1000) // inhibit break before WS span
		// for i := 2; i <= swb.matchLen; i++ {        // inhibit break between WS runes
		// 	swb.penalties = append(swb.penalties, 1000)
		// }
		// swb.penalties = append(swb.penalties, -900)
		//swb.matchLen = 0
	}
	if r != 0 && cpClass != swb.matchType {
		//CT().Debugf("word breaker: char type changed to %d", cpClass)
		swb.matchLen[swb.matchType] = 0
		swb.matchType = cpClass
		swb.matchLen[swb.matchType] = 1
	} else if r == 0 {
		swb.matchType = eofType
	}
	// else {
	// 	swb.penalties = inhibitBreakBefore // do not break between non-whitespace
	// }
}

// LongestActiveMatch is part of interface UnicodeBreaker
func (swb *SimpleWordBreaker) LongestActiveMatch() int {
	return swb.matchLen[swb.matchType]
}

// Penalties is part of interface UnicodeBreaker
func (swb *SimpleWordBreaker) Penalties() []int {
	CT().Debugf("word breaker: emitting %v\n", swb.penalties)
	return swb.penalties
}
