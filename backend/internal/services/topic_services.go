func CreateTopic(authorID int, title string){
	//return topic or error
	//repo call TopicRepository.Create, UserRepository.GetByID
}

func GetTopicByID(topicID int){
	//return topic or error
	//repo call TopicRepository.GetByID
}

func ListTopics(){
	//return []Topic
	//repo call TopicRepository.List
}

func UpdateTopic(topicID int, authorID int, title string){
	//return topic or error
	//repo call TopicRepository.GetByID, TopicRepository.Update
}

func DeleteTopic(topicID int, authorID int){
	//return nothing or error
	//repo call TopicRepository.GetByID, TopicRepository.Delete
}