package abstract

type ComputerFactory interface {
	CreateDisplay() Display
	CreateStorage() Storage
	CreateCalculate() Calculate
}

type Display interface {
	ShowDisplay()
}

type Storage interface {
	ShowStorage()
}

type Calculate interface {
	ShowCalculate()
}


