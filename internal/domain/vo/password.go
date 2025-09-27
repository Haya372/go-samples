package vo

type Password Vo[string]

func NewPassword(value string) Password {
	return &voImpl[string]{
		value: value,
	}
}
