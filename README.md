# minigram

## run
```bash
go mod vender

go run main.go
```

## test
```bash
curl --request POST   --url http://localhost:8080/v1/blorec    --header 'content-type: application/json'   --data '{"id":"1","elerecs":[]}'


curl --request GET   --url http://localhost:8080/v1/blorec/1    --header 'content-type: application/json'



curl --request POST   --url http://localhost:8080/v1/elerec    --header 'content-type: application/json'   --data '{"id":"1","elerecs":[]}'



```
