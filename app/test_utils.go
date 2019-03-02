package main

import (
	"bytes"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

type CommandMock struct {
	Command
}

func (cm *CommandMock) GetCommand() *exec.Cmd {
	return &exec.Cmd{
		ProcessState: &os.ProcessState{},
	}
}

func createFile(t *testing.T, pathToFile string) {
	testFileCreation, err := os.Create(pathToFile)
	assert.Nilf(
		t,
		err,
		"unexpected error when creating file at '%s': %s",
		pathToFile,
		err,
	)
	defer testFileCreation.Close()
}

func InitCommandMock(application string, arguments []string, logOutput *bytes.Buffer) *CommandMock {
	command := &CommandMock{}
	command.config = &CommandConfig{
		Application: application,
		Arguments:   arguments,
	}
	command.logger = InitLogger(&LoggerConfig{
		Name:   application,
		Format: "production",
		Level:  "trace",
	})
	command.logger.SetOutput(logOutput)
	return command
}

func expectError(t *testing.T) func() {
	return func() {
		assert.NotNil(t, recover(), "expected an error but none was panicked")
	}
}

func expectNoError(t *testing.T) func() {
	return func() {
		err := recover()
		assert.Nilf(t, err, "expected no errors but '%s' was panicked", err)
	}
}

func removeFile(t *testing.T, pathToFile string) {
	err := os.Remove(pathToFile)
	assert.Nilf(
		t,
		err,
		"unexpected error when removing file at '%s': %s",
		pathToFile,
		err,
	)
}
