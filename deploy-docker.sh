#!/bin/bash

# GoSpace Docker Deployment Script
# This script builds and deploys the GoSpace application using Docker Compose

set -e

echo "🚀 GoSpace Docker Deployment"
echo "=============================="
echo ""

# Check if .env file exists
if [ ! -f .env ]; then
    echo "⚠️  .env file not found. Creating from .env.example..."
    if [ -f .env.example ]; then
        cp .env.example .env
        echo "✅ Created .env file. Please update it with your configuration."
        echo "   Edit .env and run this script again."
        exit 1
    else
        echo "❌ .env.example not found. Please create .env file manually."
        exit 1
    fi
fi

echo "📦 Building Docker image..."
docker build -t gospace:latest .

echo ""
echo "🔄 Starting services with Docker Compose..."
docker-compose up -d

echo ""
echo "⏳ Waiting for services to be ready..."
sleep 5

echo ""
echo "✅ Deployment complete!"
echo ""
echo "📊 Service Status:"
docker-compose ps

echo ""
echo "🌐 Application URLs:"
echo "   - Application: http://localhost:8080"
echo "   - PostgreSQL: localhost:5432"
echo ""
echo "📝 Useful commands:"
echo "   - View logs: docker-compose logs -f"
echo "   - Stop services: docker-compose down"
echo "   - Restart: docker-compose restart"
echo "   - Remove all: docker-compose down -v"
echo ""

# Made with Bob
