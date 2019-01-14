# juridigo_server
Servidor


# Obtenção e inicialização do projeto

```
cd go/src/github.com/
mkdir juridigo
cd juridigo
git clone https://github.com/andre250/juridigo_api.git
cd juridigo_api/
dep ensure 
gin -i main.go
```

# Endpoints

## Registros
```
http://.../user/register

METHOD: POST

Descrição: Registra um usuário na plataforma

Body: 
{
	"credenciais":{
		"credencial":[ String | 
                    tipo = 0 - Hash256 de usuario:senha,
                    tipo = 1 - Facebook Access Token
                ],
	        "tipo": [ Uint | 
                    0 - Registro comum
                    1 - Facebook
                ],
                "recuperacaoLogin": [
                    Hash256 de usuario:senha - Tipo = 1
                ]
		
	},
	"cadastrais": {
		"nome": [ Strign | Nome do usuario ],
		"dataNascimento": [ String | Data de nascimento no formato 1997-05-29T00:00:00Z],
		"email": [ String | Email do usuario],
		"telefone": [ String | Telefone do usuario],
		"rg": [ String | Rg do usuario],
		"cpf": [ String | Cpf do usuario],
		"cep": [ String | Cep da residencia ],
		"cidade": [ String | Cidade],
		"bairro": [ String | Bairro ],
		"rua": [ String | (Rua|Av|Viela|etc) Endereço],
		"numero": [ String | Número da residência],
		"complemento": [ String | Complemento ],
		"longitude": [ Float64 | Valor da longitude],
		"latitude": [ Float64 | Valor da latitude]
	},
	"curriculares": {
		"formacao":[{
			"escolaridade": [ UINT | 
                            0-Sem Escolaridade, 
                            1-Ensino Básico, 
                            2-Ensino Fundamental, 
                            3-Ensino Superior 
                        ],
			"instituicao": [ String | Nome da instituição ],
			"anoConclusao": [ Strign | Ano por extenso de conclusão Ex. 2018]
		}],
		"oab":[ String | Número de OAB],
		"curriculum": [ String | Base64 do CV em PDF]
	},
	"pagamento": [ String | Criptografia CFB de {
            "numero": [String],
            "banco": [String],
            "agencia": [String],
            "conta": [String],
            "numero":[String],
            "cvv":[String],
            "anoVencimento":[String],
            "mesVencimento":[String]
        }]
}
```
```
http://.../user/facebook

METHOD: POST

Descrição: Obtem dados do usuário via Facebook

Body: 
{
	"credencial": [ String | Facebook Token Access ]
}
```

## Autenticações
```
http://.../auth/login

METHOD: POST

Descrição: Autenticação de usuario via usuario e senha

Body: 
{
	"credencial": [ String | Hash256 de usuario:senha]
}
```
```
http://.../auth/login/facebook

METHOD: POST

Descrição: Autenticação de usuario via facebook access token

Body: 
{
	"credencial": [ String | Facebook Token Access ]
}
```
```
http://.../auth/login/facebook/recovery
METHOD: POST

Body: 
```
