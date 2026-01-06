type PostRepository interface {
	Create(post *models.Post) error
	GetByID(id int) (*models.Post, error)
	GetByTopicID(topicID int) ([]*models.Post, error)
}