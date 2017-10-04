package pearl

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAuthMethodString(t *testing.T) {
	assert.Equal(t, "AUTH0003", AuthMethodEd25519SHA256RFC5705.String())
}

func TestAuthenticateCellRoundTrip(t *testing.T) {
	data := []byte{
		0x00, 0x00, 0x00, 0x00, 0x83, 0x01, 0x64, 0x00, 0x01, 0x01, 0x60, 0x41,
		0x55, 0x54, 0x48, 0x30, 0x30, 0x30, 0x31, 0x81, 0xe4, 0x71, 0x36, 0x1d,
		0x86, 0x96, 0x47, 0x49, 0x72, 0x0f, 0x6e, 0x79, 0x00, 0x0d, 0xfa, 0xa8,
		0x8f, 0x83, 0x4a, 0x67, 0x41, 0x5c, 0xc0, 0x34, 0xfe, 0xa2, 0xc1, 0x35,
		0xe7, 0x84, 0xcb, 0x87, 0x2b, 0x11, 0x3e, 0x05, 0x85, 0x31, 0x4a, 0x25,
		0x5c, 0x66, 0x94, 0x11, 0x2a, 0x18, 0xff, 0x1c, 0xcb, 0x2c, 0x5b, 0x40,
		0x26, 0xfc, 0x03, 0x2e, 0x8a, 0xa4, 0x01, 0xad, 0x92, 0xb4, 0x74, 0xb4,
		0xa0, 0xcf, 0xad, 0x1b, 0x16, 0xd5, 0x10, 0xbf, 0x67, 0x23, 0xfb, 0x3e,
		0x7a, 0x88, 0xea, 0x5b, 0x27, 0x13, 0x00, 0x65, 0x73, 0x8c, 0x14, 0xe2,
		0xca, 0x50, 0xc5, 0x6c, 0x3c, 0xa6, 0xa8, 0xbc, 0xc2, 0x41, 0x02, 0x2a,
		0xe3, 0x97, 0x32, 0x51, 0x99, 0xaa, 0xb7, 0x5d, 0x86, 0xd5, 0xc7, 0xe8,
		0x5e, 0x24, 0x22, 0xab, 0x5c, 0xaf, 0xe0, 0x1e, 0x30, 0x96, 0xe3, 0x0f,
		0x27, 0xd6, 0x5b, 0xef, 0x8f, 0x62, 0x83, 0x0a, 0x48, 0x45, 0x57, 0x45,
		0x1c, 0x3b, 0x28, 0x1f, 0x06, 0x6a, 0xbc, 0xf2, 0x38, 0xae, 0x86, 0xde,
		0xbe, 0xeb, 0x04, 0x29, 0x4c, 0xb9, 0x6b, 0x30, 0xe6, 0xad, 0x30, 0x25,
		0x5c, 0x5b, 0xec, 0xa4, 0xc4, 0x72, 0xed, 0xca, 0xbb, 0x65, 0xa8, 0x67,
		0x2c, 0x8a, 0x9d, 0xfa, 0xf0, 0x63, 0x4d, 0xd6, 0xc7, 0x8b, 0x9c, 0xba,
		0xd7, 0x80, 0xc5, 0x1c, 0xc6, 0x0b, 0xbc, 0x8b, 0x91, 0x3d, 0xf8, 0xb0,
		0x81, 0x14, 0x68, 0x3a, 0x56, 0x1a, 0xd2, 0xef, 0xfd, 0xe1, 0xea, 0x4b,
		0xcc, 0xe4, 0xf5, 0xc0, 0x01, 0x5e, 0x75, 0x3e, 0xe2, 0xdc, 0x16, 0x0f,
		0x59, 0x29, 0x30, 0x5d, 0x48, 0x3d, 0x52, 0x0b, 0x99, 0x9b, 0x26, 0x7a,
		0x19, 0x43, 0x7b, 0x73, 0x42, 0xd2, 0x0c, 0x61, 0x65, 0x39, 0xa6, 0x9d,
		0x60, 0x0d, 0xc8, 0x92, 0x40, 0x2d, 0x88, 0xbc, 0x69, 0xfd, 0xad, 0x39,
		0xcf, 0xb1, 0xa8, 0xde, 0x4d, 0x80, 0xb9, 0x5e, 0xd6, 0x2e, 0xd1, 0xb9,
		0x63, 0x97, 0x87, 0x2e, 0xe9, 0xb2, 0xeb, 0x7e, 0x99, 0xaf, 0xf5, 0xe5,
		0xe6, 0x8f, 0xb8, 0xac, 0x8b, 0xd6, 0x03, 0x6f, 0x5e, 0x8b, 0x5b, 0xa7,
		0x8e, 0x7c, 0x86, 0x33, 0x9c, 0xa1, 0x37, 0x8a, 0x5f, 0x0b, 0x0d, 0xbf,
		0x75, 0xff, 0xdc, 0xa2, 0xfd, 0x9d, 0x32, 0x55, 0x65, 0x0c, 0x8a, 0xe6,
		0x6c, 0x4e, 0xfc, 0xfa, 0xaa, 0x29, 0x80, 0x0a, 0x12, 0x93, 0xd6, 0xaa,
		0x89, 0xff, 0xdb, 0x06, 0x65, 0xf0, 0xee, 0xb2, 0x55, 0xfa, 0x84, 0x7a,
		0xdf, 0x65, 0xa7,
	}
	f := CircID4Format{}
	c := NewCellFromBuffer(f, data)
	a, err := ParseAuthenticateCell(c)
	require.NoError(t, err)
	c2, err := a.Cell(f)
	require.NoError(t, err)
	assert.Equal(t, c.Bytes(), c2.Bytes())
}