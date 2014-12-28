package main

import (
	"testing"
	"testing/quick"
)

func TestDeviceIsHashable(t *testing.T) {
	f := func(d *Device) []byte {
		return nil
	}

	g := func(d *Device) []byte {
		return nil
	}

	if err := quick.CheckEqual(f, g, nil); err != nil {
		t.Error(err)
	}
}
