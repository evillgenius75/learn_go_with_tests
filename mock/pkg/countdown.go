package countdown

import (
	"bytes"
	"fmt"
)

func Countdown(out *bytes.Buffer) {
	fmt.Fprintf(out, "3")
}
