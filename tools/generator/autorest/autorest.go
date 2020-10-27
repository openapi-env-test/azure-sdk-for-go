package autorest

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/tools/generator/model"
	"os/exec"
	"strings"
)

// Task ...
type Task struct {
	// AbsReadmeMd absolute path of the readme.md file to generate
	AbsReadmeMd string
}

type Options struct {
	// AutorestArguments are the optional flags for the autorest tool
	AutorestArguments []string
	// AfterScripts are the scripts that need to be run after the SDK is generated
	AfterScripts []string
}

// Execute executes the autorest task, and then invoke the after scripts
// the error returned will be TaskError
func (t *Task) Execute(options Options) (*model.PackageResult, error) {
	if err := t.executeAutorest(options.AutorestArguments); err != nil {
		return nil, err
	}

	if err := t.executeAfterScript(options.AfterScripts); err != nil {
		return nil, err
	}

	return t.getPackageResult()
}

func (t *Task) executeAutorest(options []string) error {
	arguments := append(options, t.AbsReadmeMd)
	c := exec.Command("autorest", arguments...)
	output, err := c.CombinedOutput()
	if err != nil {
		return &TaskError{
			AbsReadmeMd: t.AbsReadmeMd,
			Script:      "autorest",
			Message:     string(output),
		}
	}
	return nil
}

func (t *Task) executeAfterScript(afterScripts []string) error {
	for _, script := range afterScripts {
		arguments := strings.Split(script, " ")
		c := exec.Command(arguments[0], arguments[1:]...)
		output, err := c.CombinedOutput()
		if err != nil {
			return &TaskError{
				AbsReadmeMd: t.AbsReadmeMd,
				Script:      script,
				Message:     string(output),
			}
		}
	}

	return nil
}

func (t *Task) getPackageResult() (*model.PackageResult, error) {
	// get the file change list first
	// TODO -- finish this
	return nil, nil
}

// TaskError the error returned during an autorest task
type TaskError struct {
	// AbsReadmeMd relative path of the readme.md file to generate
	AbsReadmeMd string
	// Script the script running when the error is thrown
	Script string
	// Message the error message
	Message string
}

func (r *TaskError) Error() string {
	return fmt.Sprintf("autorest task failed for readme file '%s' during '%s': %s", r.AbsReadmeMd, r.Script, r.Message)
}
