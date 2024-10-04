package entities

type PersonalCredentials struct {
	ID       int `json:"id"`
	Personal *Personal `pg:"rel:has-one" json:"personal"`
	CredentialsType *CredentialsType `pg:"rel:has-one" json:"credentials_type"`
	Name    string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	AccessToken string `json:"access_token"`
	ClientID string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
