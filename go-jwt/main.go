package main

import (
    "io/ioutil"
    "log"
    "net/http"
    "encoding/json"
    "fmt"
    jwt "github.com/dgrijalva/jwt-go"
    "time"
)

// using asymmetric crypto/RSA keys
// location of the used for signing and verification
const (
    privateKeyPath = "keys/app.rsa"         // openssl genrsa -out app.rsa 1024
    pubKeyPath  = "keys/app.rsa.pub"        // openssl rsa -in app.rsa -pubout > app.rsa.pub
)

// verify key and sign key
var (
    verifyKey, signKey []byte
)
// struct User for parsing login credentials
type User struct {
    UserName string `json:"username"`
    Password string `json:"password"`
}

// read the key files before starting http handlers
func init()  {
    var err error

    signKey, err = ioutil.ReadFile(privateKeyPath)
    if err != nil {
        log.Fatal("Error reading private key")
        return
    }

    verifyKey, err = ioutil.ReadFile(pubKeyPath)
    if err != nil {
        log.Fatal("Error reading private key")
        return
    }
}

// reads the loging credentials, checks them and creates JWT the token
func loginHandler(w http.ResponseWriter, r http.Request)  {
    var user User
    // decode into User struct
    err := json.NewDecoder(r.Body).Decode(&user)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintln(w, "Error in request body")
        return
    }
    // validate user credentials
    if user.UserName != "shijuvar" && user.Password != "pass" {
        w.WriteHeader(http.StatusForbidden)
        fmt.Fprintln(w, "Wrong info")
        return
    }

    // create a signer for rsa 256
    t := jwt.New(jwt.GetSigningMethod("RS256"))

    // set our claims
    t.Claims["iss"] = "admin"
    t.Claims["CustomUserInfo"] = struct {
        Name string
        Role string
    }{user.UserName, "Member"}

    // set the expire time
    t.Claims["exp"] = time.Now().Add(time.Minute * 20).Unix()
    tokenString, err := t.SignedString(signKey)
    if err != nil {
        w.WriteHeader(http.StatusInternalServerError)
        fmt.Fprintln(w, "Sorry, error while Signing Token!")
        log.Printf("Token Signing error:%v\n", err)
        return
    }
    response := Token{tokenString}
    jsonResponse(response, w)
}

// only accessible with a valid token
func authHandler(w http.ResponseWriter, r *http.Request)  {
    // validate the token
    token, err := jwt.Parsef
}
func main()  {

}
