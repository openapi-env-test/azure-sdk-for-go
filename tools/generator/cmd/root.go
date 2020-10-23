package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/tools/generator/model"
	"github.com/spf13/cobra"
	"io/ioutil"
	"os"
)

func Command() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "generator <generate input filepath> <generate output filepath>",
		Short: "", // TODO
		Long:  ``, // TODO
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			return generate(args[0], args[1])
		},
	}

	return rootCmd
}

func generate(inputPath, outputPath string) error {
	input, err := readInputFrom(inputPath)
	if err != nil {
		return fmt.Errorf("cannot read generate input: %+v", err)
	}
	output, err := generateOutput(input)
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

func generateOutput(input *model.GenerateInput) (*model.GenerateOutput, error) {
	return &model.GenerateOutput{
		Packages: []model.PackageResult{
			{
				PackageName: "",
				Path:        nil,
				ReadmeMd:    nil,
				Changelog: model.Changelog{
					Content:           "",
					HasBreakingChange: false,
				},
				Artifacts:           nil,
				InstallInstructions: model.InstallInstructionScriptOutput{},
			},
		},
	}, nil
}
