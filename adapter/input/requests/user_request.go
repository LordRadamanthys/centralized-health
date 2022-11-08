package requests

type UserRequest struct {
	Name     string   `copier:"Name" json:"name"`
	Password string   `copier:"Password" json:"password"`
	Email    string   `copier:"Email" json:"email"`
	Birth    string   `copier:"Birth" json:"birth"`
	Phones   []string `copier:"Phone" json:"phones"`
}
