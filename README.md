# Sms Sender Api

## Requirements
1. Redis

## Usage

The application contains different configuration files for different local, dev, tst, acc, prd environments. Environment information can be passed as an environmental variable with the APP_ENV parameter. If environment information is not given, the application runs in the local environment.
```shell,
http://localhost:8050/swagger/index.html
```

You can run the application with the following command.

```shell,
docker-compose up -d
```

The application basically consists of 4 endpoints.

- /api/v1/cronjob/start - PATCH: This endpoint serves to start a cronjob.
- /api/v1/cronjob/stop - PATCH: This endpoint serves to stop a cronjob.
- /api/v1/notifications/sms - POST: This endpoint adds unsent status sms to the Redis database.
- /api/v1/notifications/sms - GET: This endpoint returns the list of sent SMSs.

![img.png](img.png)