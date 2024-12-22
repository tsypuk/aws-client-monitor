**aws-client-monitor** is created for analysing locally running applications that are interacting with **AWS**.

Very useful on local development, troubleshooting or profiling 3rd party apps.

Setup is simple:

1. for each app set following environment variables:

```shell
export AWS_CSM_ENABLED=true 
export AWS_CSM_CLIENT_ID=my-app-service-name
```

Optional:
```shell
export AWS_CSM_HOST=127.0.0.1
export AWS_CSM_PORT=31000
```

2. Start ``aws-client-monitor``

```shell
 docker run -p 8080:8080 -p 31000:31000/udp rtsypuk/aws-client-monitor:development
```

3. After that any apps will automatically send information about invocation of AWS API to the ``aws-client-monitor``:

```shell
2024-12-22 16:35:23 UTC [Client:cli-session(aws-cli/1.27.92 md/Botoco) ==========> AWS:S3:ListBuckets(s3.eu-west-1.amazonaws.com)] : 205ms eu-west-1 Code:400
2024-12-22 16:35:23 UTC [Client:cli-session(aws-cli/1.27.92 md/Botoco) ==========> AWS:S3:ListBuckets] : 208ms eu-west-1 Code:400
2024-12-22 16:35:27 UTC [Client:cli-session(aws-cli/1.27.92 md/Botoco) ==========> AWS:Lambda:ListFunctions(lambda.eu-west-1.amazonaws.com)] : 222ms eu-west-1 Code:403
2024-12-22 16:35:27 UTC [Client:cli-session(aws-cli/1.27.92 md/Botoco) ==========> AWS:Lambda:ListFunctions] : 226ms eu-west-1 Code:403
2024-12-22 16:35:23 UTC [Client:user-service(aws-cli/1.27.92 md/Botoco) ==========> AWS:S3:ListBuckets(s3.eu-west-1.amazonaws.com)] : 205ms eu-west-1 Code:400
2024-12-22 16:35:23 UTC [Client:user-service(aws-cli/1.27.92 md/Botoco) ==========> AWS:S3:ListBuckets] : 208ms eu-west-1 Code:400
2024-12-22 16:35:27 UTC [Client:subscirption-service(aws-cli/1.27.92 md/Botoco) ==========> AWS:Lambda:ListFunctions(lambda.eu-west-1.amazonaws.com)] : 222ms eu-west-1 Code:403
2024-12-22 16:35:27 UTC [Client:subscirption-service(aws-cli/1.27.92 md/Botoco) ==========> AWS:Lambda:ListFunctions] : 226ms eu-west-1 Code:403

```

## Browser UI-based mode:

Information is also available on UI (http://localhost:8080)

![ui_screen.png](ui.png)