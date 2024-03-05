package iteration

import "testing"

func TestRepeat(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, expected string) {
		t.Helper()
		if got != expected {
			t.Errorf("expected %q but got %q", expected, got)
		}
	}

	t.Run("repeat n (>0) times", func(t *testing.T) {
		repeated := Repeat("a")
		expected := "aaaaa"
		assertCorrectMessage(t, repeated, expected)
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}
