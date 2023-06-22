package in

import "message-delivery/application/domain"

type EventCollector interface {
	Collect(event domain.Event) error
}
