package sample

import (
	"github.com/google/uuid"
	"gorpc/src/pb"
	"math/rand"
)

func RandomKeyBoardLayout() pb.KeyBoard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.KeyBoard_QWERTY
	case 2:
		return pb.KeyBoard_QWERTZ
	default:
		return pb.KeyBoard_AZERTY
	}
}

func RandomBool() bool {
	return rand.Intn(2) == 1

}

func RandomBrand() string {
	return randomStringsFromSet("INTEL", "AMD", "NVIDIA")

}

func randomStringsFromSet(stringSet ...string) string {
	if len(stringSet) == 0 {
		return ""
	}
	return stringSet[rand.Intn(len(stringSet))]

}

func RandomCpuName(brand string) string {
	return brand + string(rune(rand.Intn(5)))

}

func RandomInt(min, max int) int {
	return min + rand.Int()%(max-min+1)
}

func RandomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func RandomMemory(unit pb.Memory_UNIT, minMem int, maxmem int) *pb.Memory {
	var mem pb.Memory
	if unit > 0 && unit < 7 {
		mem.Unit = unit
	} else {
		mem.Unit = pb.Memory_UNKNOWN
	}
	mem.Value = uint64(RandomInt(minMem, maxmem))
	return &mem
}

func RandomFloat32(min, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func RandomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED

}

func RandomUUID() string {
	return uuid.New().String()
}

func RandomLaptopBrand() string {
	brands := []string{"PSEUDO", "APPLE", "DELL", "HP", "ASUS", "LENOVO"}
	return randomStringsFromSet(brands...)
}

func RandomLaptopName(brand string) string {
	modelNo := rand.Intn(100)
	return brand + string(rune(modelNo))

}
