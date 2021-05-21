package models

type DbMessage struct {
	EncryptedText string
	EncryptionKey string
}

func (dbM *DbMessage) SetEncryptedText(text string) {
	dbM.EncryptedText = text
}
func (dbM *DbMessage) GetEncryptedText() string {
	return dbM.EncryptedText
}

func (dbM *DbMessage) SetEncryptionKey(text string) {
	dbM.EncryptionKey = text
}
func (dbM *DbMessage) GetEncryptionKey() string {
	return dbM.EncryptionKey
}
