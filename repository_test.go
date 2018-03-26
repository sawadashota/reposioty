package reposioty

import (
	"testing"
	"path"
	"os"
)

func TestGet(t *testing.T) {
	srcPath := path.Join(os.Getenv("GOPATH"), "src")

	dirPaths, err := Get(srcPath)

	if err != nil {
		t.Error(err.Error())
	}

	if len(dirPaths) < 1 {
		t.Error("Repository directory should be more than 1")
	}

	for _, dirPath := range dirPaths {
		if !isGitRepo(dirPath) {
			t.Errorf("%v is not Git repository\n", dirPath)
		}
	}
}

func TestGet2(t *testing.T) {
	srcPath := path.Join(os.Getenv("GOPATH"), "hoge-hoge-fuga-fuga")

	dirPaths, err := Get(srcPath)

	if err == nil {
		t.Errorf("%v is not exist. So that it should be error\n", srcPath)
	}

	if len(dirPaths) != 0 {
		t.Error("Directoies should be empty when error occured")
	}
}

func TestGet3(t *testing.T) {
	srcPath := path.Join(os.Getenv("GOPATH"), "src/github.com/sawadashota/reposioty/repository.go")

	dirPaths, err := Get(srcPath)

	if err == nil {
		t.Errorf("%v is file. So that it should be error\n", srcPath)
	}

	if len(dirPaths) != 0 {
		t.Error("Directoies should be empty when error occured")
	}
}