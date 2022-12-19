package repository

import (
	"backEndGo/models"

	"gorm.io/gorm"
)

type ChatRepository interface {
	GetChatAll() (*[]models.Chat, error)
}
type chatDB struct {
	db *gorm.DB
}

// GetChatAll implements ChatRepository
func (c chatDB) GetChatAll() (*[]models.Chat, error) {
	chats := []models.Chat{}
	result := c.db.Preload("Customer").Find(&chats)
	if result.Error != nil {
		return nil, result.Error
	}
	return &chats, nil
}

func NewChatRepository() ChatRepository {
	db, err := NewDatabaseConnection()
	if err != nil {
		return nil
	}

	return chatDB{db}
}
