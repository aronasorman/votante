package main

import (
	"encoding/binary"
	"testing"
	"testing/quick"
)

func TestDeviceIsHashable(t *testing.T) {
	f := func(d *Device) []byte {
		idBinary := make([]byte, 32)
		binary.PutUvarint(idBinary, d.Id)

		return DoubleSha256(idBinary)
	}

	g := func(d *Device) []byte {
		return d.Hash()
	}

	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
