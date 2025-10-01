package entitie

type personI interface {
	GetAge() int
	GetName() string
}

type speakerI interface {
	Speak() string
}

type living interface {
	personI
	speakerI
}

type Human interface {
	living
	From() string
}

type Pet interface {
	living
	GetOwner() string
}
