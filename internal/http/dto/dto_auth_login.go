package dto

type AuthLoginPayloadReqDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthLoginReqDTO struct {
	Payload AuthLoginPayloadReqDTO `in:"body=json" json:"payload"`
}
