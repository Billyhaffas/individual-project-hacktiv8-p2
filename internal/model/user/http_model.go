package user

type UserRespon struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type RegisterRespon struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type RegisterRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginRespon struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Token   string `json:"token"`
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetMe struct {
	Id            int     `json:"id"`
	Name          string  `json:"name"`
	Email         string  `json:"email"`
	DepositAmount float32 `json:"saldo"`
}

type GetMeRespon struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    *GetMe `json:"data"`
}

type DeleteRespon struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

type TopUpSaldo struct {
	TopUpAmount float32 `json:"top_up_amount"`
}
