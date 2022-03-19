package escape

import "testing"

func TestPointer(_ *testing.T) {
	var i int
	testPointer(&i)

}

var abcd = 123

func testPointer(i *int) *int {
	return &abcd
}
