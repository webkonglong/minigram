# minigram

## run
```bash
go mod vender

go run main.go
```

## test
```bash
# new a user, id should be unionId
curl --request POST   --url http://localhost:8080/v1/user    --header 'content-type: application/json'   --data '{"id":"1","elerecs":[""]}'

# logout:
{"message":"User created Successfully","status":201}


# get a user
curl --request GET   --url http://localhost:8080/v1/user/1    --header 'content-type: application/json'

# logout:
{"data":{"id":"1","elerecs":null},"message":"user info","status":200}


# new a electronic receipt
curl --request POST   --url http://localhost:8080/v1/elerec    --header 'content-type: application/json'   --data '{"ID":"1","user_id":"1","shop_name":"tyn品牌","total_price":-999.00,"created_at":"2020-09-24T23:07:25.678773-07:00","pay_method":"微信","ticket":-20,"serial_num":"2020042363186489198787","items": [{"name":"tyn满减商品01","amount":1,"price":333.0},{"name":"tyn满减商品02","amount":1,"price":333.0},{"name":"tyn满减商品03","amount":1,"price":313.0}],"pos_num":"0058"}'


curl --request POST   --url http://localhost:8080/v1/blorec    --header 'content-type: application/json'   --data '{"id":"1","elerecs":[]}'


curl --request GET   --url http://localhost:8080/v1/user/1    --header 'content-type: application/json'



curl --request POST   --url http://localhost:8080/v1/user    --header 'content-type: application/json'   --data '{"id":"1","elerecs":[""]}'



```
