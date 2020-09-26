# minigram

## test
terminal 0:
```bash
$ go run main.go
2020/09/26 09:41:19 Connected to db
2020/09/26 09:41:19 User table created
2020/09/26 09:41:20 EleRec table created
2020/09/26 09:41:20 BloRec table created
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:	export GIN_MODE=release
 - using code:	gin.SetMode(gin.ReleaseMode)

[GIN-debug] GET    /                         --> github.com/flyq/minigram/routes.welcome (3 handlers)
[GIN-debug] POST   /v1/user                  --> github.com/flyq/minigram/controllers.CreateUser (3 handlers)
[GIN-debug] POST   /v1/elerec                --> github.com/flyq/minigram/controllers.CreateElerec (3 handlers)
[GIN-debug] POST   /v1/blorec                --> github.com/flyq/minigram/controllers.CreateBlorec (3 handlers)
[GIN-debug] GET    /v1/user/:userId          --> github.com/flyq/minigram/controllers.GetUser (3 handlers)
[GIN-debug] GET    /v1/elerec/:recId         --> github.com/flyq/minigram/controllers.GetElerec (3 handlers)
[GIN-debug] GET    /v1/blorec/:recId         --> github.com/flyq/minigram/controllers.GetBlorec (3 handlers)
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2020/09/26 - 09:47:35 | 201 |    1.742221ms |       127.0.0.1 | POST     /v1/user





```


terminal 1:
```bash
# new a user, id should be unionId
$ curl --request POST   --url http://localhost:8080/v1/user    --header 'content-type: application/json'   --data '{"id":"3","elerecs":[""]}'
{"message":"User created Successfully","status":201}

# get a user
$ curl --request GET   --url http://localhost:8080/v1/user/3    --header 'content-type: application/json'
{"data":{"id":"3","elerecs":null},"message":"user info","status":200}

# new a electronic receipt
$ curl --request POST   --url http://localhost:8080/v1/elerec    --header 'content-type: application/json'   --data '{"ID":"1","user_id":"3","shop_name":"tyn品牌","total_price":-999.00,"created_at":"2020-09-24T23:07:25.678773-07:00","pay_method":"微信","ticket":-20,"serial_num":"2020042363186489198787","items": [{"name":"tyn满减商品01","amount":1,"price":333.0},{"name":"tyn满减商品02","amount":1,"price":333.0},{"name":"tyn满减商品03","amount":1,"price":313.0}],"pos_num":"0058"}'
{"message":"User's electric receipt update Successfully","status":200}

# get a eletronic receipt
$ curl --request GET   --url http://localhost:8080/v1/elerec/25    --header 'content-type: application/json'
{"data":{"id":"25","user_id":"3","shop_name":"tyn品牌","total_price":-999,"created_at":"2020-09-24T23:07:25.678773-07:00","pay_method":"微信","ticket":-20,"serial_num":"2020042363186489198787","items":[{"name":"tyn满减商品01","amount":1,"price":333},{"name":"tyn满减商品02","amount":1,"price":333},{"name":"tyn满减商品03","amount":1,"price":313}],"pos_num":"0058"},"message":"Electronic receipt detail","status":200}

# get a user


curl --request GET   --url http://localhost:8080/v1/user/1    --header 'content-type: application/json'


 curl --request POST   --url http://localhost:8080/v1/blorec    --header 'content-type: application/json'   --data '{"ID":"5","TxHash":"0x645d39378a7b9e569b95736fd93cec17e1715d07f375e8e3f2af966fb25ae79f","BlockNum":1425015,"CreatedAt":"2020-09-24T23:07:25.678773-07:00"}'


curl --request POST   --url http://localhost:8080/v1/user    --header 'content-type: application/json'   --data '{"id":"1","elerecs":[""]}'



```
