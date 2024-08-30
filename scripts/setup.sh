#!/bin/bash

echo "Setting up the environment..."

# Install dependencies
go get ./...

# Initialize the database
psql -h localhost -U postgres -f db/schema.sql

echo "Setup complete!"
