package main

type UserInfo struct{
	Name 		string
	Age 		int
	BDate 		int
	Login 		string 
	Password 	string 
}

type User struct{
	m map[string]UserInfo
}

func NewUserInfo(name string, age, bdate int, login, password string) UserInfo{
	return User{
		Name: name,
		Age: age,
		BDate: bdate,
		Login: login,
		Password: password,
	}
}

func NewUser(name string, ui UserInfo) *User{
	return &User{
		m: make(map[string]UserInfo),
	}
}
func main(){
	UserList := make([]User, 1)
	UserList[0] = *NewUser("Victor", )
}