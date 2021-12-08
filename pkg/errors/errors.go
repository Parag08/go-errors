package errors

import (
	"errors"
	"net/http"
)

type Op string

type Severity int8

const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel Severity = iota - 1
	// InfoLevel is the default logging priority.
	InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel

	_minLevel = DebugLevel
	_maxLevel = FatalLevel
)

type Kind int

const (
	KindNotFound            = http.StatusNotFound
	KindInternalServerError = http.StatusInternalServerError
)

func (k Kind) String() string {
	switch k {
	case KindInternalServerError:
		return "Internal Server Error"
	case KindNotFound:
		return "Not Found"
	}
	return "unknown error kind"
}

type Error struct {
	Op       Op
	Kind     Kind
	Err      error
	Severity Severity
}

func (e *Error) Error() string {
	if e.Err != nil {
		return e.Err.Error()
	}
	return ""
}

func E(args ...interface{}) error {
	e := &Error{}
	for _, arg := range args {
		if arg == nil {
			continue
		}
		switch arg := arg.(type) {
		case Op:
			e.Op = arg
		case error:
			e.Err = arg
		case Kind:
			e.Kind = arg
		case Severity:
			e.Severity = arg
		case string:
			e.Err = errors.New(arg)
		default:
			panic("Bad call to E")
		}
	}
	// Skeptical about this it essentially prevent us from sending the null errors
	if e.isZero() {
		return nil
	}
	return e
}

func (e *Error) isZero() bool {
	return e.Err == nil
}

func Ops(e Error) []Op {
	res := []Op{e.Op}

	subErr, ok := e.Err.(*Error)
	if !ok {
		return res
	}

	res = append(res, Ops(*subErr)...)

	return res
}

func Level(e Error) Severity {
	var sevirity Severity

	subErr, ok := e.Err.(*Error)
	if !ok {
		return -1
	}

	sevirity = Level(*subErr)
	if subErr.Severity > sevirity {
		sevirity = subErr.Severity
	}

	return sevirity
}

func GetKind(e Error) Kind {
	subErr, ok := e.Err.(*Error)
	if !ok {
		return -1
	}

	if subErr.Kind != 0 {
		return subErr.Kind
	}

	return GetKind(*subErr)
}
