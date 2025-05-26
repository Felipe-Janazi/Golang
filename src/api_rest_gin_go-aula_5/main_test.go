package main

import (
	"api_rest_gin/controllers"
	"api_rest_gin/database"
	"api_rest_gin/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID int

func SetupDasRotasDeTeste() *gin.Engine {
	// Apresentar os testes de uma forma mais reduzida, sem detalhar tanto cada um
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

func CriaAlunoMock()  {
	aluno := models.Aluno{Nome: "Nome do aluno teste", CPF: "12345678901", RG: "123456789"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)

}

func DeletaAlunoMock()  {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}
func TestVerificaStatusCodeDaSaudacao(t *testing.T) {

	r := SetupDasRotasDeTeste()
	r.GET("/:nome", controllers.Saudacao)

	req, _ := http.NewRequest("GET", "/felipe", nil)

	resposta := httptest.NewRecorder()
	// Onde guardar e qual a requisição
	r.ServeHTTP(resposta, req)

	// O que vamos testar, o status desejado e o status que veio e por último a mensagem de erro
	assert.Equal(t, http.StatusOK, resposta.Code, "Deveriam ser iguais")

	mockDaResposta := `{"API diz:":"E ai felipe, tudo beleza?"}`

	// Pegando o corpo da requisição
	respostaBody, _ := ioutil.ReadAll(resposta.Body)

	// Fazendo a comparação do que deveria vir e o que veio no body, que vem em bites e precisamos converter em string
	assert.Equal(t, mockDaResposta, string(respostaBody))
}


func TestListandoTodosOsAlunosHandler(t *testing.T){
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)

	req, _ := http.NewRequest("GET", "/aluno", nil)
	resposta := httptest.NewRecorder()

	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorCPFHandler(t *testing.T)  {
	
	database.ConectaComBancoDeDados()
	//  Cria o aluno e depois de fazer tudo necessário ele deleta o aluno criado
	CriaAlunoMock()
	defer DeletaAlunoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)

	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678901", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestBuscaAlunoPorIDHandler(t *testing.T)  {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)

	// CDonvertendo um inteiro pra string
	pathDaBusca := "/alunos/" + strconv.Itoa(ID)
	req, _ :=  http.NewRequest("GET", pathDaBusca, nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var alunoMock models.Aluno

	// Aqui pegamos o corpo da requisição transformada em bytes e depois convertida para json e estamos armazenando dentro do alunoMock
	json.Unmarshal(resposta.Body.Bytes(), &alunoMock)

	fmt.Println(alunoMock.Nome)

	assert.Equal(t, "Nome do aluno teste", alunoMock.Nome)
	assert.Equal(t, "12345678901", alunoMock.CPF)
	assert.Equal(t, "123456789", alunoMock.RG)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestDeletaAlunoHandler(t *testing.T)  {
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)

	pathDeBusca := "/alunos/" + strconv.Itoa(ID)

	req, _ := http.NewRequest("DELETE", pathDeBusca,  nil)

	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	assert.Equal(t, http.StatusOK, resposta.Code)
}


func TestEditaUmAlunoHandler(t *testing.T)  {
	
	database.ConectaComBancoDeDados()
	CriaAlunoMock()
	defer DeletaAlunoMock()
	r := SetupDasRotasDeTeste()
	r.PATCH("/alunos/:id", controllers.EditaAluno)

	aluno := models.Aluno{Nome: "Nome do Aluno Teste", CPF: "12345678901", RG: "123456788"}

	valorJson, _ := json.Marshal(aluno)

	pathParaEditar := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathParaEditar, bytes.NewBuffer(valorJson))

	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var alunoMockAtualizado models.Aluno

	json.Unmarshal(resposta.Body.Bytes(), &alunoMockAtualizado)

	assert.Equal(t, "Nome do Aluno Teste", alunoMockAtualizado.Nome)
	assert.Equal(t, "12345678901", alunoMockAtualizado.CPF)
	assert.Equal(t, "123456789", alunoMockAtualizado.RG)
}