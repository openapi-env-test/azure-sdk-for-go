package model

import (
	"encoding/json"
	"io"
	"io/ioutil"
)

type GenerateInput struct {
	DryRun                  bool                          `json:"dryRun,omitempty,omitempty"`
	SpecFolder              string                        `json:"specFolder,omitempty"`
	HeadSha                 string                        `json:"headSha,omitempty"`
	HeadRef                 string                        `json:"headRef,omitempty"`
	RepoHttpsUrl            string                        `json:"repoHttpsUrl,omitempty"`
	Trigger                 string                        `json:"trigger,omitempty"`
	ChangedFiles            []string                      `json:"changedFiles,omitempty"`
	RelatedReadmeMdFiles    []string                      `json:"relatedReadmeMdFiles,omitempty"`
	InstallInstructionInput InstallInstructionScriptInput `json:"installInstructionInput,omitempty"`
}

func NewGenerateInputFrom(reader io.Reader) (*GenerateInput, error) {
	b, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	var result GenerateInput
	if err := json.Unmarshal(b, &result); err != nil {
		return nil, err
	}
	return &result, nil
}

type InstallInstructionScriptInput struct {
	PackageName             string   `json:"packageName,omitempty"`
	Artifacts               []string `json:"artifacts,omitempty"`
	IsPublic                bool     `json:"isPublic,omitempty"`
	DownloadUrlPrefix       string   `json:"downloadUrlPrefix,omitempty"`
	DownloadCommandTemplate string   `json:"downloadCommandTemplate,omitempty"`
	Trigger                 string   `json:"trigger,omitempty"`
}

type GenerateOutput struct {
	Packages []PackageResult `json:"packages,omitempty"`
}

func (o GenerateOutput) WriteTo(writer io.Writer) (int64, error) {
	b, err := json.Marshal(o)
	if err != nil {
		return 0, err
	}
	i, err := writer.Write(b)
	return int64(i), err
}

func GetPackage() string {
	return ""
}

type PackageResult struct {
	PackageName         string                         `json:"packageName,omitempty"`
	Path                []string                       `json:"path,omitempty"`
	ReadmeMd            []string                       `json:"readmeMd,omitempty"`
	Changelog           Changelog                      `json:"changelog,omitempty"`
	Artifacts           []string                       `json:"artifacts,omitempty"`
	InstallInstructions InstallInstructionScriptOutput `json:"installInstructions,omitempty"`
}

type Changelog struct {
	Content           string `json:"content,omitempty"`
	HasBreakingChange bool   `json:"hasBreakingChange,omitempty"`
}

type InstallInstructionScriptOutput struct {
	Full string `json:"full,omitempty"`
	Lite string `json:"lite,omitempty"`
}
