package rez

import "time"

type Policy struct {
	retry *retryPolicy
}

func NewPolicy() Policy {
	return Policy{}
}

func BuildExecutor[ReturnType interface{}](policy Policy) PolicyExecutor[ReturnType] {
	return &executor[ReturnType]{
		policy: policy,
	}
}

type retryPolicy struct {
	betweenFailures       *func(retryCount int, err error)
	waitDurationGenerator *func(retryCount int, err error) time.Duration
	shouldRetry           func(retryCount int, err error) bool
}
