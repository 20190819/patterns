package good

import "fmt"

type PayBanker struct {
}

func (b *PayBanker) DoBiz() {
	fmt.Println("支付业务")
}
