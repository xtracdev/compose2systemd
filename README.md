# compose2systemd

This project parses a docker compose file and generates a systemd unit
file to control a docker compose service as a systemd service.

<pre>
go run main.go unit.go -compose-file-path ../es-atom-pub/cmd/proxy/docker-compose.yml -app-name myapp

[Unit]
Description=Service unit for docker service atomfeedpub
After=docker.service
BindsTo=docker.service
Conflicts=shutdown.target reboot.target halt.target

[Service]
Environment=APP='atomfeedpub'
TimeoutStartSec=0
TimeoutStopSec=30
Restart=always
RestartSec=10

WorkingDirectory=/opt/dockerapps/myapp

ExecStartPre=-/usr/local/bin/docker-compose kill $APP
ExecStartPre=-/usr/local/bin/docker-compose rm $APP
ExecStartPre=-/usr/local/bin/docker-compose rm -f $APP
ExecStart=/usr/bin/envconsul -consul={{consul_addr}} -once -prefix=dc/{{env_prefix}} env /usr/local/bin/docker-compose up --force-recreate --no-deps $APP

ExecStop=/usr/local/bin/docker-compose stop $APP

NotifyAccess=all

[Install]
WantedBy=multi-user.target
</pre>

## Dependencies

<pre>
go get github.com/docker/libcompose
</pre>

## Contributing

To contribute, you must certify you agree with the [Developer Certificate of Origin](http://developercertificate.org/)
by signing your commits via `git -s`. To create a signature, configure your user name and email address in git.
Sign with your real name, do not use pseudonyms or submit anonymous commits.


In terms of workflow:

0. For significant changes or improvement, create an issue before commencing work.
1. Fork the respository, and create a branch for your edits.
2. Add tests that cover your changes, unit tests for smaller changes, acceptance test
for more significant functionality.
3. Run gofmt on each file you change before committing your changes.
4. Run golint on each file you change before committing your changes.
5. Make sure all the tests pass before committing your changes.
6. Commit your changes and issue a pull request.

## License

(c) 2016 Fidelity Investments
Licensed under the Apache License, Version 2.0


