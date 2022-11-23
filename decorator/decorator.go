package decorator

import "fmt"

// 抽象层

type Phone interface {
	Photo() // 拍照
}

type Decorator struct {
	phone Phone
}

func (d *Decorator) Photo() {

}

// 实现层

type Huawei struct {
}

func (h *Huawei) Photo() {
	fmt.Println("huawei 手机拍照")
}

type Apple struct {
}

func (i *Apple) Photo() {
	fmt.Println("苹果手机拍照")
}

type MoDecorator struct {
	Decorator
}

func (m *MoDecorator) Photo() {
	fmt.Println("先贴个膜,再拍照")
	m.phone.Photo()
}

func NewMoDecorator(phone Phone) *MoDecorator {
	return &MoDecorator{
		Decorator{phone: phone},
	}
}

type KeDecorator struct {
	Decorator
}

func (k *KeDecorator) Photo() {
	fmt.Println("先套个壳,再拍照")
}

func NewKeDecorator(phone Phone) *KeDecorator {
	return &KeDecorator{
		Decorator{phone: phone},
	}
}

// 业务层

func TestDecorator() {
	huawei := new(Huawei)
	huawei.Photo() // 调用原构件
	fmt.Println("--------")

	mo := NewMoDecorator(huawei)
	mo.Photo()
	fmt.Println("--------")

	ke := NewKeDecorator(huawei)
	ke.Photo()
	fmt.Println("--------")

	moAndKe := NewMoDecorator(ke)
	moAndKe.Photo()
}
