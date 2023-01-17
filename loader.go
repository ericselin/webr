package webr

import (
	"html/template"
	"io/fs"
)

type Templates struct {
	templatesFS  fs.FS
	baseTemplate *template.Template
}

func InitTemplates(templatesFS fs.FS) *Templates {
	rootTemplate := template.Must(template.New("_root").Parse(`{{template "base.html" .}}`))
	baseTemplate := template.Must(rootTemplate.ParseFS(templatesFS, "views/components/*.html"))
	return &Templates{
		templatesFS:  templatesFS,
		baseTemplate: baseTemplate,
	}
}

func (t *Templates) LoadTemplate(name string) *template.Template {
	tmpl, err := template.Must(t.baseTemplate.Clone()).ParseFS(t.templatesFS, "views/"+name)
	if err != nil {
		panic(err)
	}
	return tmpl
}
