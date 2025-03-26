package db

import "time"

func (s *Session) IsExpired() bool {
	return s.ExpiresAt.Before(time.Now())
}
