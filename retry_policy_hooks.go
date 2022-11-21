package rez

import "time"

// RetryXTimes sets up a policy which retries on error based on the number of times you request.
func (p Policy) RetryXTimes(retryCount int) Policy {
	p.retry = &retryPolicy{
		shouldRetry: func(count int, err error) bool {
			return count <= retryCount
		},
	}
	return p
}

// RetryXTimesWithFailureCallback sets up a policy which retries on error based on the number of times you request, with the
// ability to run callback logic between failures.
func (p Policy) RetryXTimesWithFailureCallback(retryCount int, betweenFailures func(retryCount int, err error)) Policy {
	p.retry = &retryPolicy{
		shouldRetry: func(count int, err error) bool {
			return count <= retryCount
		},
		betweenFailures: toPtr(betweenFailures),
	}
	return p
}

// RetryOnCallback sets up a policy which will retry as long as the callback passed into it returns true.
func (p Policy) RetryOnCallback(errorCheckCallback func(retryCount int, err error) bool) Policy {
	p.retry = &retryPolicy{
		shouldRetry: errorCheckCallback,
	}
	return p
}

// RetryOnCallbackWithFailureCallback sets up a policy which will retry as long as the callback passed into it returns true, with the
// // ability to run callback logic between failures.
func (p Policy) RetryOnCallbackWithFailureCallback(errorCheckCallback func(retryCount int, err error) bool, betweenFailures func(retryCount int, err error)) Policy {
	p.retry = &retryPolicy{
		shouldRetry:     errorCheckCallback,
		betweenFailures: toPtr(betweenFailures),
	}
	return p
}

// WaitRetryXTimes sets up a policy which retries on error based on the number of times you request, while waiting a set
// duration between each attempt.
func (p Policy) WaitRetryXTimes(retryCount int, waitDuration time.Duration) Policy {
	p.retry = &retryPolicy{
		shouldRetry: func(count int, err error) bool {
			return count <= retryCount
		},
		waitDurationGenerator: toPtr(func(retryCount int, err error) time.Duration {
			return waitDuration
		}),
	}
	return p
}

// WaitRetryXTimesWithFailureCallback sets up a policy which retries on error based on the number of times you request, while waiting a set
// duration between each attempt, with the ability to run callback logic between failures.
func (p Policy) WaitRetryXTimesWithFailureCallback(retryCount int, waitDuration time.Duration, betweenFailures func(retryCount int, err error)) Policy {
	p.retry = &retryPolicy{
		shouldRetry: func(count int, err error) bool {
			return count <= retryCount
		},
		waitDurationGenerator: toPtr(func(retryCount int, err error) time.Duration {
			return waitDuration
		}),
		betweenFailures: toPtr(betweenFailures),
	}
	return p
}

// WaitRetryOnCallback sets up a policy which will retry as long as the callback passed into it returns true, while
// waiting a set duration between each attempt.
func (p Policy) WaitRetryOnCallback(waitDuration time.Duration, errorCheckCallback func(retryCount int, err error) bool) Policy {
	p.retry = &retryPolicy{
		shouldRetry: errorCheckCallback,
		waitDurationGenerator: toPtr(func(retryCount int, err error) time.Duration {
			return waitDuration
		}),
	}
	return p
}

// WaitRetryOnCallbackWithFailureCallback sets up a policy which will retry as long as the callback passed into it returns true, while
// waiting a set duration between each attempt, with the ability to run callback logic between failures.
func (p Policy) WaitRetryOnCallbackWithFailureCallback(waitDuration time.Duration, errorCheckCallback func(retryCount int, err error) bool, betweenFailures func(retryCount int, err error)) Policy {
	p.retry = &retryPolicy{
		shouldRetry: errorCheckCallback,
		waitDurationGenerator: toPtr(func(retryCount int, err error) time.Duration {
			return waitDuration
		}),
		betweenFailures: toPtr(betweenFailures),
	}
	return p
}

// ComplexWaitRetryXTimes sets up a policy which retries on error based on the number of times you request, while also
// waiting between retries a duration determined by the callback passed into the waitDurationCallback parameter.
func (p Policy) ComplexWaitRetryXTimes(retryCount int, waitDurationCallback func(retryCount int, err error) time.Duration) Policy {
	p.retry = &retryPolicy{
		shouldRetry: func(count int, err error) bool {
			return count <= retryCount
		},
		waitDurationGenerator: toPtr(waitDurationCallback),
	}
	return p
}

// ComplexWaitRetryXTimesWithFailureCallback sets up a policy which retries on error based on the number of times you request, while also
// waiting between retries a duration determined by the callback passed into the waitDurationCallback parameter, with the
// // ability to run callback logic between failures.
func (p Policy) ComplexWaitRetryXTimesWithFailureCallback(retryCount int, waitDurationCallback func(retryCount int, err error) time.Duration, betweenFailures func(retryCount int, err error)) Policy {
	p.retry = &retryPolicy{
		shouldRetry: func(count int, err error) bool {
			return count <= retryCount
		},
		waitDurationGenerator: toPtr(waitDurationCallback),
		betweenFailures:       toPtr(betweenFailures),
	}
	return p
}

// ComplexWaitRetryOnCallback sets up a policy which will retry as long as the callback passed into it returns true,
// while also waiting between retries a duration determined by the callback passed into the waitDurationCallback
// parameter.
func (p Policy) ComplexWaitRetryOnCallback(waitDurationCallback func(retryCount int, err error) time.Duration, errorCheckCallback func(retryCount int, err error) bool) Policy {
	p.retry = &retryPolicy{
		shouldRetry:           errorCheckCallback,
		waitDurationGenerator: toPtr(waitDurationCallback),
	}
	return p
}

// ComplexWaitRetryOnCallbackWithFailureCallback sets up a policy which will retry as long as the callback passed into it returns true,
// while also waiting between retries a duration determined by the callback passed into the waitDurationCallback
// parameter, with the ability to run callback logic between failures.
func (p Policy) ComplexWaitRetryOnCallbackWithFailureCallback(waitDurationCallback func(retryCount int, err error) time.Duration, errorCheckCallback func(retryCount int, err error) bool, betweenFailures func(retryCount int, err error)) Policy {
	p.retry = &retryPolicy{
		shouldRetry:           errorCheckCallback,
		waitDurationGenerator: toPtr(waitDurationCallback),
		betweenFailures:       toPtr(betweenFailures),
	}
	return p
}
