package schema

import (
	"time"
)

type HistoryLog struct {
	Id           string `json:"id" gorm:"primaryKey;type:varchar(20)"`
	UserId       string `json:"user_id"`
	Description  string `gorm:"size:50;" json:"title"`
	ErrorMessage string `gorm:"size:255;" json:"error_message"`
	Path         string `gorm:"size:255;" json:"path"`
	ClientIP     string `gorm:"size:255;" json:"client_ip"`
	Method       string `gorm:"size:255;" json:"method"`
	Application  string `gorm:"size:255;" json:"application"`
	CreatedAt    time.Time
}

func (HistoryLog) TableName() string {
	return "history_log"
}
