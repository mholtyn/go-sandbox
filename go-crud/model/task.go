package model

import "time"


type TaskCreate struct {
	Title		string			`json:"title"`
}

type TaskPublic struct {
	ID			int				`json:"id"`
	Title 		string			`json:"title"`
	Done 		bool			`json:"done"`
	CreatedAt 	time.Time		`json:"created_at"`
}