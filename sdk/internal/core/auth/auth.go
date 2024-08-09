package auth

// Token: Representa los tokens de acceso y/o refresh, necesarios para la autenticación y autorización.
// Role: Define los roles que un usuario puede tener dentro del sistema (e.g., administrador, usuario regular).
// Permission: Especifica los permisos que pueden asociarse a roles y usuarios, definiendo qué acciones pueden realizar dentro del sistema.
// Credential: Podría representar las credenciales de acceso, como un correo electrónico y una contraseña, o las credenciales de OAuth.
// Session: Representa una sesión de usuario activa, útil para gestionar la persistencia de la sesión y el manejo de tokens.

type Auth struct {
	Token      string
	Role       string
	Permission string
	Credential string
	Session    string
}
