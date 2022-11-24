package good

import "fmt"

type SaveBanker struct {
}

func (s *SaveBanker) DoBiz() {
	fmt.Println("存款业务")
}
