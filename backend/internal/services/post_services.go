func CreatePost(authorID int, topicID int, content string){
	//return post or error
	//repo call PostRepository.Create, TopicRepository.GetByID
}

func GetPostByID(postID int){
	//return post or error
	//repo call PostRepository.GetByID
}

func ListPostsByTopic(topicID int){
	//return []Post
	//repo call PostRepository.GetByTopicID
}

func UpdatePost(postID int, authorID int, content string){
	//return post or error
	//repo call PostRepository.GetByID, PostRepository.Update
}

func DeletePost(postID int, authorID int){
	//return nth or error
	//repo call PostRepository.GetByID, PostRepository.Delete
}