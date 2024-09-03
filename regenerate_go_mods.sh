#!/bin/bash

# Define an array with the paths to the directories where go.mod files need to be managed
MODULE_PATHS=(
    "internal/api-gateway"
    "internal/deployment-service"
    "internal/integration-manager"
    "internal/template-service"
    "pkg/shared"
    "pkg/shared/config"
    "pkg/shared/middlewares"
    "pkg/shared/services"
    "pkg/shared/utils"
)

# Remove existing go.mod and go.sum files
for path in "${MODULE_PATHS[@]}"; do
    cd "$path" || exit

    rm go.mod go.sum

    module_path="gitlab.com/roneeSoft/integrator/${path//\//\/}"
    go mod init "$module_path"

    cd - || exit
done

# Go back to the root directory
cd "$(dirname "$0")" || exit

# Tidy up the modules
go mod tidy
