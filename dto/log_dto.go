package dto

import (
	"crypto/tls"
	"log"
	"net"
	"time"

	"github.com/rafulinfp/gologger/models"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

//LogDAO ... The logs dataaccess
type LogDAO struct {
	Server   string
	Database string
}

var db *mgo.Database

const (
	// COLLECTION The mongoDB document name
	COLLECTION = "logs"
)

//Connect ... Connect to the database
func (m *LogDAO) Connect() {
	// DialInfo holds options for establishing a session with a MongoDB cluster.
	dialInfo := &mgo.DialInfo{
		Addrs:    []string{"host:port"}, // Get HOST + PORT
		Timeout:  60 * time.Second,
		Database: "db",       // It can be anything
		Username: "user",     // Username
		Password: "password", // PASSWORD
		DialServer: func(addr *mgo.ServerAddr) (net.Conn, error) {
			return tls.Dial("tcp", addr.String(), &tls.Config{})
		},
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

//FindAll ... Find all the elements.
func (m *LogDAO) FindAll() ([]models.LogEntry, error) {
	var logs []models.LogEntry
	err := db.C(COLLECTION).Find(bson.M{}).All(&logs)
	return logs, err
}

//Insert ... Insert a log entry
func (m *LogDAO) Insert(logEntry models.LogEntry) error {
	err := db.C(COLLECTION).Insert(&logEntry)
	return err
}
