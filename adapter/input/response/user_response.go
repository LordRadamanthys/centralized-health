package response

type UserResponse struct {
	Name   string   `copier:"Name" json:"name"`
	Email  string   `copier:"Email" json:"email"`
	Birth  string   `copier:"Birth" json:"birth"`
	Phones []string `copier:"Phone" json:"phones"`
}
