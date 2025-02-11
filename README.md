# wx-miniprogram-backend

微信小程序开发大作业的后端，使用Go语言开发。

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

WX_APP_ID= # required, 微信小程序的AppID
WX_APP_SECRET= # required, 微信小程序的AppSecret
```

run with your environment variables

```sh
source .env
set -a
go run cmd/server/main.go
```

## API

### 登录

POST /api/login

请求体：
```json
{
    "code": "wx.login()获取的code"
}
```

响应：
```json
{
    "openid": "用户唯一标识",
    "session_key": "会话密钥"
}
```

