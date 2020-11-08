package dbmodels

import (
	"time"
)

// Repository model represents the Repositories collection in database
type Repository struct {
	ID                  int       `json:"id,omitempty" bson:"id,omitempty"`
	CreatedTimestampUTC time.Time `json:"createdTimestampUTC,omitempty" bson:"CreatedTimestampUTC,omitempty"`
	UpdatedTimestampUTC time.Time `json:"updatedTimestampUTC,omitempty" bson:"UpdatedTimestampUTC,omitempty"`

	Name        string `json:"name,omitempty" bson:"name,omitempty"`
	FullName    string `json:"full_name,omitempty" bson:"full_name,omitempty"`
	Description string `json:"description,omitempty" bson:"description,omitempty"`

	URL      string `json:"html_url,omitempty" bson:"html_url,omitempty"`
	Homepage string `json:"homepage,omitempty" bson:"homepage,omitempty"`
	// Language string `json:"language,omitempty" bson:"language,omitempty"`
	Size     int `json:"size,omitempty" bson:"size,omitempty"`
	Watchers int `json:"watchers,omitempty" bson:"watchers,omitempty"`
	Forks    int `json:"forks,omitempty" bson:"forks,omitempty"`
}
