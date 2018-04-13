package coding_kata

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestGuessAllCorrect(t *testing.T) {
	t.Run("guess4BullsAndNoCows", func(t *testing.T) {
		ret := getHint("1111", "1111")
		assert.Equal(t, "4A0B", ret, "guess4BullsAndNoCows failed")
	})

	t.Run("guessNoBullsAndCows", func(t *testing.T) {
		ret := getHint("1111", "2222")
		assert.Equal(t, "0A0B", ret, "guessNoBullsAndCows failed")
	})

	t.Run("guessOneBullAndNoCow", func(t *testing.T) {
		ret := getHint("1111", "1222")
		assert.Equal(t, "1A0B", ret, "guessOneBullAndNoCow failed")
	})

	t.Run("guessOneBullAndOneCow", func(t *testing.T) {
		ret := getHint("1112", "1323")
		assert.Equal(t, "1A1B", ret, "guessOneBullAndOneCow failed")
	})

	t.Run("guessOneBullAndOneCowWithDuplicateGuessNumber", func(t *testing.T) {
		ret := getHint("1112", "1223")
		assert.Equal(t, "1A1B", ret, "guessOneBullAndOneCowWithDuplicateGuessNumber failed")
	})

	t.Run("guessNoBullAndOneCowWithDuplicateGuessNumber", func(t *testing.T) {
		ret := getHint("1234", "0111")
		assert.Equal(t, "0A1B", ret, "guessNoBullAndOneCowWithDuplicateGuessNumber failed")
	})

	t.Run("guessNoBullAndFourCowWithDuplicateDigital", func(t *testing.T) {
		ret := getHint("1122", "2211")
		assert.Equal(t, "0A4B", ret, "guessNoBullAndFourCowWithDuplicateDigital failed")
	})

	t.Run("guessNoBullAndOneCowWithDuplicatePosition", func(t *testing.T) {
		ret := getHint("1122", "0001")
		assert.Equal(t, "0A1B", ret, "guessNoBullAndOneCowWithDuplicatePosition failed")
	})
}
