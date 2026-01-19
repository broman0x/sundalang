package sundalang

type TokenType string

const (
	TOKEN_EOF     TokenType = "EOF"
	TOKEN_ILLEGAL TokenType = "ILLEGAL"
	TOKEN_IDENT   TokenType = "IDENT"
	TOKEN_INT     TokenType = "INT"
	TOKEN_STRING  TokenType = "STRING"

	TOKEN_ASSIGN   TokenType = "="
	TOKEN_PLUS     TokenType = "+"
	TOKEN_MINUS    TokenType = "-"
	TOKEN_BANG     TokenType = "!"
	TOKEN_ASTERISK TokenType = "*"
	TOKEN_SLASH    TokenType = "/"
	TOKEN_MODULO   TokenType = "%"

	TOKEN_LT  TokenType = "<"
	TOKEN_GT  TokenType = ">"
	TOKEN_LTE TokenType = "<="
	TOKEN_GTE TokenType = ">="

	TOKEN_EQ     TokenType = "=="
	TOKEN_NOT_EQ TokenType = "!="

	TOKEN_AND TokenType = "&&"
	TOKEN_OR  TokenType = "||"

	TOKEN_COMMA     TokenType = ","
	TOKEN_SEMICOLON TokenType = ";"
	TOKEN_COLON     TokenType = ":"
	TOKEN_LPAREN    TokenType = "("
	TOKEN_RPAREN    TokenType = ")"
	TOKEN_LBRACE    TokenType = "{"
	TOKEN_RBRACE    TokenType = "}"
	TOKEN_LBRACKET  TokenType = "["
	TOKEN_RBRACKET  TokenType = "]"

	TOKEN_TANDA    TokenType = "TANDA"
	TOKEN_CETAK    TokenType = "CETAK"
	TOKEN_TANYA    TokenType = "TANYA"
	TOKEN_LAMUN    TokenType = "LAMUN"
	TOKEN_LAMUNTEU TokenType = "LAMUNTEU"
	TOKEN_KEDAP    TokenType = "KEDAP"
	TOKEN_BENER    TokenType = "BENER"
	TOKEN_SALAH    TokenType = "SALAH"
	TOKEN_FUNGSI   TokenType = "FUNGSI"
	TOKEN_BALIK    TokenType = "BALIK"
	TOKEN_EUREUN   TokenType = "EUREUN"
	TOKEN_EWEHAN   TokenType = "EWEHAN"
	TOKEN_TETEP    TokenType = "TETEP"
	TOKEN_PIKEUN   TokenType = "PIKEUN"
	TOKEN_MILIH    TokenType = "MILIH"
	TOKEN_KASUS    TokenType = "KASUS"
	TOKEN_BAKU     TokenType = "BAKU"
	TOKEN_BUKA               = "BUKA"
	TOKEN_WADAH              = "WADAH"
	TOKEN_COBAAN             = "COBAAN"
	TOKEN_SANYA              = "SANYA"
)

var keywords = map[string]TokenType{
	"tanda":     TOKEN_TANDA,
	"cetakkeun": TOKEN_CETAK,
	"tanyakeun": TOKEN_TANYA,
	"lamun":     TOKEN_LAMUN,
	"lamunteu":  TOKEN_LAMUNTEU,
	"kedap":     TOKEN_KEDAP,
	"bener":     TOKEN_BENER,
	"salah":     TOKEN_SALAH,
	"fungsi":    TOKEN_FUNGSI,
	"pungsi":    TOKEN_FUNGSI,
	"balik":     TOKEN_BALIK,
	"balikkeun": TOKEN_BALIK,
	"eureun":    TOKEN_EUREUN,
	"ewehan":    TOKEN_EWEHAN,
	"tetep":     TOKEN_TETEP,
	"pikeun":    TOKEN_PIKEUN,
	"milih":     TOKEN_MILIH,
	"kasus":     TOKEN_KASUS,
	"baku":      TOKEN_BAKU,
	"buka":      TOKEN_BUKA,
	"wadah":     TOKEN_WADAH,
	"cobaan":    TOKEN_COBAAN,
	"sanya":     TOKEN_SANYA,
}

type Token struct {
	Type    TokenType
	Literal string
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return TOKEN_IDENT
}
