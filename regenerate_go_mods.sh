#!/bin/bash

# Define an array with the paths to the directories where go.mod files need to be managed
MODULE_PATHS=(
    "internal/api-gateway"
    "internal/deployment-service"
    "internal/integration-manager"
    "internal/template-service"
    "pkg/shared/config"
    "pkg/shared/middlewares"
    "pkg/shared/services"
    "pkg/shared/utils"
)

# Loop over each path
for path in "${MODULE_PATHS[@]}"; do
    # Navigate to the module directory
    cd "$path" || exit

    # Remove existing go.mod and go.sum files if they exist
    if [ -f "go.mod" ]; then
        echo "Removing existing go.mod in $path"
        rm go.mod
    fi

    if [ -f "go.sum" ]; then
        echo "Removing existing go.sum in $path"
        rm go.sum
    fi

    # Initialize a new Go module
    echo "Initializing new Go module in $path"
    go mod init "gitlab.com/roneeSoft/integrator/$path"

    # Tidy up the module's dependencies
    echo "Tidying up Go module dependencies in $path"
    go mod tidy

    # Return to the root directory
    cd - || exit
done

echo "Go module reinitialization complete."
