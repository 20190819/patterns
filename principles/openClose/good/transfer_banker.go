package good

import "fmt"

type TransferBanker struct {
}

func (t *TransferBanker) DoBiz() {
	fmt.Println("转账业务")
}
