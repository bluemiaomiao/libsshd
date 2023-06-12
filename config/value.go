package config

// Value is configuration item value.
type Value struct {
	annotation string
	value      interface{}
}

func (ctx *Value) SetAnnotation(anno string) {
	ctx.annotation = anno
}

func (ctx *Value) SetValue(v interface{}) {
	ctx.value = v
}

func (ctx *Value) GetAnnotation() string {
	return ctx.annotation
}

func (ctx *Value) GetValue() interface{} {
	return ctx.value
}
