package marine

import (
	"encoding/json"
	//"log"
	"math/big"
	"os"

	//"strconv"
	skill "github.com/hasujin/COMMClient/marineV21/skills"
	//pb "github.com/hyperledger/fabric-sdk-go/project/comm/protos"
	//"golang.org/x/net/context"
	//"google.golang.org/grpc"
	"bufio"
	"errors"

	//	"crypto/rand"
	"fmt"
	//"strings"
	//	"bytes"
)

type envelope struct {
	Pubkey    []byte `json:"pubkey"`
	Signature []byte `json:"signature"`
	Message   []byte `json:"message"`
	//Attachment 	string `json:"attachment"`
}

/*
	Paul :
		buildEnvelopeForApprove() --> deprecated!!
		However, PLZ DO NOT REMOVE !!!

*/

func buildEnvelopeForApprove(owner, spender, label, limit, expired string) ([]byte, error) {

	////////////////////////////////////////////////////////////////
	/// Obtain PrivateKey from User-input
	////////////////////////////////////////////////////////////////
	fmt.Printf(`	
		------------------------------------------------------------
		Please input your PrivateKey (base58) to check and sign 
		------------------------------------------------------------
		----->> `)

	////////////////////////////////////////////////////////////////

	scanner_ := bufio.NewScanner(os.Stdin)
	scanner_.Scan()
	t_ := scanner_.Text()

	//decodedPriv, _ := base64.StdEncoding.DecodeString(string(t))
	decodedPriv, err := Decode(string(t_)) //58-based privateKey to original type
	//validateKey(decodedPriv)
	fmt.Printf("\n\n\n\n")

	//message := []byte("I LUV U")
	message := []byte(owner + spender + label + limit + expired)

	// Sign
	sig, err := Sign(decodedPriv, message)
	if err != nil {
		fmt.Printf("buildEnvelopeForApprove() : Error in Sign()\n")
	}

	fmt.Printf("Signature has been made!! (signature size : %d)\n", len(sig))

	derivedPubKey, _ := derivePubFromPrivate([]byte(decodedPriv))
	decoded58Public, _ := Decode(string(derivedPubKey))

	//var envelope [96 + 2]byte 			//envelope(98) = pubkey(32) + 1(0x00) + sig(64) + 1 (0x00)

	var envelope_ envelope
	envelope_.Pubkey = decoded58Public
	envelope_.Signature = sig
	envelope_.Message = message

	//copy(envelope[0:32], decoded58Public[:])
	//copy(envelope[33:97], sig[:])

	envelopeBytes, err := json.Marshal(envelope_)
	if err != nil {
		fmt.Printf("buildEnvelopeForApprove() : Marshal return error %v", err)
		return nil, err
	}

	////////////////////////////////////////////////////////////////
	/// Paul
	/// 	Let us check the signature...
	////////////////////////////////////////////////////////////////
	if !Verify(envelope_.Pubkey, envelope_.Message, envelope_.Signature) { //
		fmt.Printf("---> Check Verify() : Invalid signature --> Rejected\n")
		return nil, errors.New("buildEnvelopeForApprove() : Verification Error")
	} else {
		fmt.Printf("Local Check ---> Valid envelope!! (envelope size : %d)\n", len(envelopeBytes))
	}

	return envelopeBytes, nil

}

func SumMessages(args ...string) []byte {

	len_ := len(args)
	arg_count := len_ - 1

	if len(args) <= 1 {
		fmt.Printf("More arguments required,...\n")
		return nil
	}

	////////////////////////////////////////////////////////////////////////
	/// Paul :
	///			Make sure how a message is made of !!!
	////////////////////////////////////////////////////////////////////////

	switch args[0] {

	case "transferFrom":

		if arg_count != 4 {
			fmt.Printf("transferFrom() requires [4] arguments (from, to, amount, label)\n")
			return nil
		}

	case "chargeFee":

		if arg_count != 4 {

			fmt.Printf("chargeFee() requires [4] arguments (from, to, amount, label)\n")
			return nil
		}

	case "approve":

		if arg_count != 5 {

			fmt.Printf("Approve() requires [5] arguments (owner, spender, label, limit, expired)\n")
			return nil
		}

	case "establish":

		if arg_count != 2 {
			fmt.Printf("Establish() requires [2] arguments (founder, title)\n")
			return nil
		}

	case "join":

		if arg_count != 4 {

			fmt.Printf("Join() requires [4] arguments (host, guest, title, type(0/1/2/3))\n")
			return nil

		}

	case "leave":

		if arg_count != 2 {
			fmt.Printf("Leave() requires [2] arguments (member, title)\n")
			return nil

		}

	case "modifyMember":

		if arg_count != 3 {
			fmt.Printf("modifyMember() requires [3] arguments (member, title, type)\n")
			return nil
		}

	case "status":

		if arg_count != 2 {
			fmt.Printf("status() requires [2] arguments (member, title)\n")
			return nil

		}

	case "issue":

		if arg_count != 3 {
			fmt.Printf("issue() requires [3] arguments (symbol, issuer, totalsupply)\n")
			return nil
		}

	case "mint":

		if arg_count != 3 {
			fmt.Printf("mint() requires [3] arguments (minter, symbol, mintedamount)\n")
			return nil

		}

	case "burn":

		if arg_count != 3 {
			fmt.Printf("burn() requires [3] arguments (burner, symbol, burnedamount)\n")
			return nil
		}

	case "deposit":

		if arg_count != 4 {
			fmt.Printf("DEX|deposit() requires [4] arguments (from, to, lable, amount)\n")
			return nil
		}

	case "putPrivate":

	case "login":

	case "addXrate":

	case "eXchange":

	default:
		fmt.Printf("Not defined request : " + args[0] + "?")
		return nil

	}

	//Generate message for Sign and Verify
	var message_ string
	for _, each := range args {
		message_ = message_ + each
	}

	return []byte(message_)

}

func buildEnvelope(args ...string) ([]byte, error) {

	var privKey string
	len_ := len(args)
	arg_count := len_ - 1

	if len(args) <= 1 {
		return nil, errors.New("More arguments required,...\n")
	}
	fmt.Printf("\n\n\ntype: %v\n", args[0])

	////////////////////////////////////////////////////////////////////////
	/// Paul :
	///			Make sure how a message is made of !!!
	////////////////////////////////////////////////////////////////////////

	switch args[0] {

	case "transferFrom":

		if arg_count < 4 || arg_count > 5 {
			return nil, errors.New("transferFrom() requires [4 or 5] arguments (from, to, amount, label, (privateKey))\n")
		}

		if arg_count == 5 {
			privKey = args[len(args)-1]
			args = args[:len(args)-2] // remove privkey from args
		}

	case "chargeFee":

		if arg_count != 4 {
			return nil, errors.New("chargeFee() requires [4] arguments (from, to, amount, label)\n")
		}

	case "approve":

		if arg_count < 5 || arg_count > 6 {
			return nil, errors.New("Approve() requires [5 or 6] arguments (owner, spender, label, limit, expired,(privatekey))\n")
		}
		if arg_count == 6 {
			privKey = args[len(args)-1]
			args = args[:len(args)-2] // remove privkey from args
		}

	case "establish":

		if arg_count != 2 {
			return nil, errors.New("Establish() requires [2] arguments (founder, title)\n")
		}

	case "join":

		if arg_count != 4 {
			return nil, errors.New("Join() requires [4] arguments (host, guest, title, type(0/1/2/3))\n")
		}

	case "leave":

		if arg_count != 2 {
			return nil, errors.New("Leave() requires [2] arguments (member, title)\n")
		}

	case "modifyMember":

		if arg_count != 3 {
			return nil, errors.New("modifyMember() requires [3] arguments (member, title, type)\n")
		}

	case "status":

		if arg_count != 2 {
			return nil, errors.New("status() requires [2] arguments (member, title)\n")
		}

	case "issue":
		//sj

		if arg_count < 3 || arg_count > 4 {
			return nil, errors.New("issue() requires [3 or 4] arguments (symbol, issuer, totalsupply, (privateKey))\n")
		}
		if arg_count == 4 {
			privKey = args[len(args)-1]
			args = args[:len(args)-2] // remove privkey from args
		}

	case "mint":

		if arg_count != 3 {
			return nil, errors.New("mint() requires [3] arguments (minter, symbol, mintedamount)\n")
		}

	case "burn":

		if arg_count != 3 {
			return nil, errors.New("burn() requires [3] arguments (burner, symbol, burnedamount)\n")
		}

	case "deposit":

		if arg_count != 4 {
			return nil, errors.New("DEX|deposit() requires [4] arguments (from, to, lable, amount)\n")
		}

	case "putPrivate":

	case "login":

	case "addXrate":

	case "eXchange":

	case "mintNFT":
		privKey = args[len(args)-1]
		args = args[:len(args)-2] // remove privkey from args
	default:
		return nil, errors.New("Not defined request : " + args[0] + "?")

	}

	//Generate message for Sign and Verify
	var message_ string
	for _, each := range args {
		message_ = message_ + each
	}

	/*

		Paul:
			This message will be hashed with sha256
			ComputeSHA256()


	*/

	message := skill.ComputeSHA256([]byte(message_))

	//fmt.Printf("The message to used in sig. is made of : [%s]\n", message)

	////////////////////////////////////////////////////////////////
	/// Obtain PrivateKey from User-input
	////////////////////////////////////////////////////////////////
	//GetPrivKey()

	//sj - if private key == nil
	var decodedPriv []uint8
	var err error
	if privKey == "" {
		var priv string

		fmt.Printf(`
		------------------------------------------------------------
		Please input your PrivateKey (base58) to check and sign
		------------------------------------------------------------
		----->> `)

		////////////////////////////////////////////////////////////////

		scanner_ := bufio.NewScanner(os.Stdin)
		scanner_.Scan()
		t_ := scanner_.Text()
		priv = fmt.Sprintf("%v", string(t_))

		fmt.Printf("priv: %s", priv)

		decodedPriv, err = Decode(priv) //58-based privateKey to original type
	} else {

		decodedPriv, err = Decode(privKey) //58-based privateKey to original type
	}

	//validateKey(decodedPriv)
	fmt.Printf("Decoded Private key %v(len:%d)\n\n\n\n", decodedPriv, len(decodedPriv))

	////////////////////////////////////////////////////////////////
	/// Paul
	/// 	Lets make Signature !!
	////////////////////////////////////////////////////////////////
	sig, err := Sign(decodedPriv, []byte(message))
	if err != nil {
		fmt.Printf("buildEnvelope() : %v\n", err)
	}

	fmt.Printf("Signature has been made!! (signature size : %d)\n", len(sig))

	derivedPubKey, _ := derivePubFromPrivate(decodedPriv)
	//decoded58Public, _ := Decode(string(derivedPubKey))

	////////////////////////////////////////////////////////////////
	/// Paul
	/// 	Lets make envelope!!
	////////////////////////////////////////////////////////////////

	var envelope_ envelope
	envelope_.Pubkey = derivedPubKey
	//envelope_.Pubkey = decoded58Public
	envelope_.Signature = sig
	envelope_.Message = []byte(message)

	envelopeBytes, err := json.Marshal(envelope_)
	if err != nil {
		fmt.Printf("buildEnvelope() : Marshal return error %v", err)
		return nil, err
	}

	////////////////////////////////////////////////////////////////
	/// Paul
	/// 	Let us check the signature...
	////////////////////////////////////////////////////////////////
	if !Verify(envelope_.Pubkey, envelope_.Message, envelope_.Signature) { //
		fmt.Printf("Check Verify() : Invalid signature --> Rejected\n")
		return nil, errors.New("buildEnvelope() : Verification Error")
	} else {
		fmt.Printf("Local Check ---> Valid envelope!! (envelope size : %d)\n", len(envelopeBytes))
	}

	return envelopeBytes, nil

}

/*func buildEnvelope(args ...string) ([]byte, error) {

	len_ := len(args)
	arg_count := len_ - 1

	if len(args) <= 1 {
		return nil, errors.New("More arguments required,...\n")
	}

	////////////////////////////////////////////////////////////////////////
	/// Paul :
	///			Make sure how a message is made of !!!
	////////////////////////////////////////////////////////////////////////

	switch args[0] {

	case "transferFrom":

		if arg_count != 4 {
			return nil, errors.New("transferFrom() requires [4] arguments (from, to, amount, label)\n")
		}

	case "chargeFee":

		if arg_count != 4 {
			return nil, errors.New("chargeFee() requires [4] arguments (from, to, amount, label)\n")
		}

	case "approve":

		if arg_count != 5 {
			return nil, errors.New("Approve() requires [5] arguments (owner, spender, label, limit, expired)\n")
		}

	case "establish":

		if arg_count != 2 {
			return nil, errors.New("Establish() requires [2] arguments (founder, title)\n")
		}

	case "join":

		if arg_count != 4 {
			return nil, errors.New("Join() requires [4] arguments (host, guest, title, type(0/1/2/3))\n")
		}

	case "leave":

		if arg_count != 2 {
			return nil, errors.New("Leave() requires [2] arguments (member, title)\n")
		}

	case "modifyMember":

		if arg_count != 3 {
			return nil, errors.New("modifyMember() requires [3] arguments (member, title, type)\n")
		}

	case "status":

		if arg_count != 2 {
			return nil, errors.New("status() requires [2] arguments (member, title)\n")
		}

	case "issue":
		//sj
		if arg_count != 3 {
			return nil, errors.New("issue() requires [3] arguments (symbol, issuer, totalsupply)\n")
		}

	case "mint":

		if arg_count != 3 {
			return nil, errors.New("mint() requires [3] arguments (minter, symbol, mintedamount)\n")
		}

	case "burn":

		if arg_count != 3 {
			return nil, errors.New("burn() requires [3] arguments (burner, symbol, burnedamount)\n")
		}

	case "deposit":

		if arg_count != 4 {
			return nil, errors.New("DEX|deposit() requires [4] arguments (from, to, lable, amount)\n")
		}

	case "putPrivate":

	case "login":

	case "addXrate":

	case "eXchange":

	default:
		return nil, errors.New("Not defined request : " + args[0] + "?")

	}

	//Generate message for Sign and Verify
	var message_ string
	for _, each := range args {
		message_ = message_ + each
	}


	message := skill.ComputeSHA256([]byte(message_))

	//fmt.Printf("The message to used in sig. is made of : [%s]\n", message)

	////////////////////////////////////////////////////////////////
	/// Obtain PrivateKey from User-input
	////////////////////////////////////////////////////////////////
	//GetPrivKey()

	fmt.Printf(`
	------------------------------------------------------------
	Please input your PrivateKey (base58) to check and sign
	------------------------------------------------------------
	----->> `)

	////////////////////////////////////////////////////////////////

	scanner_ := bufio.NewScanner(os.Stdin)
	scanner_.Scan()
	t_ := scanner_.Text()

	//decodedPriv, _ := base64.StdEncoding.DecodeString(string(t))
	decodedPriv, err := Decode(string(t_)) //58-based privateKey to original type

	//validateKey(decodedPriv)
	fmt.Printf("Decoded Private key %v(len:%d)\n\n\n\n", decodedPriv, len(decodedPriv))

	////////////////////////////////////////////////////////////////
	/// Paul
	/// 	Lets make Signature !!
	////////////////////////////////////////////////////////////////
	sig, err := Sign(decodedPriv, []byte(message))
	if err != nil {
		fmt.Printf("buildEnvelope() : %v\n", err)
	}

	fmt.Printf("Signature has been made!! (signature size : %d)\n", len(sig))

	derivedPubKey, _ := derivePubFromPrivate(decodedPriv)
	//decoded58Public, _ := Decode(string(derivedPubKey))

	////////////////////////////////////////////////////////////////
	/// Paul
	/// 	Lets make envelope!!
	////////////////////////////////////////////////////////////////

	var envelope_ envelope
	envelope_.Pubkey = derivedPubKey
	//envelope_.Pubkey = decoded58Public
	envelope_.Signature = sig
	envelope_.Message = []byte(message)

	envelopeBytes, err := json.Marshal(envelope_)
	if err != nil {
		fmt.Printf("buildEnvelope() : Marshal return error %v", err)
		return nil, err
	}

	////////////////////////////////////////////////////////////////
	/// Paul
	/// 	Let us check the signature...
	////////////////////////////////////////////////////////////////
	if !Verify(envelope_.Pubkey, envelope_.Message, envelope_.Signature) { //
		fmt.Printf("Check Verify() : Invalid signature --> Rejected\n")
		return nil, errors.New("buildEnvelope() : Verification Error")
	} else {
		fmt.Printf("Local Check ---> Valid envelope!! (envelope size : %d)\n", len(envelopeBytes))
	}

	return envelopeBytes, nil

}*/

func buildEnvelopeWithGivens(pubkey []byte, messsage []byte, r *big.Int, s *big.Int) ([]byte, error) {

	var envelope_ envelope

	// public key
	envelope_.Pubkey = pubkey
	//envelope_.Pubkey = decoded58Public

	// message
	envelope_.Message = skill.ComputeSHA256(messsage)

	// (r,s) --> siganture
	signature_ := signature{
		R: r.String(),
		S: s.String(),
	}

	signatureByte, err := json.Marshal(signature_)
	if err != nil {
		fmt.Printf("Marshal Error : %v\n", err)
		return nil, err
	}
	envelope_.Signature = signatureByte

	envelopeBytes, err := json.Marshal(envelope_)
	if err != nil {
		fmt.Printf("buildEnvelope() : Marshal return error %v", err)
		return nil, err
	}

	////////////////////////////////////////////////////////////////
	/// Paul
	/// 	Let us check the signature...
	////////////////////////////////////////////////////////////////
	if !Verify(envelope_.Pubkey, envelope_.Message, envelope_.Signature) { //
		fmt.Printf("Check Verify() : Invalid signature --> Rejected\n")
		return nil, errors.New("buildEnvelope() : Verification Error")
	} else {
		fmt.Printf("Local Check ---> Valid envelope!! (envelope size : %d)\n", len(envelopeBytes))
	}

	return envelopeBytes, nil

}

func buildEnvelopeAndAccount(args ...string) ([]byte, []byte) {

	var message string

	len_ := len(args)
	arg_count := len_ - 1

	if len(args) <= 1 {
		return nil, nil
	}

	////////////////////////////////////////////////////////////////////////
	/// Paul :
	///			Make sure how a message is made of !!!
	////////////////////////////////////////////////////////////////////////

	switch args[0] {

	case "registerXindex":

		if arg_count != 1 {
			return nil, nil
		}

	default:
		return nil, nil

	}

	//Generate message for Sign and Verify
	for _, each := range args {
		message = message + each
	}

	//fmt.Printf("The message to used in sig. is made of : [%s]\n", message)

	////////////////////////////////////////////////////////////////
	/// Obtain PrivateKey from User-input
	////////////////////////////////////////////////////////////////
	fmt.Printf(`	
		------------------------------------------------------------
		Please input your PrivateKey (base58) to check and sign 
		------------------------------------------------------------
		----->> `)

	////////////////////////////////////////////////////////////////

	scanner_ := bufio.NewScanner(os.Stdin)
	scanner_.Scan()
	t_ := scanner_.Text()

	//decodedPriv, _ := base64.StdEncoding.DecodeString(string(t))
	decodedPriv, err := Decode(string(t_)) //58-based privateKey to original type
	//validateKey(decodedPriv)
	fmt.Printf("Decoded Private key %v(len:%d)\n\n\n\n", decodedPriv, len(decodedPriv))

	////////////////////////////////////////////////////////////////
	/// Paul
	/// 	Lets make Signature !!
	////////////////////////////////////////////////////////////////
	sig, err := Sign(decodedPriv, []byte(message))
	if err != nil {
		fmt.Printf("buildEnvelope() : %v\n", err)
	}

	fmt.Printf("Signature has been made!! (signature size : %d)\n", len(sig))

	derivedPubKey, _ := derivePubFromPrivate(decodedPriv)
	//decoded58Public, _ := Decode(string(derivedPubKey))

	////////////////////////////////////////////////////////////////
	/// Paul
	/// 	Lets make envelope!!
	////////////////////////////////////////////////////////////////

	var envelope_ envelope
	envelope_.Pubkey = derivedPubKey
	//envelope_.Pubkey = decoded58Public
	envelope_.Signature = sig
	envelope_.Message = []byte(message)

	envelopeBytes, err := json.Marshal(envelope_)
	if err != nil {
		fmt.Printf("buildEnvelope() : Marshal return error %v", err)
		return nil, nil
	}

	////////////////////////////////////////////////////////////////
	/// Paul
	/// 	Let us check the signature...
	////////////////////////////////////////////////////////////////
	if !Verify(envelope_.Pubkey, envelope_.Message, envelope_.Signature) { //
		fmt.Printf("Check Verify() : Invalid signature --> Rejected\n")
		return nil, nil
	} else {
		fmt.Printf("Local Check ---> Valid envelope!! (envelope size : %d)\n", len(envelopeBytes))
	}

	return envelopeBytes, []byte(generateAddress(envelope_.Pubkey))

}
