package types

// InvokeFunction is a unified function type for all event handlers.
// https://yourbasic.org/golang/function-pointer-type-declaration/
type InvokeFunction func(message DataMessage) (err error)
