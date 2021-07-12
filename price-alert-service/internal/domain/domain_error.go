package domain

type ErrorDomain struct {
	message string
	key     string
	detail  string
}

func NewError(message, key, detail string) error {
	return &ErrorDomain{
		message: message,
		key:     key,
		detail:  detail,
	}
}

func (d ErrorDomain) Key() string {
	return d.key
}

func (d ErrorDomain) Detail() string {
	return d.detail
}

func (d ErrorDomain) Error() string {
	return d.message
}
