package utils

import (
	"Go-QuickStart/entity"
	"bytes"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"text/template"
)

func CreateFiles(files []entity.FileSpec) {
	for _, file := range files {
		err := CreateFile(file.Location, file.PkgName, file.ScCode, file.FileName)
		LogErrorWithPanic(err)
	}
}

func CreateFolderProject(name string, perm int) {
	err := os.Mkdir(name, fs.FileMode(perm))
	LogErrorWithPanic(err)

}

func CreateFile(location, pkgName, sourceCode, fileName string) error {
	codeTemplate, err := template.New("").Parse(sourceCode)
	if err != nil {
		return fmt.Errorf("failed to parse template: %v", err)
	}

	var tpl bytes.Buffer

	if err = codeTemplate.Execute(&tpl, struct{ PackageName string }{pkgName}); err != nil {
		return fmt.Errorf("failed to execute template: %v", err)
	}

	file, err := os.Create(fileName)
	if err != nil {
		return fmt.Errorf("failed to create file: %v", err)
	}
	defer file.Close()

	_, err = tpl.WriteTo(file)
	if err != nil {
		return fmt.Errorf("failed to write to file: %v", err)
	}

	if location != "" {
		newLocation := filepath.Join(location, filepath.Base(fileName))
		if err := os.Rename(fileName, newLocation); err != nil {
			return fmt.Errorf("failed to move file to %s: %v", newLocation, err)
		}
	}

	return nil
}
