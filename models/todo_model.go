package todo_model

import (
	"time"
	_"github.com/go-sql-driver/mysql"
)

type ACTIVITY struct {
	ID         			int32     `json:"id"`
	Email 				string    `json:"email"`
	Title      			string    `json:"title" binding:"required"`
	CreatedAt  			time.Time `json:"created_at"`
	UpdatedAt  			time.Time `json:"updated_at"`
}
type TODO struct {
	ID          		int32     `json:"id"`
	ActivityGroupId 	int32     `json:"activity_group_id"`
	Title       		string    `json:"title"`
	IsActive   			bool	  `json:"is_active"` 
	Priority       		string    `json:"priority"`
	CreatedAt   		time.Time `json:"created_at"`
	UpdatedAt   		time.Time `json:"updated_at"`
}
