package header

import "Gin_MVC/model/user"

func GetHeaderUser(usr *user.User) *user.HeaderUser {
	var headerUser user.HeaderUser
	if usr == nil {
		headerUser = user.HeaderUser{}
	} else {
		headerUser = usr.GetHeaderUser()
	}
	return &headerUser
}
