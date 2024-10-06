package dto

type Jwt struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Iat		  int64  `json:"iat"`
	ExpIn	  float64  `json:"exp_in"`
}

type RefreshToken struct {
	Token string `json:"refresh_token"`
}