package cli

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/cthain/concoct/internal/project"
	"github.com/cthain/concoct/internal/prompt"
	"github.com/cthain/concoct/internal/workflow"
)

const usage = `Usage:
  concoct init <project>
  concoct status
  concoct roadmap [--output <path>]
  concoct plan <roadmap-id> [--output <path>]
  concoct code [--output <path>]
  concoct review [--output <path>]
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
	case "roadmap", "plan", "code", "review":
		return runPrompt(args, stdout, stderr)
	default:
		fmt.Fprint(stderr, usage)
		return fmt.Errorf("unknown command %q", args[0])
	}
}

func runPrompt(args []string, stdout, stderr io.Writer) error {
	command := args[0]
	positional, output, err := parsePromptArgs(args[1:])
	if err != nil {
		fmt.Fprint(stderr, usage)
		return err
	}
	roadmapID := ""
	if command == "plan" {
		if len(positional) != 1 || strings.TrimSpace(positional[0]) == "" {
			fmt.Fprint(stderr, usage)
			return fmt.Errorf("plan requires exactly one non-empty roadmap id")
		}
		roadmapID = positional[0]
	} else if len(positional) != 0 {
		fmt.Fprint(stderr, usage)
		return fmt.Errorf("%s accepts no positional arguments", command)
	}
	base, err := callerDir()
	if err != nil {
		return err
	}
	root, err := project.Discover(base)
	if err != nil {
		return err
	}
	content, err := prompt.Render(root, prompt.Request{Command: command, RoadmapID: roadmapID})
	if err != nil {
		return err
	}
	if output == "" {
		_, err = stdout.Write(content)
		return err
	}
	if !filepath.IsAbs(output) {
		output = filepath.Join(base, output)
	}
	file, err := os.OpenFile(output, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
	if err != nil {
		return fmt.Errorf("create output %s without overwriting: %w", output, err)
	}
	wrote := false
	defer func() {
		if !wrote {
			_ = os.Remove(output)
		}
	}()
	if _, err = file.Write(content); err != nil {
		_ = file.Close()
		return fmt.Errorf("write output %s: %w", output, err)
	}
	if err = file.Close(); err != nil {
		return fmt.Errorf("close output %s: %w", output, err)
	}
	wrote = true
	return nil
}

func parsePromptArgs(args []string) ([]string, string, error) {
	var positional []string
	output := ""
	for i := 0; i < len(args); i++ {
		if args[i] != "--output" {
			positional = append(positional, args[i])
			continue
		}
		if output != "" || i+1 >= len(args) || strings.TrimSpace(args[i+1]) == "" {
			return nil, "", fmt.Errorf("--output requires exactly one non-empty path")
		}
		output = args[i+1]
		i++
	}
	return positional, output, nil
}

func callerDir() (string, error) {
	if dir := os.Getenv("CONCOCT_CALLER_DIR"); dir != "" {
		return filepath.Abs(dir)
	}
	return os.Getwd()
}
