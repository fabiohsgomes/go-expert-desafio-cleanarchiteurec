FROM mcr.microsoft.com/devcontainers/base:jammy

RUN curl -L https://packagecloud.io/golang-migrate/migrate/gpgkey | apt-key add - \
&& echo "deb https://packagecloud.io/golang-migrate/migrate/ubuntu/ $(lsb_release -sc) main" > /etc/apt/sources.list.d/migrate.list \
&& apt-get update \
&& apt-get dist-upgrade -y \
&& apt-get install -y protobuf-compiler migrate