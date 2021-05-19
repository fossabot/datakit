package kubernetes

const (
	configSample = `
[[inputs.kubernetes]]
    ## URL for the Kubernetes API
    url = "https://$HOSTIP:6443"

    ## Namespace to use. Set to "" to use all namespaces.
    # namespace = "default"

    ## Use bearer token for authorization. ('bearer_token' takes priority)
    ## If both of these are empty, we'll use the default serviceaccount:
    ## at: /run/secrets/kubernetes.io/serviceaccount/token
    # bearer_token = "/path/to/bearer/token"
    ## OR
    # bearer_token_string = "abc_123"

    ## Set http timeout (default 5 seconds)
    # timeout = "5s"

    ## Optional Resources to exclude from gathering
    ## Leave them with blank with try to gather everything available.
    ## Values can be - "daemonsets", deployments", "endpoints", "ingress", "nodes",
    ## "persistentvolumes", "persistentvolumeclaims", "pods", "services", "statefulsets"
    # resource_exclude = [ "deployments", "nodes", "statefulsets" ]

    ## Optional Resources to include when gathering
    ## Overrides resource_exclude if both set.
    # resource_include = [ "deployments", "nodes", "statefulsets" ]

    ## Optional TLS Config
    # tls_ca = "/path/to/cafile"
    # tls_cert = "/path/to/certfile"
    # tls_key = "/path/to/keyfile"
    ## Use TLS but skip chain & host verification
    # insecure_skip_verify = false

    [inputs.kube_state.tags]
    # tag1 = val1
    # tag2 = val2
`
)
