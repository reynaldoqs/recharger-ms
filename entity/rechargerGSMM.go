package entity

import "time"

// RechargerGSMM represents a arduino module recharger
type RechargerGSMM struct {
	ID          string    `json:"id" bson:"_id"`
	PhoneNumber uint      `json:"phoneNumber" validate:"required,gte=9999999,lte=100000000"`
	Company     string    `json:"company" validate:"required,oneof=entel viva tigo"`
	XDay        time.Time `json:"xDay"`
}
