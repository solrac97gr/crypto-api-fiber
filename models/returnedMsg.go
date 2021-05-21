package models

type ReturnedMsg struct {
	ID string
	Message
	DbMessage
}

func (r *ReturnedMsg) SetID(id string) {
	r.ID = id
}

func (r *ReturnedMsg) GetID() string {
	return r.ID
}
