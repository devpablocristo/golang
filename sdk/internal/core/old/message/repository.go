package message

import (
	csdgocsl "github.com/devpablocristo/golang/sdk/pkg/cassandra/gocql"
)

type CassandraRepository struct {
	csdInst csdgocsl.CassandraClientPort
}

func NewCassandraRepository(inst csdgocsl.CassandraClientPort) *CassandraRepository {
	return &CassandraRepository{
		csdInst: inst,
	}
}

func (r *CassandraRepository) Save(message *Message) error {
	return r.csdInst.GetSession().Query(
		"INSERT INTO messages (id, sender_id, recipient_id, content, timestamp) VALUES (?, ?, ?, ?, ?)",
		message.ID, message.SenderID, message.RecipientID, message.Content, message.Timestamp,
	).Exec()
}

func (r *CassandraRepository) FindByUserID(userID string) ([]Message, error) {
	var messages []Message
	iter := r.csdInst.GetSession().Query(
		"SELECT id, sender_id, recipient_id, content, timestamp FROM messages WHERE recipient_id = ?",
		userID,
	).Iter()
	var message Message
	for iter.Scan(&message.ID, &message.SenderID, &message.RecipientID, &message.Content, &message.Timestamp) {
		messages = append(messages, message)
	}
	return messages, iter.Close()
}
