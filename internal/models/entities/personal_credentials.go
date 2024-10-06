package entities

type PersonalCredentials struct {
	ID       int `json:"id"`
	Personal *Personal `json:"personal" pg:"rel:has-one, fk:personal_id, join_fk:personal_id"`
	PersonalID int `json:"personal_id" pg:"personal_id"`
	CredentialsType *CredentialsType `json:"credentials_type" pg:"rel:has-one, fk:credentials_type_id, join_fk:credentials_type_id"`
	CredentialsTypeID int `json:"credentials_type_id" pg:"credentials_type_id, credentials_type_id"`
	Name    string `json:"name" pg:"name"`
	Username string `json:"username" pg:"username"`
	Password string `json:"password" pg:"password"`
	AccessToken string `json:"access_token" pg:"access_token"`
	ClientID string `json:"client_id" pg:"client_id"`
	ClientSecret string `json:"client_secret" pg:"client_secret"`
}
