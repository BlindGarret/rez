package rez

import "time"

type executor[ReturnType interface{}] struct {
	policy Policy
}

func (e executor[ReturnType]) Execute(f func() (ReturnType, error)) (ReturnType, error) {
	r, err := f()
	if err == nil {
		// early out
		return r, nil
	}

	if e.policy.retry != nil {
		// Handle Retry Policy
		retryPolicy := *e.policy.retry
		retryCount := 1

		for retryPolicy.shouldRetry(retryCount, err) {
			// Handle failure callbacks
			if retryPolicy.betweenFailures != nil {
				(*retryPolicy.betweenFailures)(retryCount, err)
			}

			// Handle waits
			if retryPolicy.waitDurationGenerator != nil {
				time.Sleep((*retryPolicy.waitDurationGenerator)(retryCount, err))
			}

			r, err = f()
			if err == nil {
				return r, nil
			}
			retryCount++
		}
	}

	return r, err
}
