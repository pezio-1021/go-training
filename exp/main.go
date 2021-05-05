package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5433
	user     = "admin"
	password = "admin"
	dbname   = "lenslocked_dev"
)

type User struct {
	gorm.Model
	Name  string
	Email string `gorm:"not null;unique_index"`
}

type Order struct {
	gorm.Model
	UserID      uint
	Amount      int
	Description string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&User{}, &Order{})

	var user User
	db.First(&user)
	if db.Error != nil {
		panic(db.Error)
	}

	createOrder(db, user, 1001, "Fake Description #1")
	createOrder(db, user, 1001, "Fake Description #2")
	createOrder(db, user, 1001, "Fake Description #3")
}

func createOrder(db *gorm.DB, user User, amount int, desc string) {
	db.Create(&Order{
		UserID:      user.ID,
		Amount:      amount,
		Description: desc,
	})

	if db.Error != nil {
		panic(db.Error)
	}
}

// func getInfo() (name, email string) {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("What is your name?")
// 	name, _ = reader.ReadString('\n')
// 	name = strings.TrimSpace(name)
// 	fmt.Println("What is your email?")
// 	email, _ = reader.ReadString('\n')
// 	email = strings.TrimSpace(email)
// 	return name, email
// }
