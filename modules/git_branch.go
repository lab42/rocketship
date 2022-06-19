package modules

import (
	"strings"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
	"github.com/lab42/rocketship/config"
	"github.com/lab42/rocketship/formatter"
)

const GIT_BRANCH_MODULE = "git_branch"

func Git_branch_module(config *config.Config) (string, error) {
	if config.GitBranch.Disabled {
		return "", nil
	}

	repo, err := git.PlainOpen(".")
	if err != nil {
		return "", nil
	}

	ref, err := repo.Reference(plumbing.HEAD, false)
	if err != nil {
		return "", nil
	}

	formatter := formatter.NewFormatter(
		GIT_BRANCH_MODULE,
		config.GitBranch.Format,
		map[string]string{
			"git_branch": strings.Split(ref.Target().String(), "refs/heads/")[1],
			"symbol":     config.GitBranch.Symbol,
		},
	)

	return formatter.Render()
}

func init() {
	AddModule(NewModule(
		GIT_BRANCH_MODULE,
		Git_branch_module,
	))
}
