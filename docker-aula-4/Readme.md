<p align="center">
  <img width="196" height="31" alt="logo" src="https://maratona-static.s3-sa-east-1.amazonaws.com/maratona-static/static/site/img/fullcycle_avancadev.png">
</p>

# Avança Dev Desafio 3
## Requisitos a serem cumpridos
 - Gerar uma imagem Docker para cada microsserviço.
 - Subir as imagens no Docker Hub.
 - Crie um docker-compose que seja capaz de subir todo o ambiente dos microsserviços, incluindo o RabbitMQ.
 - Publicar Dockerfiles e docker-compose no repositório.

## Docker Hub
[Serviço A](https://hub.docker.com/repository/docker/rawallon/microsservicos-avancadev-a) - Porta 9090

[Serviço B](https://hub.docker.com/repository/docker/rawallon/microsservicos-avancadev-b) - Porta 9091

[Serviço C](https://hub.docker.com/repository/docker/rawallon/microsservicos-avancadev-c) - Porta 9092

## Para rodar

 ```bash
$ git clone https://github.com/Rawallon/avancadev.git

$ cd docker-aula-4

$ docker-compose up -d

# Horay!
```