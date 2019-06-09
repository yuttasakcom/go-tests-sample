# Go Tests Sample

## Simple Test

```go
package main

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		lhs      int
		rhs      int
		expected int
	}{
		{lhs: 1, rhs: 2, expected: 3},
	}

	for _, test := range tests {
		ans := add(test.lhs, test.rhs)
		if test.expected != ans {
			t.Errorf("add(%d, %d) = %d, Wanted = %d", test.lhs, test.rhs, ans, test.expected)
		}
	}
}

```

## Run Test

```bash
go test # -v -run="Regx"
```
