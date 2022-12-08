<h1> Privy - Backend Engineer Test
</h1>

####  By: [Aldi Sugiarto](https://github.com/Aldisugiarto)

## Project Description
This project is made to take technical test at Privy as Backend Engineer.
## Task
There are some task to create:
1. [GET] /cakes to get the list of cakes 
2. [GET] /cakes/:id to get a specific cake by :id 
3. [POST] /cakes to post a new cake 
4. [PATCH] /cakes/:id to update a cake by :id 
5. [DELETE] /cakes/:id to delete a cake by :id 
## Requirements
The project have some requirement:
 - The API must comply RESTFul guideline
 - Free to improvise or create your own API response object
 - Have to use the given endpoints and follow the requirements
 - Create API use GO Language
 - Store data use MariaDB or MySQL database
 - Create script to migration database.
 - Don’t use ORM
 - Applied request validation
 - Provide unit test on your project
 - Provide the proper README
 - Running in docker container
 - Adding extras like (but not limited to) fancy architectural and elegant error handling & logging

## How to run this project
### Getting started
### Setting database:
There are two type to test the API
- Running with "go test"
- Running wiith docker
#### A. Running with go test on the terminal 
1. Install [XAMPP](https://www.apachefriends.org/), [laragon](https://laragon.org/), or etc
2. Run server (apache, nginx, etc) and database (mysql)
3. Setting port refer to XAMPP, laragon, etc (Ex. Apache on port 80 and mysql on port 3306)
4. Setting database for "go test" in file config-dev.json. Fill name with existing database. In this case I used sys database
    ```
    {
        "database": {
            "host": "127.0.0.1",
            "db-container": "rest_api-db",
            "port": "3306",
            "name": "sys",
            "user": "root",
            "pass": "root",
            "timeout": 1
        }
    }
    ```
5. Setting database connection at file config.go under folder config. Uncomment dbHost and dbPort, after that uncomment connection which used host and port
    ```
    func InitDB(c *gin.Context) (db *sql.DB) {
        config := NewConfiguration()
        config.LoadConfigurationFromFile(getFilePathConfigEnvirontment())
        dbDriver := "mysql"
        dbHost := config.GetValue(`database.host`) //Uncomment this row if not use docker
        dbPort := config.GetValue(`database.port`)
        // dbContainer := config.GetValue(`database.db-container`)
        dbUser := config.GetValue(`database.user`)
        dbPass := config.GetValue(`database.pass`)
        dbName := config.GetValue(`database.name`)
        connection := fmt.Sprintf(dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName) //Uncomment this row if not use docker
        // connection := fmt.Sprintf(dbUser + ":" + dbPass + "@tcp(" + dbContainer + ")/" + dbName)
        db, err := sql.Open(dbDriver, connection)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "status": "Errors",
            })
        }
        Migrate(db)
        return db
    }
    ```
6. Open the terminal and type go test at the top of project
    ```
    >> go test
    ```
#### B. Running with docker
1. Setting database for "go test" in file config-dev.json. Fill name with existing database. In this case I used sys database
    ```
    {
        "database": {
            "host": "127.0.0.1",
            "db-container": "rest_api-db",
            "port": "3306",
            "name": "sys",
            "user": "root",
            "pass": "root",
            "timeout": 1
        }
    }
    ```
2. Setting database connection at file config.go under folder config. Uncomment dbContainer to run with docker
    ```
    func InitDB(c *gin.Context) (db *sql.DB) {
        config := NewConfiguration()
        config.LoadConfigurationFromFile(getFilePathConfigEnvirontment())
        dbDriver := "mysql"
        // dbHost := config.GetValue(`database.host`) //Uncomment this row if not use docker
        // dbPort := config.GetValue(`database.port`)
        dbContainer := config.GetValue(`database.db-container`)
        dbUser := config.GetValue(`database.user`)
        dbPass := config.GetValue(`database.pass`)
        dbName := config.GetValue(`database.name`)
        // connection := fmt.Sprintf(dbUser + ":" + dbPass + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName) //Uncomment this row if not use docker
        connection := fmt.Sprintf(dbUser + ":" + dbPass + "@tcp(" + dbContainer + ")/" + dbName)
        db, err := sql.Open(dbDriver, connection)
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{
                "status": "Errors",
            })
        }
        Migrate(db)
        return db
    }
    ```
3. Run docker daemon (in my case I used windows, so I install docker desktop) you can visit [docker](https://www.docker.com/)  website for more information. Open docker desktop with Run as Administrator
4. Open terminal on your top of project and type below command:
    ```
    >> docker-compose build
    >> docker-compose up
    ```
5. Check on your docker desktop that container for web, api, and pma is running.
6. After all of services is running please follow below steps to test
7. Server has been running at port :8080
8. Move to test case section
9. Open Postman and import apispec.json file in this project to do test
## Test Case
Refer from the task, we have created test case below:
1.  List of cakes.
 - Endpoint :   /cakes
 - method   :   GET
 - description  :   return a list of the cakes in JSON format, the cakes must be sorted by rating and alphabetically
 - URL  :   http://localhost:8080/cakes
 - Response:
    ```
    {
        "data": [
            {
                "id": 1,
                "title": "Lemon cheesecake",
                "description": "A cheesecake made of lemon",
                "rating": 10,
                "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
                "created_at": "2022-11-29T04:38:24Z",
                "updated_at": "2022-11-29T04:38:24Z"
            }
        ],
        "message": "get data cakes",
        "status": "success"
        }
    ```
2.  Detail of cake.
 - Endpoint :   /cakes/:id
 - method   :   GET
 - description  :
    ```
    return the details of a cake in JSON format
    example:
    
    {
        "id": 1,
        "title": "Lemon cheesecake",
        "description": "A cheesecake made of lemon",
        "rating": 7.0,
        "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
        "created_at": "2020-02-01 10:56:31",
        "updated_at": "2020-02-13 09:30:23"
    }
    ```
 - URL  :   http://localhost:8080/cakes/1
 - Response:
    ```
    {
        "data": [
            {
                "id": 1,
                "title": "Lemon cheesecake",
                "description": "A cheesecake made of lemon",
                "rating": 10,
                "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
                "created_at": "2022-11-29T04:38:24Z",
                "updated_at": "2022-11-29T04:38:24Z"
            }
        ],
        "message": "get data cake by ID",
        "status": "success"
    }
    ```
3. Add new cake
 - Endpoint : /cakes
 - method   :   POST
 - description  :
    ```
    Add a cake to the cakes list, the data will be sent as a JSON in the request body :
    {
        "title": "Lemon cheesecake",
        "description": "A cheesecake made of lemon",
        "rating": 10,
        "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
    }
    ```
 - URL  :  http://localhost:8080/cakes
 - Response:
    ```
    {
        "data": [
            {
                "id": 1,
                "title": "Lemon cheesecake",
                "description": "A cheesecake made of lemon",
                "rating": 10,
                "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
                "created_at": "2022-11-29T04:38:24Z",
                "updated_at": "2022-11-29T04:38:24Z"
            }
        ],
        "message": "add data cake",
        "status": "success"
    }
    ```
4. Update cake
 - Endpoint : /cakes/:id
 - method   :   PATCH
 - description  :
    ```
    Add a cake to the cakes list, the data will be sent as a JSON in the request body :
    {
        "title": "Lemon Tea",
        "description": "A cheesecake made of lemon",
        "rating": 9,
        "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg"
    }
    ```
 - URL  :  http://localhost:8080/cakes/1
 - Response:
    ```
    {
        "data": [
            {
                "id": 1,
                "title": "Lemon Tea",
                "description": "A cheesecake made of lemon",
                "rating": 9,
                "image": "https://img.taste.com.au/ynYrqkOs/w720-h480-cfill-q80/taste/2016/11/sunny-lemon-cheesecake-102220-1.jpeg",
                "created_at": "2022-11-29T04:38:24Z",
                "updated_at": "2022-11-29T04:38:24Z"
            }
        ],
        "message": "update data cake",
        "status": "success"
    }
    ```
5. Delete cake
 - Endpoint : /cakes/:id
 - method   :   DELETE
 - description  : delete a cake from database
 - URL  :  http://localhost:8080/cakes/1
 - Response:
    ```
    {
        "message": "delete id:1 is success",
        "status": "success"
    }
    ```

