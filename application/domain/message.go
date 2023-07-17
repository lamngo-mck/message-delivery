package domain

import "time"

type Message struct {
	ID         string                 `json:"id"`
	Name       string                 `json:"name"`
	Type       MessageType            `json:"type"`
	Schedule   *Schedule              `json:"schedule"`
	Endpoints  []string               `json:"endpoints"`
	TemplateId string                 `json:"template_id"`
	Attributes map[string]interface{} `json:"attributes"`
	IsBatching bool                   `json:"is_batching"`
	Source     string                 `json:"source"`
	Priority   Priority               `json:"priority"`
}

type MessageType string

var (
	MessageTriggered MessageType = "Triggered"
	MessageScheduled MessageType = "Scheduled"
)

type Priority string

var (
	PriorityHigh   Priority = "High"
	PriorityMedium Priority = "Medium"
	PriorityLow    Priority = "Low"
)

type Schedule struct {
	StartTime time.Time
	EndTime   time.Time
	Interval  time.Duration
}

func (s Schedule) IsValid() bool {
	if s.EndTime.Unix() < s.StartTime.Unix() {
		return false
	}
	return true
}
