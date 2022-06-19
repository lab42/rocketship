package shell

import "embed"

//go:embed *.shell
var ShellFiles embed.FS
