package gocelery

import "time"

//Message represents a Celery message
type Message struct {
	ID      string                 `json:"id"`
	Task    string                 `json:"task"`
	Args    []string               `json:"args"`
	Kwargs  map[string]interface{} `json:"kwargs"`
	Retries int                    `json:"retries"`
	Eta     time.Time              `json:"eta"`
	Expires time.Time              `json:"expires"`
}
