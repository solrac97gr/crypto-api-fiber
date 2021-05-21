package models

type Message struct {
	Text string
}

func (m *Message) GetText() string {
	return m.Text
}
func (m *Message) SetText(message string) {
	m.Text = message
}

func (m *Message) Validate() (bool, string) {
	if len(m.Text) > 0 {
		return true, "the message len must be more than 0"
	}
	return false, ""
}
