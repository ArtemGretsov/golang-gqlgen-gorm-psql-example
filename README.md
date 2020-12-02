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
  registrationUser(input: { login: "login", password: "password", firstName: "Artem", lastName: "Gretsov"}) {
    id
    firstName
    lastName
  }
}
```

```graphql
mutation {
  loginUser(input: { login: "login", password: "password"}) {
    id
    token
    firstName
    lastName
  }
}
```
