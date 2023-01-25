## GOLANG Mux + Postgres + RabbitMq + Cognito

## Dependencies

- Docker
- Go 1.17 +
- Make


## Important infos:

- This project is to do examples with mux server and generate a way to trigger the changes on DB made by api requests to a Queue that will be consumed and parsed by another project


   - https://github.com/morlfm/parser


- It was not created a Graceful shutdown just for purposes of who wants to use this project example to have a easy way to test and play arround wiht the api through postman, Curl or whatever :) 

- In this example it was used postgrees databse with Cloud Based: https://www.elephantsql.com but you can use with docker or local 

## Environment variables (direnv)

- Create a .env file in root and add your database credentials like that:

```
export POSTGRES_URL="${POSTGRES_URL}"
```

- Or just export them before start the application 


## Running the project locally

- export your env appropriately accordingly to your OS.

- To start Rabbit docker 

```
$ make rabbit-server-up
```

- To start mux and application:

```
$ make mux-server-up
```

- To run tests

```
$ make run-tests
```

- To remove rabbit container

```
$ make rabbit-server-removal
```

## Database used to store correct data (employees)

- Postgress


## To make requests:
 - http://localhost:8081/api/employees
 - http://localhost:8081/api/employees/{id}


## Architeture and packages in golang format of 

- Application where resided the logic and businees rules
- Repository to deal with Db part
- Cmd to gather and run main
- Ports is where to expose and treat the handlers
- Rabbit to publish our message
- Api-release the API documentation


##  Specific Bussiness rules

- PUT method only allow to update passing body: { "name" and "location"} 

- Rabbit publisher its triggered every change made on database and send all over again all employees to the queue



