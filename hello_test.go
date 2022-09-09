package main

import "testing"

// It needs to be in a file with a name like xxx_test.go
// The test function must start with the word Test
// The test function takes one argument only t *testing.T
// In order to use the *testing.T type, you need to import "testing", like we did with fmt in the other file
func TestHello(t *testing.T) {
	// For helper functions, it's a good idea to accept a testing.TB which is an interface that *testing.T and *testing.B both satisfy,
	//so you can call helper functions from a test, or a benchmark.
	assertCorrectMessage := func(t testing.TB, got, want string) {
		//t.Helper() is needed to tell the test suite that this method is a helper.
		//By doing this when it fails the line number reported will be in our function call rather than inside our test helper.
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Elodie", "French")
		want := "Bonjour, Elodie"
		assertCorrectMessage(t, got, want)
	})
}
