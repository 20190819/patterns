package facade

import "fmt"

type Action interface {
	On()
	Off()
}

type TV struct {
}

func (t *TV) On() {
	fmt.Println("TV turn on")
}

func (t *TV) Off() {
	fmt.Println("TV turn off")
}

type VoiceBox struct {
}

func (v *VoiceBox) On() {
	fmt.Println("VoiceBox turn on")
}

func (v *VoiceBox) Off() {
	fmt.Println("VoiceBox turn off")
}

type Light struct {
}

func (l *Light) On() {
	fmt.Println("Light turn on")
}

func (l *Light) Off() {
	fmt.Println("Light turn off")
}

type KTVFacade struct {
	tv       *TV
	voiceBox *VoiceBox
	light    *Light
}

func NewKTVFacade(tv *TV, box *VoiceBox, light *Light) *KTVFacade {
	return &KTVFacade{tv, box, light}
}

func (k *KTVFacade) mA() {
	k.tv.On()
	k.voiceBox.On()
}

func (k *KTVFacade) mB() {
	k.tv.On()
	k.light.On()
}

func TestKTVFacade() {
	fmt.Printf("\n")
	ktv := NewKTVFacade(new(TV), new(VoiceBox), new(Light))
	ktv.mA()
	ktv.mB()
}
