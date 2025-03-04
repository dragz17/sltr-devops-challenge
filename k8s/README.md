# Kubernetes Deployment Guide

## Generating a Self-Signed Certificate
To secure the Ingress, generate a self-signed TLS certificate:

```sh
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=welcome.local"
```

Create a Kubernetes TLS secret from the generated certificate:

```sh
kubectl create secret tls welcome-tls --cert=tls.crt --key=tls.key
```

## Applying Kubernetes Manifests

Ensure you have configured a ClusterIssuer (if using Let's Encrypt). Apply the required YAML files in the correct order:

```sh
# (Optional) Apply ClusterIssuer if using Let's Encrypt
# kubectl apply -f clusterissuer.yaml

# Apply secrets
kubectl apply -f secret.yaml

# Deploy the application
kubectl apply -f deployment.yaml 

# Expose the service
kubectl apply -f service.yaml

# Configure Ingress for external access
kubectl apply -f ingress.yaml
```

## Verification
After deployment, check the status of the resources:

```sh
kubectl get pods
kubectl get svc
kubectl get ingress
kubectl describe ingress welcome-ingress
```

If using a self-signed certificate, you may need to add the certificate to your trusted store or bypass security warnings in the browser.