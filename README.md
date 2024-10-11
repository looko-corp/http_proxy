# http_proxy

단순한 HTTP 프록시입니다.

## 실행

deploy.sh 스크립트를 super user로 실행하면 데몬까지 자동으로 등록될 겁니다.

```bash
sudo sh deploy.sh
```

그리고 다음 명령으로 잘 실행되었는지 확인해보세요.

```bash
sudo systemctl status http_proxy
```

## 사용법

다음과 같이 Proxy-Host 헤더에 최종 목적지 엔드포인트를 전달하면, 호스트와 http/https만 Proxy-Host의 값으로 치환해서 요청을 쏴주고, 응답을 반환합니다.
query-parameter, header, body 등은 그대로 포워딩합니다. Proxy-Host 헤더만 빼고요.

```
POST http://IP:8080/api/v1/health
Content-Type: application/json
Proxy-Host: https://jnv43hiecbl5c5bdsex5ncgbby0lmdcp.lambda-url.ap-northeast-2.on.aws

{
  "status": "UP"
}
```
