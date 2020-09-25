package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	orm "github.com/go-pg/pg/v9/orm"
	"github.com/go-pg/pg/v9"
	guuid "github.com/google/uuid"
)

type Todo struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	Completed string      `json:"completed"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type User struct {
	ID 			string		`json:"user_id"`
	EleRecs		[]EleRec	`json:"elerecs"`
}

type EleRec struct {
	EleRecID  	string		`json:"elerec_id"`  	// Electronic receipt 电子小票id
	ShopName	string		`json:"shop_name"` 		// Shop name店家名字
	TotalPrice	float64		`json:"total_price"`	// Total Price 总金额
	CreatedAt	time.Time	`json:"created_at"`		// 小票创建时间
	PayMethod	string		`json:"pay_method"`		// 支付方式
	Ticket		int32		`json:"ticket"`			// 抵用券
	SerialNum	string		`json:"serial_num"`		// 流水号
	Items		[]Item		`json:"item"`			// 具体商品
	PosNum		string		`json:"pos_num"`		// 收银机号
}

type Item struct {
	Name	string	`json:"name"`	// 商品名
	Amount	uint32	`json:"amount"`	// 商品数量
	Price	float64	`json:"price"`	// 商品单价
}

type BloRec struct {
	BloRecID	string		`json:"blorec_id"`	// 区块链小票 id
	TxHash		string 		`json:"tx_hash"`	// 区块链小票存证 hash
	BlockNum	uint32		`json:"block_num"`	// 所在区块
	CreatedAt	time.Time	`json:"created_at"`	// 创建时间
}




// Create Todo Table
func CreateTodoTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&Todo{}, opts)
	if createError != nil {
		log.Printf("Error while creating todo table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("Todo table created")
	return nil
}

// 创建用户表
func CreateUserTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&User{}, opts)
	if createError != nil {
		log.Printf("Error while creating user table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("User table created")
	return nil
}

// 创建电子小票表
func CreateEleRecTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&EleRec{}, opts)
	if createError != nil {
		log.Printf("Error while creating elerec table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("EleRec table created")
	return nil
}

// 创建区块链小票表
func CreateBloRecTable(db *pg.DB) error {
	opts := &orm.CreateTableOptions{
		IfNotExists: true,
	}
	createError := db.CreateTable(&BloRec{}, opts)
	if createError != nil {
		log.Printf("Error while creating blorec table, Reason: %v\n", createError)
		return createError
	}
	log.Printf("BloRec table created")
	return nil
}

// INITIALIZE DB CONNECTION (TO AVOID TOO MANY CONNECTION)
var dbConnect *pg.DB
func InitiateDB(db *pg.DB) {
	dbConnect = db
}

func GetAllTodos(c *gin.Context) {
	var todos []Todo
	err := dbConnect.Model(&todos).Select()

	if err != nil {
		log.Printf("Error while getting all todos, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Todos",
		"data": todos,
	})
	return
}

func GetUser(c *gin.Context) {
	userId := c.Param("userId")
	user := &User{ID:userId}
	err := dbConnect.Select(user)

	if  err != nil {
		log.Printf("Error while getting a user's elerecs, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Todo",
		"data": user,
	})
	return
}

func GetElerec(c *gin.Context) {
	recId := c.Param("recId")
	eleRec := &EleRec{EleRecID:recId}
	err := dbConnect.Select(eleRec)

	if  err != nil {
		log.Printf("Error while getting a elerecs's details, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Electronic receipt not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Electronic receipt detail",
		"data": eleRec,
	})
	return
}


func GetBlorec(c *gin.Context) {
	recId := c.Param("recId")
	bloRec := &BloRec{BloRecID:recId}
	err := dbConnect.Select(bloRec)

	if  err != nil {
		log.Printf("Error while getting a blorec's details, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "blockchain receipt not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "blockchain receipt detail",
		"data": bloRec,
	})
	return
}


func CreateTodo(c *gin.Context) {
	var todo Todo
	c.BindJSON(&todo)
	title := todo.Title
	body := todo.Body
	completed := todo.Completed
	id := guuid.New().String()

	insertError := dbConnect.Insert(&Todo{
		ID:         id,
		Title:       title,
		Body:      body,
		Completed:     completed,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	})
	if insertError != nil {
		log.Printf("Error while inserting new todo into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "Todo created Successfully",
	})
	return
}

func GetSingleTodo(c *gin.Context) {
	todoId := c.Param("todoId")
	todo := &Todo{ID: todoId}
	err := dbConnect.Select(todo)

	if err != nil {
		log.Printf("Error while getting a single todo, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Todo not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single Todo",
		"data": todo,
	})
	return
}

func EditTodo(c *gin.Context) {
	todoId := c.Param("todoId")
	var todo Todo
	c.BindJSON(&todo)
	completed := todo.Completed

	_, err := dbConnect.Model(&Todo{}).Set("completed = ?", completed).Where("id = ?", todoId).Update()
	if err != nil {
		log.Printf("Error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"message":  "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  200,
		"message": "Todo Edited Successfully",
	})
	return
}

func DeleteTodo(c *gin.Context) {
	todoId := c.Param("todoId")
	todo := &Todo{ID: todoId}

	err := dbConnect.Delete(todo)
	if err != nil {
		log.Printf("Error while deleting a single todo, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Todo deleted successfully",
	})
	return
}
