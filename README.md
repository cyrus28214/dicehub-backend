# wx-miniprogram-backend

后端服务器

## install dependencies

```sh
go mod tidy
```
    
## run

create `.env` file and set your environment variables

```sh
DB_HOST= # required
DB_PORT= # required
DB_USER= # required
DB_PASSWORD= # required
DB_NAME= # required

SERVER_PORT= # optional, default is 8080

LOG_PATH= # optional, default is stderr
LOG_LEVEL= # optional, default is info
```

run with your environment variables

```sh
source .env
go run cmd/server/main.go
```

