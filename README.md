# honeypot-go

### Requirements
- Git
- Port 80 needs to be available
- Docker
- Curl

### To Deploy
- Clone the Repository
- Go into the repo folder `cd honeypot-go`
- Then run the deploy command... `./bin/deploy`

### To Test
- Make sure you have deployed the honeypot-go container.
- Only then you can run the command...
  - `curl -v localhost`
  - or `curl -v localhost/autodiscover/autodiscover.json?@test.com/owa/?&Email=autodiscover/autodiscover.json%3F@test.com`