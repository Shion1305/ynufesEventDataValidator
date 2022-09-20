package model

type EventDataSet struct {
	Data []*EventData
}

func NewEventDataSet(d []*EventData) EventDataSet {
	return EventDataSet{
		Data: d,
	}
}
