type CommentRepository interface {
	Create(comment *models.Comment) error
	GetByID(id int) (*models.Comment, error)
	GetByPostID(postID int) ([]*models.Comment, error)
}