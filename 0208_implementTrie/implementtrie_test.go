package implementtrie

import (
	"testing"
)

func TestImplementTrie(t *testing.T) {
	word := "apple"
	prefix := "app"
	obj := Constructor()
	obj.Insert(word)
	if res := obj.Search(word); !res {
		t.Errorf("Output: %t: Expected: true\n", res)
	}
	if res := obj.StartsWith(prefix); !res {
		t.Errorf("Output: %t: Expected: true\n", res)
	}
}
