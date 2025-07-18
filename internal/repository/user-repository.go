// ── internal/repository/user-repository.go ───────────────────────────────────
package repository

import (
	"database/sql"
	"encoding/json"
	"time"
)

// LottoStats represents statistics for lotto entries
type LottoStats struct {
	Paid   int `json:"paid"`
	Unpaid int `json:"unpaid"`
}

// GeoStats represents geographical distribution statistics
type GeoStats struct {
	Almaty    int `json:"almaty"`
	Nursultan int `json:"nursultan"`
	Shymkent  int `json:"shymkent"`
	Karaganda int `json:"karaganda"`
	Others    int `json:"others"`
}

// AdminClientEntry represents enhanced client data for admin dashboard with geolocation
type AdminClientEntry struct {
	UserID         int64     `json:"userID"`
	UserName       string    `json:"userName"`
	Fio            string    `json:"fio"`
	Contact        string    `json:"contact"`
	Address        string    `json:"address"`
	DateRegister   string    `json:"dateRegister"`
	DatePay        string    `json:"dataPay"`
	Checks         bool      `json:"checks"`
	HasGeo         bool      `json:"hasGeo"`
	Latitude       *float64  `json:"latitude"`
	Longitude      *float64  `json:"longitude"`
	AccuracyMeters *int      `json:"accuracyMeters"`
	City           *string   `json:"city"`
	Country        string    `json:"country"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// BotSession represents bot session data
type BotSession struct {
	ID           int             `json:"id"`
	UserID       int64           `json:"userID"`
	SessionID    string          `json:"sessionID"`
	State        string          `json:"state"`
	Data         json.RawMessage `json:"data,omitempty"`
	LastActivity time.Time       `json:"lastActivity"`
	ExpiresAt    *time.Time      `json:"expiresAt,omitempty"`
	CreatedAt    time.Time       `json:"createdAt"`
	UpdatedAt    time.Time       `json:"updatedAt"`
}

// AdminLog represents admin action log
type AdminLog struct {
	ID           int             `json:"id"`
	AdminUserID  int64           `json:"adminUserID"`
	Action       string          `json:"action"`
	TargetUserID *int64          `json:"targetUserID,omitempty"`
	Details      json.RawMessage `json:"details,omitempty"`
	IPAddress    *string         `json:"ipAddress,omitempty"`
	UserAgent    *string         `json:"userAgent,omitempty"`
	CreatedAt    time.Time       `json:"createdAt"`
}

// UserRepository работает со всеми таблицами: just, client, loto, geo, bot_sessions, admin_logs.
type UserRepository struct {
	db *sql.DB
}

// NewUserRepository создаёт новый UserRepository.
func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}
