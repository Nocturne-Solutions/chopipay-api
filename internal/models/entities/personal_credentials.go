package entities

type PersonalCredentials struct {
	ID       int `json:"id"`
	Personal *Personal `json:"personal" pg:"rel:has-one, fk:personal_id, join_fk:personal_id"`
	PersonalID int `json:"personal_id" pg:"personal_id"`
	CredentialsType *CredentialsType `pg:"rel:has-one" json:"credentials_type"`
	CredentialsTypeID int `json:"credentials_type_id" pg:"credentials_type_id"`
	Name    string `json:"name"`
	Username string `json:"username"`
	Password string `json:"password"`
	AccessToken string `json:"access_token"`
	ClientID string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}
