package sample

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorpc/src/proto"
)

// NewKeyBoard returns a new sample keyboard
func NewKeyBoard() *proto.KeyBoard {
	layout := RandomKeyBoardLayout()
	backlight := RandomBool()
	return &proto.KeyBoard{Layout: layout, Backlight: backlight}
}

// NewCpu returns a new cpu
func NewCpu() *proto.CPU {
	brand := RandomBrand()
	cores := RandomInt(2, 16)
	minGhz := RandomFloat64(2.0, 3.5)
	maxGhz := RandomFloat64(minGhz, 5.0)
	return &proto.CPU{Brand: brand, Name: RandomCpuName(brand),
		NumberCores: uint32(cores), NumberThreads: uint32(RandomInt(cores, 12)),
		MinGhz: minGhz, MaxGhz: maxGhz,
	}

}

func NewGpu() *proto.GPU {
	brand := RandomBrand()
	minGhz := RandomFloat64(2.0, 3.5)
	maxGhz := RandomFloat64(minGhz, 5.0)
	return &proto.GPU{Name: RandomCpuName(brand), MinGhz: minGhz, MaxGhz: maxGhz, Brand: brand,
		Memory: RandomMemory(proto.Memory_GIGABYTE, 2, 16)}
}

func NewMemory() *proto.Memory {
	unit := proto.Memory_GIGABYTE
	return RandomMemory(unit, 2, 32)
}

func NewSSd() *proto.Storage {
	driverTye := proto.Storage_SSD
	memUnit := proto.Memory_GIGABYTE
	return &proto.Storage{DriverType: driverTye, MemorySize: RandomMemory(memUnit, 250, 2000)}

}

func NewHdd() *proto.Storage {
	driverTye := proto.Storage_HDD
	memUnit := proto.Memory_TERABYTE
	return &proto.Storage{DriverType: driverTye, MemorySize: RandomMemory(memUnit, 1, 6)}

}

func NewScreen() *proto.Screen {
	height := RandomInt(1000, 4320)
	width := height * 16 / 9
	resolution := &proto.Screen_Resolution{Height: uint32(height), Width: uint32(width)}
	size := RandomFloat32(13, 17)
	panel := RandomScreenPanel()
	multiTouch := RandomBool()
	return &proto.Screen{SizeInch: size, Resolution: resolution, Panel: panel, MultiTouch: multiTouch}
}

func NewLaptop() *proto.Laptop {
	id := RandomUUID()
	cpu := NewCpu()
	gpus := []*proto.GPU{NewGpu()}
	brand := RandomLaptopBrand()
	storages := []*proto.Storage{NewHdd(), NewSSd()}
	weight := &proto.Laptop_WeightKg{
		WeightKg: RandomFloat32(1.0, 3.0),
	}
	releaseYear := uint32(RandomInt(2015, 2019))
	lastUpdate := timestamppb.Now()
	price := RandomFloat64(300, 2000)
	return &proto.Laptop{Id: id, Brand: brand, Name: RandomLaptopName(brand),
		Screen: NewScreen(), Ram: NewMemory(), Cpu: cpu, Gpus: gpus, Keyboard: NewKeyBoard(),
		Storages: storages, Weight: weight, ReleaseYear: releaseYear, UpdatedAt: lastUpdate, Price: price}
}
