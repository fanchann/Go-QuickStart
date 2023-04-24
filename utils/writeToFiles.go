package utils

import "Go-QuickStart/entity/layer"

type IWriteToFiles interface {
	CreateFile(format []layer.FileSpec)
	CreateFolderProject(name string, perm int)
	CreateGoFile(format layer.FileSpec) error
}
