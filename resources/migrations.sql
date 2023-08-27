-- create new table for whatsapp chat daily summary
CREATE TABLE whatsapp_chat_daily_summaries (
    id            BIGINT AUTO_INCREMENT PRIMARY KEY,
    created_at    TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
    modified_at   TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL ON UPDATE CURRENT_TIMESTAMP,
    is_deleted    TINYINT   DEFAULT 0 NOT NULL,
    summary_date  DATE NOT NULL,
    user_id       BIGINT NOT NULL,
    prompt_tokens INT DEFAULT 0 NOT NULL,
    completion_tokens INT DEFAULT 0 NOT NULL,
    question_count INT,
    image_count   INT DEFAULT 0,
    chat_count    INT DEFAULT 0,
    total_count   INT DEFAULT 0,
    UNIQUE KEY unique_user_summary (user_id, summary_date)
);