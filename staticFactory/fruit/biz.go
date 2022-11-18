package fruit

import (
	american2 "patterns/staticFactory/fruit/american"
	china2 "patterns/staticFactory/fruit/china"
	japan2 "patterns/staticFactory/fruit/japan"
)

func Biz() {

	facA:=new(american2.FactoryAmerican)
	facA.CreateApple().ShowApple()
	facA.CreatePear().ShowPear()
	facA.CreateBanana().ShowBanana()

	facC:=new(china2.FactoryChina)
	facC.CreateApple().ShowApple()
	facC.CreatePear().ShowPear()
	facC.CreateBanana().ShowBanana()

	facJ:=new(japan2.FactoryJapan)
	facJ.CreateApple().ShowApple()
	facJ.CreatePear().ShowPear()
	facJ.CreateBanana().ShowBanana()

}