package card

import (
	"github.com/tcm1911/passwordcard/utils"
)

// PasswordCard struct.
type PasswordCard []struct {
	number                              int64
	digitArea, includeSpecialCharacters bool
}

// GenerateRandomCard creates a new card with a randomly chosen number.
// A secure random number generator is used so that the card number is truly
// random.
func GenerateRandomCard(number int64, digitArea, specialChars bool) {

}

/*
   Private functions.
*/
func createHeader(headerChars []byte, width int, rnd utils.Randomer) []byte {
	utils.Shufle(headerChars, rnd)
	if len(headerChars) > width {
		var shortHeader []byte
		copy(shortHeader, headerChars[:width])

	}
	return headerChars
}
