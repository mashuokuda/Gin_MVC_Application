package user

type HeaderUser struct {
	Name     string
	Location uint32
	Profile  string
	Image    string
}

func (u User) GetHeaderUser() HeaderUser {
	var l uint32
	if u.Publish {
		l = u.Location
	} else {
		l = 50
	}
	return HeaderUser{
		Name:     u.GetUserName(),
		Location: l,
		Profile:  u.Profile,
		Image:    u.Image.GetImage(),
	}
}
