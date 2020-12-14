# API Pokemon

### Installation:

```
git clone https://github.com/ablarry/golang-bootcamp-2020.git
```

### Quickstart
```
Available commands:
	make install			    Install dependencies.
	make run			        Run default command.
	make test			        Run tests.
	make coverage			    Generate coverage in html.
	make linter                 Run linter inspect
	make mocks                  Create mocks for unit tests
```
### Configuration

The path of  file yml which has the configuration of DB and server settings.
```
- config /conf/app.yml
```

### Project dirs:

| Dir | Description |
|:---|:---|
| cmd | Contains buildables files . |
| pkg/api | Contains handlers for REST. |
| pkg/controller| Contains controller to manage request and response. |
| pkg/service | Business layer service. |
| pkg/repo| Persistence layer with queries. |
| pkg/rest | REST helpers for use inside controller |
| swaggerui | dist Swagger UI|
  
### Swagger UI

Review and interact with the API with the Swagger UI.
Running the API open your browser:
```
http://server:port/swagger
```



### Wizeline requirements:
Read requirements and delivery dates [Wizeline](REQUIREMENTS.md)
