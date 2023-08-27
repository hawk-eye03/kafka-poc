package models

import "time"

type Base struct {
	ID         int64     `gorm:"type:primary_key;column:user_id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	ModifiedAt time.Time `gorm:"column:modified_at"`
	IsDeleted  int32     `gorm:"default:0"`
}

type WhatsappChatDailySummaries struct {
	Base
	SummaryDate      time.Time `gorm:"summary_date"`
	PromptTokens     int32     `gorm:"column:prompt_tokens"`
	CompletionTokens int32     `gorm:"column:completion_tokens"`
	QuestionCount    int32     `gorm:"column:question_count"`
	ImageCount       int32     `gorm:"column:image_count"`
	ChatCount        int32     `gorm:"column:chat_count"`
	TotalCount       int32     `gorm:"column:total_count"`
}
