package controllers

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v9"
	orm "github.com/go-pg/pg/v9/orm"
)

type User struct {
	ID      string   `json:"id"`
	EleRecs []string `json:"elerecs"`
}

type EleRec struct {
	ID         string    `json:"id"`          // Electronic receipt 电子小票id
	UserID     string    `json:"user_id"`     // 用户id，表示这是谁的小票
	ShopName   string    `json:"shop_name"`   // Shop name店家名字
	TotalPrice float64   `json:"total_price"` // Total Price 总金额
	CreatedAt  time.Time `json:"created_at"`  // 小票创建时间
	PayMethod  string    `json:"pay_method"`  // 支付方式
	Ticket     float64   `json:"ticket"`      // 抵用券
	SerialNum  string    `json:"serial_num"`  // 流水号
	Items      []Item    `json:"items"`       // 具体商品
	PosNum     string    `json:"pos_num"`     // 收银机号
}

type Item struct {
	Name   string  `json:"name"`   // 商品名
	Amount uint32  `json:"amount"` // 商品数量
	Price  float64 `json:"price"`  // 商品单价
}

type BloRec struct {
	ID        string    `json:"id"`         // 区块链小票 id
	TxHash    string    `json:"tx_hash"`    // 区块链小票存证 hash
	BlockNum  uint32    `json:"block_num"`  // 所在区块
	CreatedAt time.Time `json:"created_at"` // 创建时间
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

func GetUser(c *gin.Context) {
	userId := c.Param("userId")
	user := &User{ID: userId}
	err := dbConnect.Select(user)

	if err != nil {
		log.Printf("Error while getting a user's elerecs, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "user not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "user info",
		"data":    user,
	})
	return
}

func GetElerec(c *gin.Context) {
	recId := c.Param("recId")
	eleRec := &EleRec{ID: recId}
	err := dbConnect.Select(eleRec)

	if err != nil {
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
		"data":    eleRec,
	})
	return
}

func GetBlorec(c *gin.Context) {
	recId := c.Param("recId")
	bloRec := &BloRec{ID: recId}
	
	err := dbConnect.Model(bloRec).WherePK().Select()
	if err != nil {
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
		"data":    bloRec,
	})
	return
}

func CreateUser(c *gin.Context) {
	var user User
	c.BindJSON(&user)
	user_id := user.ID
	elerecs := user.EleRecs

	insertError := dbConnect.Insert(&User{
		ID:      user_id,
		EleRecs: elerecs,
	})

	if insertError != nil {
		log.Printf("Error while inserting new user into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "User created Successfully",
	})
	return
}

func CreateElerec(c *gin.Context) {
	var elerec EleRec
	var items []Item

	c.BindJSON(&elerec)
	elerec_id := elerec.ID
	user_id := elerec.UserID
	shop_name := elerec.ShopName
	total_price := elerec.TotalPrice
	created_at := elerec.CreatedAt
	pay_method := elerec.PayMethod
	ticket := elerec.Ticket
	serial_num := elerec.SerialNum
	items = elerec.Items
	pos_num := elerec.PosNum

	insertError := dbConnect.Insert(&EleRec{
		ID:         elerec_id,
		UserID:     user_id,
		ShopName:   shop_name,
		TotalPrice: total_price,
		CreatedAt:  created_at,
		PayMethod:  pay_method,
		Ticket:     ticket,
		SerialNum:  serial_num,
		Items:      items,
		PosNum:     pos_num,
	})

	if insertError != nil {
		log.Printf("Error while inserting new electronic receipt into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong: Error while inserting new electronic receipt into db",
		})
		return
	}


	user := &User{ID: user_id}
	getErr := dbConnect.Model(user).WherePK().Select()
	log.Printf("user first:: %v\n", user)

	if getErr != nil {
		log.Printf("can't get a user, so new one, Database: %v\n", getErr)
		user.EleRecs = append(user.EleRecs,elerec_id)
		insertUserErr := dbConnect.Insert(user)
		if insertUserErr != nil {
			log.Printf("can't new one user, Database: %v\n", insertUserErr)
			c.JSON(http.StatusInternalServerError, gin.H{
				"status":  http.StatusInternalServerError,
				"message": "Something went wrong: Error while inserting new user into db",
			})
			return
		}
		c.JSON(http.StatusCreated, gin.H{
			"status":  http.StatusCreated,
			"message": "New user and eletronic created Successfully",
		})
		return
	}

	user.EleRecs = append(user.EleRecs, elerec_id)

	log.Printf("updated user: %v\n", user)
	
	_, err := dbConnect.Model(user).WherePK().Update()
	if err != nil {
		log.Printf("insert electronic receipt to user error, Reason: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong: Error while add new electronic receipt to a user",
		})
		return
	}


	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "New electric and User's electric receipt update Successfully",
	})
	return
}

func CreateBlorec(c *gin.Context) {
	var blorec BloRec
	c.BindJSON(&blorec)
	blorec_id := blorec.ID
	tx_hash := blorec.TxHash
	block_num := blorec.BlockNum
	created_at := blorec.CreatedAt

	insertError := dbConnect.Insert(&BloRec{
		ID:        blorec_id,
		TxHash:    tx_hash,
		BlockNum:  block_num,
		CreatedAt: created_at,
	})

	if insertError != nil {
		log.Printf("Error while inserting new block receipt into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "receipt created Successfully",
	})
	return
}
