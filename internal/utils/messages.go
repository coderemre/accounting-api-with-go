package utils

type Message string

const (
	ErrInternalServer       Message = "Internal server error"
	ErrInvalidRequest       Message = "Invalid request data"
	ErrUnauthorized         Message = "Unauthorized"
	ErrForbidden            Message = "Forbidden"
	ErrNotFound             Message = "Resource not found"
	ErrDuplicateEntry       Message = "Duplicate entry"
	ErrEnvFileNotFound Message = "No .env file found, using system environment variables"

	ErrUsernameRequired    Message = "Username cannot be empty"
	ErrEmailRequired       Message = "Email cannot be empty"
	ErrInvalidEmailFormat  Message = "Invalid email format"
	ErrPasswordRequired    Message = "Password cannot be empty"
	ErrPasswordTooShort    Message = "Password must be at least 6 characters long"
	ErrInvalidToken		   Message = "Invalid token"

	ErrInvalidCredentials   Message = "Invalid email or password"
	ErrPasswordHashFailed   Message = "Failed to hash password"
	ErrTokenGeneration      Message = "Token generation failed"
	ErrResponseEncodingFailed Message = "Failed to encode response"
	ErrUserRetrievalFailed  Message = "Failed to retrieve user data"
	ErrDatabaseCloseFailed  Message = "Error closing database connection"
	ErrServerListenFailed   Message = "Server listen and serve failed"
	ErrServerShutdownFailed Message = "Server forced to shutdown"

	ErrDatabaseConnectionFailed Message = "Failed to connect to the database"
	ErrDatabasePingFailed       Message = "Failed to ping the database"
	ErrUserAlreadyExists        Message = "User with this email already exists"
	ErrUsernameExists           Message = "Username already exists"
	

	SuccessUserRegistered       Message = "User registered successfully"
	SuccessLogin                Message = "Login successful"
	SuccessTokenValidated       Message = "Token validated successfully"
	SuccessLoggerInitialized    Message = "Logger initialized successfully"
	SuccessServerRunning        Message = "Server running on port 4520"
	SuccessServerShutdown       Message = "Shutting down server..."
	SuccessServerExited         Message = "Server exited gracefully"
	SuccessConfigLoaded         Message = "Configuration loaded successfully"
	SuccessDatabaseConnected    Message = "Database connected successfully"
	SuccessDatabaseDisconnected Message = "Database disconnected successfully"
)

func (m Message) String() string {
	return string(m)
}