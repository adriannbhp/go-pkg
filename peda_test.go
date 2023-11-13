package peda

import (
	"fmt"
	"testing"

	"github.com/aiteung/atdb"
	"github.com/whatsauth/watoken"
	"go.mongodb.org/mongo-driver/bson"
)

func TestCreateNewUserRole(t *testing.T) {
	var userdata User
	userdata.Username = "adrian"
	userdata.Password = "bener"
	userdata.Role = "admin"
	mconn := SetConnection("MONGOSTRING", "gisajah")
	CreateNewUserRole(mconn, "users", userdata)
}

func TestDeleteUser(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "petapedia")
	var userdata User
	userdata.Username = "yyy"
	DeleteUser(mconn, "user", userdata)
}

func CreateNewUserToken(t *testing.T) {
	var userdata User
	userdata.Username = "dzikri"
	userdata.Password = "dzikri"
	userdata.Role = "admin"

	// Create a MongoDB connection
	mconn := SetConnection("MONGOSTRING", "petapedia")

	// Call the function to create a user and generate a token
	err := CreateUserAndAddToken("your_private_key_env", mconn, "user", userdata)

	if err != nil {
		t.Errorf("Error creating user and token: %v", err)
	}
}

func TestGFCPostHandlerUser(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "petapedia")
	var userdata User
	userdata.Username = "dzikri"
	userdata.Password = "dzikri"
	userdata.Role = "admin"
	CreateNewUserRole(mconn, "user", userdata)
}

func TestProduct(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "petapedia")
	var productdata Product
	productdata.Nomorid = 1
	productdata.Name = "dzikri"
	productdata.Description = "haq"
	productdata.Price = 1000
	productdata.Size = "XL"
	productdata.Stock = 100
	productdata.Image = "https://images3.alphacoders.com/165/thumb-1920-165265.jpg"
	CreateNewProduct(mconn, "product", productdata)
}

func TestAllProduct(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "petapedia")
	product := GetAllProduct(mconn, "product")
	fmt.Println(product)
}

// func TestUpdateGetData(t *testing.T) {
// 	mconn := SetConnection("MONGOULBI", "petapedia")
// 	datagedung := GetAllBangunanLineString(mconn, "petapedia")
// 	fmt.Println(datagedung)
// }

func TestGeneratePasswordHash(t *testing.T) {
	password := "adrian"
	hash, _ := HashPassword(password) // ignore error for the sake of simplicity

	fmt.Println("Password:", password)
	fmt.Println("Hash:    ", hash)

	match := CheckPasswordHash(password, hash)
	fmt.Println("Match:   ", match)
}
func TestGeneratePrivateKeyPaseto(t *testing.T) {
	privateKey, publicKey := watoken.GenerateKey()
	fmt.Println(privateKey)
	fmt.Println(publicKey)
	hasil, err := watoken.Encode("adrian", privateKey)
	fmt.Println(hasil, err)
}

func TestHashFunction(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "gisajah")
	var userdata User
	userdata.Username = "adrian"
	userdata.Password = "adrian"

	filter := bson.M{"username": userdata.Username}
	res := atdb.GetOneDoc[User](mconn, "users", filter)
	fmt.Println("Mongo User Result: ", res)
	hash, _ := HashPassword(userdata.Password)
	fmt.Println("Hash Password : ", hash)
	match := CheckPasswordHash(userdata.Password, res.Password)
	fmt.Println("Match:   ", match)

}

func TestIsPasswordValid(t *testing.T) {
	mconn := SetConnection("MONGOULBI", "petapedia")
	var userdata User
	userdata.Username = "petped"
	userdata.Password = "secret"

	anu := IsPasswordValid(mconn, "user", userdata)
	fmt.Println(anu)
}

func CreateContent(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "petapedia")
	var contentdata Content
	contentdata.ID = 1
	contentdata.Content = "dzikri"
	contentdata.Description = "haq"
	contentdata.Image = "https://images3.alphacoders.com/165/thumb-1920-165265.jpg"
	CreateNewContent(mconn, "content", contentdata)
}

func TestUserFix(t *testing.T) {
	mconn := SetConnection("MONGOSTRING", "petapedia")
	var userdata User
	userdata.Username = "dzikri"
	userdata.Password = "dzikri"
	userdata.Role = "admin"
	CreateUser(mconn, "user", userdata)
}
