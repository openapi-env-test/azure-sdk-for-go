package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/tools/apidiff/ioext"
	"github.com/Azure/azure-sdk-for-go/tools/generator/autorest"
	"github.com/Azure/azure-sdk-for-go/tools/generator/model"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Command() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "generator <generate input filepath> <generate output filepath>",
		Short: "", // TODO
		Long:  ``, // TODO
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return execute(args[0], args[1])
		},
	}

	return rootCmd
}

func execute(inputPath, outputPath string) error {
	input, err := readInputFrom(inputPath)
	if err != nil {
		return fmt.Errorf("cannot read generate input: %+v", err)
	}
	output, err := generate(input)
	if err != nil {
		return fmt.Errorf("cannot generate: %+v", err)
	}
	if err := writeOutputTo(outputPath, output); err != nil {
		return fmt.Errorf("cannot write generate output: %+v", err)
	}
	return nil
}

func readInputFrom(inputPath string) (*model.GenerateInput, error) {
	b, err := ioutil.ReadFile(inputPath)
	if err != nil {
		return nil, err
	}
	var input model.GenerateInput
	if err := json.Unmarshal(b, &input); err != nil {
		return nil, err
	}
	return &input, nil
}

func writeOutputTo(outputPath string, output *model.GenerateOutput) error {
	b, err := json.Marshal(*output)
	if err != nil {
		return err
	}
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := file.Write(b); err != nil {
		return err
	}
	return nil
}

const sdkRoot = "azure-sdk-for-go"

// TODO -- support dry run
func generate(input *model.GenerateInput) (*model.GenerateOutput, error) {
	if input.DryRun {
		return nil, fmt.Errorf("dry run not supported yet")
	}
	// backup the current sdk to temp dir
	tempDir := filepath.Join(os.TempDir(), sdkRoot)
	if err := ioext.CopyDir(".", tempDir); err != nil {
		return nil, err
	}
	defer os.RemoveAll(tempDir)
	// iterate over all the readme
	var results []model.PackageResult
	for _, readme := range input.RelatedReadmeMdFiles {
		task := autorest.Task{
			AbsReadmeMd: filepath.Join(input.SpecFolder, readme),
		}
		options := autorest.Options{
			AutorestArguments: []string{
				// TODO -- store these settings elsewhere rather than hard code here
				"--use=@microsoft.azure/autorest.go@~2.1.157",
				"--go",
				"--verbose",
				"--go-sdk-folder=.", // TODO -- if dry run, maybe we could use a temp directory to put the generated files and then delete
				"--multiapi",
				"--use-onever",
				"--preview-chk",
				"--version=V2",
			},
			AfterScripts:      []string{
				// TODO -- store these settings elsewhere rather than hard code here
				//"dep ensure", // TODO -- enable this
				"go generate ./profiles/generate.go",
				"gofmt -w ./profiles/",
				"gofmt -w ./services/",
			},
		}
		packageResult, err := task.Execute(options)
		if err != nil {
			return nil, err
		}
		results = append(results, *packageResult)
	}

	return &model.GenerateOutput{
		Packages: results,
	}, nil
}
