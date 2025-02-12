package testkube

// GitAuthType defines git auth type
type GitAuthType string

const (
	// GitAuthTypeBasic for git basic auth requests
	GitAuthTypeBasic GitAuthType = "basic"
	// GitAuthTypeHeader for git header auth requests
	GitAuthTypeHeader GitAuthType = "header"
)

// NewGitRepository is a constructor for new repository
func NewGitRepository(uri, branch string) *Repository {
	return &Repository{
		Type_:  "git",
		Branch: branch,
		Uri:    uri,
	}
}

// WithPath supplies path for repository
func (r *Repository) WithPath(path string) *Repository {
	r.Path = path
	return r
}

// WithCommit supplies commit for repository
func (r *Repository) WithCommit(commit string) *Repository {
	r.Commit = commit
	return r
}

// WithAuthType supplies auth type for repository
func (r *Repository) WithAuthType(authType GitAuthType) *Repository {
	r.AuthType = string(authType)
	return r
}
