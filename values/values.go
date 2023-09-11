package values

type RuntimeVal struct {
	Type string
}

func (v *RuntimeVal) typeof() string {
	return v.Type
}

type Value struct {
	Type string
}

type Number struct {
	*RuntimeVal
	Type  string
	Value float64
}

type Null struct {
	*RuntimeVal
	Type  string
	Value string
}

type String struct {
	*RuntimeVal
	Type  string
	Value string
}
