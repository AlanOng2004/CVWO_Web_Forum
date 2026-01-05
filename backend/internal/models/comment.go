package models 

import "time"

type Comment struct {
	ID 			int
	PostID 		int
	Content 	string
	AuthorID 	int
	CreatedAt 	time.Time
}