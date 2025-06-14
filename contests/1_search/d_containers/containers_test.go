package containers

import (
	"fmt"
	"testing"
)

func TestContainers(t *testing.T) {
	tests := []struct {
		n, k int64
		w    []int64
		exp  int64
	}{
		{5, 3, []int64{1, 2, 3, 4, 5}, 6},
		{4, 2, []int64{7, 2, 5, 10}, 14},
		{6, 4, []int64{9, 8, 7, 6, 5, 4}, 9},
	}

	for i, tt := range tests {
		res := Containers(tt.n, tt.k, tt.w)
		if res != tt.exp {
			fmt.Println("test", i+1, "failed: expected", tt.exp, "got", res)
		} else {
			fmt.Println("test", i+1, "passed")
		}
	}
}
