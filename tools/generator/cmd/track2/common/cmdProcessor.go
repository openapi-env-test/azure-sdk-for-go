// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package common

import (
	"fmt"
	"log"
	"os/exec"
)

func ExecuteGoGenerate(path string) error {
	cmd := exec.Command("go", "generate")
	log.Printf("execute path: %s", path)
	cmd.Dir = path
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute go generate '%s': %+v", string(output), err)
	}
	log.Printf("go generate run result: %s", string(output))
	return nil
}
