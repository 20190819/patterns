package flyWeight

type FlyWeightFactory struct {
	pool map[string]FlyWeight
}

func NewFactory() *FlyWeightFactory {
	return &FlyWeightFactory{pool: make(map[string]FlyWeight, 50)}
}

func (f *FlyWeightFactory) get(name string) FlyWeight {
	_, ok := f.pool[name]
	if !ok {
		f.pool[name] = NewCharacterFlyWeight(name)
	}
	return f.pool[name]
}

func (f *FlyWeightFactory) count() int {
	return len(f.pool)
}
