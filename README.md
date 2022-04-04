# easy-mail

Easy mail is a service for sending emails in your Golang application.

## How to use?

## Deploy

## Testing

```shell
[GIN-debug] Listening and serving HTTP on :5000
[GIN] 2022/04/04 - 14:02:27 | 200 |  2.291217333s |       127.0.0.1 | POST     "/mail/send"
```

```shell
default ✓ [======================================] 1 VUs  00m03.3s/10m0s  1/1 iters, 1 per VU

     data_received..................: 202 B 61 B/s
     data_sent......................: 281 B 85 B/s
     http_req_blocked...............: avg=912µs min=912µs med=912µs max=912µs p(90)=912µs p(95)=912µs
     http_req_connecting............: avg=237µs min=237µs med=237µs max=237µs p(90)=237µs p(95)=237µs
     http_req_duration..............: avg=2.29s min=2.29s med=2.29s max=2.29s p(90)=2.29s p(95)=2.29s
       { expected_response:true }...: avg=2.29s min=2.29s med=2.29s max=2.29s p(90)=2.29s p(95)=2.29s
     http_req_failed................: 0.00% ✓ 0        ✗ 1  
     http_req_receiving.............: avg=273µs min=273µs med=273µs max=273µs p(90)=273µs p(95)=273µs
     http_req_sending...............: avg=45µs  min=45µs  med=45µs  max=45µs  p(90)=45µs  p(95)=45µs 
     http_req_tls_handshaking.......: avg=0s    min=0s    med=0s    max=0s    p(90)=0s    p(95)=0s   
     http_req_waiting...............: avg=2.29s min=2.29s med=2.29s max=2.29s p(90)=2.29s p(95)=2.29s
     http_reqs......................: 1     0.303148/s
     iteration_duration.............: avg=3.29s min=3.29s med=3.29s max=3.29s p(90)=3.29s p(95)=3.29s
     iterations.....................: 1     0.303148/s
     vus............................: 1     min=1      max=1
     vus_max........................: 1     min=1      max=1
```