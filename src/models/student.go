package models

import (
	"fmt"
	config "gofiber-mongodb/src/config"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

type Student struct {
	Name  string `bson:"name"`
	Grade int    `bson:"Grade"`
}

func Find() []Student {
	db, err := config.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	csr, err := db.Collection("student").Find(config.Ctx, bson.M{"name": "Ethan Winter"})
	if err != nil {
		log.Fatal(err.Error())
	}

	result := make([]Student, 0)
	for csr.Next(config.Ctx) {
		var row Student
		err := csr.Decode(&row)
		if err != nil {
			log.Fatal(err.Error())
		}

		result = append(result, row)
	}

	if len(result) > 0 {
		fmt.Println("Name  :", result[0].Name)
		fmt.Println("Grade :", result[0].Grade)
	}

	return result
}

func Insert(item *Student) {
	db, err := config.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	_, err = db.Collection("student").InsertOne(config.Ctx, Student{item.Name, item.Grade})
	if err != nil {
		log.Fatal(err.Error())
	}

}

func Update(name string, newStudent *Student) {
	db, err := config.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	var selector = bson.M{"name": name}
	var changes = Student{newStudent.Name, newStudent.Grade}
	_, err = db.Collection("student").UpdateOne(config.Ctx, selector, bson.M{"$set": changes})
	if err != nil {
		log.Fatal(err.Error())
	}

}

func Delete(name string) {
	db, err := config.Connect()
	if err != nil {
		log.Fatal(err.Error())
	}

	var selector = bson.M{"name": name}
	_, err = db.Collection("student").DeleteOne(config.Ctx, selector)
	if err != nil {
		log.Fatal(err.Error())
	}
}
