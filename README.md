
# Recipe API
  

 API que recebe ingredientes como parâmetro de entrada em uma chamada GET e retorna uma lista de receitas.
  
 ### Dependências
 * [Docker](https://www.docker.com/)
 
 
###  Acesso
  
| endpoint | função | método|  
|---|---|---|  
|`/recipes/?i={ingredient_1},{ingredient_2}`| Busca por receitas com os ingredientes fornecidos | GET |  
  
  
### Configuração  
A configuração é feita através de um arquivo `.env`, com as chaves: 
  
| Chave | Descrição | Exemplo|  
|---|---|---|  
|`SERVER_PORT`| Porta usada pelo servidor| `9797` |  
|`GIPHY_API_KEY`| Chave de acesso à API do Giphy | wlP2JB7YbTSjIYzkOdI7Yz2AAAAAAA |  
  
  
### Execução  
Para iniciar a execução, basta fazer o build da imagem do Docker:

```bash  
docker build -t recipe-api .
```
E em seguida rodar:
```bash
docker run -d -p 9797:9797 recipe-api
```
Caso outra porta se não a `9797` seja usada, altere o comando de acordo. 

