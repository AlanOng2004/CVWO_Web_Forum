func CreateComment(authorID int, postID int, content string){
	//return comment or error
	//repo call CommentRepository.Create, PostRepository.GetByID
}

func GetCommentByID(commentID int){
	//return comment or error
	//repo call CommentRepository.GetByID
}

func ListCommentsByPost(postID int){
	//return []Comment
	//repo call CommentRepository.GetByPostID
}

func UpdateComment(commentID int, authorID int, content string){
	//return comment or error
	//repo call CommentRepository.GetByID, CommentRepository.Update
}

func DeleteComment(commentID int, authorID int){
	//return nth or error
	//repo call CommentRepository.GetByID, CommentRepository.Delete
}