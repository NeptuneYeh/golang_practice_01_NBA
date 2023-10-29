package init

import (
	"errors"
	loggerModule "github.com/NeptuneYeh/golang_practice_01_NBA/init/logger"
)

type ModuleInterFace interface {
	Run() error
	Shutdown() error
}

type MainInitProcess struct {
	loggerModule *loggerModule.LoggerModule
}

// NewMainInitProcess is a constructor for MainInitProcess
func NewMainInitProcess() *MainInitProcess {
	loggerModule := loggerModule.NewModule(nil)
	return &MainInitProcess{
		loggerModule: loggerModule,
	}
}

func (m *MainInitProcess) Run() {
	m.loggerModule.Debug("error code: 1010000")
	m.loggerModule.DebugNoSkip("error code: 1010017")
}

func (m *MainInitProcess) Shutdown() error {
	return errors.New("something went wrong")
}
