package skill

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"reflect"
	"strings"
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
	"github.com/hyperledger/fabric/bccsp"
	"github.com/hyperledger/fabric/bccsp/factory"
)

type UTXO struct {
	Address string   `json:"address,omitempty"`
	Balance *big.Rat `json:"balance,omitempty"`
	Label   string   `json:"label,omitempty"`
}
type UTXOs struct {
	Key string `json:"Key"`
	//Value   interface{}   `json:"Value"`
	Value UTXO `json:"Value"`
}

type combined struct {
	Amount *big.Rat `json:"Amount"`
	Count  int      `json:"Count"`
}

// Paul :

// 		Given a source, recursively retrieves every leaf node in a struct in depth-first fashion
// 		and aggregate the results into given string slice with format: "path.to.leaf = value"
// 		in the order of definition.
//
// Take a look on the example below
// A{
//   B{
//     C: "foo",
//     D: 42,
//   },
//   E: nil,
// }
// it should yield a slice of string containing following items:
// [
//   "B.C = \"foo\"",
//   "B.D = 42",
//   "E =",
// ]

func Flatten(i interface{}) []string {
	var res []string
	flatter("", &res, reflect.ValueOf(i))
	return res
}

const DELIMITER = "."

func flatter(k string, m *[]string, v reflect.Value) {
	delimiter := DELIMITER
	if k == "" {
		delimiter = ""
	}

	switch v.Kind() {
	case reflect.Ptr:
		if v.IsNil() {
			*m = append(*m, fmt.Sprintf("%s =", k))
			return
		}
		flatter(k, m, v.Elem())
	case reflect.Struct:
		if x, ok := v.Interface().(fmt.Stringer); ok {
			*m = append(*m, fmt.Sprintf("%s = %v", k, x))
			return
		}

		for i := 0; i < v.NumField(); i++ {
			flatter(k+delimiter+v.Type().Field(i).Name, m, v.Field(i))
		}
	case reflect.String:
		// It is useful to quote string values
		*m = append(*m, fmt.Sprintf("%s = \"%s\"", k, v))
	default:
		*m = append(*m, fmt.Sprintf("%s = %v", k, v))
	}
}

type alg struct {
	hashFun func([]byte) string
}

const defaultAlg = "sha256"

var availableIDgenAlgs = map[string]alg{
	defaultAlg: {GenerateIDfromTxSHAHash},
}

// ComputeSHA256 returns SHA2-256 on data
func ComputeSHA256(data []byte) (hash []byte) {
	hash, err := factory.GetDefault().Hash(data, &bccsp.SHA256Opts{})
	if err != nil {
		panic(fmt.Errorf("Failed computing SHA256 on [% x]", data))
	}
	return
}

// ComputeSHA3256 returns SHA3-256 on data
func ComputeSHA3256(data []byte) (hash []byte) {
	hash, err := factory.GetDefault().Hash(data, &bccsp.SHA3_256Opts{})
	if err != nil {
		panic(fmt.Errorf("Failed computing SHA3_256 on [% x]", data))
	}
	return
}

// GenerateBytesUUID returns a UUID based on RFC 4122 returning the generated bytes
func GenerateBytesUUID() []byte {
	uuid := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, uuid)
	if err != nil {
		panic(fmt.Sprintf("Error generating UUID: %s", err))
	}

	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80

	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40

	return uuid
}

// GenerateIntUUID returns a UUID based on RFC 4122 returning a big.Int
func GenerateIntUUID() *big.Int {
	uuid := GenerateBytesUUID()
	z := big.NewInt(0)
	return z.SetBytes(uuid)
}

// GenerateUUID returns a UUID based on RFC 4122
func GenerateUUID() string {
	uuid := GenerateBytesUUID()
	return idBytesToStr(uuid)
}

func idBytesToStr(id []byte) string {
	return fmt.Sprintf("%x-%x-%x-%x-%x", id[0:4], id[4:6], id[6:8], id[8:10], id[10:])
}

// CreateUtcTimestamp returns a google/protobuf/Timestamp in UTC
func CreateUtcTimestamp() *timestamp.Timestamp {
	now := time.Now().UTC()
	secs := now.Unix()
	nanos := int32(now.UnixNano() - (secs * 1000000000))
	return &(timestamp.Timestamp{Seconds: secs, Nanos: nanos})
}

//GenerateHashFromSignature returns a hash of the combined parameters
func GenerateHashFromSignature(path string, args []byte) []byte {
	return ComputeSHA256(args)
}

// GenerateIDfromTxSHAHash generates SHA256 hash using Tx payload
func GenerateIDfromTxSHAHash(payload []byte) string {
	return fmt.Sprintf("%x", ComputeSHA256(payload))
}

// GenerateIDWithAlg generates an ID using a custom algorithm
func GenerateIDWithAlg(customIDgenAlg string, payload []byte) (string, error) {
	if customIDgenAlg == "" {
		customIDgenAlg = defaultAlg
	}
	var alg = availableIDgenAlgs[customIDgenAlg]
	if alg.hashFun != nil {
		return alg.hashFun(payload), nil
	}
	return "", fmt.Errorf("Wrong ID generation algorithm was given: %s", customIDgenAlg)
}

// FindMissingElements identifies the elements of the first slice that are not present in the second
// The second slice is expected to be a subset of the first slice
func FindMissingElements(all []string, some []string) (delta []string) {
all:
	for _, v1 := range all {
		for _, v2 := range some {
			if strings.Compare(v1, v2) == 0 {
				continue all
			}
		}
		delta = append(delta, v1)
	}
	return
}

func UTXOperLabel(input string) (symbol []string, balance []string, err error) {

	input = strings.ReplaceAll(input, `\`, `\\`)

	//var checks []interface{}
	var checks []UTXOs
	json.Unmarshal([]byte(input), &checks)

	if len(checks) <= 0 {
		err = fmt.Errorf("Not Json type, or No Contents")
		return symbol, balance, err
	}

	var myWallet map[string]*big.Rat
	myWallet = make(map[string]*big.Rat)

	for _, each := range checks {

		if balance, exists := myWallet[each.Value.Label]; !exists {
			myWallet[each.Value.Label] = each.Value.Balance
		} else {
			balance.Add(balance, each.Value.Balance)
			myWallet[each.Value.Label] = balance
		}
	}

	fmt.Println("\n----------------------------------")
	fmt.Println("  Label |    Balance(Sum)        |")
	fmt.Println("----------------------------------")

	for key, val := range myWallet {
		if key != "" {
			fmt.Printf("   %s  |   %s\n", key, val.FloatString(5))
			symbol = append(symbol, key)
			balance = append(balance, val.FloatString(5))
		}

	}

	return symbol, balance, nil
}

func DisplayUTXOperLabel(input string) bool {

	input = strings.ReplaceAll(input, `\`, `\\`)

	//var checks []interface{}
	var checks []UTXOs
	json.Unmarshal([]byte(input), &checks)

	if len(checks) <= 0 {
		fmt.Println("Not Json type, or No Contents")
		return false
	}

	var myWallet map[string]*big.Rat
	myWallet = make(map[string]*big.Rat)

	for _, each := range checks {

		if balance, exists := myWallet[each.Value.Label]; !exists {
			myWallet[each.Value.Label] = each.Value.Balance
		} else {
			balance.Add(balance, each.Value.Balance)
			myWallet[each.Value.Label] = balance
		}
	}

	fmt.Println("\n----------------------------------")
	fmt.Println("  Label |    Balance(Sum)        |")
	fmt.Println("----------------------------------")

	for key, val := range myWallet {
		if key != "" {
			fmt.Printf("   %s  |   %s\n", key, val.FloatString(5))
		}

	}

	return true
}

func DisplayUTXOperLabelwithCount(input string) bool {

	input = strings.ReplaceAll(input, `\`, `\\`)

	//var checks []interface{}
	var checks []UTXOs
	json.Unmarshal([]byte(input), &checks)

	if len(checks) <= 0 {
		fmt.Println("Not Json type, or No Contents")
		return false
	}

	var myWallet map[string]combined
	myWallet = make(map[string]combined)
	//var combined_ combined
	total := 0
	total_checks := 0
	for _, each := range checks {

		//fmt.Printf("%v\n", each)

		if combined_, exists := myWallet[each.Value.Label]; !exists {
			combined_.Amount = each.Value.Balance
			combined_.Count = 1
			myWallet[each.Value.Label] = combined_
		} else if each.Value.Balance != nil{
			combined_.Amount.Add(combined_.Amount, each.Value.Balance)
			combined_.Count++
			myWallet[each.Value.Label] = combined_
		}

		total_checks++
	}
	fmt.Println("\n\n===============================================")
	fmt.Println("\t\tMY WALLET")
	fmt.Println("----------------------------------------------")
	fmt.Println("  Label |    Balance(Sum)        |\tCount ")
	fmt.Println("===============================================")

	for key, val := range myWallet {
		if key != "" {
			fmt.Printf("   %s  |   %s\t|\t%d\n", key, val.Amount.FloatString(5), val.Count)
			total = total + val.Count
		}
	}
	fmt.Println("----------------------------------------------")
	fmt.Printf("\tTotal(%d) = FT(%d) + NFT(%d) \n", total_checks, total, (total_checks-total))
	fmt.Println("===============================================\n\n")

	return true
}