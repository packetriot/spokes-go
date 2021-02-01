package spokes

import (
	"bytes"
	"encoding/hex"
	"fmt"

	"github.com/satori/go.uuid"
)

const (
	UIDLen = 16
)

type UID [UIDLen]byte

func NewUID() UID {
	uid := uuid.NewV4()
	return UID(uid)
}

func UIDFromString(s string) UID {
	uid := UID{}
	uid.Set(s)
	return uid
}

func UIDFromBytes(b []byte) (UID, error) {
	if len(b) != UIDLen {
		return UID{}, fmt.Errorf("Parameter length %d does not match UIDLen const %d",
			len(b), UIDLen)
	}
	uid := UID{}
	copy(uid[0:], b[0:])
	return uid, nil

}

func (this *UID) Bytes() []byte {
	return this[0:]
}

func (this *UID) String() string {
	return fmt.Sprintf("%x", this[0:])
}

func (this *UID) Set(s string) error {
	b, err := hex.DecodeString(s)
	if err == nil {
		copy(this[0:], b[0:])
		return nil
	}
	return err
}

func (this *UID) Equals(uid UID) bool {
	return (bytes.Compare(this[0:], uid[0:]) == 0)
}

func (this *UID) EqualsHex(hex string) bool {
	uid := UID{}
	uid.Set(hex)
	return this.Equals(uid)
}

func (this *UID) IsZero() bool {
	for _, b := range this {
		if b > 0 {
			return false
		}
	}
	return true
}

// implementing the Marshaler interface
func (this *UID) MarshalJSON() ([]byte, error) {
	// have to append quotes so the encoder writes it as a string
	return []byte("\"" + this.String() + "\""), nil
}

// implementing the Unmarshaler interface
func (this *UID) UnmarshalJSON(b []byte) error {
	n := len(b)
	s := string(b[1 : n-1]) // take off quotes b/c we don't want it
	tuid := UIDFromString(s)
	*this = tuid
	return nil
}
