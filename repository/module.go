package repository

import (
	db "task/database"
)

type RepositoryModule struct {
	database *db.DBModule
}

func RepositoryModuleInit(d *db.DBModule) *RepositoryModule {
	return &RepositoryModule{database: d}
}
