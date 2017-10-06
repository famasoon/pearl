package ntor

import (
	"testing"

	"github.com/mmcloughlin/pearl/torkeys"
	"github.com/stretchr/testify/require"
)

func TestServerHandshake(t *testing.T) {
	h := &ServerHandshake{
		ClientPK: []byte{
			0x6a, 0xd8, 0xef, 0xd7, 0xf0, 0x5e, 0x8f, 0x18, 0x80, 0xb8, 0xd6, 0xbd, 0xb9, 0xf8, 0x4a, 0xc9,
			0xb5, 0x44, 0x43, 0xde, 0x8b, 0xd0, 0x5c, 0x97, 0x45, 0xf8, 0xc8, 0x59, 0x89, 0xeb, 0xd6, 0x45,
		},
		ServerFingerprint: []byte{
			0x69, 0x54, 0x6f, 0x6c, 0x64, 0x59, 0x6f, 0x75, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x53, 0x74, 0x61,
			0x69, 0x72, 0x73, 0x2e,
		},
		ServerKeyPair: &torkeys.Curve25519KeyPair{
			Private: [32]byte{
				0xb8, 0x55, 0xce, 0x59, 0xea, 0xb7, 0x49, 0x66, 0xe9, 0x64, 0xf9, 0x1f, 0xd5, 0x0d, 0xd4, 0x42,
				0xdb, 0x2d, 0x6c, 0xaa, 0x12, 0xaa, 0x00, 0x6c, 0x52, 0xf9, 0x3f, 0xc8, 0x67, 0xd1, 0x2e, 0x5a,
			},
			Public: [32]byte{
				0x16, 0x1d, 0xe2, 0x87, 0xea, 0x55, 0xde, 0x83, 0x1f, 0xc2, 0x99, 0x9a, 0x4f, 0x59, 0x24, 0xcd,
				0xd5, 0xab, 0xde, 0x20, 0xee, 0x05, 0x24, 0x94, 0xc3, 0xe3, 0x9a, 0x80, 0x49, 0x09, 0xc7, 0x40,
			},
		},
		ServerNTORKey: &torkeys.Curve25519KeyPair{
			Private: [32]byte{
				0xb0, 0x7d, 0xd8, 0xfe, 0x3e, 0x88, 0x6f, 0x9d, 0x1d, 0xf8, 0xe5, 0x6a, 0xc5, 0x6b, 0xb8, 0x8a,
				0x00, 0x4b, 0x46, 0xe8, 0x31, 0x95, 0x14, 0xc1, 0xde, 0x85, 0xac, 0x58, 0x7a, 0xb3, 0x25, 0x57,
			},
			Public: [32]byte{
				0xb7, 0x7b, 0x71, 0xa8, 0x71, 0x75, 0xd5, 0x24, 0x3f, 0xe3, 0xc1, 0xe6, 0x3f, 0xff, 0x00, 0xee,
				0xba, 0xdb, 0xeb, 0x00, 0xce, 0xcc, 0xc8, 0x49, 0x75, 0x26, 0xf0, 0xdf, 0x26, 0xc0, 0xf8, 0x26,
			},
		},
	}

	expectSecretInput := []byte{
		0x26, 0xa7, 0x78, 0x2d, 0x22, 0xb5, 0xcc, 0x22, 0xc9, 0x0f, 0x00, 0xa6, 0x69, 0x6a, 0x16, 0x22,
		0x22, 0x27, 0x48, 0xce, 0x0b, 0x85, 0x85, 0x47, 0xc6, 0xe2, 0xa7, 0xad, 0x8b, 0x65, 0x98, 0x4f,
		0x5e, 0x58, 0xfd, 0x34, 0x79, 0x86, 0xec, 0x4c, 0xdd, 0xc9, 0xf0, 0x18, 0xe8, 0x51, 0x0e, 0xe6,
		0x46, 0x3e, 0x5f, 0xd3, 0x1e, 0xb0, 0x1f, 0x54, 0x82, 0x1d, 0x87, 0x07, 0x9a, 0xea, 0xc1, 0x50,
		0x69, 0x54, 0x6f, 0x6c, 0x64, 0x59, 0x6f, 0x75, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x53, 0x74, 0x61,
		0x69, 0x72, 0x73, 0x2e, 0xb7, 0x7b, 0x71, 0xa8, 0x71, 0x75, 0xd5, 0x24, 0x3f, 0xe3, 0xc1, 0xe6,
		0x3f, 0xff, 0x00, 0xee, 0xba, 0xdb, 0xeb, 0x00, 0xce, 0xcc, 0xc8, 0x49, 0x75, 0x26, 0xf0, 0xdf,
		0x26, 0xc0, 0xf8, 0x26, 0x6a, 0xd8, 0xef, 0xd7, 0xf0, 0x5e, 0x8f, 0x18, 0x80, 0xb8, 0xd6, 0xbd,
		0xb9, 0xf8, 0x4a, 0xc9, 0xb5, 0x44, 0x43, 0xde, 0x8b, 0xd0, 0x5c, 0x97, 0x45, 0xf8, 0xc8, 0x59,
		0x89, 0xeb, 0xd6, 0x45, 0x16, 0x1d, 0xe2, 0x87, 0xea, 0x55, 0xde, 0x83, 0x1f, 0xc2, 0x99, 0x9a,
		0x4f, 0x59, 0x24, 0xcd, 0xd5, 0xab, 0xde, 0x20, 0xee, 0x05, 0x24, 0x94, 0xc3, 0xe3, 0x9a, 0x80,
		0x49, 0x09, 0xc7, 0x40, 0x6e, 0x74, 0x6f, 0x72, 0x2d, 0x63, 0x75, 0x72, 0x76, 0x65, 0x32, 0x35,
		0x35, 0x31, 0x39, 0x2d, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x2d, 0x31,
	}
	require.Equal(t, expectSecretInput, h.SecretInput())

	expectVerify := []byte{
		0x38, 0x85, 0xf9, 0x27, 0x2d, 0x32, 0x24, 0xdf, 0x55, 0xe8, 0x95, 0xcd, 0xed, 0x37, 0x4b, 0x92,
		0x80, 0x25, 0x78, 0x64, 0x5b, 0x04, 0xbc, 0x0a, 0x2a, 0x26, 0x46, 0x2a, 0x33, 0xa8, 0xb1, 0xe5,
	}
	require.Equal(t, expectVerify, h.Verify())

	expectAuthInput := []byte{
		0x38, 0x85, 0xf9, 0x27, 0x2d, 0x32, 0x24, 0xdf, 0x55, 0xe8, 0x95, 0xcd, 0xed, 0x37, 0x4b, 0x92,
		0x80, 0x25, 0x78, 0x64, 0x5b, 0x04, 0xbc, 0x0a, 0x2a, 0x26, 0x46, 0x2a, 0x33, 0xa8, 0xb1, 0xe5,
		0x69, 0x54, 0x6f, 0x6c, 0x64, 0x59, 0x6f, 0x75, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x53, 0x74, 0x61,
		0x69, 0x72, 0x73, 0x2e, 0xb7, 0x7b, 0x71, 0xa8, 0x71, 0x75, 0xd5, 0x24, 0x3f, 0xe3, 0xc1, 0xe6,
		0x3f, 0xff, 0x00, 0xee, 0xba, 0xdb, 0xeb, 0x00, 0xce, 0xcc, 0xc8, 0x49, 0x75, 0x26, 0xf0, 0xdf,
		0x26, 0xc0, 0xf8, 0x26, 0x16, 0x1d, 0xe2, 0x87, 0xea, 0x55, 0xde, 0x83, 0x1f, 0xc2, 0x99, 0x9a,
		0x4f, 0x59, 0x24, 0xcd, 0xd5, 0xab, 0xde, 0x20, 0xee, 0x05, 0x24, 0x94, 0xc3, 0xe3, 0x9a, 0x80,
		0x49, 0x09, 0xc7, 0x40, 0x6a, 0xd8, 0xef, 0xd7, 0xf0, 0x5e, 0x8f, 0x18, 0x80, 0xb8, 0xd6, 0xbd,
		0xb9, 0xf8, 0x4a, 0xc9, 0xb5, 0x44, 0x43, 0xde, 0x8b, 0xd0, 0x5c, 0x97, 0x45, 0xf8, 0xc8, 0x59,
		0x89, 0xeb, 0xd6, 0x45, 0x6e, 0x74, 0x6f, 0x72, 0x2d, 0x63, 0x75, 0x72, 0x76, 0x65, 0x32, 0x35,
		0x35, 0x31, 0x39, 0x2d, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x2d, 0x31, 0x53, 0x65, 0x72, 0x76,
		0x65, 0x72,
	}
	require.Equal(t, expectAuthInput, h.AuthInput())

	expectAuth := []byte{
		0xa0, 0xe4, 0xae, 0xfb, 0xb1, 0x09, 0xd9, 0xad, 0x76, 0xee, 0x9e, 0x74, 0x76, 0x74, 0x5e, 0x24,
		0x33, 0x08, 0x77, 0x02, 0x54, 0x2d, 0xa8, 0xc3, 0x0e, 0x95, 0x8f, 0x81, 0x0b, 0x5c, 0x24, 0x3b,
	}
	require.Equal(t, expectAuth, h.Auth())
}

func TestClientHandshake(t *testing.T) {
	h := &ClientHandshake{
		Kx: [32]byte{
			0x68, 0x85, 0x77, 0xf7, 0x7b, 0xdc, 0xa9, 0x62, 0xf0, 0x0b, 0x4b, 0x9f, 0x82, 0x64, 0xb0, 0x67,
			0x77, 0xad, 0x7e, 0x64, 0x85, 0xe8, 0xae, 0x14, 0xa3, 0x6a, 0x00, 0x50, 0x60, 0xc8, 0xb5, 0x41,
		},
		KX: [32]byte{
			0x1b, 0x21, 0x9c, 0x2e, 0x47, 0xd6, 0x7e, 0x4e, 0xf2, 0x32, 0x37, 0xda, 0x32, 0xfa, 0xde, 0x87,
			0x40, 0x27, 0x04, 0xcc, 0x56, 0xc1, 0x95, 0x18, 0x88, 0xa1, 0xb7, 0x9d, 0xc4, 0x80, 0x87, 0x2a,
		},
		KY: [32]byte{
			0x29, 0x26, 0xf3, 0xcb, 0x94, 0x7e, 0x4c, 0x71, 0x78, 0x6e, 0xa3, 0x9e, 0x4b, 0xe6, 0x3b, 0x6c,
			0x79, 0x51, 0xb4, 0x99, 0xf3, 0xe5, 0xcf, 0x20, 0x7e, 0x20, 0xf1, 0x48, 0xb6, 0x3e, 0xab, 0x24,
		},
		KB: [32]byte{
			0x18, 0x2d, 0x19, 0x6c, 0xb8, 0x12, 0xc2, 0xa1, 0xb2, 0xd3, 0x5e, 0x9c, 0x6e, 0x8b, 0x67, 0x69,
			0xe6, 0x76, 0xcb, 0x33, 0xde, 0xaa, 0xc4, 0xaa, 0xd4, 0x04, 0x45, 0x6f, 0x90, 0xf1, 0xa0, 0x14,
		},
		ID: []byte{
			0x69, 0x54, 0x6f, 0x6c, 0x64, 0x59, 0x6f, 0x75, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x53, 0x74, 0x61,
			0x69, 0x72, 0x73, 0x2e,
		},
	}

	expectSecretInput := []byte{
		0x19, 0xa5, 0x32, 0x39, 0x90, 0x62, 0x65, 0x42, 0xd9, 0x4d, 0x92, 0x87, 0x44, 0xde, 0x90, 0x86,
		0x9c, 0xca, 0x1c, 0xd4, 0xca, 0x82, 0xc1, 0xb8, 0xa7, 0x45, 0x86, 0xf1, 0xa2, 0xee, 0x6a, 0x0b,
		0xb9, 0xab, 0xe4, 0x85, 0xe9, 0x3a, 0x7f, 0x7a, 0x29, 0x7f, 0xbf, 0xbf, 0xc8, 0x1c, 0x1b, 0x9f,
		0xcb, 0xe5, 0x89, 0x70, 0x27, 0x3a, 0x80, 0xa8, 0x5e, 0x2f, 0x0e, 0x7c, 0xa5, 0xe7, 0x6e, 0x58,
		0x69, 0x54, 0x6f, 0x6c, 0x64, 0x59, 0x6f, 0x75, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x53, 0x74, 0x61,
		0x69, 0x72, 0x73, 0x2e, 0x18, 0x2d, 0x19, 0x6c, 0xb8, 0x12, 0xc2, 0xa1, 0xb2, 0xd3, 0x5e, 0x9c,
		0x6e, 0x8b, 0x67, 0x69, 0xe6, 0x76, 0xcb, 0x33, 0xde, 0xaa, 0xc4, 0xaa, 0xd4, 0x04, 0x45, 0x6f,
		0x90, 0xf1, 0xa0, 0x14, 0x1b, 0x21, 0x9c, 0x2e, 0x47, 0xd6, 0x7e, 0x4e, 0xf2, 0x32, 0x37, 0xda,
		0x32, 0xfa, 0xde, 0x87, 0x40, 0x27, 0x04, 0xcc, 0x56, 0xc1, 0x95, 0x18, 0x88, 0xa1, 0xb7, 0x9d,
		0xc4, 0x80, 0x87, 0x2a, 0x29, 0x26, 0xf3, 0xcb, 0x94, 0x7e, 0x4c, 0x71, 0x78, 0x6e, 0xa3, 0x9e,
		0x4b, 0xe6, 0x3b, 0x6c, 0x79, 0x51, 0xb4, 0x99, 0xf3, 0xe5, 0xcf, 0x20, 0x7e, 0x20, 0xf1, 0x48,
		0xb6, 0x3e, 0xab, 0x24, 0x6e, 0x74, 0x6f, 0x72, 0x2d, 0x63, 0x75, 0x72, 0x76, 0x65, 0x32, 0x35,
		0x35, 0x31, 0x39, 0x2d, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x2d, 0x31,
	}
	require.Equal(t, expectSecretInput, h.SecretInput())

	/*
	   // verify=
	   []byte{
	   0x08, 0x63, 0x23, 0x05, 0xa2, 0xee, 0x6f, 0x83, 0xe4, 0x9a, 0xcd, 0x18, 0xaa, 0x5b, 0x7f, 0x8e,
	   0x96, 0x6f, 0x99, 0x10, 0xd5, 0xac, 0xa7, 0xa4, 0xf7, 0x3c, 0x30, 0xc3, 0xcc, 0x5a, 0xc7, 0x02,
	   }
	   // auth_input=
	   []byte{
	   0x08, 0x63, 0x23, 0x05, 0xa2, 0xee, 0x6f, 0x83, 0xe4, 0x9a, 0xcd, 0x18, 0xaa, 0x5b, 0x7f, 0x8e,
	   0x96, 0x6f, 0x99, 0x10, 0xd5, 0xac, 0xa7, 0xa4, 0xf7, 0x3c, 0x30, 0xc3, 0xcc, 0x5a, 0xc7, 0x02,
	   0x69, 0x54, 0x6f, 0x6c, 0x64, 0x59, 0x6f, 0x75, 0x41, 0x62, 0x6f, 0x75, 0x74, 0x53, 0x74, 0x61,
	   0x69, 0x72, 0x73, 0x2e, 0x18, 0x2d, 0x19, 0x6c, 0xb8, 0x12, 0xc2, 0xa1, 0xb2, 0xd3, 0x5e, 0x9c,
	   0x6e, 0x8b, 0x67, 0x69, 0xe6, 0x76, 0xcb, 0x33, 0xde, 0xaa, 0xc4, 0xaa, 0xd4, 0x04, 0x45, 0x6f,
	   0x90, 0xf1, 0xa0, 0x14, 0x29, 0x26, 0xf3, 0xcb, 0x94, 0x7e, 0x4c, 0x71, 0x78, 0x6e, 0xa3, 0x9e,
	   0x4b, 0xe6, 0x3b, 0x6c, 0x79, 0x51, 0xb4, 0x99, 0xf3, 0xe5, 0xcf, 0x20, 0x7e, 0x20, 0xf1, 0x48,
	   0xb6, 0x3e, 0xab, 0x24, 0x1b, 0x21, 0x9c, 0x2e, 0x47, 0xd6, 0x7e, 0x4e, 0xf2, 0x32, 0x37, 0xda,
	   0x32, 0xfa, 0xde, 0x87, 0x40, 0x27, 0x04, 0xcc, 0x56, 0xc1, 0x95, 0x18, 0x88, 0xa1, 0xb7, 0x9d,
	   0xc4, 0x80, 0x87, 0x2a, 0x6e, 0x74, 0x6f, 0x72, 0x2d, 0x63, 0x75, 0x72, 0x76, 0x65, 0x32, 0x35,
	   0x35, 0x31, 0x39, 0x2d, 0x73, 0x68, 0x61, 0x32, 0x35, 0x36, 0x2d, 0x31, 0x53, 0x65, 0x72, 0x76,
	   0x65, 0x72,
	   }
	   // auth=
	   []byte{
	   0x6c, 0x13, 0x07, 0x98, 0xed, 0x87, 0x1c, 0x7a, 0xc7, 0xa2, 0x9f, 0x07, 0x10, 0xff, 0x41, 0x1c,
	   0x37, 0x59, 0xe2, 0x98, 0xc4, 0x0e, 0x3a, 0x10, 0x1c, 0x82, 0xda, 0x54, 0xe4, 0x4d, 0x9a, 0xb8,
	   }
	*/
}
