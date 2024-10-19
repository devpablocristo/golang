package sdkhclnt

type genericTokenResponse struct {
	data map[string]interface{}
}

func (g *genericTokenResponse) GetAccessToken() string {
	if token, ok := g.data["access_token"].(string); ok {
		return token
	}
	return ""
}

type TokenResponse interface {
	GetAccessToken() string
}
