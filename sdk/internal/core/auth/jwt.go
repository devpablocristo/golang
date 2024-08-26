package coreauth

type JwtService any

type jwtService struct{}

func NewJwtService() {

}

func (j *jwtService) GenerateToken(claims map[string]any) error {
	return nil
}
