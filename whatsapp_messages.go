package gozoko

import (
	"context"
	"net/http"
)

const messagesBasePath = "v2/messages"

type WhatsAppMessagesService interface {
	GetMessage() error
	SendTemplateMessage(context.Context, *SendMessageRequest) (*SendMessageStatus, *Response, error)
	GetMessageHistory() error
	DeleteMessage() error
}

type Attachment struct {
	Id string `json:"id"`
}

type Message struct {
	Attachments       Attachment `json:"attachments"`
	Customer          Customer   `json:"customer"`
	DeliveryStatus    string     `json:"deliveryStatus"`
	Direction         string     `json:"direction"`
	FileCaption       string     `json:"fileCaption"`
	FileUrl           string     `json:"fileUrl"`
	Id                string     `json:"id"`
	Platform          string     `json:"platform"`
	PlatformSenderId  string     `json:"platformSenderId"`
	PlatformTimestamp string     `json:"platformTimestamp"`
	Text              string     `json:"text"`
	Type              string     `json:"type"`
}

type SendMessageRequest struct {
	Channel          string   `json:"channel"`
	Recipient        string   `json:"recipient"`
	Type             string   `json:"type"`
	Message          string   `json:"message"`
	TemplateId       string   `json:"templateId"`
	TemplateArgs     []string `json:"templateArgs"`
	TemplateLanguage string   `json:"templateLanguage"`
}

type SendMessageStatus struct {
	MessageId  string `json:"messageId"`
	Status     string `json:"status"`
	StatusText string `json:"statusText"`
}

type WhatsAppMessagesServiceImpl struct {
	client *Client
}

func (s *WhatsAppMessagesServiceImpl) GetMessage() error {
	return nil
}

func (s *WhatsAppMessagesServiceImpl) SendTemplateMessage(ctx context.Context, sendMessageRequest *SendMessageRequest) (*SendMessageStatus, *Response, error) {
	if sendMessageRequest == nil {
		return nil, nil, NewArgError("sendMessageRequest", "cannot be nil")
	}
	path := messagesBasePath
	req, err := s.client.NewRequest(ctx, http.MethodPost, path, sendMessageRequest)
	if err != nil {
		return nil, nil, err
	}

	status := new(SendMessageStatus)
	resp, err := s.client.Do(ctx, req, status)
	if err != nil {
		return nil, resp, err
	}

	return status, resp, nil
}

func (s *WhatsAppMessagesServiceImpl) GetMessageHistory() error {
	return nil
}

func (s *WhatsAppMessagesServiceImpl) DeleteMessage() error {
	return nil
}
