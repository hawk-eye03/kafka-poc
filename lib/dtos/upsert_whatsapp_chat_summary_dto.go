package dtos

type UpsertWhatsappChatSummary struct {
	UserID          int64  `json:"userId"`
	PromptTokens    int32  `json:"promptTokens"`
	CompletionToken int32  `json:"completionTokens"`
	QuestionType    string `json:"questionType"`
	ContainsImage   bool   `json:"containsImage"`
	QuestionID      string `json:"questionId"`
}

// {
//     "userId": 22,
//     "promptTokens": 333,
//     "completionTokens": 434,
//     "questionType": "NUMERICAL",
//     "containsImage": true,
//     "questionId": "ffsfdfdfdfdsfdf",
// }
