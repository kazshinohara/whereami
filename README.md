# Whereami

A test application which runs on GKE cluster or other platform, for the case you need to know where your pod runs.  
Corresponding to GET request, this app returns the following information as JSON format.

GET /kind
- Kind (You can specify as an env var)

GET /version
- Version (You can specify as an env var)
  
GET /region
- GCP Region
  
GET /cluster
- GKE cluster name
  
GET /hostname
- Pod's Hostname

GET /
- all of above

GET /headers/{key}
- A specific request header value.

```shell
❯ curl -H 'Host:www.example.com'  http://localhost:8080/headers/Host                                                                                                                                                                                            (gke_kzs-sandbox_asia-northeast1_asm-cluster-01/fortio)
www.example.com
```

## How to use

This is a public container image.  
You can use in any you like.

Here is an example of Service/Deployment manifest.
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: whereami
spec:
  replicas: 1
  selector:
    matchLabels:
      app: whereami 
  template:
    metadata:
      labels:
        app: whereami
    spec:
      containers:
        - image: asia-northeast1-docker.pkg.dev/kzs-sandbox/public/whereami:v2
          imagePullPolicy: Always
          name: whereami
          ports:
            - containerPort: 8080
          env:
            - name: PORT
              value: "8080"
            - name: VERSION
              value: "v1"
            - name: KIND
              value: "test"
```

```yaml
apiVersion: v1
kind: Service
metadata:
  name: whereami
spec:
  ports:
    - port: 80
      targetPort: 8080
      protocol: TCP
  selector:
    app: whereami
  type: LoadBalancer
```