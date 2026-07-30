package gitrepo

import (
	"os/exec"
	"strings"
	"testing"
)

func TestTaskBranch(t *testing.T) {
	tests := map[string]string{
		TaskBranch("CON-015", "Isolate and integrate tasks with Git branches"): "concoct/con-015-isolate-and-integrate-tasks-with-git-branches",
		TaskBranch("APP-1", "Spaces / punctuation !!!"):                        "concoct/app-1-spaces-punctuation",
	}
	for got, want := range tests {
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	if got := TaskBranch("APP-999", "A very long title that must be truncated deterministically without an unsafe trailing separator"); len(got) > 64 {
		t.Fatalf("branch too long: %q", got)
	}
}

func TestCreateTaskBranchAndCollision(t *testing.T) {
	root := t.TempDir()
	run(t, root, "init", "-q", "-b", "develop")
	run(t, root, "config", "user.email", "test@example.com")
	run(t, root, "config", "user.name", "Test")
	run(t, root, "commit", "--allow-empty", "-qm", "base")
	r, ok, err := Open(root)
	if err != nil || !ok {
		t.Fatalf("open: %v %v", ok, err)
	}
	start, err := r.CreateTaskBranch("APP-1", "Demo")
	if err != nil {
		t.Fatal(err)
	}
	if start.Trunk != "develop" || start.Branch != "concoct/app-1-demo" {
		t.Fatalf("start = %#v", start)
	}
	if err := r.Checkout("develop"); err != nil {
		t.Fatal(err)
	}
	before, _ := r.Head()
	if _, err := r.CreateTaskBranch("APP-1", "Demo"); err == nil || !strings.Contains(err.Error(), "collision") {
		t.Fatalf("error = %v", err)
	}
	branch, _ := r.Branch()
	head, _ := r.Head()
	if branch != "develop" || head != before {
		t.Fatal("collision changed repository")
	}
}

func run(t *testing.T, root string, args ...string) {
	t.Helper()
	c := exec.Command("git", args...)
	c.Dir = root
	if out, err := c.CombinedOutput(); err != nil {
		t.Fatalf("git %s: %v: %s", strings.Join(args, " "), err, out)
	}
}
