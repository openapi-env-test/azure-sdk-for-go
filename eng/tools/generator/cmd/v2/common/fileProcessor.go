// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package common

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/Azure/azure-sdk-for-go/eng/tools/generator/autorest/model"
	"github.com/Azure/azure-sdk-for-go/eng/tools/generator/common"
	"github.com/Azure/azure-sdk-for-go/eng/tools/internal/exports"
	"github.com/Masterminds/semver"
)

const (
	generated_file_scan_string        = "Code generated by Microsoft"
	autorest_md_file_suffix           = "readme.md"
	autorest_md_module_version_prefix = "module-version: "
	swagger_md_module_name_prefix     = "module-name: "
)

var (
	v2BeginRegex                    = regexp.MustCompile("^```\\s*yaml\\s*\\$\\(go\\)\\s*&&\\s*\\$\\((track2|v2)\\)")
	v2EndRegex                      = regexp.MustCompile("^\\s*```\\s*$")
	newClientMethodNameRegex        = regexp.MustCompile("^New.*Client$")
	versionLineRegex                = regexp.MustCompile(`moduleVersion\s*=\s*\".*v.+"`)
	changelogPosWithPreviewRegex    = regexp.MustCompile(`##\s*(?P<version>.+)\s*\((\d{4}-\d{2}-\d{2}|Unreleased)\)`)
	changelogPosWithoutPreviewRegex = regexp.MustCompile(`##\s*(?P<version>\d+\.\d+\.\d+)\s*\((\d{4}-\d{2}-\d{2}|Unreleased)\)`)
	packageConfigRegex              = regexp.MustCompile(`\$\((package-.+)\)`)
)

type PackageInfo struct {
	Name        string
	Config      string
	SpecName    string
	RequestLink string
}

// reads from readme.go.md, parses the `track2` section to get module and package name
func ReadV2ModuleNameToGetNamespace(path string) (map[string][]PackageInfo, error) {
	result := make(map[string][]PackageInfo)
	log.Printf("Reading from readme.go.md '%s'...", path)
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	log.Printf("Parsing module and package name from readme.go.md ...")
	b, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(b), "\n")

	var start []int
	var end []int
	for i, line := range lines {
		if v2BeginRegex.MatchString(line) {
			start = append(start, i)
		}
		if len(start) != len(end) && v2EndRegex.MatchString(line) {
			end = append(end, i)
		}
	}

	if len(start) == 0 {
		return nil, fmt.Errorf("cannot find any `track2` section")
	}
	if len(start) != len(end) {
		return nil, fmt.Errorf("last `track2` section does not properly end")
	}

	s := strings.ReplaceAll(path, "\\", "/")
	s1 := strings.Split(s, "/")
	specName := s1[len(s1)-3]

	for i := range start {
		// get the content of the `track2` section
		section := lines[start[i]+1 : end[i]]
		// iterate over the rest lines, get module name
		for _, line := range section {
			if strings.HasPrefix(line, swagger_md_module_name_prefix) {
				modules := strings.Split(strings.TrimSpace(line[len(swagger_md_module_name_prefix):]), "/")
				if len(modules) != 4 {
					return nil, fmt.Errorf("cannot parse module name from `track2` section")
				}
				namespaceName := strings.TrimSuffix(strings.TrimSuffix(modules[3], "\n"), "\r")
				log.Printf("RP: %s Package: %s", modules[2], namespaceName)
				packageConfig := ""
				matchResults := packageConfigRegex.FindAllStringSubmatch(lines[start[i]], -1)
				for _, matchResult := range matchResults {
					packageConfig = matchResult[1] + ": true"
				}
				result[modules[2]] = append(result[modules[2]], PackageInfo{Name: namespaceName, Config: packageConfig, SpecName: specName})
			}
		}
	}

	return result, nil
}

// remove all sdk generated files in given path
func CleanSDKGeneratedFiles(path string) error {
	log.Printf("Removing all sdk generated files in '%s'...", path)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return err
	}

	for _, file := range files {
		if strings.HasSuffix(file.Name(), ".go") {
			fileWithPath := filepath.Join(path, file.Name())
			b, err := ioutil.ReadFile(fileWithPath)
			if err != nil {
				return err
			}

			if strings.Contains(string(b), generated_file_scan_string) {
				err = os.Remove(fileWithPath)
				if err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// replace repo commit with local path in autorest.md file
func ChangeConfigWithLocalPath(path, specPath, specRPName string) error {
	log.Printf("Replacing repo commit with local path in autorest.md ...")
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	specPath, err = filepath.Abs(specPath)
	if err != nil {
		return err
	}

	lines := strings.Split(string(b), "\n")
	for i, line := range lines {
		if strings.Contains(line, autorest_md_file_suffix) {
			lines[i] = fmt.Sprintf("- %s", filepath.Join(specPath, specRPName, "resource-manager", "readme.md"))
			lines[i+1] = fmt.Sprintf("- %s", filepath.Join(specPath, specRPName, "resource-manager", "readme.go.md"))
			break
		}
	}

	return ioutil.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
}

// replace repo URL and commit id in autorest.md file
func ChangeConfigWithCommitID(path, repoURL, commitID, specRPName string) error {
	log.Printf("Replacing repo URL and commit id in autorest.md ...")
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	lines := strings.Split(string(b), "\n")
	for i, line := range lines {
		if strings.Contains(line, autorest_md_file_suffix) {
			lines[i] = fmt.Sprintf("- %s/blob/%s/specification/%s/resource-manager/readme.md", repoURL, commitID, specRPName)
			lines[i+1] = fmt.Sprintf("- %s/blob/%s/specification/%s/resource-manager/readme.go.md", repoURL, commitID, specRPName)
			break
		}
	}

	return ioutil.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
}

// RemoveTagSet delete tag set in config file
func RemoveTagSet(path string) error {
	log.Printf("Removing tag set in autorest.md ...")
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	lines := strings.Split(string(b), "\n")
	for i, line := range lines {
		if strings.Contains(line, "tag:") {
			lines = append(lines[:i], lines[i+1:]...)
			break
		}
	}

	return ioutil.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644)
}

// get swagger rp folder name from autorest.md file
func GetSpecRpName(packageRootPath string) (string, error) {
	b, err := ioutil.ReadFile(path.Join(packageRootPath, "autorest.md"))
	if err != nil {
		return "", err
	}

	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		if strings.Contains(line, autorest_md_file_suffix) {
			allParts := strings.Split(line, "/")
			for i, part := range allParts {
				if part == "specification" {
					return allParts[i+1], nil
				}
			}
		}
	}
	return "", fmt.Errorf("cannot get sepc rp name from config")
}

// replace version: use `module-version: ` prefix to locate version in autorest.md file, use version = "v*.*.*" regrex to locate version in constants.go file
func ReplaceVersion(packageRootPath string, newVersion string) error {
	path := filepath.Join(packageRootPath, "autorest.md")
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	lines := strings.Split(string(b), "\n")
	for i, line := range lines {
		if strings.HasPrefix(line, autorest_md_module_version_prefix) {
			lines[i] = line[:len(autorest_md_module_version_prefix)] + newVersion
			break
		}
	}

	if err = ioutil.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644); err != nil {
		return err
	}

	path = filepath.Join(packageRootPath, "constants.go")
	if b, err = ioutil.ReadFile(path); err != nil {
		return err
	}
	contents := versionLineRegex.ReplaceAllString(string(b), "moduleVersion = \"v"+newVersion+"\"")

	return ioutil.WriteFile(path, []byte(contents), 0644)
}

// calculate new version by changelog using semver package
func CalculateNewVersion(changelog *model.Changelog, previousVersion string, isCurrentPreview bool) (*semver.Version, error) {
	version, err := semver.NewVersion(previousVersion)
	if err != nil {
		return nil, err
	}
	log.Printf("Lastest version is: %s", version.String())

	var newVersion semver.Version
	if version.Major() == 0 {
		// preview version calculation
		if !isCurrentPreview {
			tempVersion, err := semver.NewVersion("1.0.0")
			if err != nil {
				return nil, err
			}
			newVersion = *tempVersion
		} else if changelog.HasBreakingChanges() || changelog.Modified.HasAdditiveChanges() {
			newVersion = version.IncMinor()
		} else {
			newVersion = version.IncPatch()
		}
	} else {
		if isCurrentPreview {
			if strings.Contains(previousVersion, "beta") {
				betaNumber, err := strconv.Atoi(strings.Split(version.Prerelease(), "beta.")[1])
				if err != nil {
					return nil, err
				}
				newVersion, err = version.SetPrerelease("beta." + strconv.Itoa(betaNumber+1))
				if err != nil {
					return nil, err
				}
			} else {
				if changelog.HasBreakingChanges() {
					newVersion = version.IncMajor()
				} else if changelog.Modified.HasAdditiveChanges() {
					newVersion = version.IncMinor()
				} else {
					newVersion = version.IncPatch()
				}
				newVersion, err = newVersion.SetPrerelease("beta.1")
				if err != nil {
					return nil, err
				}
			}
		} else {
			if strings.Contains(previousVersion, "beta") {
				return nil, fmt.Errorf("must have stable previous version")
			}
			// release version calculation
			if changelog.HasBreakingChanges() {
				newVersion = version.IncMajor()
			} else if changelog.Modified.HasAdditiveChanges() {
				newVersion = version.IncMinor()
			} else {
				newVersion = version.IncPatch()
			}
		}
	}

	log.Printf("New version is: %s", newVersion.String())
	return &newVersion, nil
}

// add new changelog md to changelog file
func AddChangelogToFile(changelog *model.Changelog, version *semver.Version, packageRootPath, releaseDate string) (string, error) {
	path := filepath.Join(packageRootPath, common.ChangelogFilename)
	b, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	oldChangelog := string(b)
	newChangelog := "# Release History\n\n"
	var matchResults [][]int
	if version.Prerelease() == "" {
		matchResults = changelogPosWithoutPreviewRegex.FindAllStringSubmatchIndex(oldChangelog, -1)
	} else {
		matchResults = changelogPosWithPreviewRegex.FindAllStringSubmatchIndex(oldChangelog, -1)
	}
	additionalChangelog := changelog.ToCompactMarkdown()
	if releaseDate == "" {
		releaseDate = time.Now().Format("2006-01-02")
	}

	for _, matchResult := range matchResults {
		newChangelog = newChangelog + "## " + version.String() + " (" + releaseDate + ")\r\n" + additionalChangelog + "\r\n\r\n" + oldChangelog[matchResult[0]:]
		break
	}

	err = ioutil.WriteFile(path, []byte(newChangelog), 0644)
	if err != nil {
		return "", err
	}
	return additionalChangelog, nil
}

// replace `{{NewClientName}}`` placeholder in README.md by first func name according to `^New.+Method$` pattern
func ReplaceNewClientNamePlaceholder(packageRootPath string, exports exports.Content) error {
	path := filepath.Join(packageRootPath, "README.md")
	var clientName string
	for k, v := range exports.Funcs {
		if newClientMethodNameRegex.MatchString(k) && *v.Params == "string, azcore.TokenCredential, *arm.ClientOptions" {
			clientName = k
			break
		}
	}

	b, err := ioutil.ReadFile(path)
	if err != nil {
		return fmt.Errorf("cannot read from file '%s': %+v", path, err)
	}

	var content = strings.ReplaceAll(string(b), "{{NewClientName}}", clientName)
	return ioutil.WriteFile(path, []byte(content), 0644)
}

func UpdateModuleDefinition(packageRootPath, rpName, namespaceName string, version *semver.Version) error {
	if version.Major() > 1 {
		path := filepath.Join(packageRootPath, "go.mod")

		b, err := ioutil.ReadFile(path)
		if err != nil {
			return fmt.Errorf("cannot parse version from changelog")
		}

		lines := strings.Split(string(b), "\n")
		for i, line := range lines {
			if strings.HasPrefix(line, "module") {
				line = strings.TrimRight(line, "\r")
				parts := strings.Split(line, "/")
				if parts[len(parts)-1] != fmt.Sprintf("v%d", version.Major()) {
					lines[i] = fmt.Sprintf("module github.com/Azure/azure-sdk-for-go/sdk/resourcemanager/%s/%s/v%d", rpName, namespaceName, version.Major())
				}
				break
			}
		}
		if err = ioutil.WriteFile(path, []byte(strings.Join(lines, "\n")), 0644); err != nil {
			return err
		}
	}
	return nil
}

func UpdateOnboardChangelogVersion(packageRootPath, versionNumber string) error {
	changelogPath := filepath.Join(packageRootPath, common.ChangelogFilename)
	b, err := ioutil.ReadFile(changelogPath)
	if err != nil {
		return err
	}

	newChangelog := regexp.MustCompile("\\d\\.\\d\\.\\d").ReplaceAllString(string(b), versionNumber)
	err = ioutil.WriteFile(changelogPath, []byte(newChangelog), 0644)
	if err != nil {
		return err
	}

	return nil
}
