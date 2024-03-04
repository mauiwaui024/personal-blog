package repository

import "fmt"

func (repo *Repository) AddArticleToDb(title, content string) error {
	_, err := repo.DataBase.Exec("INSERT INTO articles (title, content) VALUES ($1, $2)", title, content)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
