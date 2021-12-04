package depalpha_test

import (
	"testing"

	"github.com/onsi/depa/b"
	"github.com/onsi/depalpha"
)

func TestDepAlpha(t *testing.T) {
	if depalpha.Alpha()+b.B() != "AlphaB" {
		t.Fatal("bloop")
	}
}
