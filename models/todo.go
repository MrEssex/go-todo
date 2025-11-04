package models

import "github.com/keystonedb/sdk-go/keystone"

type Todo struct {
	keystone.BaseEntity
	ID        int    `keystone:"_entity_id" json:"id"`
	Title     string `json:"title"`
	Details   string `json:"details"`
	Completed bool   `json:"completed"`
}
