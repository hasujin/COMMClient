/*
 *
 * Copyright 2015, Google Inc.
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are
 * met:
 *
 *     * Redistributions of source code must retain the above copyright
 * notice, this list of conditions and the following disclaimer.
 *     * Redistributions in binary form must reproduce the above
 * copyright notice, this list of conditions and the following disclaimer
 * in the documentation and/or other materials provided with the
 * distribution.
 *     * Neither the name of Google Inc. nor the names of its
 * contributors may be used to endorse or promote products derived from
 * this software without specific prior written permission.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS
 * "AS IS" AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT
 * LIMITED TO, THE IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR
 * A PARTICULAR PURPOSE ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT
 * OWNER OR CONTRIBUTORS BE LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL,
 * SPECIAL, EXEMPLARY, OR CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT
 * LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR SERVICES; LOSS OF USE,
 * DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER CAUSED AND ON ANY
 * THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY, OR TORT
 * (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
 * OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 *
 */

package marine

import (
	"encoding/json"

	//"log"
	"os"
	"strconv"

	skill "github.com/hasujin/COMMClient/marineV21/skills"
	pb "github.com/hasujin/ston_common/protos"
	"google.golang.org/grpc"

	//ed "github.com/hyperledger/fabric-sdk-go/project/comm/COMMClient/marineV2/encdec"

	"crypto/rand"
	"fmt"
	"strings"

	"golang.org/x/net/context"

	//"bytes"
	"crypto/sha256"
	"time"

	/*
	   Please Do NOT Use this Color library,
	   when you put this into chaincode or actual running-machine/module(s).

	   Do not want? --> replace {color.} with {fmt. or fmt.}
	   (Paul.B)
	*/
	color "github.com/mitchellh/colorstring"
	//colorstring.Println("[blue]Hello [red]World!")
)

var myMarine Marine

func SetMode(mode string) {
	myMarine.Mode = mode
}

func GetMode() string {
	return myMarine.Mode
}

func ExecuteMarine(args []string, user string, tChannel string, tChaincode string) (mResponse *pb.ProcessResponse, e error) {

	userID_ = user
	channelID_ = tChannel
	chaincodeID_ = tChaincode

	if args != nil {
		os.Args = args
	}

	if len(os.Args) < 2 {
		Help()
		return

	}

	/*	tmp := os.Getenv("ASSET_CHANNEL")
		if tmp != "" {
			tokenChannel = tmp
		}
		//token chaincode name
		tmp = os.Getenv("ASSET_CC")
		if tmp != "" {
			tokenCC = tmp
		}*/

	type_, err := strconv.Atoi(os.Args[1])
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}

	if len(os.Args) > 2 && type_ != 88 && type_ != 28 && type_ != 82 {
		address = os.Args[2]
	}
	/*
		tls := os.Getenv("TLS")
		if tls == "true" {

			certFILE = os.Getenv("CERT_FILE")
			if certFILE == "" {
				//certFILE = "./TLScerts2/server.crt"
				certFILE = "./TLScerts3/serverED.crt"
				fmt.Printf("STARGATE[D] Cert is NOT given. Try to read key from : %s\n", certFILE)
			}
			TLS = true
			//fmt.Printf("TLS (ON) --> Cert is from (%s)\n", certFILE)
		} else {
			//fmt.Printf("TLS (OFF)\n")
		}

		// Set up a connection to the server.
		var conn *grpc.ClientConn

		if certFILE != "" && TLS == true {

			///////////////////////////////////////////////////////////////////////////
			//    Marine in TLS -- Paul
			///////////////////////////////////////////////////////////////////////////

			creds, err := credentials.NewClientTLSFromFile(certFILE, "")
			if err != nil {
				fmt.Printf("Cannot load tls cert: %s\n", err)
				return
			}
			// Initiate a connection with the server
			conn, err = grpc.Dial(address, grpc.WithTransportCredentials(creds))
			if err != nil {
				fmt.Printf("Cannot not connect: %s", err)
				return
			}

		} else {
			conn, err = grpc.Dial(address, grpc.WithInsecure())
			if err != nil {
				fmt.Printf("did not connect: %v", err)
				return
			}
		}
	*/

	var conn *grpc.ClientConn
	conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		fmt.Printf("did not connect: %v", err)
		return
	}

	defer conn.Close()

	c := pb.NewCOMMLayerClient(conn)

	var resp *pb.ProcessResponse
	var respErr error

	switch pb.MessageType(type_) {

	case pb.MessageType_QUERY_CHAINCODE:
		if len(os.Args) != 7 {

			fmt.Printf("Ex)./%s type[7] [serverAdress:port] userID channelID chaincode query_args({\"Args\":[\"query\",\"c\"]})\n", VERSION)
			return
		}
		userID_ = os.Args[3]
		channelID_ = os.Args[4]
		chaincodeID_ = os.Args[5]

		//args_.Args = append(args_.Args, os.Args[3], os.Args[4], os.Args[5], os.Args[6])
		fmt.Printf("Args : %s", os.Args[6])

		var sj []interface{}
		var fj interface{}
		sj = append(sj, userID_, channelID_, chaincodeID_)

		err := json.Unmarshal([]byte(os.Args[6]), &fj)
		if err != nil {
			fmt.Printf("Unmarshal return error: %v", err)
			return
		}
		sj = append(sj, fj)
		mj := make(map[string]interface{})
		mj["Args"] = sj

		mjBytes, err = json.Marshal(mj)
		if err != nil {
			fmt.Printf("Marshal return error %v", err)
			return
		}
	case pb.MessageType_INVOKE_CHAINCODE:
		len_arg := len(os.Args)

		if len_arg < 7 {
			fmt.Printf("Ex)./%s type[8] [serverAdress:port] userID channelID chaincode invoke_args({\"Args\":[\"invoke\",\"c\"]}) attachment(option)\n", VERSION)

			return
		}

		userID_ = os.Args[3]
		channelID_ = os.Args[4]
		chaincodeID_ = os.Args[5]

		if len_arg == 8 { // Have a Pad??
			attachment = []byte(os.Args[7])
		} else {
			attachment = nil
		}

		//args_.Args = append(args_.Args, os.Args[3], os.Args[4], os.Args[5], os.Args[6])
		fmt.Printf("Args : %s", os.Args[6])

		var sj []interface{}
		var fj interface{}
		sj = append(sj, userID_, channelID_, chaincodeID_)

		err := json.Unmarshal([]byte(os.Args[6]), &fj)
		if err != nil {
			fmt.Printf("Unmarshal return error: %v\n", err)
			return
		}
		sj = append(sj, fj)
		mj := make(map[string]interface{})
		mj["Args"] = sj

		mjBytes, err = json.Marshal(mj)
		if err != nil {
			fmt.Printf("Marshal return error %v\n", err)
			return
		}

	case 88: // Key and Address Generagtion
		if len(os.Args) != 3 && len(os.Args) != 4 {
			fmt.Printf(`

            Ex)./%s type[88] keytype[Edward25519:1, ECC:2, RSA:3] (seed)
            
            If seed not given, a random value is set to generate your key pairs.
            (Seed is recommended for the case you may lose your keys)


            Notice!!
            The Keys and address are generated inside local machine, not connecting to any server or medium in the internet!!
`, VERSION)
			return
		}

		keyTpye, err := strconv.Atoi(os.Args[2])

		if err != nil {
			fmt.Printf("Key Type Error, Plz check again!!\n")
			return
		}

		if len(os.Args) == 4 {
			seed = os.Args[3]
		}

		//marine(seed)
		var public, private []byte

		//seed := "All Roads Lead to Rome of FABRIC"  // --> 32 characters or more required!!!

		if seed == "" {
			public, private, _ = generateKey(rand.Reader, keyTpye)

		} else if len(seed) < 32 {
			color.Println("The SEED MUST be longer than 32 characters.....Try again!!\n")
			color.Printf("Your input SEED has the length of [%d]......Put more!!\n", len(seed))
			return
		} else {

			public, private, _ = generateKey(strings.NewReader(seed), keyTpye)

		}

		private_base58 := Encode(private)
		public_base58 := Encode(public)
		spaceAddr := generateAddress(public)

		ress := fmt.Sprintf(`{"private_key":"%v", "public_key":"%v", "address":"%v"}`, private_base58, public_base58, spaceAddr)
		fmt.Printf("%v", ress)

		var resultb pb.ProcessResponse
		resultb.Payload = []byte(ress)
		resultb.Status = int32(pb.Status_SUCCESS)

		if myMarine.Mode == "CLI" {
			fmt.Printf(`


				            ---------------------------------------------------------------------------------------------------------------------------------------------------------------
				            1) Public  Key (base58)     ---> %s
				                           (base58_btc) ---> %s
				                           (Raw : %x)
				                           (len : %d)

				            2) Private Key (base58)     ---> %s
				                           (base58_btc) ---> %s
				                           (Raw : %x)
				                           (len : %d)

				            3) Address -------> %s
				            ---------------------------------------------------------------------------------------------------------------------------------------------------------------




				`, public_base58, Encode_(public), public, len(public), private_base58, Encode_(private), private, len(private), spaceAddr)
		}
		return &resultb, nil

	case 888: // balanceof() from QUERY command.

		if len(os.Args) != 4 {

			//fmt.Printf(`Ex)./marine ip:port 888 '{"Args":["balanceOf",{"address":"z39UzmN7GhPqkYiQfVWY4pYADMQzv7qWC"}]}' `)
			fmt.Printf("Ex)./%s 888 ip:port Address(z39UzmN7GhPqkYiQfVWY4pYADMQzv7qWC)\n", VERSION)
			e = fmt.Errorf("args len ERROR")
			//                  0       1   2       3
			return nil, e
		}

		fmt.Printf("Given Address : %s\n", os.Args[3])

		/*
		   Paul : Need to check if the address given is correct or not
		*/

		//balance_query := `{"Args":[ADMIN_USER,tokenChannel,"generic_cc",{"Args":["balanceOf",{"address":"%s"}]}]}`
		balance_query := `{"Args":["balanceOf",{"address":"%s"}]}`
		result, _ := validateAddress(os.Args[3])

		if result {
			balance_query = fmt.Sprintf(balance_query, os.Args[3])
			fmt.Printf("Balance Query : %s\n", balance_query)

		} else {
			fmt.Printf("Please Try again,....\nThe given address is INVALID !!! : %s\n", os.Args[3])
			e = fmt.Errorf("Please Try again,....\nThe given address is INVALID !!! : %s\n", os.Args[3])
			return nil, e
		}

		//=========================================================================================================
		// Hoding,..Here!!
		//=========================================================================================================

		var sj []interface{}
		var fj interface{}
		//sj = append(sj, userID_, channelID_, chaincodeID_)  // ---> For money system, they(chaincode, channel and userID) might be fixed.

		err := json.Unmarshal([]byte(balance_query), &fj)
		if err != nil {
			fmt.Printf("Unmarshal return error: %v\n", err)
			return
		}
		sj = append(sj, fj)
		mj := make(map[string]interface{})
		mj["Args"] = sj

		mjBytes, err = json.Marshal(mj)
		if err != nil {
			fmt.Printf("Marshal return error %v\n", err)
			e = fmt.Errorf("Marshal return error %v\n", err)
			return nil, e
		}

		fmt.Printf("Request --> balanaceOf  [888] : Query --> Sign and Verification ......END\n")

	case 808: // new verion; tranferFrom2()) : type is changed from 808 to 8
		if len(os.Args) < 7 {
			fmt.Printf(`Ex)./%s 808 ip:port from to amount label`, VERSION)
			e = fmt.Errorf("args len ERROR")
			return nil, e
		}
		//    --> These need to be implanted into stargate with dockerArgumentsFile.

		from := os.Args[3]
		to := os.Args[4]
		amount := os.Args[5]
		label := os.Args[6]

		var privKey string
		if len(os.Args) >= 8 {
			privKey = os.Args[7]
		}

		result, _ := validateAddress(from)
		if !result {
			fmt.Printf("Please Try again,....\nThe given address is INVALID !!! : %s\n", from)
			e = fmt.Errorf("Please Try again,....\nThe given address is INVALID !!! : %s\n", from)
			return nil, e
		}

		result, _ = validateAddress(to)
		if !result {
			fmt.Printf("Please Try again,....\nThe given address is INVALID !!! : %s\n", to)
			e = fmt.Errorf("Please Try again,....\nThe given address is INVALID !!! : %s\n", from)
			return nil, e
		}

		transfer_args := fmt.Sprintf(`{"Args":["transferFrom",{"from":"%s", "to":"%s", "amount":"%s", "label":"%s"}]}`, from, to, amount, label)

		//args_.Args = append(args_.Args, os.Args[3], os.Args[4], os.Args[5], os.Args[6])
		//fmt.Printf("Args : %s", os.Args[6])

		var sj []interface{}
		var fj interface{}
		sj = append(sj, userID_, channelID_, chaincodeID_)

		err := json.Unmarshal([]byte(transfer_args), &fj)
		if err != nil {
			fmt.Printf("Unmarshal return error: %v\n", err)
			e = fmt.Errorf("Unmarshal return error: %v\n", err)
			return nil, e
		}
		sj = append(sj, fj)
		mj := make(map[string]interface{})
		mj["Args"] = sj

		mjBytes, err = json.Marshal(mj)
		if err != nil {
			fmt.Printf("Marshal return error %v\n", err)
			e = fmt.Errorf("Marshal return error %v\n", err)
			return nil, e
		}

		//////////////////////////////////////////////////////////////////////////////////////////////////
		///     Paul:
		///         Build a envelope including {pubkey, signature}
		//////////////////////////////////////////////////////////////////////////////////////////////////

		//attachment, err = buildEnvelopeForApprove(owner, spender, label, limit, expired)

		/*if myMarine.Mode == "CLI" {
			attachment, err = buildEnvelope("transferFrom", from, to, amount, label)
		} else {
			fmt.Printf("\nmode isn't CLI\n")
			attachment, err = buildEnvelope("transferFrom", from, to, amount, label, privKey)
		}*/

		attachment, err = buildEnvelope("transferFrom", from, to, amount, label, privKey)

		if err != nil {
			// ALready print out the error of "buildEnvelopeForApprove()"
			fmt.Printf("Error in buildEnvelope() : %v\n", err)
			e = fmt.Errorf("Error in buildEnvelope() : %v\n", err)
			return nil, e
		}

		/*		attachment, err = buildEnvelope("transferFrom", from, to, amount, label)

				if err != nil {
					fmt.Printf("Error : %v\n", err)
					// ALready print out the error of "buildEnvelopeForApprove()"
					return
				}*/

		///////////////////////////////////
		type_ = 8
		///////////////////////////////////

		/*
		   fmt.Printf("Testing,..............Done!!\n")
		   return
		*/
	case 818: //approve(), 818 --> 8 to keep the format as the INVOKE command (type : 8)
		if len(os.Args) < 9 {
			fmt.Printf("Ex)./marine type[818(8)] [serverAdress:port] owner spender label limit(positive/0/inf) expired(UTC in sec) (privatekey)\n")
			//               0                   1               2         3     4       5       6                   7                  (8)      (len : 8 or 9)
			return
		}

		///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
		owner := os.Args[3]
		spender := os.Args[4]
		label := os.Args[5]
		limit := os.Args[6]
		expired := os.Args[7]

		var privKey string
		if len(os.Args) >= 9 {
			privKey = os.Args[8]
			fmt.Printf("privateKey param: %v\n", privKey)
		}

		result, _ := validateAddress(owner)
		if !result {
			fmt.Printf("Please Try again,....\nThe given address is INVALID !!! : %s\n", owner)
			e = fmt.Errorf("Please Try again,....\nThe given address is INVALID !!! : %s\n", owner)
			return nil, e
		}

		result, _ = validateAddress(spender)
		if !result {
			fmt.Printf("Please Try again,....\nThe given address is INVALID !!! : %s\n", spender)
			e = fmt.Errorf("Please Try again,....\nThe given address is INVALID !!! : %s\n", spender)
			return nil, e
		}

		if owner == spender {
			fmt.Printf("Please Try again,....\n The given addresses (owner and spender) are the same to each other!!\n")
			e = fmt.Errorf("Please Try again,....\n The given addresses (owner and spender) are the same to each other!!\n")
			return nil, e
		}

		approve_args := fmt.Sprintf(`{"Args":["approve",{"owner":"%s", "spender":"%s", "label":"%s", "limit":"%s", "expired":"%s"}]}`, owner, spender, label, limit, expired)

		var sj []interface{}
		var fj interface{}
		sj = append(sj, userID_, channelID_, chaincodeID_)

		err := json.Unmarshal([]byte(approve_args), &fj)
		if err != nil {
			fmt.Printf("Unmarshal return error: %v\n", err)
			return
		}
		sj = append(sj, fj)
		mj := make(map[string]interface{})
		mj["Args"] = sj

		mjBytes, err = json.Marshal(mj)
		if err != nil {
			fmt.Printf("Marshal return error %v\n", err)
			e = fmt.Errorf("Marshal return error %v\n", err)
			return nil, e
		}

		/*if myMarine.Mode == "CLI" {
			attachment, err = buildEnvelope("approve", owner, spender, label, limit, expired)
		} else {
			attachment, err = buildEnvelope("approve", owner, spender, label, limit, expired, privKey)
		}*/

		attachment, err = buildEnvelope("approve", owner, spender, label, limit, expired, privKey)
		if err != nil {
			fmt.Printf("Error : %v\n", err)
			e = fmt.Errorf("Error : %v\n", err)
			// ALready print out the error of "buildEnvelopeForApprove()"
			return nil, e
		}

		///////////////////////////////////
		type_ = 8
		///////////////////////////////////

	case 828: //issue(), 828 --> 8 to keep the format as the INVOKE command (type : 8)
		if len(os.Args) < 6 {
			fmt.Printf("Ex)./%s type[828(->8)] [serverAdress:port] symbol, issuer totalsupply\n", VERSION)
			//               0        1               2              3       4        5      (total : 6)
			return
		}

		///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

		fmt.Printf("\n[ExecuteMarine] len: %v\n", len(os.Args))
		fmt.Printf("\n[ExecuteMarine] args: %v\n", os.Args)

		symbol := os.Args[3]
		issuer := os.Args[4]
		totalsupply := os.Args[5]
		//sj

		var privKey string
		if len(os.Args) >= 7 {
			privKey = os.Args[6]
			fmt.Printf("privateKey param: %v\n", privKey)
		}
		//

		result, _ := validateAddress(issuer)
		if !result {
			fmt.Printf("Please Try again,....\nThe given address(founder) is INVALID !!! : %s\n", issuer)
			e = fmt.Errorf("Please Try again,....\nThe given address(founder) is INVALID !!! : %s\n", issuer)
			return nil, e
		}

		if len(symbol) <= 1 {
			fmt.Printf("title MUST be longer than size(1) : %s\n", symbol)
			e = fmt.Errorf("title MUST be longer than size(1) : %s\n", symbol)
			return nil, e
		}

		issue_args := fmt.Sprintf(`{"Args":["issue",{"symbol":"%s", "issuer":"%s", "totalsupply":"%s"}]}`, symbol, issuer, totalsupply)
		//'{"Args":["approve",{"owner":"z12abcdef", "spender":"zlkjhgffd", "lable":"btc", "limit" :"1000000", "expired":""}]}'
		///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

		//args_.Args = append(args_.Args, os.Args[3], os.Args[4], os.Args[5], os.Args[6])
		//fmt.Printf("Args : %s", os.Args[6])

		var sj []interface{}
		var fj interface{}
		sj = append(sj, userID_, channelID_, chaincodeID_)

		err := json.Unmarshal([]byte(issue_args), &fj)
		if err != nil {
			fmt.Printf("Unmarshal return error: %v\n", err)
			e = fmt.Errorf("Unmarshal return error: %v\n", err)
			return nil, e
		}
		sj = append(sj, fj)
		mj := make(map[string]interface{})
		mj["Args"] = sj

		mjBytes, err = json.Marshal(mj)
		if err != nil {
			fmt.Printf("Marshal return error %v\n", err)
			e = fmt.Errorf("Marshal return error %v\n", err)
			return nil, e
		}

		//////////////////////////////////////////////////////////////////////////////////////////////////
		///     Paul:
		///         Build a envelope including {pubkey, signature}
		//////////////////////////////////////////////////////////////////////////////////////////////////
		/*if len(os.Args) >= 7 {
			fmt.Printf("privateKey !! os.Args >= 7: %v\n", privKey)
			attachment, err = buildEnvelope("issue", symbol, issuer, totalsupply, privKey)
		} else {

			attachment, err = buildEnvelope("issue", symbol, issuer, totalsupply)
		}*/

		attachment, err = buildEnvelope("issue", symbol, issuer, totalsupply, privKey)

		if err != nil {
			// ALready print out the error of "buildEnvelopeForApprove()"
			fmt.Printf("Error in buildEnvelope() : %v\n", err)
			e = fmt.Errorf("Error in buildEnvelope() : %v\n", err)
			return nil, e
		}

		///////////////////////////////////
		type_ = 8
		///////////////////////////////////

		/*
		   fmt.Printf("Testing,..............Done!!\n")
		   return
		*/

	case 721: //MintNFT(), 721 --> 8 to keep the format as the INVOKE command (type : 8)
		if len(os.Args) < 8 {
			fmt.Printf("Ex)./%s type[721(->8)] [serverAdress:port] cid  metaid  address  tokenName  tokenSymbol privkey(option)\n", VERSION)
			//               0        1               2             3     4       5       6           7          8      (total : 8 or 9)
			return
		}

		///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

		cid := os.Args[3]
		metaid := os.Args[4]
		address := os.Args[5]
		tokenName := os.Args[6]
		tokenSymbol := os.Args[7]

		var privKey string
		if len(os.Args) > 8 {
			privKey = os.Args[8]
			fmt.Printf("privateKey param: %v\n", privKey)
		}
		//

		result, _ := validateAddress(address)
		if !result {
			fmt.Printf("Please Try again,....\nThe given address(founder) is INVALID !!! : %s\n", address)
			return
		}

		if len(cid) != 46 || len(metaid) != 46 {
			fmt.Printf("CID(%s)/MetaID(%s) size must be greater than 46 bytes : %s\n", cid, metaid)
			return
		}

		mintNFT_args := fmt.Sprintf(`{"Args":["mintNFT",{"cid":"%s", "metaid":"%s", "address":"%s","tokenName":"%s","tokenSymbol":"%s"}]}`, cid, metaid, address, tokenName, tokenSymbol)
		//'{"Args":["approve",{"owner":"z12abcdef", "spender":"zlkjhgffd", "lable":"btc", "limit" :"1000000", "expired":""}]}'
		///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

		//args_.Args = append(args_.Args, os.Args[3], os.Args[4], os.Args[5], os.Args[6])
		//fmt.Printf("Args : %s", os.Args[6])

		var sj []interface{}
		var fj interface{}
		sj = append(sj, userID_, channelID_, chaincodeID_)

		err := json.Unmarshal([]byte(mintNFT_args), &fj)
		if err != nil {
			fmt.Printf("Unmarshal return error: %v\n", err)
			return
		}
		sj = append(sj, fj)
		mj := make(map[string]interface{})
		mj["Args"] = sj

		mjBytes, err = json.Marshal(mj)
		if err != nil {
			fmt.Printf("Marshal return error %v\n", err)
			return
		}

		//////////////////////////////////////////////////////////////////////////////////////////////////
		///     Paul:
		///         Build a envelope including {pubkey, signature}
		//////////////////////////////////////////////////////////////////////////////////////////////////
		/*if len(os.Args) > 8 {

			attachment, err = buildEnvelope("mintNFT", cid, metaid, address, tokenName, tokenSymbol, privKey)
		} else {

			attachment, err = buildEnvelope("mintNFT", cid, metaid, address, tokenName, tokenSymbol)
		}*/

		attachment, err = buildEnvelope("mintNFT", cid, metaid, address, tokenName, tokenSymbol, privKey)

		if err != nil {
			// ALready print out the error of "buildEnvelopeForApprove()"
			fmt.Printf("Error in buildEnvelope() : %v\n", err)
			return
		}

		///////////////////////////////////
		type_ = 8
		///////////////////////////////////

		/*
		   fmt.Printf("Testing,..............Done!!\n")
		   return
		*/

	default:

		fmt.Printf("Not defined yet!! Try again!!\n")
		return
	}

	var requestArgs []byte
	if pb.MessageType(type_) == pb.MessageType_QUERY_CHAINCODE || pb.MessageType(type_) == pb.MessageType_INVOKE_CHAINCODE || pb.MessageType(type_) == 21 || pb.MessageType(type_) == 888 || pb.MessageType(type_) == 808 || pb.MessageType(type_) == 899 || pb.MessageType(type_) == 900 {
		fmt.Printf("Q1 %s\n", mjBytes)
		requestArgs = mjBytes
	} else {
		requestArgs, err = json.Marshal(args_)
		if err != nil {
			fmt.Printf("Marshal return error: %v\n", err)
			return
		}
	}
	fmt.Printf("requestArgs --> %v\n", string(requestArgs))

	/*
	   var receivedArgs arguments
	   json.Unmarshal(requestArgs, &receivedArgs)
	   fmt.Printf("receivedArgs --> %v", receivedArgs)
	*/

	if pb.MessageType(type_) == pb.MessageType_INVOKE_CHAINCODE || pb.MessageType(type_) == 900 {
		resp, respErr = c.ProcessCOMM(context.Background(), &pb.ProcessRequest{RequestType: int32(type_), Args: string(requestArgs), Document: attachment})
		/*} else if pb.MessageType(type_) == 808 {// for old transferFrom()
		  resp, respErr = c.ProcessCOMM(context.Background(), &pb.ProcessRequest{RequestType: int32(type_), Args: string(requestArgs), Document: envelope[:]})
		*/
	} else {
		resp, respErr = c.ProcessCOMM(context.Background(), &pb.ProcessRequest{RequestType: int32(type_), Args: string(requestArgs)})
	}
	if respErr != nil {
		fmt.Printf("Cannot greet: %v %v\n", resp, respErr)
		return nil, respErr

	}
	if resp.Status == int32(pb.Status_ERROR) {
		fmt.Printf("Delivered Responses from Server --> %v\n", resp)
		return resp, respErr
	} else {
		fmt.Printf("\nDelivered Responses from Server --> Status : %d\n%s\n", resp.Status, string(resp.Payload))
	}

	if strings.Contains(string(requestArgs), "balanceOf") && myMarine.Mode == "CLI" {
		//skill.DisplayUTXOperLabel(string(resp.Payload))
		skill.DisplayUTXOperLabelwithCount(string(resp.Payload))
		//pay := fmt.Sprintf("%v", string(resp.Payload))

	}

	return resp, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////
//  End of Marine
////////////////////////////////////////////////////////////////////////////////////////////////////

func withdrawable(resp *pb.ProcessResponse, contract *contractDetails, secret string) bool {
	//if the requested contract doesn't exist
	if resp.Status == int32(pb.Status_ERROR) {
		fmt.Printf("Contract does not exist. \n\nDelivered Responses from Server --> Status : %d  %v", resp.Status, resp)
		return false
	}

	//get json
	err := json.Unmarshal(resp.Payload, &contract)
	if err != nil {
		fmt.Printf("Unmarshal return error: %v", err)
		return false
	}

	//if the requested contract is already settled
	if contract.Withdrawn == true {
		fmt.Printf("\n already withdrawn \n\n")
		return false
	}

	//if the hashlock hash does not match
	hashlock := sha256.Sum256([]byte(secret))
	if contract.Hashlock != fmt.Sprintf("%x", hashlock) {
		fmt.Printf("\n hash does not match \n\n")
		return false
	}

	//if the timelock time is not in the future
	timelock, _ := strconv.Atoi(contract.Timelock)
	if timelock > 0 && int64(timelock) < time.Now().Unix() {
		fmt.Printf("\n timelock time must be in the future \n\n")
		return false
	}

	return true
}

func refundable(resp *pb.ProcessResponse, contract *contractDetails) bool {
	//if the requested contract doesn't exist
	if resp.Status == int32(pb.Status_ERROR) {
		fmt.Printf("Contract does not exist. \n\nDelivered Responses from Server --> Status : %d  %v", resp.Status, resp)
		return false
	}

	//get json
	err := json.Unmarshal(resp.Payload, &contract)
	if err != nil {
		fmt.Printf("Unmarshal return error: %v", err)
		return false
	}

	//if the requested contract is already settled
	if contract.Withdrawn == true {
		fmt.Printf("\n already withdrawn \n\n")
		return false
	}

	//if the requested contract is already refunded
	if contract.Refunded == true {
		fmt.Printf("\n already refunded \n\n")
		return false
	}

	//if the timelock time is not passed
	timelock, _ := strconv.Atoi(contract.Timelock)
	if timelock > 0 && int64(timelock) > time.Now().Unix() {
		fmt.Printf("\n timelock not yet passed \n\n")
		return false
	}

	return true
}
