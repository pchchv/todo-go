# Todo list backend
# Running the application without Docker
```
go run .
```
## HTTP Methods
```
"GET" /ping — Checking the server connection

    example: 
        "GET" :8000/ping
```
```
"POST" /todo — Create a todo
    options:
        title — Name of todo
        text — Extra text
        completed — Is the task completed (true or false)

    example: 
        "POST" :8000/todo?title=Buy coffee
        "POST" :8000/todo?title=Pay the electricity bill&text=Bill in the inbox on the email
        "POST" :8000/todo?title=Go to the gym&completed=true
```
```
"PATCH" /todo — Update a todo
    options:
        id — Id of todo
        title — Name of todo
        text — Extra text
        completed — Is the task completed (true or false)

    example: 
        "PATCH" :8000/todo?id=111&completed=true
        "PATCH" :8000/todo?title=Go to the gym&completed=true
```
```
"GET" /todo — Get a todo
    options:
        id — Id of todo
        title — Name of todo

    example: 
        "GET" :8000/todo?id=111
        "GET" :8000/todo?title=Buy coffee
```
```
"DELETE" /todo — Delete a todo
    options:
        id — Id of todo
        title — Name of todo

    example: 
        "DELETE" :8000/todo?id=111
        "DELETE" :8000/todo?title=Buy coffee
```
```
"GET" /todos — Get all todos

    example: 
        "GET" :8000/todos
```
```
"DELETE" /todos — Delete all todos

    example: 
        "DELETE" :8000/todos
```