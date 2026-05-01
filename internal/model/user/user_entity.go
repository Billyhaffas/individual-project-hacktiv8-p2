package user

type User struct {
	UserId        int
	Name          string
	Email         string
	Password      string
	DepositAmount float32
	JwtToken      string
}
