
package marine

const (
    ADMIN_USER = "Admin"
    VERSION = "MARINE VER-21.3"
)

var (

    address         = "localhost:50061"
    TLS             = false
    certFILE        = ""
   //org_       string
    userID_      string
    channelID_   string
    chaincodeID_ string
    index_       string
    args_        arguments
    mjBytes      []byte

    type_ int
    seed  string

    attachment  []byte

    //sig         []byte


)

type Marine struct{
    Mode            string

    UserID          string
    TokenChannel    string
    TokenCC         string
}

type arguments struct {
    Args []string
}

type contractDetails struct {
    DocType string `json:"doctype"`
    From      string `json:"from"`
    To      string `json:"to"`
    Amount string `json:"amount"`
    Label string `json:"label"`
    Hashlock string `json:"hashlock"`
    Timelock string `json:"timelock"`
    Withdrawn bool `json:"withdrawn"`
    Refunded bool `json:"refunded"`
}



