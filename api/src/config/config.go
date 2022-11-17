package config // pacote para abrir a conexão com o banco de dados. A conexão vai depender de variaveis de ambiente. As variaveis vao depedenr do pacote config
import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	// StringConexaoBanco é a string de conexão com o mysql
	StringConexaoBanco = ""
	// Porta onde a aAPI vai rodar
	Porta = 0
	// Container indica a porta do container do banco de dados
	Container = 0
	// SecretKey é a chave que vai ser usada para assinar o token
	SecretKey []byte
)

// Carregar vai inicializar as variaveis de ambiente. Vai mexer com variaveis que estarão foras dela
func Carregar() {
	var erro error

	if erro = godotenv.Load(); erro != nil { // dotenv é o pacote utilizado para ler arquivos env e colocar dentro do ambiente
		log.Fatal(erro) // neste caso se tiver algum problema para carregar as variaveis de ambeinte, não tem como a API rodar e por isso tem que parar a aplicação mesmo
	}

	Porta, erro = strconv.Atoi(os.Getenv("API_PORT")) // Basicamente, isso é um Parseint, pois eu preciso primeiramente colocar a nossa string da porta convertendo para inteiro
	if erro != nil {
		Porta = 9000
	}
	Container, erro = strconv.Atoi(os.Getenv("CONTAINER_PORT"))
	if erro != nil {
		Container = 3307
	}

	StringConexaoBanco = fmt.Sprintf("%s:%s@tcp(localhost:%v)/%s?charset=utf8&parseTime=True&loc=Local",
		os.Getenv("DB_USUARIO"),
		os.Getenv("DB_SENHA"),
		Container,
		os.Getenv("DB_NOME"),
	)

	SecretKey = []byte(os.Getenv("SECRET_KEY"))

}