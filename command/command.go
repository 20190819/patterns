package command

import "fmt"

// 抽象层

type Command interface {
	Treat()
}

// 实现层

type Doctor struct{}

func (d *Doctor) treatEye() {
	fmt.Println("医生治疗眼睛")
}

func (d *Doctor) treatNose() {
	fmt.Println("医生治疗鼻子")
}

type TreatEye struct {
	doctor *Doctor
	Command
}

func NewTreatEye(doctor *Doctor) *TreatEye {
	return &TreatEye{doctor: doctor}
}

func (t *TreatEye) Treat() {
	t.doctor.treatEye()
}

type TreatNose struct {
	doctor *Doctor
	Command
}

func NewTreatNose(doctor *Doctor) *TreatNose {
	return &TreatNose{doctor: doctor}
}

func (t *TreatNose) Treat() {
	t.doctor.treatNose()
}

// Nurse 护士
type Nurse struct {
	cmds []Command
}

func (n *Nurse) Notify() {

	//
	if len(n.cmds) == 0 {
		return
	}

	// 发布命令
	for _, cmd := range n.cmds {
		cmd.Treat()
	}
}

// 业务逻辑层

func TestCommand() {
	fmt.Printf("\n")

	doc := new(Doctor)
	te := NewTreatEye(doc)
	tn := NewTreatNose(doc)

	nurse := new(Nurse)
	nurse.cmds = append(nurse.cmds, te)
	nurse.cmds = append(nurse.cmds, tn)
	nurse.Notify()

}
