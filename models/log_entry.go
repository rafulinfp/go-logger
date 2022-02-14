package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// LogEntry ... model for the log entry.
type LogEntry struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	ServiceName string        `bson:"service_name" json:"serviceName"`
	Type        string        `bson:"type" json:"type"`
	Message     string        `bson:"message" json:"message"`
	Timestamp   time.Time     `bson:"timestamp" json:"timestamp"`
}
