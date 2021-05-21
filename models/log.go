package models

type Log struct {
	Method string
	URL    string
}

func (l *Log) SetMethod(method string) {
	l.Method = method
}

func (l *Log) GetMethod() string {
	return l.Method
}

func (l *Log) SetURL(method string) {
	l.URL = method
}

func (l *Log) GetURL() string {
	return l.URL
}
