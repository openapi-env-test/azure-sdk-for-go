package changelog

import (
	"github.com/Azure/azure-sdk-for-go/tools/apidiff/exports"
	"log"
	"os"
	"os/exec"
)

func NewChangelogForPackage(pkgDir string) (c *Changelog, err error) {
	// first we need to get the current status of the package
	rhs, err := getExportForPackage(pkgDir)
	if err != nil {
		return nil, err
	}
	log.Printf("Exports of package '%s': %+v", pkgDir, rhs)
	// stash everything and get the previous status of the package
	if err := stashEverything(); err != nil {
		return nil, err
	}
	// reset everything when we are done
	defer func() {
		err = resetEverything()
	}()
	// get the original state of the package
	lhs, err := getExportForPackage(pkgDir)
	if err != nil {
		return nil, err
	}
	log.Printf("Exports of original package '%s': %+v", pkgDir, lhs)
	return &Changelog{}, nil
}

func stashEverything() error {
	if err := gitAddAll(); err != nil {
		return err
	}
	if err := gitStash(); err != nil {
		return err
	}
	return nil
}

func resetEverything() error {
	if err := gitStashPop(); err != nil {
		return err
	}
	if err := gitResetHead(); err != nil {
		return err
	}
	return nil
}

func gitAddAll() error {
	log.Printf("Executing `git add *`")
	c := exec.Command("git", "add", "*")
	return c.Run()
}

func gitStash() error {
	log.Printf("Executing `git stash`")
	c := exec.Command("git", "stash")
	return c.Run()
}

func gitStashPop() error {
	log.Printf("Executing `git stash pop`")
	c := exec.Command("git", "stash", "pop")
	return c.Run()
}

func gitResetHead() error {
	log.Printf("Executing `git reset HEAD`")
	c := exec.Command("git", "reset", "HEAD")
	return c.Run()
}

func getChangelogForPackage(lhs, rhs *exports.Content) {

}

func getExportForPackage(pkgDir string) (*exports.Content, error) {
	// The function exports.Get does not handle the circumstance that the package does not exist
	// therefore we have to check if it exists and if not exit early to ensure we do not return an error
	if _, err := os.Stat(pkgDir); os.IsNotExist(err) {
		return nil, nil
	}
	exp, err := exports.Get(pkgDir)
	if err != nil {
		return nil, err
	}
	return &exp, nil
}
