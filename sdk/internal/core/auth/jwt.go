package auth

type JwtService interface{}

type jwtService struct{}

func NewJwtService() {

}

func (j *jwtService) GenerateToken(claims map[string]interface{}) error {
	return nil
}
