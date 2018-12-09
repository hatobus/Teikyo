# TeikyoGenerator
"提供目"(Teikyo-me) is japanese traditional funny picture for TV captured things.

Japanese TV has fund provider. And japanese TV maker introduce the cumpany for appreciation. So sometimes that caption be coverd with some people eyes.

ex)
![](./picture/document/TeikyomeAIKATSU.jpg)

This repository generate this "提供目".

## Installation
To install Teikyo generator, you need to install Go and Go workspaces.

1. `git clone` this repository
    ```
        $ git clone https://github.com/hatobus/TeikyoGenerator.git
    ```
2. To write `.env` file
    .env file is like this
    ```
        URL=[Your request URL with Azure]
        KEY1=[Azure Key1]
        KEY2=[Azure Key2]
    ```
3. To start `server.go`
   ```
        go run server.go
   ```
4. Do POST Request
   
   - curl
        ```
            curl http://localhost:8080/detect \ 
            -F "upload[]=@path/to/1.jpg" \
            -F "upload[]=@path/to/2.jpg" \
            -H Content-Type: multipart/form-data"
        ```
   - httpie

## Used
- Azure 
    - FaceAPI
- Golang
    - godotenv
    - gin