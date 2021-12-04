package depalpha

import "github.com/onsi/depa/b"

const ALPHA_VERSION = "1.2.0"

func Alpha() string {
	return "Alpha" + b.B()
}
