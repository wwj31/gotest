package escape

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	var i = new(int)
	fmt.Println(i) // i发生逃逸，fmt.Println调用了refect.Valueof()
}
