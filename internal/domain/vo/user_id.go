package vo

type UserId Vo[string]

func NewUserId(value string) UserId {
	return &voImpl[string]{
		value: value,
	}
}
