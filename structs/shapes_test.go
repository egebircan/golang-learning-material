package structs

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{5, 7})
	want := 24.0

	// Notice the new format string? The f is for our float64 and the .2 means print 2 decimal places.
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		shape Shape
		want  float64
	}{
		// The test speaks to us more clearly, as if it were an assertion of truth, not a sequence of operations
		// so construct with names of the fields
		{Rectangle{Width: 12, Height: 6}, 72.0},
		{Circle{Radius: 10}, 314.1592653589793},
		{Triangle{Base: 12, H: 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}

func TestAreaV2(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, H: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name

		// One final tip with table driven tests is to use t.Run and to name the test cases.
		// By wrapping each case in a t.Run you will have clearer test output on failures as it will print the name of the case
		// And you can run specific tests within your table with go test -run TestArea/Rectangle
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				// We can change our error message into %#v got %g want %g.
				// The %#v format string will print out our struct with the values in its field,
				// so the developer can see at a glance the properties that are being tested.
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}
}
