# High-Level Architecture

## Overview
This project is designed to deploy a simple web application in a cloud environment using Google Cloud Platform (GCP).

![Architecture](internal/storage/architecture.png)

## Architecture Components
- **GCP VPC**: Network segmentation for isolating resources.
- **GKE (Google Kubernetes Engine)**: Manages the deployment of containers.
- **Cloud Load Balancer**: Distributes traffic across multiple instances.
- **Cloud Armor**: Provides Web Application Firewall (WAF) functionality.
- **Cloud Storage**: Stores static assets and backups.
- **Cloud SQL**: Managed relational database service.
- **Secret Manager**: Stores and manages sensitive information securely.
- **Cloud Monitoring & Logging**: Ensures observability and monitoring for system health.

---

# CI/CD Pipeline

![Pipeline](internal/storage/pipeline.png)

## Workflow
1. **Code Push**: Developers push code to the GitHub repository.
2. **GitHub Actions**: Triggers the CI/CD pipeline.
3. **Build Process**:
   - Builds the Docker image.
   - Tags the image with `latest` and release version.
   - Pushes the image to a public Docker registry.
4. **Kubernetes Deployment**:
   - Updates the deployment manifest.
   - Applies the new deployment to GKE using `kubectl`.
5. **Verification**:
   - Access the application.
   - Checks the health of the deployed application.

---

# Evidence

## Deployment Logs
*(Attach screenshots or logs from the CI/CD execution and successful deployment verification)*

## Application Access
![Evidence-1](internal/storage/evidence1.png)
- URL: `http(s)://<IP>/welcome/{name}`
- API Endpoints:
  - `/welcome/{name}`: Returns `"Selamat datang {name}"` if name is provided; otherwise, returns `"Anonymous"`.

## Containerization
- Docker Run: `docker build -t testing/welcome .`
- Docker Compose: `docker-compose up -d`

## Container Registry
![Evidence-2](internal/storage/evidence2.png)
- DockerHub: `dragz17/sltr-devops-boy`

## GitHub Actions
![Evidence-3](internal/storage/evidence3.png)
- URL: `https://github.com/dragz17/sltr-devops-challenge/actions/workflows/docker-ci.yml`
![Evidence-4](internal/storage/evidence4.png)
- Deployment to VM: `http://13.229.123.148:5000/`

## Orchestration
- Please refer to [this  link](./k8s/README.md)
- Using Minikube in Local

---

# Conclusion
This document outlines the high-level architecture, CI/CD pipeline, and deployment evidence for the application.
