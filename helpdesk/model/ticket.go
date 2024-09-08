package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ticket details
type Ticket struct {
	ID          primitive.ObjectID `json:"id" bson:"_id"`
	Author      string             `json:"author" bson:"author"`
	Subject     string             `json:"subject" bson:"subject"`
	Description string             `json:"description" bson:"description"`
	Assignee    string             `json:"assignee" bson:"assignee"`
	Priority    string             `json:"priority" bson:"priority"`
	IssueType   string             `json:"issue_type" bson:"issue_type"`
	Status      string             `json:"status" bson:"status"`
	Progress    int                `json:"progress" bson:"progress"`
	Remarks     string             `json:"remarks" bson:"remarks"`
	CreatedAt   time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt   *time.Time         `json:"updated_at" bson:"updated_at"`
	DeadLine    time.Time          `json:"dead_line" bson:"dead_line"` // number of days
}
