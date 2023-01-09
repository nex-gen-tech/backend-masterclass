// package main

// import "fmt"

type Namer interface {
	GetFullName() string
	SetAge(age int) error
	GetAge() int
}

type person struct {
	FirstName string
	LastName  string
	Age       int
}

func NewPerson(firstName, lastName string, age int) Namer {
	return &person{
		FirstName: firstName,
		LastName:  lastName,
	}
}

// // GetFullName returns the full name of the person
// func (p person) GetFullName() string {
// 	return p.FirstName + " " + p.LastName
// }

// // SetAge sets the age of the person
// func (p *person) SetAge(age int) error {
// 	p.Age = age
// 	return nil
// }

// // GetAge -
// func (p person) GetAge() int {
// 	return p.Age
// }

// func main() {
// 	namer := NewPerson("John", "Doe", 30)
// 	namer.SetAge(40)
// 	fmt.Println(namer.GetFullName())
// 	fmt.Println(namer.GetAge())
// }

// package main

// import "fmt"

// // Usages with Database
// // Get
// // GetAll
// // Create
// // Update
// // Delete

// // type Repository interface {
// // 	Get(id int) (interface{}, error)
// // 	GetAll() ([]interface{}, error)
// // 	Create(interface{}) error
// // 	Update(interface{}) error
// // 	Delete(id int) error
// // }

// // Token Based Authentication

// // JWT (JSON Web Token)
// // OAuth 2.0
// // Pesto Token

// type Auth interface {
// 	GenerateToken(id string) (string, error)
// 	ValidateToken(token string) (bool, error)
// }

// type JWTToken struct {
// 	Secret         string
// 	ExpirationTime int
// 	ID             string
// }
// type PestoToken struct {
// 	Secret         string
// 	ExpirationTime int
// 	ID             string
// }

// func NewJWTToken(secret string, expirationTime int, id string) Auth {
// 	return &JWTToken{
// 		Secret:         secret,
// 		ExpirationTime: expirationTime,
// 		ID:             id,
// 	}
// }

// func NewPestoToken(secret string, expirationTime int, id string) Auth {
// 	return &PestoToken{
// 		Secret:         secret,
// 		ExpirationTime: expirationTime,
// 		ID:             id,
// 	}
// }

// func (j *JWTToken) GenerateToken(id string) (string, error) {
// 	// Not Implemented
// 	return "", nil
// }

// func (j *JWTToken) ValidateToken(token string) (bool, error) {
// 	// Not Implemented
// 	return false, nil
// }

// func (j *PestoToken) GenerateToken(id string) (string, error) {
// 	// Not Implemented
// 	return "", nil
// }

// func (j *PestoToken) ValidateToken(token string) (bool, error) {
// 	// Not Implemented
// 	return false, nil
// }

// func main() {
// 	fmt.Println("Hello, World!")
// }

// fmt.Println("Hello, World!")
// x := interface{}("hello")
// str, ok := x.(string)

// if ok {
// 	fmt.Println(str)
// } else {
// 	fmt.Println("not a string")
// }

// switch x := interface{}("hello").(type) {
// case string:
// 	fmt.Printf("%s is a string", x)
// case int:
// 	fmt.Printf("%d is an int", x)
// default:
// 	fmt.Printf("%v is of another type", x)
// }

// package main

// import "encoding/json"

// type Person struct {
// 	ID   int    `json:"id"`
// 	Name string `json:"name"`
// 	Age  bool   `json:"age"`
// }

// func main() {
// 	// typ := reflect.TypeOf(Person{})
// 	// fld, _ := typ.FieldByName("Name")
// 	// tag := fld.Tag

// 	// println(tag.Get("android"))
// 	// println(tag.Get("ios"))
// 	// println(tag.Get("web"))

// 	p := Person{ID: 1, Name: "John", Age: true}
// 	b, _ := json.Marshal(p)

// 	println(string(b))
// }
