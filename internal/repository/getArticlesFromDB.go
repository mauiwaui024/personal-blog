package repository

import "day_06/internal/model"

func (r *Repository) GetArticlesFromDB() ([]model.Article, error) {
	rows, err := r.DataBase.Query("SELECT id, title, content FROM articles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var myArticles []model.Article

	for rows.Next() {
		var article model.Article
		// Сканирование данных текущей строки в переменные Go
		if err := rows.Scan(&article.ID, &article.Title, &article.Content); err != nil {
			return nil, err
		}
		myArticles = append(myArticles, article)

	}

	// Проверка наличия ошибок после завершения итерации
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return myArticles, nil
}
func (r *Repository) GetArticleFromDBById(articleID int) (model.Article, error) {
	row := r.DataBase.QueryRow("SELECT id, title, content FROM articles WHERE id = $1", articleID)
	var article model.Article
	err := row.Scan(&article.ID, &article.Title, &article.Content)
	if err != nil {
		return model.Article{}, err
	}
	return article, nil
}
