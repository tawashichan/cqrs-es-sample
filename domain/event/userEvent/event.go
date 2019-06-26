package userEvent

// eventにはバージョン番号を割り当て,楽観ロックをしている
type UpdateEmail struct {
	Email   string
	Version uint
}

type UpdateName struct {
	Name    string
	Version uint
}
