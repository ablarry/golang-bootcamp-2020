# API Pokemon

### Installation:

```
git clone https://github.com/ablarry/golang-bootcamp-2020.git
```

### Quickstart
```
Available commands:
	make install			 Install dependencies.
	make run			     Run default command.
	make test			     Run tests.
	make coverage			 Generate coverage in html.
	make linter              Run linter inspect
	make mocks               Create mocks for unit tests
```
### Configuration

The path of  file yml which has the configuration of DB and server settings.
```
- config /conf/app.yml
```

### Project dirs:

| Dir | Description |
|:---|:---|
| cmd               | Contains buildables files . |
| pkg/api           | Contains handlers for REST. |
| pkg/controller    | Contains controller to manage request and response. |
| pkg/service       | Business layer service. |
| pkg/repo          | Persistence layer with queries. |
| pkg/rest          | REST helpers for use inside controller |
| swaggerui         | dist Swagger UI|
  
### Swagger UI

Review and interact with the API with the Swagger UI.
Running the API open your browser:
```
http://server:port/swagger
```
### Data Base
The information of this API is store in Memory Data Base [po](https://www.elephantsql.com/)

* Schema ddl - 
``` sql
CREATE TABLE pokemon (
  id INT PRIMARY KEY,
  name VARCHAR(100) NOT NULL,
  type VARCHAR(100) NOT NULL,
  category VARCHAR(100) NOT NULL,
  weakness VARCHAR(100) DEFAULT NULL
);
``` 
* Data dml - 

``` sql
INSERT INTO pokemon (id, name, type, category, weakness) VALUES
    (1,'Bulbasaur', 'Grass', 'Seed', 'Fire'),
    (2,'Ivysaur', 'Grass', 'Seed', 'Fire'),
    (3,'Venusaur','Grass', 'Seed', 'Fire'),
    (4, 'Charmander', 'Fire', 'Lizard', 'Water'),
    (5, 'Charmeleon', 'Fire', 'Lizard', 'Water'),
    (6, 'Charizard', 'Fire', 'Lizard', 'Water');
``` 
Configure data connection in /config/app.yaml, review example /config/app.yaml.dist

### Wizeline requirements:
Read requirements and delivery dates [Wizeline](REQUIREMENTS.md)
