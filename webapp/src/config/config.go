package config // pacote para abrir a conexão com o banco de dados. A conexão vai depender de variaveis de ambiente. As variaveis vao depedenr do pacote config
import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	APIURL   = ""
	Porta    = 0
	HashKey  []byte
	BlockKey []byte
)

// Carregar vai inicializar as variaveis de ambiente. Vai mexer com variaveis que estarão foras dela
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil { // dotenv é o pacote utilizado para ler arquivos env e colocar dentro do ambiente
		log.Fatal(erro) // neste caso se tiver algum problema para carregar as variaveis de ambeinte, não tem como a API rodar e por isso tem que parar a aplicação mesmo
	}

	Porta, erro = strconv.Atoi(os.Getenv("APP_PORT"))
	if erro != nil {
		log.Fatal(erro)
	}
	APIURL = os.Getenv("API_URL")

	HashKey = []byte(os.Getenv("HASH_KEY"))
	BlockKey = []byte(os.Getenv("BLOCK_KEY"))

	// StringConexaoBanco = fmt.Sprintf("%s:%s@tcp(localhost:%v)/%s?charset=utf8&parseTime=True&loc=Local",
	// 	os.Getenv("DB_USUARIO"),
	// 	os.Getenv("DB_SENHA"),
	// 	Container,
	// 	os.Getenv("DB_NOME"),
	// )

}
