# Stress CLI (desafio-goexpert-2)

Uma ferramenta de linha de comando simples para realizar testes de carga (stress test) em um endpoint HTTP. Ela envia múltiplas requisições GET de forma concorrente para uma URL informada e, ao final, exibe um pequeno relatório com a duração total e a contagem de respostas por código HTTP.

Principais características:
- Envia N requisições HTTP GET para uma URL alvo.
- Controla o nível de concorrência (quantos workers simultâneos).
- Tempo limite (timeout) de 5 segundos por requisição.
- Relatório final com duração total e contagem por status HTTP.


## Requisitos
- Docker instalado.


## Como usar com Docker

A imagem é construída a partir do Dockerfile do projeto e o binário final é exposto como `stress-cli` (ENTRYPOINT).

1) Construir a imagem localmente na raiz do projeto:

```
docker build -t stress-cli:latest .
```

2) Executar o contêiner passando os parâmetros obrigatórios:

```
docker run --rm stress-cli:latest \
  -u http://example.com\
  -r 1000 \
  -c 10
```

- `-u, --url` URL alvo (deve incluir o protocolo http/https)
- `-r, --requests` Quantidade total de requisições a enviar
- `-c, --concurrency` Número de workers concorrentes

Exemplo usando um endpoint público de teste:

```
docker run --rm stress-cli:latest \
  -u https://httpbin.org/status/200 \
  -r 50 \
  -c 10
```

Saída típica:

```
Testing URL: https://httpbin.org/status/200
Requests: 50
Concurrency: 10
⠋ Stress testing...
✔ Stress testing completed!

-----RESULTS-----
Duration: 1.234567s
Total requests: 50
HTTP 200 - 50
```

Observações importantes:
- O comando executa apenas requisições GET.
- Se houver falha de conexão/timeout, essas tentativas são contadas sob o "status" 0 e aparecerão como `HTTP 0 - X` no relatório.
- O timeout por requisição é de 5 segundos.


## Execução direta (opcional, sem Docker)
Se preferir executar localmente (requer Go >= 1.20+ instalado), você pode rodar:

```
go run . -u https://exemplo.com -r 100 -c 10
```
