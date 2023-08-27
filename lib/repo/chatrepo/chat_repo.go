package chatrepo

import (
	"time"

	"github.com/hawk-eye03/kafka-poc/lib/models"
	"github.com/jinzhu/gorm"
)

type ChatRepo struct {
	db *gorm.DB
}

func NewChatRepo(db *gorm.DB) *ChatRepo {
	return &ChatRepo{
		db: db,
	}
}

// TODO: handle error during updates
func (c *ChatRepo) UpsertWhatsappDailySummary(dailySummaries models.WhatsappChatDailySummaries) {
	var existingSummary models.WhatsappChatDailySummaries

	res := c.db.First(&existingSummary)
	loc, _ := time.LoadLocation("Asia/Kolkata")
	currentTime := time.Now().In(loc)
	currentDate := currentTime.Truncate(24 * time.Hour)
	if res.Error != nil {
		// Data doesn't exist, create a new row
		if gorm.IsRecordNotFoundError(res.Error) {
			dailySummaries.SummaryDate = currentDate
			dailySummaries.CreatedAt = currentTime
			dailySummaries.ModifiedAt = currentTime
			c.db.Create(&dailySummaries)
		}
	} else {
		// Data exists, update the row according to pre-existing data
		dailySummaries.ChatCount = existingSummary.ChatCount + dailySummaries.ChatCount
		dailySummaries.ImageCount = existingSummary.ImageCount + dailySummaries.ImageCount
		dailySummaries.QuestionCount = existingSummary.QuestionCount + dailySummaries.QuestionCount
		dailySummaries.TotalCount = existingSummary.QuestionCount + 1
		dailySummaries.ModifiedAt = currentTime
		c.db.Model(&existingSummary).Updates(dailySummaries)
	}
}
