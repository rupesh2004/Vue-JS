package model

import (
	"sync"
	"time"
)

type Student struct {
	Username  string    `bson:"username" json:"username" validate:"required"`
	FirstName string    `bson:"first_name" json:"first_name" validate:"required"`
	LastName  string    `bson:"last_name" json:"last_name" validate:"required"`
	Email     string    `bson:"email" json:"email" validate:"required,email"`
	Password  string    `bson:"password" json:"password" validate:"required,min=6"`
	CreatedOn time.Time `bson:"created_on" json:"created_on"`
	UpdatedOn time.Time `bson:"updated_on" json:"updated_on"`
	IsActive  bool      `bson:"is_active" json:"is_active"`
	Roles     string    `bson:"roles" json:"roles" validate:"required,oneof=admin user"`
	CreatedBy string    `bson:"created_by" json:"created_by"`
}

type Exam struct {
	ExamID               string     `bson:"exam_id" json:"examId" validate:"required"`
	ExamName             string     `json:"examName" bson:"exam_name"`
	ExamDate             *time.Time `json:"examDate" bson:"exam_date"`
	Duration             string     `json:"examDuration" bson:"exam_duration"`
	ExamType             string     `json:"examType" bson:"exam_type"`
	CreatedBy            string     `json:"createdBy" bson:"created_by"`
	CreatedAt            time.Time  `json:"createdAt" bson:"created_at"`
	ApplicationCloseDate *time.Time `json:"applicationCloseDate" bson:"application_close_date"`
}

type ExamApplication struct {
	ExamID    string    `bson:"exam_id" json:"examId"`
	ExamName  string    `bson:"exam_name" json:"examName"`
	Email     string    `bson:"email" json:"email"`
	AppliedAt time.Time `bson:"applied_at" json:"appliedAt"`
}

type DatabaseConfig struct {
	ServerIP string `json:"server_ip"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Port     int    `json:"port"`
}

type DatabaseConfiguration struct {
	Port            string         `json:"port"`
	MongoConnection DatabaseConfig `json:"mongoConnections"`
}


var (
	TokenStore = make(map[string]string)
	mu         sync.Mutex
)

// SaveToken stores the token in memory
func SaveToken(token, email string) {
	mu.Lock()
	defer mu.Unlock()
	TokenStore[token] = email
}

// GetEmail retrieves email from token
func GetEmail(token string) (string, bool) {
	mu.Lock()
	defer mu.Unlock()
	email, exists := TokenStore[token]
	return email, exists
}
func DeleteToken(token string) {
	mu.Lock()
	defer mu.Unlock()
	// Delete token from the map
	delete(TokenStore, token)
}
