#language: pt
#encoding: iso-8859-1
#author: msalcantara
#Version 1.0
Funcionalidade: Validar resposta da API do github

  @ID12345
  Cenario: Validar resposta 200 ao consultar os repositorios na API do github
    Dado que realizei um request para o end point "https://api.github.com/users/msalcantara/repos"
    Entao o response da API deve ser "200"
