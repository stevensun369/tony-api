package models

type Account struct {
  ID string `json:"ID"`
  FirstName string `json:"firstName"`
  LastName string `json:"lastNamel"`
  Phone string `json:"phone"`
  WalletID string `json:"wallet"`
  Password string `json:"password"`
}

func (a Account) CreateUser() {
  
}