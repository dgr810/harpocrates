# Harpocrates
This will be the home of the master of all secrets.


When using a ServiceAccount in Kubernetes, the jwt token can be retrieved by reading the file `/var/run/secrets/kubernetes.io/serviceaccount/token`

And then it can be exchanged to a Vault token by posting it to `/auth/kubernetes/login`

Example of a secret file:
```yaml
format: json
dirPath: path/to/dir/to/save/secret/to
secrets:
  - path/to/secret1
  - path/to/secret2:
    - key1:
        saveAsFile: true
    - key2
```
At the moment it takes a json file as input, you can convert your secret to json by doing:
`yq read secret.yml -j`

## How to use
You set the following flags using environment variables:

| Flag          | Environment variable       |
| ------------- | -------------------------- |
| vault_address | HARPOCRATES_VAULT_ADDRESS  |
| cluster_name  | HARPOCRATES_CLUSTER_NAME   |
| token_path    | HARPOCRATES_TOKEN_PATH     |
| vault_token   | HARPOCRATES_VAULT_TOKEN    |
| prefix        | HARPOCRATES_ PREFIX        |


```bash
harpocrates \
  --cluster_name="cluster01-dev" \
  --token_path="./cluster01-dev.token" \
  --vault_address="http://127.0.0.1:8200" \
  --file="./secret.yml"
```


## Deployment.yml
To use this, you can add harpocrates as an initContainers like so:
```yaml
initContainers:
  - name: secret-dumper
    image: harbor.bestsellerit.com/library/harpocrates:68
    args:
      - '{"format":"env","dirPath":"/secrets","prefix":"alfeios_","secrets":["ES/data/alfeios/prod"]}'
    volumeMounts:
      - name: secrets
        mountPath: /secrets
    env:
      - name: VAULT_ADDRESS
        value: $VAULT_ADDR
      - name: CLUSTER_NAME
        value: es03-prod
volumes:
  - name: secrets
    emptyDir: {}
```

CircleCI steps:
```yaml
- secret-injector:
    app-name: alfeios
    file: deployment.yml
    secretFile: alfeios-secrets.yml
- secret-injector:
    app-name: alfeios-db
    file: deployment.yml
    secretFile: alfeios-db-secrets.yml
```

## How to run locally
```bash
export
```


## TO-DO
- [X] harpocrates --inline '{}'
- [X] harpocrates --file /path/to/yaml
- [ ] harpocrates --format=env --dirPath=/tmp/secrets.env --prefix=K8S_CLUSTER_ --secret=ES/data/someSecret
- [ ] harpocrates --format=env --dirPath=/tmp/secrets.env --secrets=ES/data/someSecret:DOCKER_,ES/data/something:K8S_CLUSTER_
- [ ] Should we support more login option?