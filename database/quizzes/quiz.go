package quizzes

import (
	dbInstance "backend/database"
	quizzesModel "backend/models/quizzes"
	"time"
)

// CreateQuiz inserts a new quiz to the database
func CreateQuiz(quiz *quizzesModel.Quiz) {
	dbInstance.GetDBConnection().Create(quiz)
}

//GetPastQuizzes retrieves list of past quizzes
func GetPastQuizzes(year int) []quizzesModel.Quiz {
	var pastQuizzes []quizzesModel.Quiz
	dbInstance.GetDBConnection().Where("year = ? AND end_time < ?", year, time.Now()).Find(&pastQuizzes)
	return pastQuizzes
}

//GetFutureQuizzes retrieves list of future quizzes
func GetFutureQuizzes(year int) []quizzesModel.Quiz {
	var futureQuizzes []quizzesModel.Quiz
	dbInstance.GetDBConnection().Where("year = ? AND start_time > ?", year, time.Now()).Find(&futureQuizzes)
	return futureQuizzes
}

//GetCurrentQuizzes retrieves list of current quizzes
func GetCurrentQuizzes(year int) []quizzesModel.Quiz {
	var currentQuizzes []quizzesModel.Quiz
	currentTime := time.Now()
	dbInstance.GetDBConnection().Where("year = ? AND start_time < ? AND end_time > ?", year, currentTime, currentTime).Find(&currentQuizzes)
	return currentQuizzes
}
