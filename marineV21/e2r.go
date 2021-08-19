
package marine

import (
	
	"fmt"
	"io"
	"math/big"
	//"os"
	//"reflect"
	cryptorand "crypto/rand"
	//"encoding/base64"
    //"crypto/sha256"
    "encoding/json"
    //color "github.com/mitchellh/colorstring"
    "errors"
)



func generateKey(random io.Reader, type_ int) (publicKey []byte, privateKey []byte, err error) {

	if random == nil {
		random = cryptorand.Reader
	}


	switch type_ {

		case 1: //EW
			//fmt.Printf("Generating........Edward Keys\n")
			publicKey, privateKey, err = generateKeyEW(random)

		case 2: //
			//fmt.Printf("Generating........ECDSA(ECC) Keys\n")
			publicKey, privateKey, err = generateKeysECC("P-256", random)

		case 3:
			//fmt.Printf("Generating........RSA Keys\n")
			publicKey, privateKey, err = generateKeysRSA(random, 2048)

		default :
			//fmt.Printf("generateKey() : Not supported type, plz check again\n")
			return nil, nil, errors.New("Not supported type, plz check again\n") 

	}


	return publicKey, privateKey, err


}


func Sign(privateKey []byte,message []byte) (signature []byte, err error) {


	type_ := detectKeyTypeFromPrivKey(privateKey)

	//fmt.Printf("Sign() : Key Type  --> %d\n", type_)

	switch type_ {

		case 1: //EW
			signature, err = SignEW(privateKey, message), nil

		case 2: //
			signature, err = SignECC(privateKey, message)

		case 3:
			//signature, err := rsa.SignPSS(rng, privateKey, newhash, hashed, nil)
			signature, err = SignRSA(privateKey, message)			

		default :
			//fmt.Printf("Sign() : Not supported type, plz check again\n")
			return nil, errors.New("Not supported type, plz check again\n") 

	}

	return signature, err

}

func Verify(publicKey []byte, message []byte, signature []byte) bool {

	type_ := detectKeyTypeFromPubKey(publicKey)
	
	//fmt.Printf("Verify() : Key Type --> %d\n", type_)

	var result bool

	switch type_ {

		case 1: //EW
			result = VerifyEW(publicKey, message, signature)

		case 2: //ECC
			result = VerifyECC(publicKey, message, signature)

		case 3: // RSA
			//signature, err := rsa.SignPSS(rng, privateKey, newhash, hashed, nil)
			result = VerifyRSA(publicKey, message, signature)			

		default :
			fmt.Printf("Verify() : Not supported type [%d], plz check again\n", type_)
			return false 

	}

	return result


}


func VerifyWithRS(publicKey []byte, message []byte, r *big.Int, s *big.Int) bool {

	type_ := detectKeyTypeFromPubKey(publicKey)
	
	//fmt.Printf("Verify() : Key Type --> %d\n", type_)


	if type_ != 2 {
		fmt.Printf("VerifyWithRS() : only ECDSA publick key type is supported, but the input was [%d] (1:RSA, 2:ECDSA, 3:Edward25519)\n", type_)
		return false

	}

	signature_ := signature{
        R: r.String(),
        S: s.String(),
    }

    signatureByte, err := json.Marshal(signature_)
    if err != nil {
        fmt.Printf("Marshal Error : %v\n", err)
        return false
    }

	result := VerifyECC(publicKey, message, signatureByte)	


	return result


}


func VerifyWithType(publicKey []byte, message []byte, signature []byte, type_ int) bool {

	if type_ < 1 || type_ >3 {

		type_ = detectKeyTypeFromPubKey(publicKey)

	}


	//fmt.Printf("Verify() : Key Type --> %d\n", type_)

	var result bool

	switch type_ {

		case 1: //EW
			result = VerifyEW(publicKey, message, signature)

		case 2: //ECC
			result = VerifyECC(publicKey, message, signature)

		case 3: // RSA
			//signature, err := rsa.SignPSS(rng, privateKey, newhash, hashed, nil)
			result = VerifyRSA(publicKey, message, signature)			

		default :
			fmt.Printf("VerifyWithType() : Not supported type [%d], plz check again\n", type_)
			return false 

	}

	return result


}



func derivePubFromPrivate(privateKey []byte)(publicKey []byte, err error) {

	//func derivePubFromPrivateECC(private_key_byte []byte) ([]byte, error)
	//func derivePubFromPrivateRSA(private_key_byte []byte) ([]byte, error)
	//func derivePubFromPrivateEW(givenprivateKey PrivateKey) (derivedprivateKey string,err error)

	type_ := detectKeyTypeFromPrivKey(privateKey)
	err = nil
	switch type_ {

		case 1: //EW
			publicKey, err = derivePubFromPrivateEW(privateKey)

		case 2: //
			publicKey, err = derivePubFromPrivateECC(privateKey)

		case 3:
			publicKey, err = derivePubFromPrivateRSA(privateKey)			

		default :
			//fmt.Printf("Not supported type, plz check again\n")
			return nil, errors.New("Not supported type, plz check again\n") 

	}

	return publicKey, err


}


func derivePubFromPrivateWithType(privateKey []byte, type_ int) (publicKey []byte, err error) {
	err = nil
	switch type_ {

		case 1: //EW
			publicKey, err = derivePubFromPrivateEW(privateKey)

		case 2: //
			publicKey, err = derivePubFromPrivateECC(privateKey)

		case 3:
			publicKey, err = derivePubFromPrivateRSA(privateKey)			

		default :
			//fmt.Printf("Not supported type, plz check again\n")
			return nil, errors.New("Not supported type, plz check again\n") 

	}

	return publicKey, err


}

