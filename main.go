package main

import (
	"time"
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/davecgh/go-spew/spew"
	_ "github.com/lib/pq"
)

func main() {
	db, err := gorm.Open("postgres", "user=postgres password=admin dbname=user sslmode=disable")
	if err != nil {
		ping := db.DB().Ping()
		fmt.Println(ping)
		fmt.Println(err.Error())
	}

	defer db.Close()

	// create tables
	db.DropTableIfExists(&Client{}, &Product{}, &Author{})
	db.CreateTable(&Client{}, &Product{}, &Author{})

	//create a record
	products := []Product{
        {
            Name: "bolo do caco",
            Type: "pao",
            Weight: "100g",
						Quantity: "12",
        },
    }

	client := Client {
		FirstName: "SuperFrutas Almeida",
		LastName: "Correia, Almeida e Ribeiro LDA",
		Address: "Avenida Antonio Agusto Aguiar, 58H, 1050-017 Lisboa",
		Nif: "111111111",
		Products: products,
	}

	spew.Dump(client)
	db.Create(&client)

	// change a record
	client.FirstName = "Super Frutas"
	db.Debug().Save(&client)

}

type Client struct {
	gorm.Model
	FirstName string
	LastName  string
	Address		string
	Nif				string
	Products  []Product
}

type Product struct {
	gorm.Model
	Name        string
	Type 				string
	Weight			string
	Quantity		string
	PublishDate time.Time
	ClientID    uint     `sql:"index"`
	Authors     []Author `gorm:"many2many:products_authors"`
}

type Author struct {
	gorm.Model
	FirstName string
	LastName  string
}

// define name of tables
func (c *Client) TableName() string {
	return "Client"
}

func ( *Product) TableName() string {
	return "Product"
}

func (a *Author) TableName() string {
	return "Author"
}
