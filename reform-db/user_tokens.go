package main

//go:generate reform

// UserTokens represents a row in user_tokens table.
//reform:user_tokens
type UserTokens struct {
	User            int32  `reform:"user,pk" json:"user"`
	MainToken       string `reform:"main_token" json:"main_token"`
	RefreshToken    string `reform:"refresh_token" json:"refresh_token"`
	TokenKey        []byte `reform:"token_key" json:"token_key"`
	RefreshtokenKey []byte `reform:"refreshtoken_key" json:"refreshtoken_key"`
}
