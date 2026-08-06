#!/usr/bin/env bash

set -euo pipefail

echo "======================================"
echo " DockARR Bootstrap"
echo "======================================"

echo ""
echo "Creating project directories..."

mkdir -p \
backend/cmd/server \
backend/internal/{app,config,database,docker,events,handlers,logging,middleware,models,registry,router,server,services,settings,version,websocket} \
backend/pkg/sdk \
backend/migrations \
frontend/src \
frontend/public \
docker/dev \
docs/ADR \
.github/workflows

echo "Creating project files..."

touch \
README.md \
LICENSE \
.gitignore \
Makefile \
docker-compose.yml \
backend/Dockerfile \
backend/.air.toml \
backend/.env.example

cd backend

if [ ! -f go.mod ]; then
    echo "Initializing Go module..."
    go mod init github.com/z354prance/DockARR/backend
fi

echo "Installing Go dependencies..."

go get \
github.com/gin-gonic/gin \
go.uber.org/zap \
github.com/spf13/viper \
gorm.io/gorm \
gorm.io/driver/postgres \
github.com/jackc/pgx/v5

echo ""
echo "Bootstrap complete."