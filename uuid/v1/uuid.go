package uuid

import (
	"encoding/binary"
	"errors"

	"github.com/google/uuid"
)

func New() *UUID { return (&UUID{}).FromUUID(uuid.New()) }

func NewUUID() (*UUID, error) {
	val, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}
	return (&UUID{}).FromUUID(val), nil
}

func (x *UUID) FromUUID(u uuid.UUID) *UUID {
	v := x
	if v == nil {
		v = new(UUID)
	}
	v.Hi, v.Lo = binary.BigEndian.Uint64(u[:8]), binary.BigEndian.Uint64(u[8:])
	return v
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
