package computer

import (
	intelpkg "patterns/staticFactory/computer/intel"
	"patterns/staticFactory/computer/kingston"
	nvidiapkg "patterns/staticFactory/computer/nvidia"
)

func Biz() {

	intel := new(intelpkg.ComputerFactoryIntel)
	intel.CreateStorage().ShowStorage()
	intel.CreateDisplay().ShowDisplay()
	intel.CreateCalculate().ShowCalculate()

	kston := new(kingston.ComputerFactoryKingston)
	kston.CreateStorage().ShowStorage()
	kston.CreateDisplay().ShowDisplay()
	kston.CreateCalculate().ShowCalculate()

	nvidia := new(nvidiapkg.ComputerFactoryNvidia)
	nvidia.CreateStorage().ShowStorage()
	nvidia.CreateDisplay().ShowDisplay()
	nvidia.CreateCalculate().ShowCalculate()

}
