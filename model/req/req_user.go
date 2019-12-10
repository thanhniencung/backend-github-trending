package req

type ReqUpdateUser struct {
	FullName string `json:"fullName,omitempty" validate:"required"` // tags
	Email    string `json:"email,omitempty" validate:"required"`
}
