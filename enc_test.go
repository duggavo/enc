package enc

import (
	"bytes"
	"testing"
)

func TestB64(t *testing.T) {
	x := B64("test value")

	str := x.String()

	t.Logf("base64 value is %s", str)

	x2 := B64{}

	err := x2.UnmarshalText([]byte(str))

	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(x, x2) {
		t.Errorf("base64 values do not match: %x, %x", x, x2)
	}
}

func TestHex(t *testing.T) {
	x := Hex("test value")

	str := x.String()

	t.Logf("hex value is %s", str)

	x2 := Hex{}

	err := x2.UnmarshalText([]byte(str))

	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(x, x2) {
		t.Errorf("hex values do not match: %x, %x", x, x2)
	}
}
