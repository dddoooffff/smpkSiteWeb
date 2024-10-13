package service


var ID int = 0

func AddUser(name, email string) func() (string, string, int) {

	return func() (string, string, int) {
		ID++
		return name, email, ID
	}
}

