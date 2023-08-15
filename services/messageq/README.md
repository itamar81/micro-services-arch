 helm show values bitnami/rabbitmq > rabbitmq-default-value.yaml




helm upgrade -i rabbitmq bitnami/rabbitmq -n rabbitmq  --create-namespace # -f rabbitmq-values.yaml