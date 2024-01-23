package sessions

import "time"

type Session struct {
	ID          string    `json:"id"`
	UserID      string    `json:"user_id"`
	ExpiredAt   time.Time `json:"expired_at"`
	XSRFToken   string    `json:"xsrf_token"`
	MobileToken string    `json:"mobile_token"`
}

func NewSessionModel(id string, userID string, expiredAt time.Time, xsrfToken string, mobileToken string) *Session {
	return &Session{
		ID:          id,
		UserID:      userID,
		ExpiredAt:   expiredAt,
		XSRFToken:   xsrfToken,
		MobileToken: mobileToken,
	}
}

func (s *Session) ToEntity() *SessionEntity {
	return &SessionEntity{
		ID:          s.ID,
		UserID:      s.UserID,
		ExpiredAt:   s.ExpiredAt,
		XSRFToken:   s.XSRFToken,
		MobileToken: s.MobileToken,
	}
}

func ToSessionModel(se *SessionEntity) *Session {
	return &Session{
		ID:          se.ID,
		UserID:      se.UserID,
		ExpiredAt:   se.ExpiredAt,
		XSRFToken:   se.XSRFToken,
		MobileToken: se.MobileToken,
	}
}
