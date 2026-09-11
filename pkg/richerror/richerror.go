package richerror

type Kind int

const (
	KindInvalid Kind = iota + 1
	KindForbidden
	KindNotFound
	KindUnexpected
)

type Op string

type RichError struct {
	wrappedError error
	message      string
	kind         Kind
	operation    Op
	meta         map[string]interface{}
}

func New(op Op) RichError {
	return RichError{operation: op}
}

//func New(wrappedError error, operation, message string, kind Kind, meta map[string]interface{}) RichError {
//	return RichError{wrappedError: wrappedError,
//		operation: operation,
//		message:   message,
//		kind:      kind,
//		meta:      meta,
//	}
//}

func (r RichError) WithErr(err error) RichError {
	r.wrappedError = err
	return r
}

func (r RichError) WithOp(op Op) RichError {
	r.operation = op
	return r
}

func (r RichError) WithMessage(message string) RichError {
	r.message = message
	return r
}

func (r RichError) WithKind(kind Kind) RichError {
	r.kind = kind
	return r
}
func (r RichError) WithMeta(meta map[string]interface{}) RichError {
	r.meta = meta
	return r
}

func (r RichError) Error() string {

	if r.message == "" {
		return r.wrappedError.Error()
	}

	return r.message
}

func (r RichError) Kind() Kind {

	if r.kind != 0 {
		return r.kind
	}

	re, ok := r.wrappedError.(RichError)
	if !ok {
		return 0
	}

	return re.Kind()
}

func (r RichError) Message() string {
	if r.message != "" {
		return r.message
	}

	re, ok := r.wrappedError.(RichError)
	if !ok {
		return r.wrappedError.Error()
	}

	return re.Message()
}
