package quizzes

import (
	"gorm.io/gorm"
)

// Choice struct to store the choice data
type Choice struct {
	gorm.Model
	Text  string `json:"text"`
	MCQ   MCQ    `json:"mcq"`
	MCQID uint   `json:"mcqID"`
}
