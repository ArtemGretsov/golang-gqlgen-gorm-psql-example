# golang-gqlgen-gorm-psql-example

Example project. 

Monitoring system. Every day the system calls a third-party API.

- You can get information about the day via GraphQL API
- You can track statistical changes (object `difference`). 
- You can add #hashtags to each day.


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
mutation createTag {
  createTag(input: { text: "tag", dayId: 1 }) {
    id
    text
  }
}
```

```graphql
query findDays {
  days {
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
query findDays {
  days(day: "2020-12-12") {
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
