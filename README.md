# Youtube API Monitoring APP

## Build & Run (Locally)
### Prerequisites
- go 1.16+
- docker
- migrate ( https://github.com/golang-migrate/migrate )

Create .env file in root directory and add following values:
```dotenv
YOUTUBE_KEY= ##< Your Youtube API key >
CHANNEL_ID= ##< YouTube channel identifiers, separated by commas (example use: UC1CchA0SjApw4T-AYkN7ytg,UClMe70teNK5LpBE3Im_JRSQ ) >
PART=statistics,snippe ##Make sure to add these parameters >
```

Requests:
```
GET localhost:8000/list - get list of channels
GET localhost:8000/stat - update channels with Youtube API request
```

Use `make build && make run` to build&run project.\
If the project is launched for the first time, use `make migrate`.