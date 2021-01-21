// Package sample generated sample data
package sample

import "github.com/Akshit8/go-grpc/pb"

// NewKeyboard returns a new sample keyboard
func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout: randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

// NewCPU returns a new sample cpu
func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	cpu := &pb.CPU{
		Brand: brand,
	}
	return cpu
}
