package abstract

// 产品族 同一个地域产地 同一个厂商 （理解为：系列产品）
// 产品等级结构 （相同功能，产地不同，厂商不同） （理解为：竞品）

// 抽象层

type AbstractApple interface {
	ShowApple()
}

type AbstractPear interface {
	ShowPear()
}

type AbstractBanana interface {
	ShowBanana()
}

type AbstractFactoryInterface interface {
	CreateApple() AbstractApple
	CreatePear() AbstractPear
	CreateBanana() AbstractBanana
}