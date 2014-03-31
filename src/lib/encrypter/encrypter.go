/*
Encapsulation of password algorithm.
*/

package encrypter

import "crypto/md5"
import "crypto/sha1"
import "encoding/hex"

//Return encrypted string.
func DoEncrypt(username string, password string) string {
	md5enc := md5.New()
	md5enc.Write([]byte(password))

	bytesUsername := []byte(username)

	bytesToSha1 := append(bytesUsername, md5enc.Sum([]byte(""))...)

	sha1enc := sha1.New()
	sha1enc.Write(bytesToSha1)
	result := sha1enc.Sum([]byte(""))
	return hex.EncodeToString(result)
}
