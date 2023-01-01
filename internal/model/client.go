package model

type Client struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Lastname string `json:"lastname"`
	Weight   uint16 `json:"weight"`
	Age      uint8  `json:"age"`
	Height   uint8  `json:"height"`
}

type GetClientRequest struct {
	ID string `json:"id"`
}

type CreateClientResponse struct {
	ID string `json:"id"`
}

type DeleteClientResponse struct {
	ID string `json:"id"`
}
