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
	/*
	   grid = new char[HEIGHT][WIDTH];
	   Random random = new Random(number);
	   char[] headerChars = HEADER_CHARS.toCharArray();
	   shuffle(headerChars, random);
	   if (headerChars.length > WIDTH) {
	       char[] tmp = headerChars;
	       headerChars = new char[WIDTH];
	       System.arraycopy(tmp, 0, headerChars, 0, WIDTH);
	   }
	   grid[0] = headerChars;
	*/
}
