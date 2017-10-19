package lesson24

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if add(1, 2) != 3 {
		t.Error("error")
	}
}

func BenchmarkAdd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		add(1, 2)
	}

}
