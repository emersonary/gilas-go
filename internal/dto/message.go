package dto

import "gitgub.com/emersonary/gilasw/go/pkg/model"

type CreateMessage struct {
	CategoryID  model.ID `json:"categoryid"`
	MessageText string   `json:"messagetext"`
}
