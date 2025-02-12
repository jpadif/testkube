# Overview

This sections provides guidance to get Testkube to production, in particular to enable access to Testkube without requiring k8s privileges to the Testkube k8s namespace and workloads:
1. Exposing Testkube Dashboard [externally with ingresses](exposing-testkube/overview.md).
2. Overall [Testkube deployment on AWS](aws.md).
3. Add [OAuth authentication to the Testkube dashboard](authentication/oauth-ui.md). 
4. Add [OAuth authentication to the Testkube api-server used by the CLI](authentication/oauth-cli.md) as an alternative to default `proxy` mode leveraging [kube apiserver proxy](https://kubernetes.io/docs/concepts/cluster-administration/proxies/).   
