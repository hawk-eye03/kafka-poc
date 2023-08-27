package consumers

import (
	"encoding/json"
	"strings"

	"github.com/hawk-eye03/kafka-poc/lib/config"
	"github.com/hawk-eye03/kafka-poc/lib/constants"
	"github.com/hawk-eye03/kafka-poc/lib/db"
	"github.com/hawk-eye03/kafka-poc/lib/dtos"
	"github.com/hawk-eye03/kafka-poc/lib/models"
	"github.com/hawk-eye03/kafka-poc/lib/repo/chatrepo"
	"go.uber.org/zap"
)

type DailySummaryConsumer struct {
	dbConnect db.MySQLConnection
	chatRepo  chatrepo.ChatRepo
}

func NewDailySummaryConsumer(config *config.ConfigMap) *DailySummaryConsumer {
	mysqlConnect := *db.NewMainDBConnection(config)
	return &DailySummaryConsumer{
		chatRepo: *chatrepo.NewChatRepo(mysqlConnect.DB),
	}
}

func (d *DailySummaryConsumer) GetTopicName() []string {
	return []string{"foo"}
}

func (d *DailySummaryConsumer) ProcessMessage(msg string) {

	msg = strings.Replace(msg, "\\", "", -1)
	msg = msg[1 : len(msg)-1]
	// fmt.Println("msg=", msg)
	chatSummary := dtos.UpsertWhatsappChatSummary{}
	err := json.Unmarshal([]byte(msg), &chatSummary)
	if err != nil {
		zap.L().Error("Error unmarshalling kafka event message", zap.Error(err))
		return
	}

	// var whatsappDailySummaries models.WhatsappChatDailySummaries
	var imageCount int32
	if chatSummary.ContainsImage {
		imageCount += 1
	}

	var questionCount int32
	var chatCount int32

	if chatSummary.QuestionType == constants.NUMERICAL || chatSummary.QuestionType == constants.THEORETICAL {
		questionCount = 1
	} else if chatSummary.QuestionType == constants.CHAT {
		chatCount = 1
	}

	whatsappDailySummaries := models.WhatsappChatDailySummaries{
		Base: models.Base{
			ID: chatSummary.UserID,
		},
		PromptTokens:     chatSummary.PromptTokens,
		CompletionTokens: chatSummary.CompletionToken,
		ImageCount:       imageCount,
		QuestionCount:    questionCount,
		ChatCount:        chatCount,
		TotalCount:       1,
	}

	d.chatRepo.UpsertWhatsappDailySummary(whatsappDailySummaries)
}

// {
//     "userId": 22,
//     "promptTokens": 333,
//     "completionTokens": 434,
//     "questionType": "NUMERICAL",
//     "containsImage": true,
//     "questionId": "ffsfdfdfdfdsfdf",
// }

// val Numerical: Value = Value("NUMERICAL")
//   val Theoretical: Value = Value("THEORETICAL")
//   val Unknown: Value = Value("UNKNOWN")
//   val NoQuestion: Value = Value("NO_QUESTION")
//   val Feedback: Value = Value("FEEDBACK")
//   val NonAcademic: Value = Value("NON_ACADEMIC")
//   val Profane: Value = Value("PROFANE")
//   val Chat: Value = Value("CHAT")
//   val SHARE: Value = Value("SHARE")
