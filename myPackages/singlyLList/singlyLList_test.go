package singlyLList

import (
	"testing"
)

func TestCreate(t *testing.T) {
	got := *create(1)
	want := Node{1, nil}
	if got != want {
		t.Errorf("Expected %v, got %v\n", want, got)
	}
}

func TestAdd(t *testing.T) {
	n := create(1)
	n = add(n, 2)
	if n.Next.Value != 2 && n.Value == 1 {
		t.Errorf("add() did not work\n")
	}
}

func BenchmarkFor(b *testing.B) {
	s := make([]int, 2000000000)
	for i := 0; i < b.N; i++ {
		s[i] = s[i] * s[i]
	}
}

func BenchmarkRange(b *testing.B) {
	s := make([]int, 2000000000)
	for i, _ := range s {
		s[i] = s[i] * s[i]
	}
}

func TestRemove(t *testing.T) {}

func TestReverse(t *testing.T) {}
