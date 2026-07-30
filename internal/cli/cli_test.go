package cli

import (
	"bytes"
	"os"
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
