# VALIDATE EXERCISE

Create a simple website to allow user to validate their implementation for Go exercise during the time they learn Go

Admin create/update/delete exercise with:
- Title
- Description
- Upload solution_test.go

Candidate upload solution.go, system validate their solution by running with solution_test.go

## Requirement
- Lang: Go
- Web library: gorialla/mux
- UI: HTML - Go server side rendering (using template/html)
- DB: MongoDB
- Deployment: Docker
- Run/build go code: use API from [play.golang.org](https://play.golang.org)

## Structure
### Routing:
- `/`: show list all exercises
- `/admin`: show list all exercises
- `/exercise`:
    + `/{id}`: show an exercise with `id` in database
    + `/create/`: create an new exercise
    + `/edit/{id}`: edit an exercise
    + `/delete/{id}`: delete an exercise

### API:
- `/api/exercise` GET: get all exercises
- `/api/exercise` POST: create a new exercise
    + `/{id}` GET: get an exercise with `id` in database
    + `/{id}` PUT: update an exercise
    + `/{id}` DELETE: delete an exercise

### Execute: 
- Using API "/fmt" and "/complie" on webstie [play.golang.org](https://play.golang.org) to reformat and run test code

### Library:
- [github.com/google/uuid](https://github.com/google/uuid) v1.1.1
- [github.com/gorilla/mux](https://github.com/gorilla/mux) v1.7.3
- [github.com/kelseyhightower/envconfig](https://github.com/kelseyhightower/envconfig) v1.4.0
- [go.mongodb.org/mongo-driver](https://github.com/mongodb/mongo-go-driver) v1.0.4

## Build
### Import libraries
Enable and using go module in current terminal
```bash
$ export GO111MODULE=on
$ go mod vendor
```
### Docker
Build the docker image
```bash
$ make docker
```
Run project in docker. Open your browser and access to link [localhost:8080](http://localhost:8080). Press `Ctrl+C` to stop
```bash
$ make compose
```
Clean temporary data from docker
```bash
$ make docker_prune
```