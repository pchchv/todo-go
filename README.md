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