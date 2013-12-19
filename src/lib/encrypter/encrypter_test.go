package encrypter

import "testing"

func TestDoEncrypt(t *testing.T){
        username := "admin"
        password := "admin"
        result := DoEncrypt(username, password)
        t.Log("Result: ", result)
}
