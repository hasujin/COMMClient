package marine

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/x509"
	"io"
	"math/big"
	"encoding/json"
	"fmt"
)


func generateKeysECC(curveName string, random io.Reader) (public_key_bytes []byte, private_key_bytes []byte,  err error) {

	var curve elliptic.Curve

	switch curveName {
	case "P-256":
		curve = elliptic.P256() //default
	case "P-384":
		curve = elliptic.P384()
	case "P-521":
		curve = elliptic.P521()
	}

	private_key, err := ecdsa.GenerateKey(curve, random)
	if err!= nil{
		return nil, nil, err
	}
	//The prime # size could be {192, 224, 256, 384, 512}. Definitely, the longer means stronger
	private_key_bytes, err = x509.MarshalECPrivateKey(private_key)
	if err!= nil{
		return nil,nil, err
	}

	public_key_bytes, err = x509.MarshalPKIXPublicKey(&private_key.PublicKey)
	if err!= nil{
		return nil,nil, err
	}
	return public_key_bytes, private_key_bytes,  nil
}

func derivePubFromPrivateECC(private_key_byte []byte) ([]byte, error) {

	//func ParseECPrivateKey(der []byte) (*ecdsa.PrivateKey, error)

	private_key, err := x509.ParseECPrivateKey(private_key_byte)

	if err != nil {
		return nil, err
	}

	public_key_bytes, err := x509.MarshalPKIXPublicKey(&private_key.PublicKey)
	if err != nil {
		return nil, err
	}

	return public_key_bytes, nil

}

type signature struct {
	R 		  string `json:"r"`
	S 		  string `json:"s"`
}

func SignECC(private_key_bytes []byte, hash []byte) (signautre []byte, err error) {

	var zero []byte

	//SignECC(message, []byte(back2_privatekey))
	r, s, err := SignECC_(hash, private_key_bytes)
	
	if err != nil {
		return zero, err
	}

	fmt.Printf("R : %v\nS : %v\n", r, s)


	var signature_ signature

	signature_ = signature{
		R: r.String(),
		S: s.String(),
	}

	fmt.Printf("Signature : %v\n", signature_)

	marshaledSignature, err := json.Marshal(signature_)
	if err != nil {
		return zero, err
	}

	
	return marshaledSignature, nil

}


func SignECC_(hash []byte, private_key_bytes []byte) (r, s *big.Int, err error) {
	zero := big.NewInt(0)
	private_key, err := x509.ParseECPrivateKey(private_key_bytes)
	if err != nil {
		return zero, zero, err
	}

	r, s, err = ecdsa.Sign(rand.Reader, private_key, hash)
	if err != nil {
		return zero, zero, err
	}
	return r, s, nil
}

func VerifyECC(public_key_bytes []byte, hash []byte, signautre []byte) (result bool) {


	var signature_ signature
	err := json.Unmarshal(signautre, &signature_)
    if err != nil {
        //fmt.Printf("VerifyECC() : Unmarshal return error: %v", err)
        return false
    }

	var r_, _ = new(big.Int).SetString(signature_.R, 10)
	var s_, _ = new(big.Int).SetString(signature_.S, 10)

	return VerifyECC_(hash, public_key_bytes, r_, s_) 
}

func VerifyECC_(hash []byte, public_key_bytes []byte, r *big.Int, s *big.Int) (result bool) {
	public_key, err := x509.ParsePKIXPublicKey(public_key_bytes)
	if err != nil {
		return false
	}

	switch public_key := public_key.(type) {
		case *ecdsa.PublicKey:
			return ecdsa.Verify(public_key, hash, r, s)
		default:
			fmt.Printf("The given public key is not ECDSA,..plz check again\n")
			return false
	}
}


func generateAddressECC(public []byte) string {


	/*
	Paul.B
		Issue AccountID from Public Key
	*/

	accountID := issueAccountID(public)
	//fmt.Printf("Issued Account ID : %x (size:%d)\n", accountID, len(accountID))

	/*
	Paul.B
		Set prefix with "z", supported by HJ
	*/
	
	//accountID[0] = 0x00
	accountID[0] = 0x00

	//accountID[1] = 0x01
	//accountID[1] = 0x17 //seconde for 'E'
	//accountID[1] = 0x0e //seconde for 'e'

	//encoded58AccountID := Encode(accountID)
	//log.Println("Encoded58AccountID : ", encoded58AccountID)

	/*
	Paul.B
		Checksum
	*/

	checkSum := []byte(doubleCheckSum4(accountID[0:19]))

	//log.Printf("checksum  value (first 4buytes) --> %x", checkSum[0:3])

	/*
		Finally Address composed of {accountID, checksum}; the first 4 bytes comes from the checksum calcualted from accountID
	*/

	var rawAddress [24]byte
	//accountID_ := accountID[0:19]
	//checkSum_ := checkSum[0:4]
	copy(rawAddress[0:19], accountID[0:19])
	copy(rawAddress[20:23], checkSum[0:3])

	//spaceAddr := Encode(rawAddress[:])

	spaceAddr := Encoding58withType(rawAddress[:], 2) // type(2) indicates "ECC/ECDSA"
	
	//log.Printf ("spaceAddr (raw)    : %x", rawAddress)
	//log.Println("spaceAddr (base58) : ", spaceAddr)

	return spaceAddr
}





