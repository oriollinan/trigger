package gmail

import (
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/oauth2"
	"trigger.com/api/src/auth"
	"trigger.com/api/src/endpoints/user"
	"trigger.com/api/src/service"
)

type Service interface {
	auth.Authenticator
	service.Service
	Send(client *http.Client, from string, to string, subject string, body string) error
	GetUserFromGoogle(*oauth2.Token) (*GmailUser, error)
	GetUserFromDbByEmail(string) (*user.User, error)
	AddUserToDb(string, *oauth2.Token) error
}

type Handler struct {
	Service
}

type Model struct {
	auth.Authenticator
	Mongo *mongo.Client
}

type GmailUser struct {
	EmailAddress  string `json:"emailAddress" bson:"email"`
	MessagesTotal int64  `json:"messagesTotal"`
	ThreadsTotal  int64  `json:"threadsTotal"`
	HistoryId     string `json:"historyId"`
}

type DbUser struct {
	Email        string    `bson:"email"`
	AccessToken  string    `bson:"access_token"`
	RefreshToken string    `bson:"refresh_token"`
	Expiry       time.Time `bson:"expiry"`
	TokenType    string    `bson:"token_type"`
}

type Event struct {
	Message struct {
		Data         string `json:"data"`
		MessageId    string `json:"messageId"`
		Message_id   string `json:"message_id"`
		PublishTime  string `json:"publishTime"`
		Publish_time string `json:"publish_time"`
	} `json:"message"`
	Subscription string `json:"subscription"`
}

type EventData struct {
	EmailAddress string `json:"emailAddress"`
	HistoryId    int64  `json:"historyId"`
}

type HistoryList struct {
	History       []History `json:"history"`
	NextPageToken string    `json:"nextPageToken"`
	HistoryId     string    `json:"historyId"`
}

type History struct {
	Id              string           `json:"id"`
	Messages        []Message        `json:"messages,omitempty"`
	MessagesAdded   []MessageAdded   `json:"messagesAdded,omitempty"`
	MessagesDeleted []MessageDeleted `json:"messagesDeleted,omitempty"`
	LabelsAdded     []LabelAdded     `json:"labelsAdded,omitempty"`
	LabelsRemoved   []LabelRemoved   `json:"labelsRemoved,omitempty"`
}

type Message struct {
	Id           string      `json:"id"`
	ThreadId     string      `json:"threadId"`
	LabelIds     []string    `json:"labelIds"`
	Snippet      string      `json:"snippet"`
	HistoryId    string      `json:"historyId"`
	InternalDate string      `json:"internalDate"`
	Payload      MessagePart `json:"payload"`
	SizeEstimate int64       `json:"sizeEstimate"`
	Raw          string      `json:"raw"`
}

type MessagePart struct {
	PartId   string          `json:"partId"`
	MimeType string          `json:"mimeType"`
	Filename string          `json:"filename"`
	Headers  []MessageHeader `json:"headers"`
	Body     MessagePartBody `json:"body"`
	Parts    []MessagePart   `json:"parts"`
}

type MessageHeader struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type MessagePartBody struct {
	AttachmentId string `json:"attachmentId"`
	Size         int64  `json:"size"`
	Data         string `json:"data"`
}

type MessageAdded struct {
	Message Message `json:"message"`
}

type MessageDeleted struct {
	Message Message `json:"message"`
}

type LabelAdded struct {
	Message  Message  `json:"message"`
	LabelIds []string `json:"labelIds"`
}

type LabelRemoved struct {
	Message  Message  `json:"message"`
	LabelIds []string `json:"labelIds"`
}
