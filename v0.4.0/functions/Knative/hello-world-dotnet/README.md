# Sample Function Dotnet

## Run on OpenFunction

1. [Install OpenFunction](https://github.com/OpenFunction/OpenFunction#quickstart)
2. [Refer to the go function sample](../hello-world-go/README.md)

Definition of a ```Function``` for ```dotnet``` is shown below:

```yaml
apiVersion: core.openfunction.io/v1alpha2
kind: Function
metadata:
  name: dotnet-sample
spec:
  version: "v1.0.0"
  image: "<your registry name>/sample-dotnet-func:v0.4.0"
  imageCredentials:
    name: push-secret
  port: 8080 # default to 8080
  build:
    builder: "openfunction/gcp-builder:v1"
    env:
      GOOGLE_FUNCTION_TARGET: "helloworld"
      GOOGLE_FUNCTION_SIGNATURE_TYPE: "http"
    srcRepo:
      url: "https://github.com/OpenFunction/samples.git"
      sourceSubPath: "v0.4.0/functions/Knative/hello-world-dotnet"
  serving:
    runtime: "Knative" # default to Knative
```
