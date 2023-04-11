# BR Validate
<a href="https://pkg.go.dev/github.com/walteravelino/brvalidate/">
  <img alt="GoMod" src="https://img.shields.io/github/go-mod/go-version/walteravelino/brvalidate">
</a>

[![Build Status](https://travis-ci.com/walteravelino/Projetos.svg?branch=master)](https://travis-ci.com/walteravelino/Projetos)
<img src = "https://img.shields.io/github/languages/top/walteravelino/brvalidate">
<a href="https://github.com/walteravelino/Projetos/blob/master/LICENSE"><img src = "https://img.shields.io/github/license/walteravelino/Projetos"></a>


## Autor

👤 **Walter Avelino**

- StackOverFlow [@walteravelino](https://stackoverflow.com/users/13001807/walter-avelino)
- Github: [@walteravelino](https://github.com/walteravelino)
- DEV: [@walteravelino](https://dev.to/walteravelino)


## 📝 Licença

Copyright © 2023 [Walter Avelino](https://github.com/walteravelino). <br />
Os projetos estão sob a licença [MIT](https://github.com/walteravelino/Projetos/blob/master/LICENSE).

Pacote para validação de documentos

## Instalação do BR Validate

Para instalar o pacote:

```bash
go get github.com/walteravelino/brvalidate
```

# Funções de Validação
### Documentos
#### Pessoa
- `func person.CPF(value string) bool`
- `func person.CNPJ(value string) bool`
- `func person.CNH(value string) bool`
#### Veículo
- `func vehicle.Renavan(value string) bool`
- `func vehicle.Plate(value string) bool`
### Extras
#### EBCT
- `func extras.CEP(value string, ufs ...FederativeUnit) bool`

------------
## Validação
Valida o documento passado como argumento. Retorna um `bool`, `True` caso seja válido, `False` caso contrário . Recebe os parâmetros:

| Parâmetro | Tipo | Valor Padrão | Requerido | Descrição |
| --------- | ---- | ----------- | ------------ | --------- |
| `value` | `str`| `''` | X | O documento a ser validado. |

```go
import (
"github.com/walteravelino/brvalidate/person"
)

cpf = CPF()

// Validar CPF
person.CPF("01234567890") // True
person.CPF("01234567891") // False
```

### Caso Especial de CPF
Os CPFs de 000.000.000-00 até 999.999.999-99 são considerados válidos pois, em alguns casos, existem pessoas vinculadas a eles. Foi utilizada a base de dados [Coleção de CNPJs e CPFs brasileiros do Brasil.IO](https://brasil.io/dataset/documentos-brasil/documents) para verificar esses documentos:

| CPF | Pessoa | Consulta |
| --- | ------ | -------- |
| 000.000.000-00 | - | `https://brasil.io/dataset/documentos-brasil/documents?search=00000000000&document_type=CPF&document=&name=&sources=` |
| 111.111.111-11 | AKA CENTRAL PARK - NEW YORK | `https://brasil.io/dataset/documentos-brasil/documents?search=11111111111&document_type=CPF&document=&name=&sources=` |
| 222.222.222-22 | - | `https://brasil.io/dataset/documentos-brasil/documents?search=22222222222&document_type=CPF&document=&name=&sources=` |
| 333.333.333-33 | - | `https://brasil.io/dataset/documentos-brasil/documents?search=33333333333&document_type=CPF&document=&name=&sources=` |
| 444.444.444-44 | - | `https://brasil.io/dataset/documentos-brasil/documents?search=44444444444&document_type=CPF&document=&name=&sources=` |
| 555.555.555-55 | ISAEL HERMINIO DA SILVA | `https://brasil.io/dataset/documentos-brasil/documents?search=55555555555&document_type=CPF&document=&name=&sources=` |
| 666.666.666-66 | - | `https://brasil.io/dataset/documentos-brasil/documents?search=66666666666&document_type=CPF&document=&name=&sources=` |
| 777.777.777-77 | ANTONIO GONÇALO DA SILVA | `https://brasil.io/dataset/documentos-brasil/documents?search=77777777777&document_type=CPF&document=&name=&sources=` |
| 888.888.888-88 | - | `https://brasil.io/dataset/documentos-brasil/documents?search=88888888888&document_type=CPF&document=&name=&sources=` |
| 999.999.999-99 | JOAQUIM ROCHA MATOS | `https://brasil.io/dataset/documentos-brasil/documents?search=99999999999&document_type=CPF&document=&name=&sources=` |