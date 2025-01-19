package domain

type LoginRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type Result struct {
	Username string `json:"username"`
	Nama     string `json:"nama"`
	Password string `json:"password"`
	IdRole   string `json:"id_role"`
	Token    string `json:"token"`
}

type HasilLogin struct {
	Username string `json:"username"`
	Nama     string `json:"nama"`
	Password string `json:"password"`
	IdRole   string `json:"id_role"`
}

type LoginRepository interface {
	ActLogin(username string) (*HasilLogin, error)
}
