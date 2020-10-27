package model

type GenerateInput struct {
	DryRun                  bool                          `json:"dryRun"`
	SpecFolder              string                        `json:"specFolder"`
	HeadSha                 string                        `json:"headSha"`
	HeadRef                 string                        `json:"headRef"`
	RepoHttpsUrl            string                        `json:"repoHttpsUrl"`
	Trigger                 string                        `json:"trigger"`
	ChangedFiles            []string                      `json:"changedFiles"`
	RelatedReadmeMdFiles    []string                      `json:"relatedReadmeMdFiles"`
	InstallInstructionInput InstallInstructionScriptInput `json:"installInstructionInput"`
}

type InstallInstructionScriptInput struct {
	PackageName             string   `json:"packageName"`
	Artifacts               []string `json:"artifacts"`
	IsPublic                bool     `json:"isPublic"`
	DownloadUrlPrefix       string   `json:"downloadUrlPrefix"`
	DownloadCommandTemplate string   `json:"downloadCommandTemplate"`
	Trigger                 string   `json:"trigger"`
}

type GenerateOutput struct {
	Packages []PackageResult `json:"packages"`
}

type PackageResult struct {
	PackageName         string                         `json:"packageName"`
	Path                []string                       `json:"path"`
	ReadmeMd            []string                       `json:"readmeMd"`
	Changelog           Changelog                      `json:"changelog"`
	Artifacts           []string                       `json:"artifacts"`
	InstallInstructions InstallInstructionScriptOutput `json:"installInstructions"`
}

type Changelog struct {
	Content           string `json:"content"`
	HasBreakingChange bool   `json:"hasBreakingChange"`
}

type InstallInstructionScriptOutput struct {
	Full string `json:"full"`
	Lite string `json:"lite"`
}
