package integers

import (
	"fmt"
	"testing"
)

func TestAddition(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

func ExampleAddition() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}