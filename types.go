package client

import "time"

const (
	DefaultOAuthURL      = "https://ngw.devices.sberbank.ru:9443/api/v2/oauth"
	DefaultRefreshMargin = 1 * time.Minute
	DefaultMinRefreshInt = 30 * time.Second
	DefaultTimeout       = 30 * time.Second
)

// Token represents OAuth token
type Token struct {
	Value     string
	ExpiresAt time.Time
}

// IsValid checks if token is valid with given margin
func (t *Token) IsValid(margin time.Duration) bool {
	if t == nil || t.Value == "" {
		return false
	}
	return time.Now().Add(margin).Before(t.ExpiresAt)
}

// TokenResponse represents OAuth token response
type TokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type Scope string

const (
	ScopeSaluteSpeechPers Scope = "SALUTE_SPEECH_PERS"
	ScopeSaluteSpeechCorp Scope = "SALUTE_SPEECH_CORP"
	ScopeSaluteSpeechB2B  Scope = "SALUTE_SPEECH_B2B"
	ScopeSberSpeech       Scope = "SBER_SPEECH"
	ScopeGigaChatAPIPers  Scope = "GIGACHAT_API_PERS"
)
