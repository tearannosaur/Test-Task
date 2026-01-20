package repository

func (r *RepositoryModule) SaveNumber(num int) error {
	query := `INSERT INTO numbers(number)
	VALUES ($1)`
	_, err := r.Database.Db.Exec(query, num)
	if err != nil {
		return err
	}
	return nil
}

func (r *RepositoryModule) GetNumbers() ([]int, error) {
	var nums []int
	query := `SELECT * FROM numbers`
	_ = query
	err := r.Database.Db.Select(&nums, query)
	if err != nil {
		return nil, err
	}
	return nums, nil
}
