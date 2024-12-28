package iteration

import "testing"

func TestRepeat(t *testing.T) {
	t.Run("Repeats characters as many as specified", func(t *testing.T){
		repeated := Repeat("a", 4)
		expected := "aaaa"
		
		if repeated != expected {
			t.Errorf("expected %q bu got %q", expected, repeated)
		}
	})
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 4)
	}
}