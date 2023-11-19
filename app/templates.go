package app

import (
	htmpl "html/template"
	"io/fs"
	"path"
)

func compileHTMLTemplates(templateRoot fs.FS, dir string) (*htmpl.Template, error) {
	return htmpl.ParseFS(templateRoot, path.Join(dir, "*.html"))
}
