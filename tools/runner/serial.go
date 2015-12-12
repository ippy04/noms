package runner

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/attic-labs/noms/d"
)

// Env is a map of env vars, mapping key string to value string.
type Env map[string]string

func (e Env) toStrings() (out []string) {
	out = os.Environ()
	// Sadly, it seems like we need to force-set GOROOT in the environment to handle some funky runtime environments (e.g. on our Travis setup)
	if _, overridden := e["GOROOT"]; !overridden {
		e["GOROOT"] = runtime.GOROOT()
	}
	for n, v := range e {
		out = append(out, fmt.Sprintf("%s=%s", n, v))
	}
	return
}

// ForceRun runs 'exe [args...]' in current working directory, and d.Chk()s on failure. Inherits the environment of the current process.
func ForceRun(exe string, args ...string) {
	err := runEnvDir(os.Stdout, os.Stderr, Env{}, "", exe, args...)
	d.Chk.NoError(err)
}

// runEnvDir 'exe [args...]' in dir with the environment env overlaid on that of the current process. If dir == "", use the current working directory.
func runEnvDir(out, err io.Writer, env Env, dir, exe string, args ...string) error {
	cmd := exec.Command(exe, args...)
	cmd.Dir = dir
	cmd.Env = env.toStrings()
	cmd.Stdout = out
	cmd.Stderr = err
	return cmd.Run()
}

// Serial serially runs all instances of filename found under dir, mapping stdout and stderr to each subprocess in the obvious way. env is overlaid on the environment of the current process. If args are provided, they're passed en masse to each subprocess.
func Serial(stdout, stderr io.Writer, env Env, dir, filename string, args ...string) bool {
	success := true
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		d.Exp.NoError(err, "Failed directory traversal at %s", path)
		if !info.IsDir() && filepath.Base(path) == filename {
			scriptAndArgs := append([]string{filepath.Base(path)}, args...)
			runErr := runEnvDir(stdout, stderr, env, filepath.Dir(path), "python", scriptAndArgs...)
			if runErr != nil {
				success = false
				fmt.Fprintf(stderr, "Running %s failed with %v", path, runErr)
			}
		}
		return nil
	})
	d.Exp.NoError(err)
	return success
}
