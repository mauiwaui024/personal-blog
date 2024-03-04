package service

import "day_06/internal/model"

func (s Service) GetArticles() ([]model.Article, error) {
	////значит здесь объявить слайс со статьями
	// var article []model.Article
	return s.Repo.GetArticlesFromDB()
}
func (s Service) GetArticleById(articleID int) (model.Article, error) {
	////значит здесь объявить слайс со статьями
	// var article []model.Article
	return s.Repo.GetArticleFromDBById(articleID)
}
