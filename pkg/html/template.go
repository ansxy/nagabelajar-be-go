package html

import (
	"bytes"
	"html/template"
	"os"
	"path/filepath"
)

func ParseTemplateHTML(templateFileName string, data interface{}) (*bytes.Buffer, error) {
	basePath, _ := os.Getwd()
	fileFullPath := filepath.Join(basePath, "pkg/html/template", templateFileName)
	t, err := template.ParseFiles(fileFullPath)
	if err != nil {
		return nil, err
	}

	buf := new(bytes.Buffer)
	if err := t.Execute(buf, data); err != nil {
		return nil, err
	}

	return buf, nil
}
