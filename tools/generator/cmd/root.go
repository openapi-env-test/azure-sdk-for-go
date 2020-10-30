package cmd

import (
	"fmt"
	"github.com/Azure/azure-sdk-for-go/tools/apidiff/ioext"
	"github.com/Azure/azure-sdk-for-go/tools/generator/autorest"
	"github.com/Azure/azure-sdk-for-go/tools/generator/model"
	"github.com/spf13/cobra"
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
	inputFile, err := os.Open(inputPath)
	if err != nil {
		return nil, err
	}
	return model.NewGenerateInputFrom(inputFile)
}

func writeOutputTo(outputPath string, output *model.GenerateOutput) error {
	file, err := os.Create(outputPath)
	if err != nil {
		return err
	}
	defer file.Close()
	if _, err := output.WriteTo(file); err != nil {
		return err
	}
	return nil
}

func getTempDir() string {
	tempDirRoot := os.Getenv("TMPDIR")
	if tempDirRoot == "" {
		tempDirRoot = os.TempDir()
	}
	return filepath.Join(tempDirRoot, sdkRoot)
}

const sdkRoot = "azure-sdk-for-go"

// TODO -- support dry run
func generate(input *model.GenerateInput) (*model.GenerateOutput, error) {
	if input.DryRun {
		return nil, fmt.Errorf("dry run not supported yet")
	}
	// backup the current sdk to temp dir
	tempDir := getTempDir()
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
		if err := task.Execute(options); err != nil {
			return nil, err
		}
		// get changed file list
		changedFiles, err := getChangedFiles()
		if err != nil {
			return nil, err
		}
		// get packages using the changed file list
		packages, err := autorest.GetChangedPackages(changedFiles)
		if err != nil {
			return nil, err
		}
		// key is package path, value is files that have changed
		for p := range packages {
			pp := p
			results = append(results, model.PackageResult{
				PackageName: pp, // TODO -- this is the package identifier
				Path: []string{pp}, // TODO -- this is the package path relative to the root of SDK
				ReadmeMd: []string{readme},
			})
		}
	}

	return &model.GenerateOutput{
		Packages: results,
	}, nil
}

func getChangedFiles() ([]string, error) {
	return []string{}, nil
}
