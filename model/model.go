package model

import (
	"time"

	"github.com/google/uuid"
)

type Payment struct {
	// ID               primitive.ObjectID
	Sender           string
	Payment_id       uuid.UUID
	Chain_id         int64
	Initiated_at     string
	Status           string
	Created_by       string
	Payment_type     string
	Updated_at       time.Time
	Created_at       time.Time
	Completed_at     string
	Transaction_hash string
	Task_id          string
	Currency         string
	Amount           string
	Dao              string
	Receiver         string
}

// type provider struct {
// 	ID            primitive.ObjectID
// 	dao_id        string
// 	provider_type string
// 	address       string
// 	created_by    string
// 	created_at    string
// 	updated_at    string
// 	chain_id      string
// 	is_default    string
// 	name          string
// }
