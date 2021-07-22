package templ

import (
	"bytes"
	"fmt"
	"text/template"

	"github.com/0xma12k/graylog-line-notify-gateway/internal/config"
	"github.com/0xma12k/graylog-line-notify-gateway/internal/entity"
	"github.com/sirupsen/logrus"
)

var defaultTempl string = "{{.EventDefinitionTitle}}"

func DefaultExecute(graylog *entity.GraylogJSON) (string, error) {
	var buf bytes.Buffer
	logrus.Debug("using build-in template")
	t := template.Must(template.New("default").Parse(defaultTempl))
	if err := t.Execute(&buf, graylog); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func ExecuteTemplate(templName string, graylog *entity.GraylogJSON) (string, error) {

	var buf bytes.Buffer

	if templName == "" {
		templName = config.Get().DefaultTemplate
	}

	path, err := findTemplatePath(templName)
	if err != nil {
		return "", err
	}

	tmpl, err := template.ParseFiles(path)
	if err != nil {
		return "", err
	}
	t := template.Must(tmpl, err)

	if err := t.Execute(&buf, graylog); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func findTemplatePath(name string) (string, error) {
	for _, templ := range config.Get().Templates {
		if name == templ.Name {
			return templ.Path, nil
		}
	}
	return "", fmt.Errorf("template not found")
}
