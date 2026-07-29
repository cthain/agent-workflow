package concoct

import "embed"

// Templates contains the complete project distribution. The all: prefix includes
// dotfiles, which are part of Concoct's installed contract.
//
//go:embed all:templates
var Templates embed.FS
