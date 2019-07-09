# VALIDATE EXERCISE

Create a simple website to allow user to validate their implementation for Go exercise during the time they learn Go.

Admin create/update/delete exercise with:
- Title
- Description
- Upload solution_test.go

Candidate upload solution.go, system validate their solution by running with solution_test.go

Requirements:
- Lang: Go
- Web library: gorialla/mux
- UI: HTML - Go server side rendering (using template/html)
- DB: MongoDB
- Deployment: Docker
- Run/build go code: use API from [play.golang.org](https://play.golang.org)

## DESIGN
### ROUTING:
- "/": homepage, show list all exercises with title
- "/admin": the site for admin manage exercises
    + "/create": create new exercise
    + "/edit": edit exist exercise
- "/exercise/{id}": show a exercise have id in database

### API:
- "/api" GET: get all exercise
- "/api" PUT: create new exercise
- "/api/{id}" GET: get a exercise have id in database
- "/api/{id}" POST: run function to test solution
- "/api/{id}" PUT: update a exercise
- "/api/{id}" DELETE: remove a exercise

### EXECUTE:
- Using API "/complie" on webstie [play.golang.org](https://play.golang.org)

### TEMPLATE:
- "template/html" library

### DATABASE:
- MongoDB

### DEPLOY:
- Docker

## BUILD
You can clone project and type command line below to run it:

```bash
go run server.go
```
Open your browser and access to link [localhost:8080](http://localhost:8080)