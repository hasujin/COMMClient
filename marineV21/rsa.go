
package marine

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"os"
	"io/ioutil"
	"io"
	"crypto/x509"
	"encoding/asn1"
	"encoding/gob"
	"encoding/pem"
	//"strings"
	


)

/************************************************

Paul RSA - Enc./Dec./Signing


실제 필드에서는 이런 기본 RSA연산 뿐만아니라 input 데이터를 인코딩하는 방식을 정한 규격이 존재

RSA PKCS#1 이라는 규격이 있다. PKCS#1 v1.5, OAEP, PSS 등 인코딩 방식에 따라 정해진 규격이 있음

인코딩에 난수값이나 해시값 등을 이용해서 원문을 더욱 유추하기 어렵게하는 방식
​
또한, 실제 필드에서는 기본 RSA 보다는 CRT를 선호하기도 한다. 이유는 속도 때문인데 CRT의 경우

속도는 빠르지만, 키 이외에도 관리해야할 것들이 많아지기 때문에 오히려 보안에 취약해 질 가능성이 높다.

그렇기 때문에 RSA보다 키 길이는 짧지만 비슷하거나 오히려 강한 보안강도를 가지는 

공개키 암호의 다른 방식인 타원 곡선을 이용한 ECC 암호가 늘어나고 있는 추세임


////////////////////////////////////////////////
	NOTICE : According to golang.org (google)
////////////////////////////////////////////////

func EncryptPKCS1v15(rand io.Reader, pub *PublicKey, msg []byte) ([]byte, error)
	EncryptPKCS1v15 encrypts the given message with RSA and the padding scheme from PKCS#1 v1.5. The message must be no longer than the length of the public modulus minus 11 bytes.
	The rand parameter is used as a source of entropy to ensure that encrypting the same message twice doesn't result in the same ciphertext.
	
	--> WARNING: Use of this function to encrypt plaintexts other than session keys is dangerous. Use RSA OAEP in new protocols.



************************************************/

/**********************************************************

	Keep the keys in the official ways with .pem and .key


************************************************************/

	//************************************************************************************
	// Generate RSA Keys
	//************************************************************************************

	// Creating RSA Key (type : 3), defualt size : 2048
	/*
		Paul.B notes : 
		////////////////////////////////////////////////////////////////////////

		type PrecomputedValues struct {
	        Dp, Dq *big.Int // D mod (P-1) (or mod Q-1)
	        Qinv   *big.Int // Q^-1 mod P

	        // CRTValues is used for the 3rd and subsequent primes. Due to a
	        // historical accident, the CRT for the first two primes is handled
	        // differently in PKCS#1 and interoperability is sufficiently
	        // important that we mirror this.
	        CRTValues []CRTValue
		}

type PrivateKey struct {
    PublicKey            // public part.
    D         *big.Int   // private exponent
    Primes    []*big.Int // prime factors of N, has >= 2 elements.

    // Precomputed contains precomputed values that speed up private
    // operations, if available.
    Precomputed rsa.PrecomputedValues
}

type PublicKey struct {
    N *big.Int // modulus
    E int      // public exponent
}
*/
////////////////////////////////////////////////////////////////////////

	

func generateKeysRSA(random io.Reader, size int) (PublicKey []byte, PrivateKey []byte, err error) {

	privateKey, err := rsa.GenerateKey(random, size)
	
	if err != nil {
		return nil, nil, err
	}

	publicKey := &privateKey.PublicKey

	return MarshalPKCS1PublicKey(publicKey), MarshalPKCS1PrivateKey(privateKey),  nil

}

func SignRSA(privateKey []byte, message []byte) (signature []byte, err error){

	//************************************************************************************
	//Encrypt Miryan Message
	//************************************************************************************

	rng := rand.Reader

	//message := []byte("the code must be like a piece of music")
	//label := []byte("")  // for only encryption
	//hash := sha256.New()


	//fmt.Printf("\n\n1) rng --> %v\n", rng)

	//ciphertext, err := rsa.EncryptOAEP(hash, rng, paulPublicKey, message, label)
	//randomSource := strings.NewReader("Mypassword")

	//Skip Encryption
	/*
	ciphertext, err := rsa.EncryptOAEP(hash, randomSource, paulPublicKey, message, label)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	*/
	//mt.Printf("OAEP encrypted \n\t[%x] to \n\t[%x]\n", string(message), ciphertext )
	

	//************************************************************************************
	// Message - Signing (Signature)
	//************************************************************************************
	var opts rsa.PSSOptions
	opts.SaltLength = rsa.PSSSaltLengthAuto // for simple example
	//PSSmessage := ciphertext
	newhash := crypto.SHA256
	
	/*
		Paul : 
			Other options can be used as below,
				Signing the string using the {md5; SHA256/512/SHA1} hash algorithm.
				Generally, SHA256 is used.

	newhash := crypto.SHA512
	newhash := crypto.MD5
	newhash := crypto.SHA1

	*/

	pssh := newhash.New()
	//pssh.Write(PSSmessage)
	pssh.Write(message)
	hashedMessage := pssh.Sum(nil)

	//fmt.Printf("\n\n2) rng --> %v\n", rng)

	//signature, err := rsa.SignPSS(rng, bewPrivateKey, newhash, hashed, &opts) //"hashed" is from the ciphertext


	privateKey_, err := ParsePKCS1PrivateKey(privateKey)

	if err != nil {
		return nil, err
	}

	return rsa.SignPSS(rng, privateKey_, newhash, hashedMessage, nil) //"hashed" is from the ciphertext
	
	// signature --> []byte (by paul)
	/*
	if err != nil {
		return nil, err
		
	}
*/
	//fmt.Printf("PSS Signature : %x\n", signature)

}


func VerifyRSA(publicKey []byte, message []byte, signature []byte) bool {


	publicKey_, err := ParsePKCS1PublicKey(publicKey)
	if err != nil {
		return false
	}

	newhash := crypto.SHA256
	pssh := newhash.New()
	pssh.Write(message)
	hashed := pssh.Sum(nil)


	//err = rsa.VerifyPSS(bewPublicKey, newhash, hashed, signature, nil)
	err = rsa.VerifyPSS(publicKey_, newhash, hashed, signature, nil)


	if err != nil {
		fmt.Printf("%v\n", err)
		return false

	} else {
		fmt.Printf("\n\n--------------------------------------\nVerify Signature successful\n--------------------------------------\n")
	}

	//************************************************************************************
	// Decrypt Message
	//************************************************************************************

	//SKipped.....

	//The label parameter must match with the value given when encrypting. - Paul


	return true


}


func EncryptOAEP(random io.Reader, publicKey []byte, message []byte, label []byte) (ciphertext []byte, err error){

	
	label_ := []byte(label)
	hash := sha256.New()

	/*
		Paul : 
			You can put "passwd" instead of a random value
			randomSource := strings.NewReader("Mypassword")

	*/
	publicKey_, err := ParsePKCS1PublicKey(publicKey)
	if err != nil {
		return nil, err
	}


	ciphertext, err = rsa.EncryptOAEP(hash, random, publicKey_, message, label_)
	if err != nil {
		return nil, err
	}

	return ciphertext, nil

}

func DecryptOAEP(privateKey []byte, ciphertext []byte, label []byte) (plaintext []byte, err error){

	label_ := []byte(label)
	hash := sha256.New()

	privateKey_, err := ParsePKCS1PrivateKey(privateKey)
	if err != nil {
		return nil, err
	}

	plaintext, err = rsa.DecryptOAEP(hash, nil, privateKey_, ciphertext, label_)

	if err != nil {
		return nil, err
	}

	return plaintext, nil
}


func derivePubFromPrivateRSA(private_key_byte []byte) ([]byte, error) {

	privateKey, err := ParsePKCS1PrivateKey(private_key_byte)

	if err != nil {
		return nil, err
	}

	publicKey := &privateKey.PublicKey

	return MarshalPKCS1PublicKey(publicKey), nil


}




func saveGobKey(fileName string, key interface{}) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	encoder := gob.NewEncoder(outFile)
	err = encoder.Encode(key)
	checkError(err)
}

func loadGobKey(fileName string, key interface{}) {
    inFile, err := os.Open(fileName)
    checkError(err)
    decoder := gob.NewDecoder(inFile)
    err = decoder.Decode(key)
    checkError(err)
    inFile.Close()
}

func savePrivateKeyInPEM(fileName string, key *rsa.PrivateKey) {
	outFile, err := os.Create(fileName)
	checkError(err)
	defer outFile.Close()

	var privateKey = &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: MarshalPKCS1PrivateKey(key),
	}

	err = pem.Encode(outFile, privateKey)
	checkError(err)
}

func savePublicKeyInPEM(fileName string, pubkey rsa.PublicKey) {
	asn1Bytes, err := asn1.Marshal(pubkey)
	checkError(err)

	var pemkey = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: asn1Bytes,
	}

	pemfile, err := os.Create(fileName)
	checkError(err)
	defer pemfile.Close()

	err = pem.Encode(pemfile, pemkey)
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}

func GetPrivateKeyFromPemFile(fileName string){
	dat, err := ioutil.ReadFile(fileName)
    checkError(err)
    


    block, _ := pem.Decode(dat)
    privatekey, err := ParsePKCS1PrivateKey(block.Bytes)


    if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   fmt.Println("\n\n************************************************************************************************************")
   fmt.Print(string(dat))
   fmt.Println("Private Key (from pem file) : ", privatekey)
   fmt.Println("Public Key (from pem file) : ", privatekey.PublicKey)
   fmt.Println("************************************************************************************************************")
   //fmt.Println("\n\nPublic Key : ", &privatekey.PublicKey)

   /*
   fmt.Println("\n---------------------------------\n")
   fmt.Printf("Public Key :  : %s\n\n", privatekey.PublicKey)

   fmt.Printf("Private Key D :  : %d\n\n", privatekey.D)

   fmt.Printf("Private Key Primes :  : %d\n\n", privatekey.Primes[0])

   fmt.Printf("Precomputed :  : %v\n\n", privatekey.Precomputed)
	*/



}

func GetPublicKeyFromPemFile(fileName string){
	dat, err := ioutil.ReadFile(fileName)
    checkError(err)
    


    block, _ := pem.Decode(dat)
    
    // X509
    // func ParsePKCS1PublicKey(der []byte) (*rsa.PublicKey, error)
    publickey, err := ParsePKCS1PublicKey(block.Bytes)


    if err != nil {
      fmt.Println(err)
      os.Exit(1)
   }

   fmt.Println("\n\n************************************************************************************************************")
   fmt.Print(string(dat))
   fmt.Println("Public Key (from file) : ", publickey)
   fmt.Println("************************************************************************************************************")
   


   /*
   fmt.Println("\n---------------------------------\n")
   fmt.Printf("Public Key :  : %s\n\n", privatekey.PublicKey)

   fmt.Printf("Private Key D :  : %d\n\n", privatekey.D)

   fmt.Printf("Private Key Primes :  : %d\n\n", privatekey.Primes[0])

   fmt.Printf("Precomputed :  : %v\n\n", privatekey.Precomputed)
	*/

}

func MarshalPKCS1PublicKey(key *rsa.PublicKey) []byte {
	return x509.MarshalPKCS1PublicKey(key)
}

func MarshalPKCS1PrivateKey(key *rsa.PrivateKey) []byte {
	return x509.MarshalPKCS1PrivateKey(key)

}


func ParsePKCS1PublicKey(pub []byte) (*rsa.PublicKey, error) {
	return x509.ParsePKCS1PublicKey(pub)
}

func ParsePKCS1PrivateKey(priv []byte) (*rsa.PrivateKey, error) {
	return x509.ParsePKCS1PrivateKey(priv)

}


