

# My Golang API Playground Project

*Simple Task Management API*,

Simple CRUD API for task management,
with status and files as properties,
status can be complete and incomplte, 
and can have multiple files.

just to test most common use case of API development.  

## Components
* Golang for API server,
* Mysql as database ( only as DB and used as docker services )
* postman/newman cli for automated testing using docker


The project means to be worked together, 
* one can work on the API ( adding model, service, logic)
* one can work on the API Test  ( test scenario using postmane)
* or else, one can work on database initialization ( DB infra should not be dockerized, may need to get optimum configuration)

## Actions plans:

* select the framework to follow,
* create the model, DB and http request and response 
* create orm database layer and persisntence
* create service class for logic 
* create controller for http handler
* create routing for routes handle
* create application for wrapping and initialization
* create swagger API file from the code
* use swagger API json file to create postman/newman automated test plan.
* use docker compose to run it all.



## framework  used

* gin for http framework
* gorm for ORM and database access framework
* swaggo for swagger API file generation
* postman/newman for automated test
* docker compose for auto deployment.


### 
