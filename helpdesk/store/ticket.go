package store

import (
	"context"
	"time"

	"github.com/selvamtech08/helpdesk/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type ticketImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

var Ticket = ticketImpl{
	collection: GetCollection("tickets"),
	ctx:        context.TODO(),
}

func (u *ticketImpl) NewTicket(ticket *model.Ticket) (*model.Ticket, error) {
	ticket.ID = primitive.NewObjectID()
	_, err := u.collection.InsertOne(u.ctx, ticket)
	if err != nil {
		return nil, err
	}

	return ticket, nil
}

func (u *ticketImpl) GetTicketByName(userName string) (*model.Ticket, error) {
	var ticket model.Ticket
	filer := bson.M{"author": userName}
	if err := u.collection.FindOne(u.ctx, filer).Decode(&ticket); err != nil {
		return nil, err
	}

	return &ticket, nil
}

func (u *ticketImpl) GetTicketByID(id string) (*model.Ticket, error) {
	var ticket model.Ticket
	ticketID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filer := bson.M{"_id": ticketID}
	if err := u.collection.FindOne(u.ctx, filer).Decode(&ticket); err != nil {
		return nil, err
	}

	return &ticket, nil
}

func (u *ticketImpl) GetAllTicketByName(userName string) ([]model.Ticket, error) {
	var tickets []model.Ticket
	filer := bson.M{"author": userName}
	cursor, err := u.collection.Find(u.ctx, filer)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var ticket model.Ticket
		if err := cursor.Decode(&ticket); err != nil {
			return nil, err
		}
		tickets = append(tickets, ticket)
	}

	return tickets, nil
}

func (u *ticketImpl) UpdateTicket(superUserName string, updateDetail model.UpdateTicket) error {
	ticket, err := u.GetTicketByID(updateDetail.ID)
	if err != nil {
		return err
	}

	updates := bson.D{}
	updates = append(updates, primitive.E{Key: "assignee", Value: superUserName})
	if updateDetail.IssueType != "" {
		updates = append(updates, primitive.E{Key: "issue_type", Value: updateDetail.IssueType})
	}
	if updateDetail.Priority != "" {
		updates = append(updates, primitive.E{Key: "priority", Value: updateDetail.Priority})
	}
	if updateDetail.Status != "" {
		updates = append(updates, primitive.E{Key: "status", Value: updateDetail.Status})
	}
	if updateDetail.Progress != 0 {
		updates = append(updates, primitive.E{Key: "progress", Value: updateDetail.Progress})
	}
	if updateDetail.Remarks != "" {
		updates = append(updates, primitive.E{Key: "remarks", Value: updateDetail.Remarks})
	}
	if updateDetail.DeadLine != 0 {
		updates = append(updates, primitive.E{Key: "dead_line", Value: time.Now().AddDate(0, 0, updateDetail.DeadLine)})
	}
	updates = append(updates, primitive.E{Key: "updated_at", Value: time.Now()})

	filter := bson.M{"_id": ticket.ID}
	_, err = u.collection.UpdateOne(context.TODO(), filter, bson.M{"$set": updates})
	if err != nil {
		return err
	}

	return nil
}
