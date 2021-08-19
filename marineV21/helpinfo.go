package marine

import (
	"fmt"
	"os"

	color "github.com/mitchellh/colorstring"
	//colorstring.Println("[blue]Hello [red]World!")
)

func Help() {

	color.Printf(`
        
        [green]
        ********************************************************************
        ********************************************************************        
        	COMMAND COLLECTION FOR MARINE
        ********************************************************************
        ********************************************************************

        "CHANNEL_LIST":                 0,
        "CHANNEL_CONFIG":               1,
        "BLOCKCHAIN_INFO":              2,
        "BLOCK_INFO_NUMBER":            3,
        "TRANSACTION_INFO":             4,
        "INSTALLED_CHAINCODE_LIST":     5,
        "INSTANTIATED_CHAINCODE_LIST":  6,
        "QUERY_CHAINCODE":              7,
        "INVOKE_CHAINCODE":             8,
            \--> {create|query|update|delete|queryByPass|queryAttachment|queryByKey|queryHistoryByKey}
        "COMMIT_STATUS":                9,
        "GET_NOTI_SERVER":              10,     
        "SET_NOTI_SERVER":              11,     
        "HEALTH_CHECK_FOR_PEERS":       12,     
        "GET_TPS_INFO":                 13,     
        "GET_NODE_INFO":                20,     
        "SET_NODE_INFO":                21,     
        
        [yellow]
        ***************************************************
        	AUTHENTICATION CONTROL ENGINE (ACE)
        ***************************************************
        "ESTABLISH (CREATE)":           101, (~>8)
        "JOIN":                         102, (~>8)
        "LEAVE":                        103, (~>8)
        "MODIFY":                       104, (~>8), only for "type"=[public(0), private(1), memberOnly(2), left(4)]
        "STATUS":                       105, (~>8), partial info with given title and condition
        "TITLE":                        106, (~>8), all info with given title

        [white]
        ***************************************************
        	Single Sign over Blockchain (SSB)
        ***************************************************
        "LOGIN"                         108, (~>8), 
        "LOGOUT"                        108, (~>8), FLIPFLOP --> {login|logout}
        "STAUS"                         (undergoing) --> the permissioned can check the status, based on "ACE"

        [yellow]
        ***************************************************
        	ACE CLOACKING - 111
        ***************************************************
        "putPrivate"                    111, (~8),({..."putPrivate",..........}@json)
        "getPrivate"                    111, (~8),({..."getPrivate",byKey.....}@json)
           \--> "queryPrivate"          111, (~8),({..."queryPrivate",........}@json)  -- undergoing
           \--> "getPrivateHistory"     111, (~8),({..."getPrivateHistory",...}@json)  -- undergoing

        [white]
        ***************************************************
        	KEY & ID
        ***************************************************
        "GENERATE KEY&ADDRESS":         88,
        "OBTAIN ID/ADDRESS":            87, get ID(1)/Address(2) from pubKey
        "SIGN"                          871
        "VERIFY"                        872
        "CHECK Pubkey  TYPE"            873 (0:None 1:ED 2:ECC 3:RSA)
        "CHECK Privkey TYPE"            874 (0:None 1:ED 2:ECC 3:RSA)
        "Extract Pubkey from PrivKey"   8741 (support only Ed25519)
        "Extract {R,S} from sig"        875

        [yellow]
        ***************************************************
        	Cross Assets Transfer (CAT)
        ***************************************************
        "BLANACEOF":                    888,
        "TRANSFERFROM":                 808,
        "APPROVE":                      818 (~>8)
        "ALLOWANCE":                    7   (use "query" as Admin)
        "QUERY_UTXOINFO":               899,
        "ISSUE_":                       8   (Deprecated!! only for test and admin, please DO NOT USE ME)
        "ISSUE":                        828 (~>8), official version
        "MINT":                         868 (~>8), Done!!
        "BURN":                         848 (~>8), Done!!
        ----------------------->
        "LOCKIN"                        877 (~>8) -- undergoing
        "RELEASE"                       878 (~>8) -- undergoing

        [red]
        ***************************************************
        	FEE POLICY (under construction,....)
        ***************************************************
        "CHARGEFEE"                     807,(~>8) --> NEVER TRY ME!, ONLY FOR SYSTEM!
        "PUTFEE"                        --> 111
        "GETFEE"                        --> 111

        [white]
        ***************************************************
        	eXchange (under construction,....)
        ***************************************************
        "registerXindex":                600 -- Done.

        "addXrate":                      601 -- Done.
        "updateXrate":                   602 -- Done.
        "removeXrate":                   604 -- Not implemented, wrapping from "updateXrate()"

        "queryXrate":                    7 (using QUERY_CHAINCODE)                            
                                            \-- (index, label, std.....)
                                            \-- By index
                                            \-- By label


        "eXchange":                   	 626  (transferFrom() adding events to ICBM)
                                                - Not implemented yet

        ***************************************************
            DEX (under construction,....)
        ***************************************************
        DEPOSIT :                        900 
        WITHDRAW :                       910 (~> 900)
        REFUND :                         920 (~> 900)


        [yellow]
        ***************************************************
        	"KEY-CHAINING" (undergoing):                 
        ***************************************************
        "KEY-CHAIN"                     [KEY-KEY..KEY],[KEY-KEY],[KEY]

        [white]
        ***************************************************
        	Converting Time
        ***************************************************
        "DATE  --> EPOCH":              28,
        "EPOCH --> DATE":               82
        ***************************************************

        [yellow]
        ***************************************************
       		ENCRYPT/DECRYPT (Asymmetric)
        ***************************************************
        "encrypt":                      201, (undergoing,....)
        "decrypt":                      202, (undergoing,....)

        
        ***************************************************
        	ENCRYPT/DECRYPT (Symmetric/PKCS7)
        ***************************************************
        "encrypt":                      301, 
        "decrypt":                      302, 

        [blue]
        ***************************************************
        	BENCHMARCH - PERFORMANCE ON KEY
        ***************************************************
        KEY PERFORMANCE TEST:           1000, (Case/N/KeyType) -- Now, only working for KeyGeneration!!
        
        ***************************************************
            SYSTEM CALL
        ***************************************************
        UUID :                          68
        getCreator :                    7/8 (using QUERY_CHAINCODE)  


`)

	tls := os.Getenv("TLS")
	if tls != "true" && tls != "false" {
		fmt.Printf("TLS ENV format error\n")
		return
	}

	fmt.Printf("\n\n\t********************************************************************\n\tASSET_CHANNEL: \t%s\n", myMarine.TokenChannel)
	fmt.Printf("\tTLS:\t\t%v\n", tls)
	fmt.Printf("\tASSET_CC: \t%s\n\t********************************************************************\n", myMarine.TokenCC)

}
