package deleteuserhandler

type DeleteUserOptions struct {
	Email   string `json:"email"`
	Idtoken string `json:"idtoken"`
}
