package hystrixgo

import (
	"fmt"
)

type HystrixConfig struct {
	Name                   string
	Timeout                int
	MaxConcurrentRequests  int
	ErrorPercentThreshold  int
	RequestVolumeThreshold int
	SleepWindow            int
}

func (c HystrixConfig) Validate() error {
	if c.Name == "" || c.Timeout == 0 || c.MaxConcurrentRequests == 0 || c.ErrorPercentThreshold == 0 || c.RequestVolumeThreshold == 0 || c.SleepWindow == 0 {
		return fmt.Errorf("incomplete Hystrix configuration")
	}
	return nil
}
