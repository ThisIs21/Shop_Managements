package dto
// LoginRequest represents the request body for user login
type LoginRequest struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=6"`
}

// UpdateUserRequest represents the request body for updating a user
type UpdateUserRequest struct {
    Name     string `json:"name" validate:"omitempty"`         // Opsional
    Email    string `json:"email" validate:"omitempty,email"`  // Opsional, harus unik
    Password string `json:"password" validate:"omitempty,min=6"` // Opsional
    Role     string `json:"role" validate:"required,oneof=OWNER KASIR PEMBELIAN GUDANG KEPALA_GUDANG"` // Nama role
    Status   string `json:"status" validate:"required,oneof=active inactive"` // Hanya "active" atau "inactive"
}

// CreateUserRequest represents the request body for creating a new user
type CreateUserRequest struct {
    Name     string `json:"name" validate:"required"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required,min=6"`
    RoleID   uint   `json:"role_id" validate:"required"` // ID role
    Status   string `json:"status" validate:"required,oneof=active inactive"`
}

// LoginResponse represents the response body for user login
type LoginResponse struct {
    Token string `json:"token"`
    Name  string `json:"name"`
    Role  string `json:"role"`
}

