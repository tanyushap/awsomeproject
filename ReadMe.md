#ToDo task creator API

##How To Run

###Prerequisites 
- Go programming language
- Make (not mandatory)

##Running the test 
navigate to the project directory e.g. <some_path>/awesomeProject

Using make file
```bash
make test
```

alternatively use the command line

```bash
go test -v ./test
```

##Running the application
Using make file
```bash
make run
```

alternatively use the command line

```bash
go run main/main.go
```
the server should run on http://127.0.0.1:8081

##What Endpoints Are Available

###GET
```
    /task - takes a parameter called id e.g. /tasks?id=1
    /task/get-all
    /task/get-all-todo
    /task/get-all-completed
```
###POST
```
    /task/add
```
###TBD
```
    /task/complete
    /task/delete
```

##Example Json

```json
[
  {
    "id": 1,
    "taskName": "Task 1",
    "description": "Task One",
    "status": "Pending",
    "creator": "Bob"
  },
  {
    "id": 2,
    "taskName": "Task 2",
    "description": "Task Two",
    "status": "Completed",
    "creator": "Alice"
  }
]
```