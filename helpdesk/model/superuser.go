package model

// ticket details updated by super user
type UpdateTicket struct {
	ID        string `json:"id" bson:"_id"`
	Priority  string `json:"priority" bson:"priority"`
	IssueType string `json:"issue_type" bson:"issue_type"`
	Status    string `json:"status" bson:"status"`
	Progress  int    `json:"progress" bson:"progress"`
	Remarks   string `json:"remarks" bson:"remarks"`
	DeadLine  int    `json:"dead_line" bson:"dead_line"` // number of days
}
