package walk

import (
	"fmt"
	"os"
	"path/filepath"

	"golang.org/x/sync/errgroup"
)

// FileAction is an interface for an action to be done on each file that isnt
// a directory
// FileAction has to implement its own thread safety
type FileAction interface {
	Do(filePath string) error
}

// Walk: recursively walks through file paths and doing an action on each file
// that isnt a directory
// walk.Walk calls walk.step in a go routine
func Walk(pathName string, fa FileAction) error {
	// errgroup for go routines and error propegation
	eg := new(errgroup.Group)
	// launch step in its own routine
	eg.Go(func() error {
		return step(pathName, eg, fa)
	})

	// wait for all go routines to finish
	err := eg.Wait()
	if err != nil {
		return fmt.Errorf("walk.Walk error: %w", err)
	}

	return nil
}

// step: run FileAction.Do on each file, recursively call walk.step() with each
// new directory
func step(pathName string, eg *errgroup.Group, fa FileAction) error {
	// open directory file descriptor
	descriptor, err := os.Open(pathName)
	if err != nil {
		return fmt.Errorf("walk.step Error: %w", err)
	}

	// read all the the files inside of current directory
	files, err := descriptor.ReadDir(-1)
	if err != nil {
		return fmt.Errorf("walk.step Error: %w", err)
	}

	// loop over each file
	for _, file := range files {

		// get the absolute path name
		absoluteName := filepath.Join(pathName, file.Name())

		// launch fa.Do in go routine
		eg.Go(func() error {
			return fa.Do(absoluteName)
		})

		// recursively call so it crawls the entire path
		if file.IsDir() {
			err = step(absoluteName, eg, fa)
			if err != nil {
				return fmt.Errorf("walk.step Error: %w", err)
			}
		}
	}

	// close file descriptor
	err = descriptor.Close()
	// checking error because closing can err
	if err != nil {
		return fmt.Errorf("walk.Walk Error: %w", err)
	}

	return nil
}
