package formatter

import (
	"bytes"
	"text/template"

	"github.com/gookit/color"
)

type Formatter struct {
	Name     string
	Template string
	Data     interface{}
}

func NewFormatter(name string, template string, data interface{}) *Formatter {
	return &Formatter{
		Name:     name,
		Template: template,
		Data:     data,
	}
}

func (f *Formatter) Render() (string, error) {
	var result bytes.Buffer

	tmpl, err := template.New(f.Name).Parse(f.Template)

	if err != nil {
		return "", err
	}

	err = tmpl.Execute(&result, f.Data)

	if err != nil {
		return "", err
	}

	return color.Sprint(result.String()), err
}
