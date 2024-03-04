package service

func (s *Service) AddArticle(title, content string) error {
	return s.Repo.AddArticleToDb(title, content)
}
