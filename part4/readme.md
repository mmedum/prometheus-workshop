# Alerting

(Still under review!)

In this session, alerting will be done through Grafana. 

Graph:

```
sum(up{job="go-service"})
```

Alert:

```
When sum() query(A, 10s, now) is below() 1
```

This will generate an alert.
