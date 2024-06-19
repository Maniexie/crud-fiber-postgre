package req

type MahasiswaReq struct {
	Nim   string `json:"nim" validate:"required"`
	Name  string `json:"name" validate:"required"`
	Email string `json:"email" validate:"required,email" `
}
