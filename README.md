# Whereami

A test application which runs on GKE cluster or other platform, for the case you need to know where your pod runs.  
Corresponding to GET request, this app returns the following information as JSON format.

GET /kind
- Kind (You can specify as an env var)
```shell
❯ curl -s http://example.com/kind | jq
{
  "kind": "test"
}
```

GET /version
- Version (You can specify as an env var)
```shell
❯ curl -s http://example.com/version | jq
{
  "version": "v1"
}
```

GET /region
- GCP Region
```shell
❯ curl -s http://example.com/region | jq
{
  "region": "asia-northeast1"
}
```  

GET /cluster
- GKE cluster name
```shell
❯ curl -s http://example.com/cluster | jq
{
  "cluster": "asm-cluster-01"
}
```
GET /hostname
- Pod's Hostname
```shell
❯ curl -s http://example.com/hostname | jq
{
  "hostname": "whereami-687fbc6846-plrnj"
}
```

GET /sourceip
- Source IP Address of your request, if you set multiple X-Forwarded-For, this app returns the first one.
```shell
❯ curl -s http://example.com/sourceip | jq
{
  "source_ip": "X.X.X.X"
}
```
```shell
❯ curl -s -H "X-Forwarded-For: 1.1.1.1" http://example.com/sourceip | jq
{
  "source_ip": "1.1.1.1"
}
```


GET /
- all of above
- you can get a single field wth a query param
```shell
❯ curl -s http://example.com/ | jq
{
  "kind": "test",
  "version": "v1",
  "region": "asia-northeast1",
  "cluster": "asm-cluster-01",
  "hostname": "whereami-687fbc6846-plrnj"
  "source_ip": "X.X.X.X"
}

❯ curl -s http://example.com/?param=hostname | jq
{
  "hostname": "whereami-687fbc6846-plrnj"
}
```

GET /headers/{key}
- A specific request header value.
```shell
❯ curl http://example.com/headers/Host
example.com
```
```shell
❯ curl http://example.com/headers/User-Agent
curl/7.64.1
```

## How to use

This is a public container image, please take it if you like.  
Image url is ***asia-northeast1-docker.pkg.dev/kzs-sandbox/public/whereami:1.0.2***  
The version tag might change per future release.


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
        - image: asia-northeast1-docker.pkg.dev/kzs-sandbox/public/whereami:1.0.2
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