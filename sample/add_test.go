package sample

import "testing"

func Test3Plus2ShouldBe5(t *testing.T) {
	actual := Add(3, 2)
	expected := 5
	if actual != expected {
		t.Errorf("Add(3, 2) = %d; expected %d", actual, expected)
	}
}
