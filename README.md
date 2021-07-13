# Whereami

A test application which runs on GKE cluster, for the case you need to know where your pod runs.  
Corresponding to GET "/", this app returns the following information as JSON format.

- Kind (You can specify as an env var)
- Version (You can specify as an env var)
- GCP Region
- GKE cluster name
- Pod's Hostname

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
        - image: asia-northeast1-docker.pkg.dev/kzs-sandbox/public/whereami:v1
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