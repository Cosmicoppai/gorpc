package sample

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorpc/pb"
)

// NewKeyBoard returns a new sample keyboard
func NewKeyBoard() *pb.KeyBoard {
	layout := RandomKeyBoardLayout()
	backlight := RandomBool()
	return &pb.KeyBoard{Layout: layout, Backlight: backlight}
}

// NewCpu returns a new cpu
func NewCpu() *pb.CPU {
	brand := RandomBrand()
	cores := RandomInt(2, 16)
	minGhz := RandomFloat64(2.0, 3.5)
	maxGhz := RandomFloat64(minGhz, 5.0)
	return &pb.CPU{Brand: brand, Name: RandomCpuName(brand),
		NumberCores: uint32(cores), NumberThreads: uint32(RandomInt(cores, 12)),
		MinGhz: minGhz, MaxGhz: maxGhz,
	}

}

func NewGpu() *pb.GPU {
	brand := RandomBrand()
	minGhz := RandomFloat64(2.0, 3.5)
	maxGhz := RandomFloat64(minGhz, 5.0)
	return &pb.GPU{Name: RandomCpuName(brand), MinGhz: minGhz, MaxGhz: maxGhz, Brand: brand,
		Memory: RandomMemory(pb.Memory_GIGABYTE, 2, 16)}
}

func NewMemory() *pb.Memory {
	unit := pb.Memory_GIGABYTE
	return RandomMemory(unit, 2, 32)
}

func NewSSd() *pb.Storage {
	driverTye := pb.Storage_SSD
	memUnit := pb.Memory_GIGABYTE
	return &pb.Storage{DriverType: driverTye, MemorySize: RandomMemory(memUnit, 250, 2000)}

}

func NewHdd() *pb.Storage {
	driverTye := pb.Storage_HDD
	memUnit := pb.Memory_TERABYTE
	return &pb.Storage{DriverType: driverTye, MemorySize: RandomMemory(memUnit, 1, 6)}

}

func NewScreen() *pb.Screen {
	height := RandomInt(1000, 4320)
	width := height * 16 / 9
	resolution := &pb.Screen_Resolution{Height: uint32(height), Width: uint32(width)}
	size := RandomFloat32(13, 17)
	panel := RandomScreenPanel()
	multiTouch := RandomBool()
	return &pb.Screen{SizeInch: size, Resolution: resolution, Panel: panel, MultiTouch: multiTouch}
}

func NewLaptop() *pb.Laptop {
	id := RandomUUID()
	cpu := NewCpu()
	gpus := []*pb.GPU{NewGpu()}
	brand := RandomLaptopBrand()
	storages := []*pb.Storage{NewHdd(), NewSSd()}
	weight := &pb.Laptop_WeightKg{
		WeightKg: RandomFloat32(1.0, 3.0),
	}
	releaseYear := uint32(RandomInt(2015, 2019))
	lastUpdate := timestamppb.Now()
	price := RandomFloat64(300, 2000)
	return &pb.Laptop{Id: id, Brand: brand, Name: RandomLaptopName(brand),
		Screen: NewScreen(), Ram: NewMemory(), Cpu: cpu, Gpus: gpus, Keyboard: NewKeyBoard(),
		Storages: storages, Weight: weight, ReleaseYear: releaseYear, UpdatedAt: lastUpdate, Price: price}
}
