# minigram

## 目录
- [minigram](#minigram)
  - [目录](#目录)
  - [环境](#环境)
  - [接口文档](#接口文档)
    - [测试接口](#测试接口)
    - [创建用户](#创建用户)
    - [创建电子小票](#创建电子小票)
    - [创建区块链小票](#创建区块链小票)
    - [获取用户数据](#获取用户数据)
    - [获取电子小票](#获取电子小票)
    - [获取区块链小票](#获取区块链小票)
  - [启动](#启动)
  - [操作数据库【学习笔记】](#操作数据库学习笔记)

## 环境
```bash
$ go version 
go version go1.12.9 linux/amd64

$ psql --version
psql (PostgreSQL) 12.4 (Ubuntu 12.4-1.pgdg18.04+1)
```

## 接口文档

### 测试接口
1. 接口 `/`
2. 示例：
```bash
{"message":"Welcome To API","status":200}
```

### 创建用户
1. 接口 `/v1/user`
2. 数据格式 `{"id":"1","elerecs":[""]}`
3. 示例：
```bash
$ curl --request POST   --url http://localhost:8080/v1/user    --header 'content-type: application/json'   --data '{"id":"1","elerecs":[""]}'

{"message":"User created Successfully","status":201}
```

### 创建电子小票
1. 接口 `/v1/elerec`
2. 数据格式 `{"id":"1","user_id":"3","shop_name":"tyn品牌","total_price":-999.00,"created_at":"2020-09-24T23:07:25.678773-07:00","pay_method":"微信","ticket":-20,"serial_num":"2020042363186489198787","items": [{"name":"tyn满减商品01","amount":1,"price":333.0},{"name":"tyn满减商品02","amount":1,"price":333.0},{"name":"tyn满减商品03","amount":1,"price":313.0}],"pos_num":"0058"}`
3. 示例：
```bash
$ curl --request POST   --url http://localhost:8080/v1/elerec    --header 'content-type: application/json'   --data '{"ID":"1","user_id":"3","shop_name":"tyn品牌","total_price":-999.00,"created_at":"2020-09-24T23:07:25.678773-07:00","pay_method":"微信","ticket":-20,"serial_num":"2020042363186489198787","items": [{"name":"tyn满减商品01","amount":1,"price":333.0},{"name":"tyn满减商品02","amount":1,"price":333.0},{"name":"tyn满减商品03","amount":1,"price":313.0}],"pos_num":"0058"}'

{"message":"User's electric receipt update Successfully","status":201}
```
### 创建区块链小票
1. 接口 `/v1/blorec`
2. 数据格式 `{"id":"5","tx_hash":"0x645d39378a7b9e569b95736fd93cec17e1715d07f375e8e3f2af966fb25ae79f","block_num":1425015,"created_at":"2020-09-24T23:07:25.678773-07:00"}`
3. 示例
```bash
$ curl --request POST   --url http://localhost:8080/v1/blorec    --header 'content-type: application/json'   --data '{"id":"5","tx_hash":"0x645d39378a7b9e569b95736fd93cec17e1715d07f375e8e3f2af966fb25ae79f","block_num":1425015,"created_at":"2020-09-24T23:07:25.678773-07:00"}'

{"message":"receipt created Successfully","status":201}
```

### 获取用户数据
1. 接口 `/v1/user/:userId`
2. 示例
```bash
$ curl --request GET   --url http://localhost:8080/v1/user/3    --header 'content-type: application/json'

{"data":{"id":"3","elerecs":null},"message":"user info","status":200}
```

### 获取电子小票
1. 接口 `/v1/elerec/:recId`
2. 示例
```bash
$ curl --request GET   --url http://localhost:8080/v1/elerec/25    --header 'content-type: application/json'

{"data":{"id":"25","user_id":"3","shop_name":"tyn品牌","total_price":-999,"created_at":"2020-09-24T23:07:25.678773-07:00","pay_method":"微信","ticket":-20,"serial_num":"2020042363186489198787","items":[{"name":"tyn满减商品01","amount":1,"price":333},{"name":"tyn满减商品02","amount":1,"price":333},{"name":"tyn满减商品03","amount":1,"price":313}],"pos_num":"0058"},"message":"Electronic receipt detail","status":200}
```

### 获取区块链小票
1. 接口 `/v1/blorec/:recId`
2. 示例
```bash
$ curl --request GET   --url http://localhost:8080/v1/blorec/5    --header 'content-type: application/json'

{"data":{"id":"5","tx_hash":"","block_num":0,"created_at":"0001-01-01T00:00:00Z"},"message":"blockchain receipt detail","status":200}
```

## 启动
```bash
$ go mod vendor
...


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
```


## 操作数据库【学习笔记】
```sql
# 进入数据库
$ sudo -i -u postgres
postgres@ubuntu:~$

postgres@ubuntu:~$ psql
psql (12.4 (Ubuntu 12.4-1.pgdg18.04+1), server 10.14 (Ubuntu 10.14-1.pgdg18.04+1))
Type "help" for help.
postgres=# 

# 等价于
sudo -i -u postgres psql

# 查看数据库
\l

# 进入数据库
\c middleware_hn

# 查看数据库里面有哪些表
\d

# 查看表详细信息
\d ele_recs

# 创建表
CREATE TABLE table_name(
   column1 datatype,\\
   column2 datatype,
   column3 datatype,
   .....
   columnN datatype,
   PRIMARY KEY( 一个或多个列 )
);

CREATE TABLE COMPANY(
   ID             INT PRIMARY KEY     NOT NULL,
   NAME           TEXT    NOT NULL,
   AGE            INT     NOT NULL,
   ADDRESS        CHAR(50),
   SALARY         REAL
);

# 删除表
DROP TABLE company;

# 创建新用户
CREATE ROLE ubuntu WITH LOGIN CREATEDB ENCRYPTED PASSWORD 'xxxxxxxxx';  

# 创建新数据库
sudo -u postgres createdb -O ubuntu ubuntu

# 重置密码
ALTER USER ubuntu PASSWORD 'newpassword';
```