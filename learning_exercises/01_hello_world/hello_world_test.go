package greeting

import "testing"

func TestHelloWorld(t *testing.T) {
	want := "Hello, World!"
	got := HelloWorld()

	if want != got {
		t.Errorf("want %s, got %s", want, got)
	}
}
