package account

type IAccountController interface {
	LoginOut(session string) error
}
