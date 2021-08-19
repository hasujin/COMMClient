// Copyright 2016 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package ed25519 implements the Ed25519 signature algorithm. See
// https://ed25519.cr.yp.to/.
//
// These functions are also compatible with the “Ed25519” function defined in
// https://tools.ietf.org/html/draft-irtf-cfrg-eddsa-05.
package marine

// This code is a port of the public domain, “ref10” implementation of ed25519
// from SUPERCOP.

import (
	"bytes"
	"crypto"
	cryptorand "crypto/rand"
	"crypto/sha512"
	"errors"
	"io"
	"strconv"
	"log"

//	"golang.org/x/crypto/ed25519/internal/edwards25519"
)
const (
	// PublicKeySize is the size, in bytes, of public keys as used in this package.
	PublicKeySize = 32
	// PrivateKeySize is the size, in bytes, of private keys as used in this package.
	PrivateKeySize = 64
	// SignatureSize is the size, in bytes, of signatures generated and verified by this package.
	SignatureSize = 64
)
// Public returns the PublicKey corresponding to priv.
// PublicKey is the type of Ed25519 public keys.
type PublicKey []byte

// PrivateKey is the type of Ed25519 private keys. It implements crypto.Signer.
type PrivateKey []byte

func (priv PrivateKey) PublicEW() crypto.PublicKey {
	publicKey := make([]byte, PublicKeySize)
	copy(publicKey, priv[32:])
	return PublicKey(publicKey)
}

// Sign signs the given message with priv.
// Ed25519 performs two passes over messages to be signed and therefore cannot
// handle pre-hashed messages. Thus opts.HashFunc() must return zero to
// indicate the message hasn't been hashed. This can be achieved by passing
// crypto.Hash(0) as the value for opts.
func (priv PrivateKey) SignEW_(rand io.Reader, message []byte, opts crypto.SignerOpts) (signature []byte, err error) {
	if opts.HashFunc() != crypto.Hash(0) {
		return nil, errors.New("ed25519: cannot sign hashed message")
	}

	return SignEW(priv, message), nil
}

// GenerateKey generates a public/private key pair using entropy from rand.
// If rand is nil, crypto/rand.Reader will be used.
func generateKeyEW(rand io.Reader) (publicKey PublicKey, privateKey PrivateKey, err error) {
	if rand == nil {
		rand = cryptorand.Reader
	}

	/*
	Paul:
		Prepare the two boxes for public and private keys, which are 32 and 64 respectively.

	*/

	privateKey = make([]byte, PrivateKeySize)
	publicKey = make([]byte, PublicKeySize)
	_, err = io.ReadFull(rand, privateKey[:32])
	if err != nil {
		return nil, nil, err
	}

	/*
		Paul : Digesting,.....
	*/

	var hBytes [32]byte
	var publicKeyBytes [32]byte
	var A ExtendedGroupElement

	//log.Printf("Before private Area : %x", privateKey[:32])

	digest := sha512.Sum512(privateKey[:32])
	digest[0] &= 248
	digest[31] &= 127
	digest[31] |= 64

	copy(hBytes[:], digest[:])
	//log.Printf("After digest, private Area : %x", privateKey[:32])  --> No changes!!


	/*
		Paul : Put some cal. on the hashed one,.....
	*/
	GeScalarMultBase(&A, &hBytes)
	A.ToBytes(&publicKeyBytes)

	copy(privateKey[32:], publicKeyBytes[:])
	copy(publicKey, publicKeyBytes[:])

	return publicKey, privateKey, nil

	/*

		privateKey =  |---------a-----------|---------b-----------|

		publicKey  =  |---------b-----------|


	*/


}

func generatePubKeyFromPrivateKeyEW(privateKey PrivateKey) (publicKey PublicKey, err error) {
	
	type_ := detectKeyTypeFromPrivKey(privateKey)
	if type_ != 1{
		return nil, errors.New("The given key is not Ed25519, plz try again!!\n")
	}

	//var publicKeyBytes [32]byte
	//privateKey = make([]byte, PrivateKeySize)
	publicKey = make([]byte, PublicKeySize)

	copy(publicKey[:], privateKey[32:])
	//copy(publicKey, publicKeyBytes[:])

	return publicKey, nil

	/*

		privateKey =  |---------a-----------|---------b-----------|

		publicKey  =  |---------b-----------|


	*/


}


func validateKeyEW(givenPrivateKey PrivateKey) bool {

	//First length check
	if (len([]byte(givenPrivateKey)) != PrivateKeySize){
		log.Printf("Error : Invalid Address Key Type (Length)")
		return false
	}

	/*
		Paul:
			Extracting pubkey from the given private key		
	*/

	//derivedPublicKey := make([]byte, PrivateKeySize)
	var hBytes [32]byte
	var derivedPublicKeyBytes [32]byte
	var originalPublicKeyBytes [32]byte
	var A ExtendedGroupElement


	digest := sha512.Sum512(givenPrivateKey[:32])
	digest[0] &= 248
	digest[31] &= 127
	digest[31] |= 64

	copy(hBytes[:], digest[:])
	//log.Printf("After digest, private Area : %x", privateKey[:32])  --> No changes!!


	/*
		Paul : Put some cal. on the hashed one,.....
	*/
	GeScalarMultBase(&A, &hBytes)
	A.ToBytes(&derivedPublicKeyBytes)


	copy(originalPublicKeyBytes[:], givenPrivateKey[32:])
	if !bytes.Equal(derivedPublicKeyBytes[:], originalPublicKeyBytes[:]){
		log.Println("Error : Invalid Address and Key")
		return false
	}


	log.Println("OK, Valid Address and Key")
//	givenPrivateKey


	return true

}
/*
func derivePubFromPrivate(givenPrivateKey PrivateKey) bool {

	//First length check
	if (len([]byte(givenPrivateKey)) != PrivateKeySize){
		log.Printf("Error : Invalid Address Key Type (Length)")
		return false
	}


	//derivedPublicKey := make([]byte, PrivateKeySize)
	var hBytes [32]byte
	var derivedPublicKeyBytes [32]byte
	var originalPublicKeyBytes [32]byte
	var A ExtendedGroupElement


	digest := sha512.Sum512(givenPrivateKey[:32])
	digest[0] &= 248
	digest[31] &= 127
	digest[31] |= 64

	copy(hBytes[:], digest[:])
	//log.Printf("After digest, private Area : %x", privateKey[:32])  --> No changes!!


	GeScalarMultBase(&A, &hBytes)
	A.ToBytes(&derivedPublicKeyBytes)

	copy(originalPublicKeyBytes[:], givenPrivateKey[32:])
	if !bytes.Equal(derivedPublicKeyBytes[:], originalPublicKeyBytes[:]){
		log.Println("Error : Invalid Address and Key")
		return false
	}


	log.Println("Extracting PublicKey......")
	log.Printf("	--> %x", derivedPublicKeyBytes)
//	givenPrivateKey


	return true

}
*/
func derivePubFromPrivateEW(givenprivateKey PrivateKey) (derivedprivateKey []byte, err error) {

	//First length check
	if (len([]byte(givenprivateKey)) != PrivateKeySize){
		log.Printf("Error : Invalid Address Key Type (Length)")
		return derivedprivateKey, errors.New("Error 1 : derivePubFromPrivate()")
	}

	/*
		Paul:
			Extracting pubkey from the given private key		
	*/

	//derivedprivateKey := make([]byte, privateKeySize)
	var hBytes [32]byte
	var derivedprivateKeyBytes [32]byte
	var originalprivateKeyBytes [32]byte
	var A ExtendedGroupElement


	digest := sha512.Sum512(givenprivateKey[:32])
	digest[0] &= 248
	digest[31] &= 127
	digest[31] |= 64

	copy(hBytes[:], digest[:])
	//log.Printf("After digest, private Area : %x", privateKey[:32])  --> No changes!!


	/*
		Paul.B put some cal. on the hashed one,.....
	*/
	GeScalarMultBase(&A, &hBytes)
	A.ToBytes(&derivedprivateKeyBytes)

	copy(originalprivateKeyBytes[:], givenprivateKey[32:])
	if !bytes.Equal(derivedprivateKeyBytes[:], originalprivateKeyBytes[:]){
		log.Println("Error : Invalid Address and Key")
		return derivedprivateKey, errors.New("Error 2 : derivePubFromPrivate()")
	}


	//log.Println("Extracting privateKey......")
	//log.Printf("	--> %x", derivedprivateKeyBytes)

	//derivedprivateKey = string(Encode(derivedprivateKeyBytes[:]))

	//log.Printf("	--> %s", derivedprivateKey)
//	givenprivateKey
	
	derivedprivateKey = derivedprivateKeyBytes[:]


	return derivedprivateKey, nil

}

// Sign signs the message with privateKey and returns a signature. It will
// panic if len(privateKey) is not PrivateKeySize.
func SignEW(privateKey PrivateKey, message []byte) []byte {
	if l := len(privateKey); l != PrivateKeySize {
		panic("pled25519: bad private key length: " + strconv.Itoa(l))
	}

	h := sha512.New()
	h.Write(privateKey[:32])

	var digest1, messageDigest, hramDigest [64]byte
	var expandedSecretKey [32]byte
	h.Sum(digest1[:0])
	copy(expandedSecretKey[:], digest1[:])
	expandedSecretKey[0] &= 248
	expandedSecretKey[31] &= 63
	expandedSecretKey[31] |= 64

	h.Reset()
	h.Write(digest1[32:])
	h.Write(message)
	h.Sum(messageDigest[:0])

	var messageDigestReduced [32]byte
	ScReduce(&messageDigestReduced, &messageDigest)
	var R ExtendedGroupElement
	GeScalarMultBase(&R, &messageDigestReduced)

	var encodedR [32]byte
	R.ToBytes(&encodedR)

	h.Reset()
	h.Write(encodedR[:])
	h.Write(privateKey[32:])
	h.Write(message)
	h.Sum(hramDigest[:0])
	var hramDigestReduced [32]byte
	ScReduce(&hramDigestReduced, &hramDigest)

	var s [32]byte
	ScMulAdd(&s, &hramDigestReduced, &expandedSecretKey, &messageDigestReduced)

	signature := make([]byte, SignatureSize)
	copy(signature[:], encodedR[:])
	copy(signature[32:], s[:])

	return signature
}

// Verify reports whether sig is a valid signature of message by publicKey. It
// will panic if len(publicKey) is not PublicKeySize.
func VerifyEW(publicKey PublicKey, message, sig []byte) bool {
	if l := len(publicKey); l != PublicKeySize {
		panic("pled25519: bad public key length: " + strconv.Itoa(l))
	}

	if len(sig) != SignatureSize || sig[63]&224 != 0 {
		return false
	}

	var A ExtendedGroupElement
	var publicKeyBytes [32]byte
	copy(publicKeyBytes[:], publicKey)
	if !A.FromBytes(&publicKeyBytes) {
		return false
	}
	FeNeg(&A.X, &A.X)
	FeNeg(&A.T, &A.T)

	h := sha512.New()
	h.Write(sig[:32])
	h.Write(publicKey[:])
	h.Write(message)
	var digest [64]byte
	h.Sum(digest[:0])

	var hReduced [32]byte
	ScReduce(&hReduced, &digest)

	var R ProjectiveGroupElement
	var b [32]byte
	copy(b[:], sig[32:])
	GeDoubleScalarMultVartime(&R, &hReduced, &A, &b)

	var checkR [32]byte
	R.ToBytes(&checkR)
	return bytes.Equal(sig[:32], checkR[:])
}
