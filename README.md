# PoC gRPC

Esse projeto consiste em uma prova de conceito e exemplificação da aplicabilidade do gRPC como modo de comunicação entre o TCNews e o TCMediaAPI.

## Rodando a PoC

Para rodar é muito simples, basta vc rodar o server/server.go em um terminal e logo em seguida, rodar o client/client.go.

## Debug

Caso queira sugerir alterações, é recomendado instalar o `proto` para modificar o contrato de comunicação. A preparação é muito simples, basta seguir os passos abaixo:

- Necessário ter Golang, qualquer uma das ultimas três releases mais recentes do Go.

- Necessário ter o compilador de buffer de protocolo, `protoc`, versão 3.

Caso queira instalar o compilador, segue os comandos abaixo:

Instalando o `proto`:

Linux:

```sh
  $ apt install -y protobuf-compiler
  $ protoc --version
```

MacOS, usando o Homebrew:

```sh
  $ brew install protobuf
  $ protoc --version
```

Instalando os plug-ins:

```sh
  $ go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
  $ go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
```

Atualize seu PATH para que o compilador protoc possa encontrar os plug-ins:

```sh
  $ export PATH="$PATH:$(go env GOPATH)/bin"
```

## Referencias:

- [Preparação gRPC]("https://grpc.io/docs/languages/go/quickstart/#prerequisites")
- [Proto]("https://developers.google.com/protocol-buffers")
- [Instalação Proto]("https://grpc.io/docs/protoc-installation/")
