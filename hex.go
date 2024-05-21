package enc

import (
	"encoding/hex"
	"errors"
)

type Hex []byte

func (m Hex) String() string {
	return hex.EncodeToString(m)
}
func (m *Hex) UnmarshalText(c []byte) error {
	dst := make([]byte, hex.EncodedLen(len(c)))

	n, err := hex.Decode(dst, c)

	*m = append((*m)[0:0], dst[:n]...)

	return err
}

func (m Hex) MarshalJSON() ([]byte, error) {
	return []byte(`"` + hex.EncodeToString(m) + `"`), nil
}
func (m *Hex) UnmarshalJSON(c []byte) error {
	if c == nil || len(c) < 2 {
		return errors.New("value is too short")
	} else if len(c) == 2 {
		*m = append((*m)[0:0], []byte{}...)
		return nil
	}

	if c[0] != '"' || c[len(c)-1] != '"' {
		return errors.New("invalid string literal")
	}

	dst := make([]byte, hex.EncodedLen(len(c)))

	n, err := hex.Decode(dst, c[1:len(c)-1])

	*m = append((*m)[0:0], dst[:n]...)

	return err
}
