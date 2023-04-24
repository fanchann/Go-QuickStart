package utils

import (
	"Go-QuickStart/entity/layer"
	"Go-QuickStart/utils/helpers"
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"
)

type WriteToFiles struct{}

func NewWriteToFiles() IWriteToFiles {
	return &WriteToFiles{}
}

func (write *WriteToFiles) CreateFile(format []layer.FileSpec) {
	for _, file := range format {
		err := write.CreateGoFile(file)
		helpers.LogErrorWithPanic(err)
	}
}

func (write *WriteToFiles) CreateFolderProject(name string, perm int) {
	err := os.Mkdir(name, fs.FileMode(perm))
	helpers.LogErrorWithPanic(err)
}

func (write *WriteToFiles) CreateGoFile(format layer.FileSpec) error {
	codeTemplate, err := template.New("").Parse(format.ScCode)
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	var tpl bytes.Buffer

	if err = codeTemplate.Execute(&tpl, struct {
		GoVersion   string
		PackageName string
	}{
		GoVersion:   format.GoVersion,
		PackageName: format.PackageName,
	}); err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	file, err := os.Create(format.FileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = tpl.WriteTo(file)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	if format.Location != "" {
		newLocation := filepath.Join(format.Location, filepath.Base(format.FileName))
		if err := os.Rename(format.FileName, newLocation); err != nil {
			return fmt.Errorf("failed to move file to %s: %v", newLocation, err)
		}
	}

	return nil
}
