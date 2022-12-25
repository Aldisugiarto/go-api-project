package controller

import (
	"database/sql"
	"net/http"
	"time"

	todo_model "rest_api/models"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

const layoutDateTime = "2006-01-02 15:04:05"

type Repo struct {
	DB *sql.DB
}

type empty struct{}

// ********************** Create Function to handle activity routes ********************** //
// ************ Function handle get all activity ************ //
func (db *Repo) GetActivity(c *gin.Context) {
	row, err := db.DB.Query("SELECT * FROM todo4.activities")
	CheckErr(err, c)
	defer row.Close()
	if row.Next() {
		selDB, err := db.DB.Query("SELECT id, email, title, created_at, updated_at FROM todo4.activities ORDER By id DESC")
		CheckErr(err, c)
		activity := todo_model.ACTIVITY{}
		activities := []todo_model.ACTIVITY{}
		var (
			id									int32
			title, email						string
			created_at, updated_at				mysql.NullTime
		)
		for selDB.Next() {
			err = selDB.Scan(&id, &email, &title, &created_at, &updated_at)
			CheckErr(err, c)
			activity.ID = id
			activity.Email = email
			activity.Title = title
			
			if created_at.Valid {
				activity.CreatedAt = created_at.Time
			} 
			
			if updated_at.Valid {
				activity.UpdatedAt = updated_at.Time
			} 
			

			activities = append(activities, activity)
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "Success",
			"message": "Success",
			"data":    activities,
		})
		defer selDB.Close()
		
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Errors",
			"message": "There is no data in database",
		})
		
		return
	}

}

// ************ Function handle get one activity ************ //
func (db *Repo) GetActivityById(c *gin.Context) {
	_id := c.Param("id")
	selDB, err := db.DB.Query("SELECT id, email, title, created_at, updated_at FROM todo4.activities WHERE id=?", _id)
	CheckErr(err, c)
	activity := todo_model.ACTIVITY{}
	var (
		id									int32
		title, email						string
		created_at, updated_at				mysql.NullTime
	)
	if selDB.Next() {
		err = selDB.Scan(&id, &email, &title, &created_at, &updated_at)
		CheckErr(err, c)
		activity.ID = id
		activity.Email = email
		activity.Title = title
		
		if created_at.Valid {
			activity.CreatedAt = created_at.Time
		} 
		
		if updated_at.Valid {
			activity.UpdatedAt = updated_at.Time
		} 

		c.JSON(http.StatusOK, gin.H{
			"status":  "Success",
			"message": "Success",
			"data":    activity,
		})
		defer selDB.Close()
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Activity with ID " + _id + " Not Found",
		})
		return
	}
}

// ************ Function handle create activity ************ //
func (db *Repo) AddActivity(c *gin.Context) {
	activity := todo_model.ACTIVITY{}
	err := c.ShouldBindJSON(&activity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad Request",
			"message": "title cannot be null",
		})
		return
	} 

	insert, err := db.DB.Query("INSERT INTO todo4.activities (email,title,created_at,updated_at) VALUES (?,?,?,?)",
		activity.Email, activity.Title, time.Now().Format(layoutDateTime), time.Now().Format(layoutDateTime))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Errors",
			"message": err.Error(),
			"data": empty{},
		})
		return
	} 
		defer insert.Close()
		selDB, err := db.DB.Query("SELECT id, email, title, created_at, updated_at FROM todo4.activities ORDER By id DESC LIMIT 1")
		CheckErr(err, c)
		var (
			id									int32
			title, email						string
			created_at, updated_at				mysql.NullTime
		)
		for selDB.Next() {
			err = selDB.Scan(&id, &email, &title, &created_at, &updated_at)
			CheckErr(err, c)
			activity.ID = id
			activity.Email = email
			activity.Title = title
			
			if created_at.Valid {
				activity.CreatedAt = created_at.Time
			} 
			
			if updated_at.Valid {
				activity.UpdatedAt = updated_at.Time
			} 
			
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  "Success",
			"message": "Success",
			"data":    activity,
		})
		return
		defer selDB.Close()
}

// ************ Function handle update activity ************ //
func (db *Repo) UpdateActivity(c *gin.Context) {
	_id := c.Param("id")
	row, err := db.DB.Query("SELECT * FROM todo4.activities WHERE id=?", _id)
	CheckErr(err, c)
	activity := todo_model.ACTIVITY{}

	if row.Next() {
		if err := c.ShouldBindJSON(&activity); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "Bad Request",
				"message": "title cannot be null",
			})
			return
		}

		update, err := db.DB.Query("UPDATE todo4.activities SET title=?, updated_at=? WHERE id=?",
			activity.Title, time.Now().Format(layoutDateTime), _id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Errors",
				"message": err.Error(),
			})
			return
		}
		defer update.Close()

		selDB, err := db.DB.Query("SELECT id, email, title, created_at, updated_at FROM todo4.activities WHERE id=?", _id)
		CheckErr(err, c)

		var (
			id									int32
			title, email						string
			created_at, updated_at				mysql.NullTime
		)
		for selDB.Next() {
			err = selDB.Scan(&id, &email, &title, &created_at, &updated_at)
			CheckErr(err, c)
			activity.ID = id
			activity.Email = email
			activity.Title = title
			
			if created_at.Valid {
				activity.CreatedAt = created_at.Time
			} 
			
			if updated_at.Valid {
				activity.UpdatedAt = updated_at.Time
			} 
			
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "Success",
			"message": "Success",
			"data":    activity,
		})

		defer selDB.Close()
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Activity with ID " + _id + " Not Found",
		})
		return
	}

}

// ************ Function handle delete activity ************ //
func (db *Repo) DeleteActivityById(c *gin.Context) {
	_id := c.Param("id")
	selDB, err := db.DB.Query("SELECT * FROM todo4.activities WHERE id=?", _id)
	CheckErr(err, c)
	if selDB.Next() {
		delete, err := db.DB.Query("DELETE From todo4.activities WHERE id=?", _id)
		CheckErr(err, c)
		defer delete.Close()

		c.JSON(http.StatusOK, gin.H{
			"status":  "Success",
			"message": "Success",
			"data": empty{},
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Activity with ID " + _id + " Not Found",
		})

		return
	}
}

// ********************** Create Function to handle todo routes ********************** //
// ************ Function handle get all todo ************ //
func (db *Repo) GetTodo(c *gin.Context) {
	query := c.Query("activity_group_id")
	if len(query) > 0 {
		row, err := db.DB.Query("SELECT * FROM todo4.todos WHERE activity_group_id=?",query)
		CheckErr(err, c)
		todo := todo_model.TODO{}
		todos := []todo_model.TODO{}
		var (
			id, activity_group_id				int32
			is_active							bool
			title, priority						string
			created_at, updated_at			 	mysql.NullTime
		)
		defer row.Close()
		if row.Next(){
			selDB, err := db.DB.Query("SELECT * FROM todo4.todos WHERE activity_group_id=?",query)
			for selDB.Next() {
				err = selDB.Scan(&id, &activity_group_id, &title, &is_active, &priority, &created_at, &updated_at)
				CheckErr(err, c)
				todo.ID = id
				todo.ActivityGroupId = activity_group_id
				todo.Title = title
				todo.IsActive = is_active
				todo.Priority = priority
				
				if created_at.Valid {
					todo.CreatedAt = created_at.Time
				} 
				
				if updated_at.Valid {
					todo.UpdatedAt = updated_at.Time
				} 

				todos = append(todos, todo)
			}
		
			c.JSON(http.StatusOK, gin.H{
				"status":  "Success",
				"message": "Success",
				"data":    todos,
			})
			defer selDB.Close()
			
		} else {
			c.JSON(http.StatusNotFound, gin.H{
				"status":  "Not Found",
				"message": "Todo with Activity Group ID " + query + " Not Found",
			})
	
			return
		}
	} else {
		row, err := db.DB.Query("SELECT * FROM todo4.todos")
		CheckErr(err, c)
		defer row.Close()
		if row.Next() {
			selDB, err := db.DB.Query("SELECT * FROM todo4.todos ORDER By id DESC")
			CheckErr(err, c)
			todo := todo_model.TODO{}
			todos := []todo_model.TODO{}
			var (
				id, activity_group_id				int32
				is_active							bool
				title, priority						string
				created_at, updated_at				mysql.NullTime
			)
			for selDB.Next() {
				err = selDB.Scan(&id, &activity_group_id, &title, &is_active, &priority, &created_at, &updated_at)
				CheckErr(err, c)
				todo.ID = id
				todo.ActivityGroupId = activity_group_id
				todo.Title = title
				todo.IsActive = is_active
				todo.Priority = priority
				
				if created_at.Valid {
					todo.CreatedAt = created_at.Time
				} 
				
				if updated_at.Valid {
					todo.UpdatedAt = updated_at.Time
				} 
	
				todos = append(todos, todo)
			}
	
			c.JSON(http.StatusOK, gin.H{
				"status":  "Success",
				"message": "Success",
				"data":    todos,
			})
			defer selDB.Close()
			
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Errors",
				"message": "There is no data in database",
			})
			
			return
		}
	}


}

// ************ Function handle get one todo ************ //
func (db *Repo) GetTodoById(c *gin.Context) {
	_id := c.Param("id")
	row, err := db.DB.Query("SELECT * FROM todo4.todos WHERE id=?", _id)
	CheckErr(err, c)
	todo := todo_model.TODO{}
	var (
		id, activity_group_id				int32
		is_active							bool
		title, priority						string
		created_at, updated_at				mysql.NullTime
	)
	defer row.Close()
	if row.Next(){
		selDB, err := db.DB.Query("SELECT * FROM todo4.todos WHERE id=?", _id)
		for selDB.Next() {
			err = selDB.Scan(&id, &activity_group_id, &title, &is_active, &priority, &created_at, &updated_at)
			CheckErr(err, c)
			todo.ID = id
			todo.ActivityGroupId = activity_group_id
			todo.Title = title
			todo.IsActive = is_active
			todo.Priority = priority
			
			if created_at.Valid {
				todo.CreatedAt = created_at.Time
			} 
			
			if updated_at.Valid {
				todo.UpdatedAt = updated_at.Time
			} 
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "Success",
			"message": "Success",
			"data":    todo,
		})
		defer selDB.Close()
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Todo with ID " + _id + " Not Found",
		})
		return
	}
}

// ************ Function handle add todo ************ //
func (db *Repo) AddTodo(c *gin.Context) {
	todo := todo_model.TODO{}
	err := c.ShouldBindJSON(&todo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Bad request",
			"message": "Errors",
			"data": empty{},
		})
		return
	} 
	if todo.ActivityGroupId > 0 && len(todo.Title) > 0{
		insert, err := db.DB.Query("INSERT INTO todo4.todos (activity_group_id,title,is_active,priority,created_at,updated_at) VALUES (?,?,?,?,?,?)",
			todo.ActivityGroupId, todo.Title,todo.IsActive,"very-high", time.Now().Format(layoutDateTime), time.Now().Format(layoutDateTime))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  "Internal server error",
				"message": err.Error(),
			})
			return
		} 
		defer insert.Close()
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "Bad request",
			"message":"title cannot be null",
		})
		return
	}
	var (
		id, activity_group_id				int32
		is_active							bool
		title, priority						string
		created_at, updated_at				mysql.NullTime
	)
	selDB, err := db.DB.Query("SELECT * FROM todo4.todos ORDER By id DESC LIMIT 1")
	for selDB.Next() {
		err = selDB.Scan(&id, &activity_group_id, &title, &is_active, &priority, &created_at, &updated_at)
		CheckErr(err, c)
		todo.ID = id
		todo.ActivityGroupId = activity_group_id
		todo.Title = title
		todo.IsActive = is_active
		todo.Priority = priority
		
		if created_at.Valid {
			todo.CreatedAt = created_at.Time
		} 
		
		if updated_at.Valid {
			todo.UpdatedAt = updated_at.Time
		} 
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  "Success",
		"message": "Success",
		"data":    todo,
	})
	return
	defer selDB.Close()
}

// ************ Function handle update todo ************ //
func (db *Repo) UpdateTodo(c *gin.Context) {
	_id := c.Param("id")
	row, err := db.DB.Query("SELECT * FROM todo4.todos WHERE id=?", _id)
	CheckErr(err, c)
	todo := todo_model.TODO{}

	if row.Next() {
		if err := c.ShouldBindJSON(&todo); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  "Errors",
				"message": err.Error(),
				"data": empty{},
			})
			return
		}
		if len(todo.Title) > 0{
			update, err := db.DB.Query("UPDATE todo4.todos SET title=?,is_active=?,priority=?,updated_at=? WHERE id=?",
			todo.Title,todo.IsActive,todo.Priority,time.Now().Format(layoutDateTime),_id)
			
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"status":  "Errors",
					"message": err.Error(),
				})
				return
			}
			defer update.Close()
		}

		var (
			id, activity_group_id				int32
			is_active							bool
			title, priority						string
			created_at, updated_at				mysql.NullTime
		)
		selDB, err := db.DB.Query("SELECT * FROM todo4.todos WHERE id=?",_id)
		for selDB.Next() {
			err = selDB.Scan(&id, &activity_group_id, &title, &is_active, &priority, &created_at, &updated_at)
			CheckErr(err, c)
			todo.ID = id
			todo.ActivityGroupId = activity_group_id
			todo.Title = title
			todo.IsActive = is_active
			todo.Priority = priority
			
			if created_at.Valid {
				todo.CreatedAt = created_at.Time
			} 
			
			if updated_at.Valid {
				todo.UpdatedAt = updated_at.Time
			} 
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  "Success",
			"message": "Success",
			"data":    todo,
		})

		defer selDB.Close()
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Todo with ID " + _id + " Not Found",
		})
		return
	}

}

// ************ Function handle delete todo ************ //
func (db *Repo) DeleteTodoById(c *gin.Context) {
	_id := c.Param("id")
	selDB, err := db.DB.Query("SELECT * FROM todo4.todos WHERE id=?", _id)
	CheckErr(err, c)
	if selDB.Next() {
		delete, err := db.DB.Query("DELETE From todo4.todos WHERE id=?", _id)
		CheckErr(err, c)
		defer delete.Close()

		c.JSON(http.StatusOK, gin.H{
			"status":  "Success",
			"message": "Success",
			"data": empty{},
		})
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "Not Found",
			"message": "Todo with ID " + _id + " Not Found",
		})

		return
	}
}

func CheckErr(err error, c *gin.Context){
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  "Errors",
			"message": err.Error(),
		})
		return 
	}
}