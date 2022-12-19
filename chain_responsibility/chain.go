package chainResponsibility

import (
	"fmt"
	"log"
)

func ChainStart() {
	patient := &patient{
		Name:             "Jone",
		RegistrationDone: true,
	}

	start := new(Start)
	start.SetNext(new(Reception)).SetNext(new(Clinic)).SetNext(new(Cashier)).SetNext(new(Pharmacy))
	if err := start.Execute(patient); err != nil {
		log.Fatalln("程序异常")
	}
	fmt.Println("看病成功")
}
