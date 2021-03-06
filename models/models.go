package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

//const : ""
const (
	COLLECTIONNAME string = "meetings"
)

//Participants : ""
type Participants struct {
	Name  string `json:"name" bson:"name,omitempty"`
	Email string `json:"email" bson:"email,omitempty"`
	RSVP  string `json:"rsvp" bson:"rsvp,omitempty"`
}

//Meetings : ""
type Meetings struct {
	ID                bson.ObjectId `json:"id" form:"id," bson:"_id,omitempty"`
	UniqueID          string        `json:"uniqueID" bson:"uniqueID,omitempty"`
	Title             string        `json:"title" bson:"title,omitempty"`
	Participants      Participants  `json:"participants" bson:"participants,omitempty"`
	StartTime         string        `json:"startTime" bson:"startTime,omitempty"`
	EndTime           string        `json:"endTime" bson:"endTime,omitempty"`
	CreationTimestamp time.Time     `json:"creationTimestamp" bson:"creationTimestamp,omitempty"`
}

//ResponseSruct : ""
type ResponseSruct struct {
	Response Response `json:"response,omitempty"`
}

// Pagination : Pagination struct
type Pagination struct {
	PageNum   int    `json:"pageNum" bson:"pageNum"`
	Limit     int    `json:"limit" bson:"limit"`
	SortBy    string `json:"sortBy" bson:"sortBy"`
	SortOrder int    `json:"sortOrder" bson:"sortOrder"`
	Count     int    `json:"count" bson:"count"`
	NextPage  int    `json:"nextPage" bson:"nextPage"`
	PrevPage  int    `json:"prevPage" bson:"prevPage"`
	TotalPage int    `json:"totalPage" bson:"totalPage"`
}

//Response : ""
type Response struct {
	StatusCode int                    `json:"statusCode,omitempty"`
	Message    string                 `json:"message,omitempty"`
	Data       map[string]interface{} `json:"data,omitempty"`
}
