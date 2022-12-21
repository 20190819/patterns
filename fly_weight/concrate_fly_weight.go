package flyWeight

import "fmt"

type CharacterFlyWeight struct {
	name string
}

func NewCharacterFlyWeight(name string) FlyWeight {
	return &CharacterFlyWeight{name}
}

func (c *CharacterFlyWeight) Render(extrinsicState string) string {
	return fmt.Sprintf("character %s with font %s", c.name, extrinsicState)
}
