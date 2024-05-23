# Go Challenge Google cloud Run: Weather API

Este projeto implementa um serviço que recebe um CEP (Código de Endereçamento Postal) brasileiro, identifica a cidade e retorna a previsão do tempo atual (temperatura em Celsius, Fahrenheit e Kelvin). O serviço é projetado para ser implantado no Google Cloud Run.

## Requisitos

- Go 1.21.3
- Docker
- Docker Compose
- Task (opcional)
- Air (opcional)

## Setup

1. **Clone o repositório:**

   ```sh
   git clone https://github.com/rzeradev/google-cloud-run.git
   cd google-cloud-run
   ```

2. **Crie e configure o aquivo `.env`:**

   ```sh
   cp .env.example .env
   # Modifique o arquivo .env como necessário
   ```

3. **Inicie o Servidor Go com Docker Compose**

   ```sh
   docker-compose up -d
   ```

4. **O servidor vai rodar no endereço `http://localhost:8080`**

## Endpoints da API

- **GET /weather/:zipcode**
  - Success Response:
    ```json
    {
    	"temp_C": 28.5,
    	"temp_F": 83.3,
    	"temp_K": 301.5
    }
    ```
  - Invalid ZIP code Response:
    ```json
    {
    	"message": "invalid zipcode"
    }
    ```
  - ZIP code not found Response:
    ```json
    {
    	"message": "can not find zipcode"
    }
    ```

## URL do Projeto rodando na Google Cloud Run

1. Basta acessar a URL abaixo e substituir o `zipcode`pelo CEP desejado:
   ```sh
   https://google-cloud-run-golang-hmyixbzgba-rj.a.run.app/weather/`zipcode`
   ```

## Exemplos que podem ser testados

1. CEP Existente

   ```sh
   https://google-cloud-run-golang-hmyixbzgba-rj.a.run.app/weather/26572070
   ```

2. CEP Inexistente

   ```sh
   https://google-cloud-run-golang-hmyixbzgba-rj.a.run.app/weather/16572070
   ```

3. CEP Inválido

   ```sh
   https://google-cloud-run-golang-hmyixbzgba-rj.a.run.app/weather/123
   ```

4. Previsão do tempo não encontrada para a cidade
   ```sh
   https://google-cloud-run-golang-hmyixbzgba-rj.a.run.app/weather/70150900
   ```

## Licença

Este projeto está licenciado sob a [MIT License](LICENSE).
