package rotas

import (
	"net/http"
	"webapp/src/controllers"
)

// rotaPaginaPrincipal vai redirecionar rota
var rotaPaginaPrincipal = Rota{
	URI:                "/home",
	Metodo:             http.MethodGet,
	Funcao:             controllers.CarregarPaginaPrincipal,
	RequerAutenticacao: true,
}
