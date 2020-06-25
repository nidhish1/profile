package entity

type Person struct {
	FirstName string `json:"firstname" binding:"required"`
	LastName  string `json:"lastname" binding:"required"`
	Age       string `json:"age" binding:"gte=1,lte=1"`
	Email     string `json:"email" validate: "required,email"`
}

type Video struct {
	Title       string `json:"title" binding:"min=2,max=10" validate:"is-cool"`
	Description string `json:desc`
	URL         string `json:url binding:"required,url"`
	Author      Person `json:"author" binding:"required"`
}
