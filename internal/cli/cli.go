package cli

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/cthain/concoct/internal/project"
	"github.com/cthain/concoct/internal/workflow"
)

const usage = `Usage:
  concoct init <project>
  concoct status
  concoct help
`

func Run(args []string, stdout, stderr io.Writer) error {
	if len(args) == 0 || args[0] == "help" || args[0] == "-h" || args[0] == "--help" {
		fmt.Fprint(stdout, usage)
		return nil
	}
	switch args[0] {
	case "init":
		if len(args) != 2 || strings.TrimSpace(args[1]) == "" {
			fmt.Fprint(stderr, usage)
			return fmt.Errorf("init requires exactly one non-empty project target")
		}
		base, err := callerDir()
		if err != nil {
			return err
		}
		return project.Initialize(base, args[1], stdout)
	case "status":
		if len(args) != 1 {
			fmt.Fprint(stderr, usage)
			return fmt.Errorf("status accepts no positional arguments")
		}
		base, err := callerDir()
		if err != nil {
			return err
		}
		root, err := project.Discover(base)
		if err != nil {
			return err
		}
		report := workflow.Detect(root)
		fmt.Fprint(stdout, report.String())
		if report.OperationalError != nil {
			return report.OperationalError
		}
		return nil
	default:
		fmt.Fprint(stderr, usage)
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func callerDir() (string, error) {
	if dir := os.Getenv("CONCOCT_CALLER_DIR"); dir != "" {
		return filepath.Abs(dir)
	}
	return os.Getwd()
}
