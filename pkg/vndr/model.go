package vndr

import "fmt"

type DepEntry struct {
	ImportPath string
	Rev        string
	RepoPath   string
}

func (d DepEntry) String() string {
	if d.RepoPath != "" {
		return fmt.Sprintf("%s %s %s\n", d.ImportPath, d.Rev, d.RepoPath)
	}
	return fmt.Sprintf("%s %s\n", d.ImportPath, d.Rev)
}
