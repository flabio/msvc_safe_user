package dto

type UserDTO struct {
	Id           uint   `json:"id"`
	FirstName    string `json:"first_name"`
	FirstSurName string `json:"first_sur_name"`
	SeconSurName string `json:"secon_sur_name"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	ZipCode      string `json:"zip_code"`
	StateId      uint   `json:"state_id"`
	RolId        uint   `json:"rol_id"`
	Password     string `json:"password"`
	Active       bool   `json:"active"`
}

type UserResponseDTO struct {
	Id           uint   `json:"id"`
	FirstName    string `json:"first_name"`
	FirstSurName string `json:"first_sur_name"`
	SeconSurName string `json:"secon_sur_name"`
	Email        string `json:"email"`
	Address      string `json:"address"`
	Phone        string `json:"phone"`
	ZipCode      string `json:"zip_code"`
	StateId      uint   `json:"state_id"`
	RolId        uint   `json:"rol_id"`
	RolName      string `json:"rol_name"`
	StateName    string `json:"state_name"`
	Active       bool   `json:"active"`
}

type LoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
