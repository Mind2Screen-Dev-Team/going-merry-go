package auth_dto

type AuthLoginDTO struct {
	Payload AuthLoginJsonDTO `in:"body=json"`
}

type AuthLoginJsonDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
