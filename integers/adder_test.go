package integers

import (
	"fmt"
	"testing"
)

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// Often code examples that can be found outside the codebase,
// such as a readme file often become out of date and incorrect compared to the actual code because they don't get checked.
//
// Go examples are executed just like tests so you can be confident examples reflect what the code actually does.
//
// Examples are compiled (and optionally executed) as part of a package's test suite.
//
// As with typical tests, examples are functions that reside in a package's _test.go files.

// Please note that the example function will not be executed if you remove the comment
// Output: 6. Although the function will be compiled, it won't be executed.

// By adding this code the example will appear in the documentation inside godoc, making your code even more accessible.
// To try this out, run godoc -http=:8000 and navigate to http://localhost:8000/pkg/
func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
