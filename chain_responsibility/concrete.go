package chainResponsibility

import "fmt"

type Start struct {
	Next
}

func (s *Start) Do(p *patient) (err error) {
	return
}

// Reception 挂号
type Reception struct {
	Next
}

func (r *Reception) Do(p *patient) (err error) {

	if p.RegistrationDone {
		fmt.Println("已完成病人的登记")
		return
	}
	fmt.Println("接待处正在登记病人")
	p.RegistrationDone = true
	return
}

// Clinic 诊室
type Clinic struct {
	Next
}

func (c *Clinic) Do(p *patient) (err error) {
	if p.DoctorCheckUpDone {
		fmt.Println("医生已经检查完毕")
		return
	}
	fmt.Println("医生正在检查病人")
	p.DoctorCheckUpDone = true
	return
}

// Cashier 收费处
type Cashier struct {
	Next
}

func (c *Cashier) Do(p *patient) (err error) {
	if p.PaymentDone {
		fmt.Println("收银员已经收费完成")
		return
	}
	fmt.Println("收银员正在向病人收费")
	p.PaymentDone = true
	return
}

// Pharmacy 药房
type Pharmacy struct {
	Next
}

func (ph *Pharmacy) Do(p *patient) (err error) {
	if p.MedicineDone {
		fmt.Println("药品已经发给病人")
		return
	}

	fmt.Println("药房正在给病人拿药")
	return
}
