package domain

type DocumentsDomain struct {
	Title       string `copier:"Title" json:"title"`
	URLDocument string `copier:"URLDocument" json:"url_document"`
}
