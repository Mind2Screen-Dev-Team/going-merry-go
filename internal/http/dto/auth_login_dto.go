package auth_dto

type AuthLoginPayloadDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthLoginDTO struct {
	Payload AuthLoginPayloadDTO `in:"body=json"`
}
