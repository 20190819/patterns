package chainResponsibility

type PatientHandler interface {
	Execute(*patient) error
	SetNext(PatientHandler) PatientHandler
	Do(*patient) error
}

// Next 充当抽象类，实现公共方法
type Next struct {
	nextHandler PatientHandler
}

func (n *Next) SetNext(handler PatientHandler) PatientHandler {
	n.nextHandler = handler
	return n.nextHandler
}

func (n *Next) Execute(patient *patient) error {
	if n.nextHandler != nil {

		err := n.nextHandler.Do(patient)
		if err != nil {
			return err
		}

		return n.nextHandler.Execute(patient)
	}
	return nil
}
