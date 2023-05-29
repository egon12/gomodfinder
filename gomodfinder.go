// Package gomodfinder package that contain function to find absolute path for go.mod
package gomodfinder

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

const maxRecursive = 100

// Find find absolutepath that have go.mod in parent folder.
// only test works in UNIX. Maybe wouldn't work in windows.
func Find() (string, error) {
	p, err := filepath.Abs(".")
	if err != nil {
		return "", fmt.Errorf("can't find go.mod in parent ancestor: cannot find absolute path of '.'")
	}

	return findRecursiveGoMod(p, 0)
}

func findRecursiveGoMod(currentPath string, recursive int) (string, error) {
	// if to much recursive
	if recursive > maxRecursive {
		return "", fmt.Errorf("can't find go.mod in parent ancestor: '%s' nested more than %d level", currentPath, maxRecursive)
	}

	// if you already in root
	if currentPath == "/" {
		return "", fmt.Errorf("can't find go.mod in parent ancestor")
	}

	files, err := ioutil.ReadDir(currentPath)
	if err != nil {
		return "", fmt.Errorf("can't find go.mod in parent ancestor: cannot read in '%s'", currentPath)
	}

	for _, f := range files {
		if f.Name() == "go.mod" {
			return currentPath, nil
		}
	}

	newPath, err := filepath.Abs(currentPath + "/../")
	if err != nil {
		return "", fmt.Errorf("can't find go.mod in parent ancestor: cannot find absolute path of %s", currentPath+"/../")
	}

	// maybe looping directory
	if currentPath == newPath {
		return "", fmt.Errorf("can't find go.mod in parent ancestor: stuck in %s", newPath)
	}

	return findRecursiveGoMod(newPath, recursive+1)
}
