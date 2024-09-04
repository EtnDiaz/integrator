# Integrator Project

The **Integrator** project is a microservices-based integration tool designed to connect disparate services through their APIs. This tool simplifies service integrations by providing a user-friendly interface for configuring integrations, scaling Kubernetes deployments, and handling backup and restoration tasks seamlessly.

## Features

- **Integration Management**: Configure integrations between different services through their APIs.
- **Microservices Architecture**: The system is built as a set of microservices including:
  - **API Gateway**
  - **Deployment Service**
  - **Integration Manager**
  - **Template Service**
  - **TODO**
- **Scalable Infrastructure**: Deployed on Kubernetes, leveraging Helm charts and Docker containers.
- **Best Practices**: Clean architecture principles for Go microservices with a clear separation of concerns.
- **Kubernetes Support**: Seamless Kubernetes deployment, backup, and recovery functionalities.
- **Logging and Monitoring**: Integrated with custom logging middleware and authentication mechanisms.

## Microservices
API Gateway
Responsible for routing incoming requests to the appropriate microservices. It uses the Gin framework and includes authentication and logging middleware.

Deployment Service
Handles the deployment of services within Kubernetes, including scaling and updating configurations.

Integration Manager
Manages the lifecycle of service integrations. Coordinates between different microservices for API-based communication.

Template Service
Handles template management (create, read, update, delete operations for templates).

## Technologies Used
Go 1.22.5: Primary language for all services.
Gin: Lightweight web framework for building fast and scalable web services.
JWT: For securing endpoints with token-based authentication.
Kubernetes: For orchestration and scaling.
Helm: For managing Kubernetes deployment configurations.
Docker: Containerizes each microservice for easy deployment.
Prometheus & Grafana: For monitoring and alerting.
GitLab CI/CD: Automated pipelines for building, testing, and deploying the services.


## Contributing
Contributions are welcome! Please open an issue or submit a pull request to help improve this project.

## License
This project is licensed under the AFFERO GENERAL PUBLIC LICENSE - see the LICENSE file for details.