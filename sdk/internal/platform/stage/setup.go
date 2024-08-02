package stagesetup

import (
	stage "github.com/devpablocristo/golang-sdk/pkg/stage"
)

func NewStageInstance() (stage.StageClientPort, error) {
	if err := stage.InitializeStageClient(); err != nil {
		return nil, err
	}

	return stage.GetStageInstance()
}
