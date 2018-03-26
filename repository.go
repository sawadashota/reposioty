package reposioty

import (
	"path"
	"os"
	"io/ioutil"
)

func Get(baseDir string) ([]string, error) {
	_, err := ioutil.ReadDir(baseDir)
	if err != nil {
		return []string{}, err
	}

	return repositories(baseDir), nil
}

// Get all Git Repositories under base directory
func repositories(dir string) []string {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	var paths []string
	for _, file := range files {
		abs := path.Join(dir, file.Name())

		if !file.IsDir() {
			continue
		}

		if isGitRepo(abs) {
			paths = append(paths, abs)
			continue
		}

		paths = append(paths, repositories(abs)...)
	}

	return paths
}

// The directory is Git repository or not
// This means the directory has .git directory or not
func isGitRepo(dir string) bool {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	for _, file := range files {
		if isDotGitDir(file) {
			return true
		}
	}

	return false
}

// The directory's name is .git or not
func isDotGitDir(dir os.FileInfo) bool {
	return dir.IsDir() && dir.Name() == ".git"
}
