package modules

import (
	"github.com/go-git/go-git/v5"
	"github.com/lab42/rocketship/config"
	"github.com/lab42/rocketship/formatter"
)

const GIT_BRANCH_MODULE = "git_branch"

func git_branch_module(config *config.Config) (string, error) {
	if config.GitBranch.Disabled {
		return "", nil
	}

	repo, err := git.PlainOpen(".")
	if err != git.ErrRepositoryNotExists {
		return "", err
	}

	head, err := repo.Head()
	if err != nil {
		return "", err
	}

	formatter := formatter.NewFormatter(
		GIT_BRANCH_MODULE,
		config.Directory.Format,
		map[string]string{
			"git_branch": head.Name().String(),
			"symbol":     config.Directory.Symbol,
		},
	)

	return formatter.Render()
}

func init() {
	AddModule(NewModule(
		GIT_BRANCH_MODULE,
		git_branch_module,
	))
}
