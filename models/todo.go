package models

import "github.com/kubex/keystone-go/keystone"

type Todo struct {
	keystone.BaseEntity
	ID        int    `keystone:",lookup" json:"id"`
	Title     string `json:"title"`
	Details   string `json:"details"`
	Completed bool   `json:"completed"`
}
