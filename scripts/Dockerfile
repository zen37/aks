# Build Image for jenkins
FROM ubuntu:latest

RUN apt-get update

RUN apt-get install \
    ca-certificates \
    curl wget -y \
    gnupg \
    lsb-release

ENV TZ=America/New_York
RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone


RUN apt-get update


# Install go
RUN wget https://go.dev/dl/go1.17.5.linux-amd64.tar.gz
RUN rm -rf /usr/local/go && tar -C /usr/local -xzf go1.17.5.linux-amd64.tar.gz
RUN echo "export PATH=$PATH:/usr/local/go/bin" | tee --append /etc/profile


RUN mkdir -p app

WORKDIR /app

COPY entrypoint.sh .

COPY go-webservice-gin .

RUN chmod +x entrypoint.sh

RUN export PATH=$PATH:/usr/local/go/bin && \
    go get .


ENTRYPOINT ./entrypoint.sh