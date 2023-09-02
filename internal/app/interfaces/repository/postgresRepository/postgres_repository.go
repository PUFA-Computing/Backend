package postgresRepository

import (
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"time"
)

type SessionData struct {
	UserID       uuid.UUID `gorm:"PRIMARY_KEY"`
	SessionToken string    `gorm:"unique"`
	ExpiredAt    time.Time
}

type PostgresRepository struct {
	DB *gorm.DB
}

func NewPostgresRepository() (*PostgresRepository, error) {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	dsn := "host=" + dbHost + " port=" + dbPort + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&SessionData{}); err != nil {
		return nil, err
	}
	return &PostgresRepository{DB: db}, nil
}

func (r *PostgresRepository) Close() {
	db, _ := r.DB.DB()
	err := db.Close()
	if err != nil {
		return
	}
}

func (r *PostgresRepository) StoreSessionData(userID uuid.UUID, sessionToken string, expirationTime time.Time) error {
	sessionData := SessionData{
		UserID:       userID,
		SessionToken: sessionToken,
		ExpiredAt:    expirationTime,
	}
	return r.DB.Create(&sessionData).Error
}

func (r *PostgresRepository) GetSessionData(userID uuid.UUID) (*SessionData, error) {
	var sessionData SessionData
	err := r.DB.Where("user_id = ?", userID).First(&sessionData).Error
	if err != nil {
		return nil, err
	}
	return &sessionData, nil
}
