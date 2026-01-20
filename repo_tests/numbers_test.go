package repo_tests

import (
	"task/database"
	m "task/models"
	"task/repository"
	"testing"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestGetNumbers(t *testing.T) {
	db, err := sqlx.Connect("postgres", "user=admin password=222444 dbname=database host=localhost port=5434 sslmode=disable")
	assert.NoError(t, err)
	repo := &repository.RepositoryModule{
		Database: &database.DBModule{Db: db},
	}
	_, err = repo.Database.Db.DB.Exec("DELETE FROM numbers")

	assert.NoError(t, err)

	numbersToInsert := []int{3, 1, 2, -3}
	for _, n := range numbersToInsert {
		err = repo.SaveNumber(n)
		assert.NoError(t, err, "insert error %d", n)
	}

	results, err := repo.GetNumbers()
	assert.NoError(t, err)

	sortedResults := m.ValidateNumbers(results)

	expected := []int{-3, 1, 2, 3}
	assert.Equal(t, expected, sortedResults, "numbers are not sorted correctly")

	_, _ = repo.Database.Db.DB.Exec("DELETE FROM numbers")

}
