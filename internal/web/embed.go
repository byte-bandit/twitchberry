package web

import "embed"

//go:embed static
//go:embed index.html
var Content embed.FS
