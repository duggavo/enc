package enc

import (
	"encoding/base64"
	"errors"
)

type B64 []byte

func (m B64) String() string {
	return base64.StdEncoding.EncodeToString(m)
}
func (m *B64) UnmarshalText(c []byte) error {
	dst := make([]byte, base64.StdEncoding.EncodedLen(len(c)))

	n, err := base64.StdEncoding.Decode(dst, c)

	*m = append((*m)[0:0], dst[:n]...)

	return err
}

func (m B64) MarshalJSON() ([]byte, error) {
	return []byte(`"` + base64.StdEncoding.EncodeToString(m) + `"`), nil
}
func (m *B64) UnmarshalJSON(c []byte) error {
	if c == nil || len(c) < 2 {
		return errors.New("value is too short")
	} else if len(c) == 2 {
		*m = append((*m)[0:0], []byte{}...)
		return nil
	}

	if c[0] != '"' || c[len(c)-1] != '"' {
		return errors.New("invalid string literal")
	}

	dst := make([]byte, base64.StdEncoding.EncodedLen(len(c)))

	n, err := base64.StdEncoding.Decode(dst, c[1:len(c)-1])

	*m = append((*m)[0:0], dst[:n]...)

	return err
}
