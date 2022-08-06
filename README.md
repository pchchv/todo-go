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
"POST" /todo — Create a poll with answer options
    options:
        title — Name of poll
        options — Answer options

    example: 
        "POST" :8000/todo?title=Buy coffee
        "POST" :8000/todo?title=Pay the electricity bill&text=Bill in the inbox on the email
        "POST" :8000/todo?title=Go to the gym&completed=true
```