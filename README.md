# Event Log Service

Service for logging cloudEvents using keywords.

> keywords are **case-insensitive**!

## configure cloudEvent protocol and keywords with yaml file
```yaml
protocol: nats
nats:
  url: http://localhost:4222
  timeoutInSec: 10 #optional
  queueGroup: testing #optional
keywords:
  - exception
  - failure
```

> the configuration process of the cloudevent protocol is described [here](https://gitlab.eclipse.org/eclipse/xfsc/libraries/messaging/cloudeventprovider) in more detail.