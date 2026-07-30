package cli

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/cthain/concoct/internal/project"
)

func TestPromptStdoutAndFileOutputAreIdenticalAndNonDestructive(t *testing.T) {
	parent := t.TempDir()
	var initOutput bytes.Buffer
	if err := project.Initialize(parent, "demo", &initOutput); err != nil {
		t.Fatal(err)
	}
	root := filepath.Join(parent, "demo")
	nested := filepath.Join(root, "doc", "nested")
	if err := os.MkdirAll(nested, 0o755); err != nil {
		t.Fatal(err)
	}
	t.Setenv("CONCOCT_CALLER_DIR", nested)

	before := workflowSnapshot(t, root)
	var stdout, stderr bytes.Buffer
	if err := Run([]string{"roadmap"}, &stdout, &stderr); err != nil {
		t.Fatalf("render stdout: %v (%s)", err, stderr.String())
	}
	output := filepath.Join(t.TempDir(), "roadmap-prompt.md")
	if err := Run([]string{"roadmap", "--output", output}, &bytes.Buffer{}, &stderr); err != nil {
		t.Fatal(err)
	}
	fileBytes, err := os.ReadFile(output)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(stdout.Bytes(), fileBytes) {
		t.Fatal("stdout and file output differ")
	}
	if before != workflowSnapshot(t, root) {
		t.Fatal("prompt rendering changed workflow artifacts")
	}
	if err := Run([]string{"roadmap", "--output", output}, &bytes.Buffer{}, &stderr); err == nil || !strings.Contains(err.Error(), "without overwriting") {
		t.Fatalf("existing output error = %v", err)
	}
	if got, _ := os.ReadFile(output); !bytes.Equal(got, fileBytes) {
		t.Fatal("existing output was modified")
	}
}

func TestPromptArgumentValidation(t *testing.T) {
	tests := [][]string{{"plan"}, {"code", "extra"}, {"review", "--output"}, {"roadmap", "--output", "a", "--output", "b"}}
	for _, args := range tests {
		if err := Run(args, &bytes.Buffer{}, &bytes.Buffer{}); err == nil {
			t.Errorf("Run(%v) succeeded", args)
		}
	}
}

func TestPlanCreatesDeterministicTaskBranchAndRefusesCollision(t *testing.T) {
	parent := t.TempDir()
	if err := project.Initialize(parent, "demo", &bytes.Buffer{}); err != nil {
		t.Fatal(err)
	}
	root := filepath.Join(parent, "demo")
	runGit(t, root, "config", "user.email", "test@example.com")
	runGit(t, root, "config", "user.name", "Test")
	road := filepath.Join(root, ".concoct/roadmap.md")
	f, err := os.OpenFile(road, os.O_APPEND|os.O_WRONLY, 0)
	if err != nil {
		t.Fatal(err)
	}
	if _, err := f.WriteString("\n## APP-001 — Branch Demo\n\n- Status: `planned`\n- Depends on: `none`\n"); err != nil {
		t.Fatal(err)
	}
	_ = f.Close()
	runGit(t, root, "add", "-A")
	runGit(t, root, "commit", "-qm", "initial")
	trunk := gitOutput(t, root, "branch", "--show-current")
	t.Setenv("CONCOCT_CALLER_DIR", root)
	var out bytes.Buffer
	if err := Run([]string{"plan", "APP-001"}, &out, &bytes.Buffer{}); err != nil {
		t.Fatal(err)
	}
	if branch := gitOutput(t, root, "branch", "--show-current"); branch != "concoct/app-001-branch-demo" {
		t.Fatalf("branch = %s", branch)
	}
	if !strings.Contains(out.String(), "Git trunk:") || !strings.Contains(out.String(), "Git task base:") {
		t.Fatal("prompt lacks recorded Git start")
	}
	runGit(t, root, "checkout", trunk)
	before := gitOutput(t, root, "rev-parse", "HEAD")
	if err := Run([]string{"plan", "APP-001"}, &bytes.Buffer{}, &bytes.Buffer{}); err == nil || !strings.Contains(err.Error(), "collision") {
		t.Fatalf("error = %v", err)
	}
	if gitOutput(t, root, "branch", "--show-current") != trunk || gitOutput(t, root, "rev-parse", "HEAD") != before {
		t.Fatal("collision changed checkout")
	}
}

func runGit(t *testing.T, root string, args ...string) { t.Helper(); _ = gitOutput(t, root, args...) }
func gitOutput(t *testing.T, root string, args ...string) string {
	t.Helper()
	c := exec.Command("git", args...)
	c.Dir = root
	out, err := c.CombinedOutput()
	if err != nil {
		t.Fatalf("git %s: %v: %s", strings.Join(args, " "), err, out)
	}
	return strings.TrimSpace(string(out))
}

func workflowSnapshot(t *testing.T, root string) string {
	t.Helper()
	var result strings.Builder
	base := filepath.Join(root, ".concoct")
	if err := filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.Mode().IsRegular() {
			data, readErr := os.ReadFile(path)
			if readErr != nil {
				return readErr
			}
			result.WriteString(path)
			result.WriteByte(':')
			result.Write(data)
			result.WriteByte('\n')
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	return result.String()
}
