package dto

type LoginResult struct {
	ID       int    `json:"id"`
	Avatar   string `json:"avatar"`
	Account  string `json:"account"`
	Nickname string `json:"nickname,omitempty"`
	Mobile   string `json:"mobile"`
	Token    string `json:"token"`
}
