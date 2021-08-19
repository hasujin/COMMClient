
package marine

import(
	"log"
	//"reflect"
	//"strings"
	"bufio"
	"os"
	//"encoding/hex"
	"encoding/base64"
	"time"
	"crypto/sha256"
	"golang.org/x/crypto/ripemd160"
	"math/big"
	"fmt"
	"errors"
	"crypto/rand"
	"bytes"
	"crypto/x509"

	
	/*
		Please Do NOT Use this Color library,
		when you put this into chaincode or actual running-machine/module(s).

		Do not want? --> replace {color.} with {fmt. or log.} 
		(Paul.B)
	*/
	//color "github.com/mitchellh/colorstring"


)



func Base64ToBigInt(s string) (*big.Int, error) {
    data, err := base64.StdEncoding.DecodeString(s)
    if err != nil {
        return nil, err
    }
    i := new(big.Int)
    i.SetBytes(data)
    return i, nil
}

/*******************************************************************************

	Paul.B 
		Benchmark implemented as below,
		Actually, the performance of sign() and verify() looks quite significant,
		since a lot of TXs can be generated and proccessed in func calls in realtime (cases)


***************************************************************************************************/


type zeroReader struct{}

func (zeroReader) Read(buf []byte) (int, error) {
	for i := range buf {
		buf[i] = 0
	}
	return len(buf), nil
}

func Benchmarking(case_ int, N int, keyType int) {

	if case_ == 1 {
		BenchmarkKeyGeneration(N, keyType)
	}else if case_ == 2 {
		BenchmarkSigning(N, keyType)
	}else if case_ == 3 {
		BenchmarkVerification(N, keyType)
	}else{
		fmt.Printf("Not defined Testcase.....plz try again\n")
	}

}

func BenchmarkKeyGeneration(N int, keyType int) {
	//var zero zeroReader
	start := time.Now()

	for i := 0; i < N; i++ {
		if _, _, err := generateKey(rand.Reader, keyType); err != nil {
			log.Fatal(err)
		}
		//log.Printf("111111\n")	
	}
	end := time.Now()
	elapsed := end.Sub(start)
	log.Printf("elapsed time : %v (key Type : %d)\n", elapsed, keyType)
}

func BenchmarkSigning(N int, keyType int) {
	var zero zeroReader
	_, priv, err := generateKey(zero, keyType)
	if err != nil {
		log.Fatal(err)
	}
	message := []byte("Luv you!")
	//b.ResetTimer()
	start := time.Now()
	for i := 0; i < N; i++ {
		SignEW(priv, message)
	}
	end := time.Now()
	elapsed := end.Sub(start)
	log.Printf("elapsed time : %v (key Type : %d\n", elapsed, keyType)
}

func BenchmarkVerification(N int, keyType int) {
	var zero zeroReader
	pub, priv, err := generateKey(zero, keyType)
	if err != nil {
		log.Fatal(err)
	}
	message := []byte("luv you2!")
	signature := SignEW(priv, message)
	//b.ResetTimer()
	start := time.Now()
	for i := 0; i < N; i++ {
		Verify(pub, message, signature)
		//VerifyWthType(pub, message, signature, 1)
	}
	end := time.Now()
	elapsed := end.Sub(start)
	log.Printf("elapsed time : %v (key Type : %d\n", elapsed, keyType)
}

func GetKeyFromUser() []byte {

	scanner_ := bufio.NewScanner(os.Stdin)
	scanner_.Scan()
	t_ := scanner_.Text()

	//decodedPriv, _ := base64.StdEncoding.DecodeString(string(t))
	decodedKey, err := Decode(string(t_)) //58-based privateKey to original type
	
	if err != nil{
		fmt.Printf("Error : %v\n", err)
		return nil
	}

	return decodedKey

}

func obtainIDandAddressFromPubKey(public []byte, sort int) string { // 1 (ID) and 2 (Address)

	if sort == 1{ // ID
		return Encode(issueAccountID(public))
	}else if sort == 2{ // Address
		return generateAddress(public)
	}

	return ""

}

func obtainIDandAddressFromPrivKey(private []byte, sort int) string { // 1 (ID) and 2 (Address)

	public, err := derivePubFromPrivate(private)

	if err != nil{
		fmt.Printf("Error : %v\n", err)
		return ""
	}

	if sort == 1{ // ID
		return Encode(issueAccountID(public))
	}else if sort == 2{ // Address
		return generateAddress(public)
	}
	
	return "" // For porting, originally should return with error (multiple returns)
	
}


func issueAccountID(public []byte) []byte{

	// pubkey inner hash : SHA256
	sha256sum := sha256.Sum256(public)
	//fmt.Printf("1) sha256 	   sum : %x size : %d\n", sha256sum, len(sha256sum))
	//fmt.Printf("1) sha256[:31] sum : %x size : %d\n", sha256sum[:31], len(sha256sum[:31]))

	

	// pubkey outer hash : RIPEMD160
	h := ripemd160.New()
	///////////////////////////////////////
    h.Write(sha256sum[:31])  //original
    //h.Write(sha256sum[:])
    ///////////////////////////////////////

	accountID := h.Sum(nil)
    //fmt.Printf("RIPEMD160 : %x  (len:%d)\n", h.Sum(nil), len(accountID))

	return accountID
}


func generateAddress(public []byte) string {


	/*
	Paul.B
		Issue AccountID from Public Key
	*/

	accountID := issueAccountID(public)
	
	/*
	encoded58_ID := Encode(accountID)
	decoded_ID, _ := Decode(encoded58_ID)
	log.Printf("Isuued Account ID(encoded58) : %s (size:%d, %d)\n", encoded58_ID, len(encoded58_ID), len(decoded_ID))
	*/


	/*
	Paul.B
		Set prefix with "z", supported by HJ
	*/
	
	accountID[0] = 0x00
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

	/*
		Paul
			Finally, An address to be genereated has 24byte-long array.
			However, we extract partial information from the account and checksum, so we can put it some features to make differnece.

	*/

	var rawAddress [24]byte
	//accountID_ := accountID[0:20]
	//checkSum_ := checkSum[0:4]
	copy(rawAddress[0:19], accountID[0:19])  // --> extract 19 bytes (was 20)
	copy(rawAddress[20:23], checkSum[0:3])	 // --> extract 3 bytes  (was 4)


	spaceAddr := Encode(rawAddress[:])

	//log.Printf ("spaceAddr (raw)    : %x", rawAddress)2
	//log.Println("spaceAddr (base58) : ", spaceAddr)

	return spaceAddr
}



func doubleCheckSum4(source []byte) []byte{


	sum := sha256.Sum256(source)
	//fmt.Printf("1) sha256 sum : %x size : %d\n", sum, len(sum))
	sum = sha256.Sum256(sum[:31])
	//sum := sha256.Sum256(public_base64)

	//fmt.Printf("2) sha256 sum : %x size : %d\n", sum, len(sum))

	return sum[:31] 

}

func validateAddress(address string) (bool, error)  {

	/*
		Paul.B
			0) check first : "z", length 24 bytes
			1) Decode(address)
			2) Get the first 20 bytes, Get checksumbytes and validate witht the given the checksumbytes
	*/

	firstChar := string([]rune(address)[0])

	//log.Println("the address first letter : ", firstChar, "length : ", len(firstChar), "Address Length : ", len(address))

	
	if firstChar != string('z') {
		return false, errors.New(`validateAddress() : The address MUST should start with "z"`)
	}
	
	decodedAddr, _ := Decode(address)

	//log.Printf ("spaceAddr (raw)    : %x", decodedAddr)
	//log.Println("spaceAddr (base58) : ", address)


	if(len(decodedAddr) != 24){
		return false, errors.New(`validateAddress() : The address MUST should be the length of 24`)	
	}


	//var extractFirst20 [20]byte
	//copy(extractFirst20[:], decodedAddr[0:19])

	//log.Printf("decoded Address")

	newlChecksum := []byte(doubleCheckSum4(decodedAddr[0:19]))

	if !bytes.Equal(newlChecksum[0:3], decodedAddr[20:23] ) {

		//log.Printf("newlChecksum --> %x", newlChecksum[0:3])
		//log.Printf("Original --> %x", decodedAddr[20:23])

		return false, errors.New(`validateAddress() : Address checksum value does NOT match with the given one`)

	}


	return true, nil

}


func validateAddressAll(address string) (bool, error)  {

	/*
		Paul.B
			0) check first : "z", length 24 bytes
			1) Decode(address)
			2) Get the first 20 bytes, Get checksumbytes and validate witht the given the checksumbytes
	*/

	firstChar := string([]rune(address)[0])

	//log.Println("the address first letter : ", firstChar, "length : ", len(firstChar), "Address Length : ", len(address))

	var decodedAddr []byte

	
	if firstChar == string('z') { //edward
		decodedAddr, _ = Decoding58withType(address, 1)
	}else if firstChar == string('e') { // ecc
		decodedAddr, _ = Decoding58withType(address, 2)
	}else if firstChar == string('r') { // rsa

		//decodedAddr, _ = Decoding58withType(address, 3)
		return false, errors.New(`INVALID : TYPE(RSA) is not supported yet, Try later~.`)

	}else { // 

		return false, errors.New(`INVALID : The given type is not supported type.`)
	}
	
	

	//log.Printf ("spaceAddr (raw)    : %x", decodedAddr)
	//log.Println("spaceAddr (base58) : ", address)


	if(len(decodedAddr) != 24){
		return false, errors.New(`validateAddress() : The address MUST should be the length of 24`)	
	}


	//var extractFirst20 [20]byte
	//copy(extractFirst20[:], decodedAddr[0:19])

	//log.Printf("decoded Address")

	newlChecksum := []byte(doubleCheckSum4(decodedAddr[0:19]))

	if !bytes.Equal(newlChecksum[0:3], decodedAddr[20:23] ) {

		log.Printf("newlChecksum --> %x", newlChecksum[0:3])
		log.Printf("Original --> %x", decodedAddr[20:23])

		return false, errors.New(`validateAddress() : Address checksum value does NOT match with the given one`)

	}


	return true, nil

}

func detectKeyTypeFromPrivKey(priv []byte) int {


	_, err := ParsePKCS1PrivateKey(priv)
	if err == nil {
		return 3 //RSA
	}


	_, err = x509.ParseECPrivateKey(priv)
	if err == nil {
		return 2 //ECC
	}

	if l := len(priv); l == 64 {
		return 1
	}

	return 0

}

func detectKeyTypeFromPubKey(pub []byte) int {


	_, err := ParsePKCS1PublicKey(pub)
	if err == nil {
		return 3 //RSA
	}


	_, err = x509.ParsePKIXPublicKey(pub)
	if err == nil {
		return 2 //ECC
	}

	if l := len(pub); l == 32 {
		return 1 //edward
	}

	return 0

}





/**************************************************************
Paul.B
	The functions below are only for "number-based strings"
***************************************************************

func base58Encoding(input []byte) []byte{

	encoding := PaulEncoding // or RippleEncoding or BitcoinEncoding
	//encoding := RippleEncoding

	encoded, err := encoding.Encode(input)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	log.Println("Paul-base58 Encoding : ", string(encoded))

	return encoded
}

func base58Decoding(encoded []byte) []byte{

	encoding := PaulEncoding // or RippleEncoding or BitcoinEncoding
	//encoding := RippleEncoding

	decoded, err := encoding.Decode(encoded)
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	log.Println("Paul-base58 Decoding : ", string(decoded))

	return decoded
}

*/

