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
curl --request POST   --url http://localhost:8080/v1/elerec    --header 'content-type: application/json'   --data '{"id":"1","elerecs":[""]}'


curl --request POST   --url http://localhost:8080/v1/blorec    --header 'content-type: application/json'   --data '{"id":"1","elerecs":[]}'


curl --request GET   --url http://localhost:8080/v1/user/1    --header 'content-type: application/json'



curl --request POST   --url http://localhost:8080/v1/user    --header 'content-type: application/json'   --data '{"id":"1","elerecs":[""]}'



```
