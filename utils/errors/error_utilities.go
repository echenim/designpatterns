package errors

type ErrorBuilder struct {
	attach string
}

func (eb *ErrorBuilder) Affixed(s string) {
	eb.attach = eb.attach + "\n" + s
}

func (eb *ErrorBuilder) Print() string {
	return eb.attach
}
