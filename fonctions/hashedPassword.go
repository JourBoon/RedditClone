package fonction_go

import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
)

//Hash

func hashedPassword(password string) []byte {

    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        fmt.Println(err)
    }

    return hashedPassword;
}

//Check

func checkPassword(hashedPassword, passwordLogin string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(passwordLogin))
    return err == nil
}

func checkHashedPassword(storedPass string, password string) {
    if checkPassword(storedPass, password) {
        fmt.Println("The password is correct")
    } else {
        fmt.Println("The password is incorrect")
    }
}
