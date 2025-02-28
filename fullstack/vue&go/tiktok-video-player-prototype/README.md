# Tiktok Video Player Prototype
 - BE: golang, redis & net/http
 - FE: vue, fortawesome & bootstrap

demo: https://kucingliar.xyz/toktok/

# Backend
I design the BE service with MVC structure,
This service will handle for video search & suggestions

Prefix: /api/v1.0/
| Method  | Route |Description|
| ------------- |-------------|-------------|
| GET      | /ping     | check the service is up/down|
| GET      | /search/videos     | searching video|
| GET      | /search/suggestions     | get list keyword suggestions|


service will running on port :8000, you can access page directly like this
http://localhost:8000/api/v1.0/ping

sample : config.json
```
{
    "app" :{
        "name": "TokTok Prototype",
        "version": "1.0",
        "port": "8000",
        "prefix": "/api/v1.0",
        "expiredVideo" : 7
    },
    "security" : {
        "enabledCors": true,
        "whitelistHost" : "http://localhost:5173"
    },
    "service": {
        "redis": {
            "host": "127.0.0.1",
            "port": 6379,
            "password": ""
        }
    },
    "integration" : {
        "pexels" : {
            "host" : "https://api.pexels.com",
            "secretKey" : "************",
            "size": 5
        }
    }
}
```

## How to Run
```
$ go run main.go

Starting server at port 8000
```

# Frontend
I design the FE with SPA (single page application)

## How to Run
```
$ yarn run dev

Port 5173 is in use, trying another one...

  VITE v5.4.10  ready in 179 ms

  ➜  Local:   http://localhost:5174/
  ➜  Network: use --host to expose
  ➜  press h + enter to show help
```

