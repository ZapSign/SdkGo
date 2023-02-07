
# SDK Golang API Zapsign

SDK das requisições de API na Zapsign em Golang.

## API Zapsign

 - [API](https://docs.zapsign.com.br/)

## Autores

- [@flubyGit](https://www.github.com/flubyGit)
- [@DouglasZapSign](https://www.github.com/DouglasZapSign)


## Funcionalidades

- Criar documento via upload
- Criar documento via modelo
- Adicionar documento extra a um documento
- Detalhar documento
- Listar documentos
- Excluir documento

- Detalhar signatário
- Excluir signatário
- Adicionar signatário
- Atualizar singatário
- Assinar em lote via API
## Rodando localmente

Clone o projeto

```bash
  git clone git@github.com:ZapSign/SdkGo.git
```

Entre no diretório do projeto

```bash
  cd SdkGo
```

Instale as dependências

```bash
  go get
```

Rode os testes

```bash
  go test ./tests/{{FOLDER_TEST}}/{{NOME_TESTE}}
```


## Rodando os testes

Para rodar os testes, rode o seguinte comando

```bash
  go test ./tests/{{FOLDER_TEST}}/{{NOME_TESTE}}
```

Exemplo:
```bash
  go test ./tests/docs/get_docs_test.go
```
## Feedback

Se você tiver algum feedback, por favor nos deixe saber por meio de:
- felipe@zapsign.com.br 
- andre@zapsign.com.br 
- douglas@zapsign.com.br

