package gogitlab

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProjects(t *testing.T) {
	ts, gitlab := Stub("stubs/projects/index.json")
	projects, err := gitlab.Projects()

	assert.Equal(t, err, nil)
	assert.Equal(t, len(projects), 2)
	defer ts.Close()
}

func TestProject(t *testing.T) {
	ts, gitlab := Stub("stubs/projects/show.json")
	project, err := gitlab.Project("1")

	assert.Equal(t, err, nil)
	assert.IsType(t, new(Project), project)
	assert.Equal(t, project.SshRepoUrl, "git@example.com:diaspora/diaspora-project-site.git")
	assert.Equal(t, project.HttpRepoUrl, "http://example.com/diaspora/diaspora-project-site.git")
	defer ts.Close()
}

func TestProjectCreation(t *testing.T) {
	ts, gitlab := Stub("stubs/projects/post.json")
	defer ts.Close()

	req := &ProjectRequest{

		Name:      "test-project",
		Namespace: 1,
	}
	project, err := gitlab.AddProject(req)
	assert.Equal(t, err, nil)
	assert.Equal(t, project.Id, 3)
	assert.Equal(t, project.Namespace.Id, 3)
}

func TestProjectBranches(t *testing.T) {
	ts, gitlab := Stub("stubs/projects/branches/index.json")
	branches, err := gitlab.ProjectBranches("1")

	assert.Equal(t, err, nil)
	assert.Equal(t, len(branches), 2)
	defer ts.Close()
}

func TestProjectBranchCreation(t *testing.T) {

	ts, gitlab := Stub("stubs/projects/branches/post.json")
	req := &ProjectBranchRequest{
		BranchName: "production",
		Ref:        "master",
	}

	commit, err := gitlab.CreateBranchForProject("1", req)
	assert.Equal(t, err, nil)
	assert.NotEmpty(t, commit)
	defer ts.Close()

}

func TestProtectBranch(t *testing.T) {
	ts, gitlab := Stub("stubs/projects/branches/protect.json")

	commit, err := gitlab.ProtectBranch("1", "production")
	assert.Equal(t, err, nil)
	assert.NotEmpty(t, commit)

	defer ts.Close()

}
