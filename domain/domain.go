package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IchibanKuji struct {
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	UserID    string             `json:"userID" bson:"userID"`
	Name      string             `json:"name" bson:"name"`
	Plan      int                `json:"plan" bson:"plan"`
	Cost      int                `json:"cost" bson:"cost"`
	HashTag   []string           `json:"hashTag" bson:"hashTag"`
	Status    int                `json:"status" bson:"status"`
	SellAt    time.Time          `json:"sellAt" bson:"sellAt"`
	CreatedAt time.Time          `json:"createdAt" bson:"createdAt"`
	UpatedAt  time.Time          `json:"updatedAt" bson:"updatedAt"`
}

type IchibanKujiReward struct {
	ID            primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	IchibanKujiID string             `json:"ichibanKujiID" bson:"ichibanKujiID"`
	IsLast        int                `json:"isLast" bson:"isLast"`
	UserID        string             `json:"userID" bson:"userID"`
	Name          string             `json:"name" bson:"name"`
	Description   string             `json:"description" bson:"description"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`
	UpatedAt      time.Time          `json:"updatedAt" bson:"updatedAt"`
}

// -------- 以下為 http request 及 response 結構
// -------- request --------

type CreateIchibanKujiReq struct {
	UserID  string    `form:"userID" validate:"required"`
	Name    string    `form:"name" validate:"required"`
	Plan    int       `form:"plan" validate:"required"`
	Cost    int       `form:"cost" validate:"required"`
	HashTag []string  `form:"hashTag" bson:"hashTag,omitempty"`
	SellAt  time.Time `form:"sellAt" validate:"required"`
}

type UpdateIchibanKujiReq struct {
	Name    string    `form:"name" bson:"name,omitempty"`
	Plan    int       `form:"plan" bson:"plan,omitempty"`
	Cost    int       `form:"cost" bson:"cost,omitempty"`
	HashTag []string  `form:"hashTag" bson:"hashTag,omitempty"`
	Status  int       `form:"status" bson:"status,omitempty"`
	SellAt  time.Time `form:"sellAt" bson:"sellAt,omitempty"`
}

type CreateIchibanKujiRewardReq struct {
	IsLast      []int    `form:"isLast" validate:"required,dive,required"`
	Name        []string `form:"name" validate:"eqcsfield=IsLast,dive,required"`
	Description []string `form:"description" validate:"eqcsfield=IsLast,dive,required"`
}

type UpdateIchibanKujiRewardReq struct {
	IsLast      int    `form:"isLast" bson:"isLast,omitempty"`
	UserID      string `form:"userID" bson:"userID,omitempty"`
	Name        string `form:"name" bson:"name,omitempty"`
	Description string `form:"description" bson:"description,omitempty"`
}
