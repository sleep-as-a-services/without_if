package strcon

import "testing"

func TestNtot_ShouldConvertNumber_To_EnglishNumber_CorrectResult(t *testing.T) {
	actual := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := []string{
		"zero",
		"one",
		"two",
		"three",
		"four",
		"five",
		"six",
		"seven",
		"eight",
		"nine",
	}

	for index := 0; index < len(actual); index++ {
		t.Run("expected should equal actual", func(t *testing.T) {
			if expected[index] != Ntot(actual[index]) {
				t.Errorf("expected: %s should  equal actual but got %d", expected[index], actual[index])
			}
		})
	}
}
