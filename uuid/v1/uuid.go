package uuid

import (
	"encoding/binary"
	"errors"

	"github.com/google/uuid"
)

func (x *UUID) FromUUID(u uuid.UUID) {
	x.Hi, x.Lo = binary.BigEndian.Uint64(u[:8]), binary.BigEndian.Uint64(u[8:])
}

func (x *UUID) ToUUID() (uuid.UUID, error) {
	b := x.ToBytes()
	if b == nil {
		return uuid.UUID{}, errors.New("invalid UUID")
	}
	return uuid.FromBytes(b)
}

func (x *UUID) ToBytes() []byte {
	if x == nil {
		return nil
	}
	return []byte{
		byte(x.Hi >> 56), byte(x.Hi >> 48), byte(x.Hi >> 40), byte(x.Hi >> 32),
		byte(x.Hi >> 24), byte(x.Hi >> 16), byte(x.Hi >> 8), byte(x.Hi),

		byte(x.Lo >> 56), byte(x.Lo >> 48), byte(x.Lo >> 40), byte(x.Lo >> 32),
		byte(x.Lo >> 24), byte(x.Lo >> 16), byte(x.Lo >> 8), byte(x.Lo),
	}
}
