# golang-gqlgen-gorm-psql-example

Example project. 

Monitoring system. Every day the system calls a third-party API.

- You can get information about the day via GraphQL API.
- You can track statistical changes (object `difference`). 
- You can add #hashtags to each day. To do this, you need to log in.


#### Collected data

  - Currency rates
  - Temperature and pressure

## Technologies:
  - Golang
  - gqlgen (GraphQL)
  - Gorm
  - PostgreSQL
  - Cron jobs

### Launch of the project
1. Install GO 1.13 or greater
2. If you want the application to collect weather data - register on https://openweathermap.org/ and get your personal API-KEY
3. In the root of the project create a file named `.env` with the contents 
```
WEATHER_API_KEY="The key obtained in step 1"
PORT=5000
DB_NAME=go
DB_PORT=5432
DB_PASS=go
DB_USER=go
DB_HOST=localhost
JWT_KEY=KJdq1LOOkdm23eofoef8wewqoOjonwdn0weWQEpqe2oom3
```
4. Start the database. This is easy to do via `docker`. Run `docker-compose up -d postgres`
5. Run `go run server.go`


## Sample Queries

```graphql
query {
  days(day: "2020-12-12", offset: 0, limit: 10) {
    id
    date
    rate {
      id
      USD
      EUR
      difference {
        EUR
        USD
      }
    }
    weather {
      id
      pressure
      temperature
      difference {
        pressure
        temperature
      }
    }
    tags {
      id
      text
    }
  }
}
```

```graphql
mutation {
  createTag(input: { text: "tag", dayId: 1 }) {
    id
    text
  }
}
```

```graphql
mutation {
  registrationUser(input: { login: "login", password: "password", firstName: "Artem", lastName: "Gretsov" }) {
    id
    firstName
    lastName
  }
}
```

```graphql
mutation {
  loginUser(input: { login: "login", password: "password" }) {
    id
    token
    firstName
    lastName
  }
}
```
