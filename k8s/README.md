# Generate Self-Signed Certificate
openssl req -x509 -nodes -days 365 -newkey rsa:2048 -keyout tls.key -out tls.crt -subj "/CN=welcome.local"

kubectl create secret tls welcome-tls --cert=tls.crt --key=tls.key

# Apply Yaml
kubectl apply -f secret.yaml
kubectl apply -f deployment.yaml 
kubectl apply -f service.yaml
kubectl apply -f ingress.yaml