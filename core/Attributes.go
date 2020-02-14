package core

type Attributes struct {
	StringAttributes  map[string]string
	Float64Attributes map[string]float64
	IntAttributes     map[string]int
}

func NewAttributes() *Attributes {
	return &Attributes{
		make(map[string]string),
		make(map[string]float64),
		make(map[string]int),
	}
}

func (a *Attributes) PutString(key string, value string) *Attributes {
	a.StringAttributes[key] = value
	return a
}
func (a *Attributes) PutFloat64(key string, value float64) *Attributes {
	a.Float64Attributes[key] = value
	return a
}
func (a *Attributes) PutInt(key string, value int) *Attributes {
	a.IntAttributes[key] = value
	return a
}
