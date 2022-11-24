package good

type BankerBusiness struct {
}

func (b BankerBusiness) Do(banker AbstractBanker) {
	banker.DoBiz()
}

func TestBanker() {
	biz := new(BankerBusiness)
	biz.Do(new(SaveBanker))
	biz.Do(new(PayBanker))
	biz.Do(new(TransferBanker))
}
