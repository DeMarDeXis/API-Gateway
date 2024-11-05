#!/bin/bash

# Функция для выполнения main команды
main_cmd() {
    echo "Running main command..."
    go run ./cmd/main/main.go
}

# Функция для сборки и запуска Docker контейнеров
docker_compose_build() {
    echo "Building Docker images and starting containers..."
    docker-compose up --build
}

# Функция для запуска Docker контейнеров
docker_compose_up() {
    echo "Starting Docker containers..."
    docker-compose up -d
}

# Функция для остановки и удаления Docker контейнеров
docker_compose_down() {
    echo "Stopping and removing Docker containers..."
    docker-compose down
}

# Функция для проверки работы Docker
docker_check() {
    echo "Checking Docker status..."
    docker ps
    sleep 2
    docker ps -a
    sleep 2
    docker image ls
}

# Основной блок скрипта
case "$1" in
    main)
        main_cmd
        ;;
    build)
        docker_compose_build
        ;;
    up)
        docker_compose_up
        ;;
    down)
        docker_compose_down
        ;;
    check)
        docker_check
        ;;
    *)
        echo "Usage: $0 {main|build|up|down|check}"
        exit 1
esac