package rez

type PolicyExecutor[ReturnType interface{}] interface {
	Execute(f func() (ReturnType, error)) (ReturnType, error)
}
