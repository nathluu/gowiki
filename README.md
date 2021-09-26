## Steps to deployment with ClientIP preservation
**Step 1:** Install istio  
**Step 2:** Deploy application
```bash
kubectl create namespace app
kubectl label ns app istio-injection=enabled
kubectl apply -f statefulset.yaml -n app
kubectl apply -f istio-gateway.yaml -n app
```
**Step 3:** Mandate the istio-gateway to preserve clientIP
```bash
kubectl patch svc istio-ingressgateway -n istio-system -p '{"spec":{"externalTrafficPolicy":"Local"}}'
```