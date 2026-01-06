type TopicRepository interface {
	Create(topic *models.Topic) error
	GetByID(id int) (*models.Topic, error)
	List() ([]models.Topic, error)
}