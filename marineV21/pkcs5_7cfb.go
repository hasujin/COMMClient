
package marine

import (
    "bytes"
    "crypto/aes"
    "crypto/cipher"
    "crypto/rand"
    "encoding/base64"
    "errors"
//    "fmt"
    "io"
    "strings"
)
/*

    Paul : You can check as below
        언제 패딩을 해야하고 언제 안해야 하나?
        1. 일반적인 방법으로, 만약 프로그래머로서 당신이 제어할 수 없는 가변 길이의 데이터를 암호화한다면, CBC 모드로 패딩해라. 모호함을 피하기위해서, 패딩을 추가하는 표준 방법을 항상 사용해라.
        2. 만약 평문이 항상 고정 길이를 가진다면, 패딩 사용을 피할 수 있다. -> 고정 길이와 블록 길이가 같다는 조건이 필요
        3. 암호화에 CFB 또는 OFB 모드가 사용되거나 RC4 등 스트림 암호 알고리즘이 사용된다면, 패딩을 할 필요가 없다.

*/




/*
func main() {
    //key := []byte("LKHlhb899Y09olUi")
    key := []byte("shosddfsdfgsdfsdfgsdfgsdfgfddddd") // ok = 32,
    text := "hello"

    encryptMsg, _ := encrypt(key, text)
    msg, err := decrypt(key, encryptMsg)
    if err != nil {
        fmt.Printf("%v\n", err)
    }
    fmt.Println(msg) // Hello World
}
*/

func addBase64Padding(value string) string {
    m := len(value) % 4
    if m != 0 {
        value += strings.Repeat("=", 4-m)
    }

    return value
}

func removeBase64Padding(value string) string {
    return strings.Replace(value, "=", "", -1)
}

func Pad(src []byte) []byte {
    padding := aes.BlockSize - len(src)%aes.BlockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(src, padtext...)
}

func Unpad(src []byte) ([]byte, error) {
    length := len(src)
    unpadding := int(src[length-1])

    if unpadding > length {
        return nil, errors.New("unpad error. This could happen when incorrect encryption key is used")
    }

    return src[:(length - unpadding)], nil
}

func CheckAndPad(key []byte) []byte {

    len_ := len(key)
     if len_%aes.BlockSize !=0 && len_ <= 32 && len_ > 0{
        key = Pad(key)
    }

    return key
}

func encryptPKCS7CFB(key []byte, text string) (string, error) {

    key = CheckAndPad(key)

    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

    msg := Pad([]byte(text))
    ciphertext := make([]byte, aes.BlockSize+len(msg))
    iv := ciphertext[:aes.BlockSize]
    if _, err := io.ReadFull(rand.Reader, iv); err != nil {
        return "", err
    }

    cfb := cipher.NewCFBEncrypter(block, iv)
    //cipher.NewCBCEncrypter(block, iv)
    cfb.XORKeyStream(ciphertext[aes.BlockSize:], []byte(msg))
    finalMsg := removeBase64Padding(base64.URLEncoding.EncodeToString(ciphertext))
    return finalMsg, nil
}

func decryptPKCS7CFB(key []byte, text string) (string, error) {

    key = CheckAndPad(key)

    block, err := aes.NewCipher(key)
    if err != nil {
        return "", err
    }

    decodedMsg, err := base64.URLEncoding.DecodeString(addBase64Padding(text))
    if err != nil {
        return "", err
    }

    if (len(decodedMsg) % aes.BlockSize) != 0 {
        return "", errors.New("blocksize must be multipe of decoded message length")
    }

    iv := decodedMsg[:aes.BlockSize]
    msg := decodedMsg[aes.BlockSize:]

    cfb := cipher.NewCFBDecrypter(block, iv)
    //cipher.NewCBCDecrypter(block, iv)
    cfb.XORKeyStream(msg, msg)

    unpadMsg, err := Unpad(msg)
    if err != nil {
        return "", err
    }

    return string(unpadMsg), nil
}

