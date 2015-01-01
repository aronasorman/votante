package main

import (
	"github.com/satori/go.uuid"
	"math/rand"
	"reflect"
)

type Device struct {
	Id uuid.UUID
}

func (d *Device) Generate(rand *rand.Rand, size int) reflect.Value {
	d = &Device{Id: uuid.NewV4()}

	return reflect.ValueOf(d)
}

func (d *Device) ComputeHash() []byte {
	return DoubleSha256([]byte(d.Id.String()))
}
