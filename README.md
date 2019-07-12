# VALIDATE EXERCISE

Create a simple website to allow user to validate their implementation for Go exercise during the time they learn Go

Admin create/update/delete exercise with:
- Title
- Description
- Upload solution_test.go

Candidate upload solution.go, system validate their solution by running with solution_test.go

## REQUIREMENT
- Lang: Go
- Web library: gorialla/mux
- UI: HTML - Go server side rendering (using template/html)
- DB: MongoDB
- Deployment: Docker
- Run/build go code: use API from [play.golang.org](https://play.golang.org)

## STRUCTURE
### ROUTING:
- "/": show list all exercises with title
- "/admin/": show list all exercises with title
    + "/delete/{id}/": delete a exercise
    + "/create/": create new exercise
    + "/{id}/": edit exercise
- "/exercise/{id}": show a exercise have id in database

### API:
- "/api/" GET: get all exercise
- "/api/" POST: create new exercise
    + "/{id}/" GET: get a exercise have id in database
    + "/{id}/" PUT: update a exercise
    + "/{id}/" DELETE: remove a exercise

### EXECUTE:
- Using API "/fmt" and "/complie" on webstie [play.golang.org](https://play.golang.org) to reformat and run test code

### TEMPLATE:
- "template/html" library

### DATABASE:
- MongoDB

### DEPLOY:
- Docker

### LIBRARY:
- [github.com/gorilla/mux](https://github.com/gorilla/mux)
- [go.mongodb.org/mongo-driver](https://github.com/mongodb/mongo-go-driver)

## BUILD
You can clone project and type command line below to run it in Docker:

```bash
$ docker-compose up -d
```
Open your browser and access to link [localhost:8080](http://localhost:8080)