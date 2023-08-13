 helm show values bitnami/external-dns > deploy/helm/external-dns-default-value.yaml
 helm show values  bitnami/cert-manager  > deploy/helm/cert-manager-default-values.yaml


kubectl create ns external-dns
helm upgrade -i external-dns bitnami/external-dns -n external-dns -f deploy/helm/external-dns-values.yaml --create-namespace
helm upgrade -i cert-manager bitnami/cert-manager -n cert-manager --create-namespace -f deploy/helm/cert-manager-values.yaml

