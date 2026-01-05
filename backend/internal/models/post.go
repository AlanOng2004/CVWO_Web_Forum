package models 

import "time"

type Post struct {
	ID 			int
	TopicID 	int
	Title 		string
	Content 	string
	AuthorID 	int
	CreatedAt 	time.Time
}