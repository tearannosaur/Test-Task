package repository

import (
	db "task/database"
)

type RepositoryModule struct {
	Database *db.DBModule
}

func RepositoryModuleInit(d *db.DBModule) *RepositoryModule {
	return &RepositoryModule{Database: d}
}
