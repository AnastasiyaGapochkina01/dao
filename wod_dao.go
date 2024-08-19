package dao

import (
	"log"
	. "github.com/AnastasiyaGapochkina01/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type WodsDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	COLLECTION = "wods"
)

// Establish a connection to database
func (m *WodsDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

// Find list of wods
func (m *WodsDAO) FindAll() ([]Wod, error) {
	var wods []Wod
	err := db.C(COLLECTION).Find(bson.M{}).All(&wods)
	return wods, err
}

// Find a wod by its id
func (m *WodsDAO) FindById(id string) (Wod, error) {
	var wod Wod
	err := db.C(COLLECTION).FindId(bson.ObjectIdHex(id)).One(&wod)
	return wod, err
}

// Insert a wod into database
func (m *WodsDAO) Insert(wod Wod) error {
	err := db.C(COLLECTION).Insert(&wod)
	return err
}

// Delete an existing wod
func (m *WodsDAO) Delete(wod Wod) error {
	err := db.C(COLLECTION).Remove(&wod)
	return err
}

// Update an existing wod
func (m *WodsDAO) Update(wod Wod) error {
	err := db.C(COLLECTION).UpdateId(wod.ID, &wod)
	return err
}