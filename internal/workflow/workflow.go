package workflow

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type State string

const (
	Ready            State = "ready"
	Planned          State = "planned"
	InProgress       State = "implementation-in-progress"
	Complete         State = "implementation-complete"
	ChangesRequested State = "review-changes-requested"
	Approved         State = "review-approved"
	Blocked          State = "review-blocked"
	Invalid          State = "invalid"
)

type Report struct {
	Project, RoadmapItem, TaskStatus, LatestReview, ReviewOutcome, CapabilityImpact, Next string
	State                                                                                 State
	Diagnostics                                                                           []string
	OperationalError                                                                      error
}

func (r Report) String() string {
	var b strings.Builder
	field := func(k, v string) {
		if v != "" {
			fmt.Fprintf(&b, "%s: %s\n", k, v)
		}
	}
	field("Project", r.Project)
	field("Active roadmap item", r.RoadmapItem)
	field("Phase", string(r.State))
	field("Task status", r.TaskStatus)
	field("Latest review", r.LatestReview)
	field("Review outcome", r.ReviewOutcome)
	field("Capability impact", r.CapabilityImpact)
	for _, d := range r.Diagnostics {
		fmt.Fprintf(&b, "Diagnostic: %s\n", d)
	}
	field("Next", r.Next)
	return b.String()
}

type taskMeta struct {
	ID         string     `yaml:"id"`
	Title      string     `yaml:"title"`
	RoadmapID  string     `yaml:"roadmap-id"`
	Status     string     `yaml:"status"`
	Created    string     `yaml:"created"`
	Updated    string     `yaml:"updated"`
	Remediates string     `yaml:"remediates-review"`
	Impact     impact     `yaml:"capability-impact"`
	Resolution resolution `yaml:"blocked-review-resolution"`
}
type impact struct {
	Type      string   `yaml:"type"`
	IDs       []string `yaml:"ids"`
	Rationale string   `yaml:"rationale"`
}
type resolution struct {
	Review     string   `yaml:"review"`
	Route      string   `yaml:"route"`
	RecordedBy string   `yaml:"recorded-by"`
	Evidence   []string `yaml:"evidence"`
}
type reviewMeta struct {
	TaskID  string `yaml:"task-id"`
	Review  int    `yaml:"review"`
	Status  string `yaml:"status"`
	Created string `yaml:"created"`
	Persona string `yaml:"persona"`
}
type roadItem struct{ Status string }

var itemHeading = regexp.MustCompile(`(?m)^## ([A-Z][A-Z0-9-]*-[0-9]+)\s+—`)
var reviewName = regexp.MustCompile(`^review-([0-9]{2})\.md$`)

func Detect(root string) Report {
	r := Report{State: Invalid, Next: "repair the reported artifacts, then run concoct status"}
	roadData, err := os.ReadFile(filepath.Join(root, ".concoct", "roadmap.md"))
	if err != nil {
		r.OperationalError = fmt.Errorf("read .concoct/roadmap.md: %w", err)
		return r
	}
	capData, err := os.ReadFile(filepath.Join(root, ".concoct", "capabilities.md"))
	if err != nil {
		r.OperationalError = fmt.Errorf("read .concoct/capabilities.md: %w", err)
		return r
	}
	var roadHead, capHead map[string]any
	if err = parseFront(roadData, &roadHead); err != nil {
		r.Diagnostics = append(r.Diagnostics, ".concoct/roadmap.md: "+err.Error())
		return r
	}
	if err = parseFront(capData, &capHead); err != nil {
		r.Diagnostics = append(r.Diagnostics, ".concoct/capabilities.md: "+err.Error())
		return r
	}
	r.Project = stringValue(roadHead["project"])
	if r.Project == "" {
		r.Diagnostics = append(r.Diagnostics, ".concoct/roadmap.md: missing project metadata")
		return r
	}
	if cp := stringValue(capHead["project"]); cp == "" || cp != r.Project {
		r.Diagnostics = append(r.Diagnostics, ".concoct/capabilities.md: project metadata must match roadmap project")
		return r
	}
	items, diags := parseRoadmap(string(roadData))
	if len(diags) > 0 {
		r.Diagnostics = append(r.Diagnostics, diags...)
		return r
	}
	cur := filepath.Join(root, ".concoct", "current")
	taskPath := filepath.Join(cur, "task-plan.md")
	notesPath := filepath.Join(cur, "notes.md")
	taskData, taskPop, err := readPopulated(taskPath)
	if err != nil {
		r.OperationalError = err
		return r
	}
	_, notesPop, err := readNotes(notesPath)
	if err != nil {
		r.OperationalError = err
		return r
	}
	reviews, revDiags, opErr := readReviews(cur)
	if opErr != nil {
		r.OperationalError = opErr
		return r
	}
	if len(revDiags) > 0 {
		r.Diagnostics = append(r.Diagnostics, revDiags...)
		return r
	}
	if !taskPop && !notesPop {
		if len(reviews) > 0 {
			r.Diagnostics = append(r.Diagnostics, ".concoct/current: reviews exist without an active task")
			return r
		}
		r.State = Ready
		r.Next = "concoct roadmap or concoct plan <roadmap-id>"
		return r
	}
	if taskPop != notesPop {
		r.Diagnostics = append(r.Diagnostics, ".concoct/current/task-plan.md and notes.md must both be populated")
		return r
	}
	var task taskMeta
	if err = parseFront(taskData, &task); err != nil {
		r.Diagnostics = append(r.Diagnostics, ".concoct/current/task-plan.md: "+err.Error())
		return r
	}
	r.RoadmapItem = task.RoadmapID
	r.TaskStatus = task.Status
	r.CapabilityImpact = task.Impact.Type
	if task.ID == "" || task.RoadmapID == "" || task.ID != task.RoadmapID {
		r.Diagnostics = append(r.Diagnostics, ".concoct/current/task-plan.md: fields id and roadmap-id must be present and equal")
		return r
	}
	if task.Title == "" {
		r.Diagnostics = append(r.Diagnostics, ".concoct/current/task-plan.md: missing required field title")
		return r
	}
	if task.Created == "" {
		r.Diagnostics = append(r.Diagnostics, ".concoct/current/task-plan.md: missing required field created")
		return r
	}
	if task.Updated == "" {
		r.Diagnostics = append(r.Diagnostics, ".concoct/current/task-plan.md: missing required field updated")
		return r
	}
	item, ok := items[task.RoadmapID]
	if !ok {
		r.Diagnostics = append(r.Diagnostics, "roadmap item "+task.RoadmapID+" does not exist")
		return r
	}
	if item.Status != "active" {
		r.Diagnostics = append(r.Diagnostics, "roadmap item "+task.RoadmapID+" must have Status active while its task is current (found "+item.Status+")")
		return r
	}
	if !validImpact(task.Impact) {
		r.Diagnostics = append(r.Diagnostics, ".concoct/current/task-plan.md: capability-impact must use add, update, remove, or none and include IDs when applicable")
		return r
	}
	if strings.TrimSpace(task.Impact.Rationale) == "" {
		r.Diagnostics = append(r.Diagnostics, ".concoct/current/task-plan.md: missing required field capability-impact.rationale")
		return r
	}
	if task.Status != "planned" && task.Status != "implementation-in-progress" && task.Status != "implementation-complete" {
		r.Diagnostics = append(r.Diagnostics, "unknown task status "+task.Status)
		return r
	}
	if len(reviews) == 0 {
		r.State = mapTask(task.Status)
		r.Next = next(r.State)
		return r
	}
	latest := reviews[len(reviews)-1]
	r.LatestReview = latest.name
	r.ReviewOutcome = latest.meta.Status
	if latest.meta.TaskID != task.ID {
		r.Diagnostics = append(r.Diagnostics, latest.name+": task-id does not match active task")
		return r
	}
	if task.Status == "planned" {
		r.Diagnostics = append(r.Diagnostics, "task status planned cannot coexist with reviews")
		return r
	}
	if task.Remediates != "" {
		if task.Remediates != latest.name {
			if !historicalReview(reviews, task.Remediates, "changes-requested") {
				r.Diagnostics = append(r.Diagnostics, "remediates-review must name the latest changes-requested review")
				return r
			}
			// A later review supersedes retained remediation history. Continue so
			// recovery evidence for that later review can still be evaluated.
		} else if latest.meta.Status != "changes-requested" {
			r.Diagnostics = append(r.Diagnostics, "remediates-review must name the latest changes-requested review")
			return r
		} else if task.Status == "implementation-complete" && !hasDispositions(string(mustRead(notesPath)), latest.body) {
			r.Diagnostics = append(r.Diagnostics, "notes.md lacks completed dispositions for findings in "+latest.name)
			return r
		} else {
			r.State = mapTask(task.Status)
			r.Next = next(r.State)
			return r
		}
	}
	if task.Resolution.Review != "" {
		if task.Resolution.Review != latest.name {
			if !historicalReview(reviews, task.Resolution.Review, "blocked") {
				r.Diagnostics = append(r.Diagnostics, "blocked-review-resolution must name the latest blocked review")
				return r
			}
			// A later completed review supersedes retained blocker-resolution history.
		} else if d := validateResolution(root, task, latest, string(mustRead(notesPath))); d != "" {
			r.Diagnostics = append(r.Diagnostics, d)
			return r
		} else {
			r.State = mapTask(task.Status)
			r.Next = next(r.State)
			return r
		}
	}

	if task.Status != "implementation-complete" {
		r.Diagnostics = append(r.Diagnostics, "task status must remain implementation-complete while latest review is authoritative")
		return r
	}
	switch latest.meta.Status {
	case "changes-requested":
		r.State = ChangesRequested
	case "approved":
		r.State = Approved
	case "blocked":
		r.State = Blocked
	default:
		r.Diagnostics = append(r.Diagnostics, latest.name+": unknown review outcome "+latest.meta.Status)
		return r
	}
	r.Next = next(r.State)
	return r
}

func historicalReview(reviews []review, name, outcome string) bool {
	for _, r := range reviews[:len(reviews)-1] {
		if r.name == name && r.meta.Status == outcome {
			return true
		}
	}
	return false
}

type review struct {
	name, body string
	meta       reviewMeta
}

func readReviews(cur string) ([]review, []string, error) {
	ents, err := os.ReadDir(cur)
	if err != nil {
		return nil, nil, fmt.Errorf("read .concoct/current: %w", err)
	}
	var names []string
	var diags []string
	for _, e := range ents {
		if e.IsDir() {
			continue
		}
		if strings.HasPrefix(e.Name(), "review-") {
			if !reviewName.MatchString(e.Name()) {
				diags = append(diags, e.Name()+": review filename must be zero-padded review-NN.md")
			} else {
				names = append(names, e.Name())
			}
		}
	}
	sort.Strings(names)
	out := make([]review, 0, len(names))
	for i, n := range names {
		want := i + 1
		m := reviewName.FindStringSubmatch(n)
		num, _ := strconv.Atoi(m[1])
		if num != want {
			diags = append(diags, fmt.Sprintf("%s: review sequence gap; expected review-%02d.md", n, want))
		}
		data, err := os.ReadFile(filepath.Join(cur, n))
		if err != nil {
			return nil, nil, err
		}
		var meta reviewMeta
		if err = parseFront(data, &meta); err != nil {
			diags = append(diags, n+": "+err.Error())
			continue
		}
		if meta.Review != num {
			diags = append(diags, n+": internal review number does not match filename")
		}
		if meta.TaskID == "" {
			diags = append(diags, n+": missing task-id")
		}
		if meta.Created == "" {
			diags = append(diags, n+": missing required field created")
		}
		if meta.Persona == "" {
			diags = append(diags, n+": missing required field persona")
		} else if meta.Persona != "reviewer" {
			diags = append(diags, n+": field persona must be reviewer")
		}
		if meta.Status != "approved" && meta.Status != "changes-requested" && meta.Status != "blocked" {
			diags = append(diags, n+": outcome must be approved, changes-requested, or blocked")
		}
		if outcome, count := documentedOutcome(string(data)); count != 1 || outcome != meta.Status {
			diags = append(diags, n+": must document exactly one Outcome matching front matter status")
		}
		out = append(out, review{name: n, body: string(data), meta: meta})
	}
	return out, diags, nil
}

func documentedOutcome(body string) (string, int) {
	start := strings.Index(body, "## Outcome")
	if start < 0 {
		return "", 0
	}
	section := body[start+len("## Outcome"):]
	if end := strings.Index(section, "\n## "); end >= 0 {
		section = section[:end]
	}
	found := ""
	count := 0
	for _, outcome := range []string{"approved", "changes-requested", "blocked"} {
		n := strings.Count(section, "`"+outcome+"`")
		if n > 0 {
			found = outcome
			count += n
		}
	}
	return found, count
}
func parseFront(data []byte, out any) error {
	if !bytes.HasPrefix(data, []byte("---\n")) {
		return fmt.Errorf("missing YAML front matter")
	}
	end := bytes.Index(data[4:], []byte("\n---\n"))
	if end < 0 {
		return fmt.Errorf("unterminated YAML front matter")
	}
	dec := yaml.NewDecoder(bytes.NewReader(data[4 : 4+end]))
	dec.KnownFields(true)
	if err := dec.Decode(out); err != nil {
		return fmt.Errorf("malformed YAML front matter: %w", err)
	}
	return nil
}
func parseRoadmap(s string) (map[string]roadItem, []string) {
	matches := itemHeading.FindAllStringSubmatchIndex(s, -1)
	items := map[string]roadItem{}
	var d []string
	for i, m := range matches {
		id := s[m[2]:m[3]]
		end := len(s)
		if i+1 < len(matches) {
			end = matches[i+1][0]
		}
		section := s[m[0]:end]
		re := regexp.MustCompile("(?m)^- Status: `?([^`\\n]+)`?\\s*$")
		sm := re.FindStringSubmatch(section)
		if len(sm) != 2 {
			d = append(d, ".concoct/roadmap.md: "+id+" missing Status")
			continue
		}
		if _, exists := items[id]; exists {
			d = append(d, ".concoct/roadmap.md: duplicate item "+id)
		}
		items[id] = roadItem{Status: strings.TrimSpace(sm[1])}
	}
	active := 0
	validStatuses := map[string]bool{"candidate": true, "planned": true, "active": true, "blocked": true, "delivered": true, "deferred": true, "cancelled": true}
	for id, v := range items {
		if !validStatuses[v.Status] {
			d = append(d, ".concoct/roadmap.md: "+id+" has unknown Status "+v.Status)
		}
		if v.Status == "active" {
			active++
		}
	}
	if active > 1 {
		d = append(d, ".concoct/roadmap.md: multiple active roadmap items")
	}
	return items, d
}
func readPopulated(path string) ([]byte, bool, error) {
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, fmt.Errorf("read %s: %w", path, err)
	}
	trim := strings.TrimSpace(string(data))
	return data, strings.HasPrefix(trim, "---"), nil
}

func readNotes(path string) ([]byte, bool, error) {
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return nil, false, nil
	}
	if err != nil {
		return nil, false, fmt.Errorf("read %s: %w", path, err)
	}
	text := strings.TrimSpace(string(data))
	placeholder := strings.HasPrefix(text, "# notes.md") && strings.Contains(text, "_Add decisions here._") && strings.Contains(text, "_Record meaningful verification results here._")
	return data, text != "" && !placeholder, nil
}
func validImpact(i impact) bool {
	switch i.Type {
	case "none":
		return len(i.IDs) == 0
	case "add", "update", "remove":
		return len(i.IDs) > 0
	default:
		return false
	}
}
func mapTask(s string) State {
	if s == "planned" {
		return Planned
	}
	if s == "implementation-in-progress" {
		return InProgress
	}
	return Complete
}
func next(s State) string {
	switch s {
	case Planned, InProgress, ChangesRequested:
		return "concoct code"
	case Complete:
		return "concoct review"
	case Approved:
		return "concoct archive"
	case Blocked:
		return "route the blocker to the responsible role or human"
	default:
		return ""
	}
}
func validateResolution(root string, t taskMeta, r review, notes string) string {
	x := t.Resolution
	if x.Review != r.name || r.meta.Status != "blocked" {
		return "blocked-review-resolution must name the latest blocked review"
	}
	if x.Route != "code" && x.Route != "review" {
		return "blocked-review-resolution route must be code or review"
	}
	if x.RecordedBy != "task-planner" && x.RecordedBy != "developer" {
		return "blocked-review-resolution recorded-by is unauthorized"
	}
	if (x.Route == "code" && t.Status != "implementation-in-progress") || (x.Route == "review" && t.Status != "implementation-complete") {
		return "blocked-review-resolution route disagrees with task status"
	}
	if len(x.Evidence) == 0 {
		return "blocked-review-resolution evidence must not be empty"
	}
	for _, p := range x.Evidence {
		if filepath.IsAbs(p) || strings.Contains(p, "..") || strings.ContainsAny(p, "*?#") || strings.Contains(p, "://") {
			return "blocked-review-resolution contains an unsafe evidence path"
		}
		info, err := os.Stat(filepath.Join(root, filepath.FromSlash(p)))
		if err != nil || !info.Mode().IsRegular() {
			return "blocked-review-resolution evidence does not name a readable file: " + p
		}
	}
	if !strings.Contains(strings.ToLower(notes), "block") || !strings.Contains(strings.ToLower(notes), "handoff to reviewer") && x.Route == "review" {
		return "notes.md lacks the required blocker disposition or fresh reviewer handoff"
	}
	return ""
}
func hasDispositions(notes, review string) bool {
	count := strings.Count(review, "### Finding ")
	if count == 0 {
		return true
	}
	lower := strings.ToLower(notes)
	words := []string{"fixed", "partially fixed", "disputed", "obsolete", "blocked"}
	n := 0
	for _, w := range words {
		n += strings.Count(lower, w)
	}
	return n >= count
}
func mustRead(path string) []byte { b, _ := os.ReadFile(path); return b }
func stringValue(v any) string    { s, _ := v.(string); return s }
