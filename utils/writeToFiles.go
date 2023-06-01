package utils

import "github.com/Go-QuickStart/entity/layer"

type IWriteToFiles interface {
	CreateFile(format []layer.FileSpec)
	CreateFolderProject(name string, perm int)
	CreateGoFile(format layer.FileSpec) error
}
