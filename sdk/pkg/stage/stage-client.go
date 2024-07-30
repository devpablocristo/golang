package stage

import (
	"fmt"
	"strings"
	"sync"

	"github.com/iancoleman/strcase"
)

var (
	instance StageClientPort
	once     sync.Once
	errInit  error
)

type StageClient struct{}

func InitializeStageClient() error {
	once.Do(func() {
		instance = &StageClient{}
	})
	return errInit
}

func GetStageInstance() (StageClientPort, error) {
	if instance == nil {
		return nil, fmt.Errorf("stage client is not initialized")
	}
	return instance, nil
}

func (c *StageClient) String(s Stage) string {
	if name, exists := stageNames[s]; exists {
		return name
	}
	return "UNKNOWN"
}

func (c *StageClient) GetFromString(str string) Stage {
	upperStr := strings.ToUpper(str)
	for key, value := range stageNames {
		if value == upperStr {
			return key
		}
	}
	return Unknown
}

func (c *StageClient) GetFromCamelCase(str string) Stage {
	upperStr := strcase.ToScreamingSnake(str)
	return c.GetFromString(upperStr)
}

func (c *StageClient) GetFromKebabCase(str string) Stage {
	upperStr := strcase.ToScreamingKebab(str)
	return c.GetFromString(upperStr)
}
