package rez_test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"rez"
	"testing"
)

// RetryXTimes
func Test_RetryXTimes_DoesNotRetryIfNoErrorIsThrown(t *testing.T) {
	a := assert.New(t)
	expectedCallCount := 1
	callCount := 0
	executor := rez.BuildExecutor[int](rez.NewPolicy().RetryXTimes(42))

	executor.Execute(func() (int, error) {
		callCount++
		return 2, nil
	})

	a.Equal(expectedCallCount, callCount)
}

func Test_RetryXTimes_ReturnsExpectedOnNonError(t *testing.T) {
	a := assert.New(t)
	expected := 42
	executor := rez.BuildExecutor[int](rez.NewPolicy().RetryXTimes(9999))

	actual, err := executor.Execute(func() (int, error) {
		return expected, nil
	})

	a.Nil(err)
	a.Equal(expected, actual)
}

func Test_RetryXTimes_ReturnsExpectedOnNonErrorAfterError(t *testing.T) {
	a := assert.New(t)
	hasErrored := false
	expected := 42
	executor := rez.BuildExecutor[int](rez.NewPolicy().RetryXTimes(9999))

	actual, err := executor.Execute(func() (int, error) {
		if !hasErrored {
			hasErrored = true
			return 0, errors.New("an error")
		}
		return expected, nil
	})

	a.Nil(err)
	a.Equal(expected, actual)
}

func Test_RetryXTimes_RetriesExpectedAmountOfTimes(t *testing.T) {
	a := assert.New(t)
	expectedCallCount := 24 // Initial run + 23 retries
	callCount := 0
	executor := rez.BuildExecutor[int](rez.NewPolicy().RetryXTimes(23))

	executor.Execute(func() (int, error) {
		callCount++
		return -1, errors.New("some error")
	})

	a.Equal(expectedCallCount, callCount)
}

func Test_RetryXTimes_ReturnsExpectedOnFullFailure(t *testing.T) {
	a := assert.New(t)
	expectedError := errors.New("some error being thrown")
	expectedValue := -400
	executor := rez.BuildExecutor[int](rez.NewPolicy().RetryXTimes(4))

	value, err := executor.Execute(func() (int, error) {
		return expectedValue, expectedError
	})

	a.Equal(expectedError, err)
	a.Equal(expectedValue, value)
}

// RetryXTimesWithFailureCallback
func Test_RetryXTimesWithFailureCallback_DoesNotRetryIfNoErrorIsThrown(t *testing.T) {
	a := assert.New(t)
	expectedCallCount := 1
	callCount := 0
	executor := rez.BuildExecutor[int](rez.NewPolicy().RetryXTimesWithFailureCallback(42, retryDoNothing))

	executor.Execute(func() (int, error) {
		callCount++
		return 2, nil
	})

	a.Equal(expectedCallCount, callCount)
}

func Test_RetryXTimesWithFailureCallback_ReturnsExpectedOnNonError(t *testing.T) {
	a := assert.New(t)
	expected := 42
	executor := rez.BuildExecutor[int](rez.NewPolicy().RetryXTimesWithFailureCallback(9999, retryDoNothing))

	actual, err := executor.Execute(func() (int, error) {
		return expected, nil
	})

	a.Nil(err)
	a.Equal(expected, actual)
}

func Test_RetryXTimesWithFailureCallback_RetriesExpectedAmountOfTimes(t *testing.T) {
	a := assert.New(t)
	expectedCallCount := 24 // Initial run + 23 retries
	callCount := 0
	executor := rez.BuildExecutor[int](rez.NewPolicy().RetryXTimesWithFailureCallback(23, retryDoNothing))

	executor.Execute(func() (int, error) {
		callCount++
		return -1, errors.New("some error")
	})

	a.Equal(expectedCallCount, callCount)
}

func Test_RetryXTimesWithFailureCallback_ReturnsExpectedOnFullFailure(t *testing.T) {
	a := assert.New(t)
	expectedError := errors.New("some error being thrown")
	expectedValue := -400
	executor := rez.BuildExecutor[int](rez.NewPolicy().RetryXTimesWithFailureCallback(4, retryDoNothing))

	value, err := executor.Execute(func() (int, error) {
		return expectedValue, expectedError
	})

	a.Equal(expectedError, err)
	a.Equal(expectedValue, value)
}

func Test_RetryXTimesWithFailureCallback_CallsFailureCallbackExpectedTimes(t *testing.T) {
	a := assert.New(t)
	retries := 3
	expectedCallbackCalls := 3 // once before each retry, but not after the 4th call.
	callbackCount := 0
	executor := rez.BuildExecutor[int](rez.NewPolicy().RetryXTimesWithFailureCallback(retries, func(retryCount int, err error) {
		callbackCount++
	}))

	executor.Execute(func() (int, error) {
		return 0, errors.New("")
	})

	a.Equal(expectedCallbackCalls, callbackCount)
}

// RetryOnCallback
func Test_RetryOnCallback_DoesNotRetryIfNoErrorIsThrown(t *testing.T) {
	a := assert.New(t)
	expectedCallCount := 1
	callCount := 0
	executor := rez.BuildExecutor[int](rez.NewPolicy().RetryOnCallback(func(retryCount int, err error) bool {
		return true
	}))

	executor.Execute(func() (int, error) {
		callCount++
		return 2, nil
	})

	a.Equal(expectedCallCount, callCount)
}

func Test_RetryOnCallback_ReturnsExpectedOnNonError(t *testing.T) {
	a := assert.New(t)
	expected := 42
	executor := rez.BuildExecutor[int](rez.NewPolicy().RetryOnCallback(func(retryCount int, err error) bool {
		return true
	}))

	actual, err := executor.Execute(func() (int, error) {
		return expected, nil
	})

	a.Nil(err)
	a.Equal(expected, actual)
}

func Test_RetryOnCallback_RetriesExpectedAmountOfTimes(t *testing.T) {
	a := assert.New(t)
	expectedCallCount := 24 // Initial run + 23 retries
	callCount := 0
	executor := rez.BuildExecutor[int](rez.NewPolicy().RetryOnCallback(func(retryCount int, err error) bool {
		return retryCount <= 23
	}))

	executor.Execute(func() (int, error) {
		callCount++
		return -1, errors.New("some error")
	})

	a.Equal(expectedCallCount, callCount)
}

func Test_RetryOnCallback_ReturnsExpectedOnFullFailure(t *testing.T) {
	a := assert.New(t)
	expectedError := errors.New("some error being thrown")
	expectedValue := -400
	executor := rez.BuildExecutor[int](rez.NewPolicy().RetryOnCallback(func(retryCount int, err error) bool {
		return false
	}))

	value, err := executor.Execute(func() (int, error) {
		return expectedValue, expectedError
	})

	a.Equal(expectedError, err)
	a.Equal(expectedValue, value)
}

// RetryOnCallbackWithFailureCallback
func Test_RetryOnCallbackWithFailureCallback_DoesNotRetryIfNoErrorIsThrown(t *testing.T) {
	a := assert.New(t)
	expectedCallCount := 1
	callCount := 0
	executor := rez.BuildExecutor[int](rez.NewPolicy().RetryOnCallbackWithFailureCallback(func(retryCount int, err error) bool {
		return true
	}, retryDoNothing))

	executor.Execute(func() (int, error) {
		callCount++
		return 2, nil
	})

	a.Equal(expectedCallCount, callCount)
}

func Test_RetryOnCallbackWithFailureCallback_ReturnsExpectedOnNonError(t *testing.T) {
	a := assert.New(t)
	expected := 42
	executor := rez.BuildExecutor[int](rez.NewPolicy().RetryOnCallbackWithFailureCallback(func(retryCount int, err error) bool {
		return true
	}, retryDoNothing))

	actual, err := executor.Execute(func() (int, error) {
		return expected, nil
	})

	a.Nil(err)
	a.Equal(expected, actual)
}

func Test_RetryOnCallbackWithFailureCallback_RetriesExpectedAmountOfTimes(t *testing.T) {
	a := assert.New(t)
	expectedCallCount := 24 // Initial run + 23 retries
	callCount := 0
	executor := rez.BuildExecutor[int](rez.NewPolicy().RetryOnCallbackWithFailureCallback(func(retryCount int, err error) bool {
		return retryCount <= 23
	}, retryDoNothing))

	executor.Execute(func() (int, error) {
		callCount++
		return -1, errors.New("some error")
	})

	a.Equal(expectedCallCount, callCount)
}

func Test_RetryOnCallbackWithFailureCallback_ReturnsExpectedOnFullFailure(t *testing.T) {
	a := assert.New(t)
	expectedError := errors.New("some error being thrown")
	expectedValue := -400
	executor := rez.BuildExecutor[int](rez.NewPolicy().RetryOnCallbackWithFailureCallback(func(retryCount int, err error) bool {
		return false
	}, retryDoNothing))

	value, err := executor.Execute(func() (int, error) {
		return expectedValue, expectedError
	})

	a.Equal(expectedError, err)
	a.Equal(expectedValue, value)
}

func Test_RetryOnCallbackWithFailureCallback_CallsFailureCallbackExpectedTimes(t *testing.T) {
	a := assert.New(t)
	expectedCallbackCalls := 3 // once before each retry, but not after the 4th call.
	callbackCount := 0
	executor := rez.BuildExecutor[int](rez.NewPolicy().RetryOnCallbackWithFailureCallback(
		func(retryCount int, err error) bool {
			return retryCount <= 3
		}, func(retryCount int, err error) {
			callbackCount++
		}))

	executor.Execute(func() (int, error) {
		return 0, errors.New("")
	})

	a.Equal(expectedCallbackCalls, callbackCount)
}
