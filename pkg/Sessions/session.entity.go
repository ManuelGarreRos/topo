package sessions

import "time"

const (
	SessionTable       = "users_sessions"
	SessionID          = "id"
	SessionUserID      = "user_id"
	SessionExpiredAt   = "expired_at"
	SessionXSRFToken   = "xsrf_token"
	SessionMobileToken = "mobile_token"
)

type SessionEntity struct {
	ID          string    `gorm:"primaryKey"`
	UserID      string    `gorm:"not null"`
	ExpiredAt   time.Time `gorm:"not null"`
	XSRFToken   string
	MobileToken string
}
