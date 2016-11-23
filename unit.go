package main

//TODO - expand template to process a list of dependencies

type Unit struct {
	Service string
	Dependency string
	AppName string
}

const UnitTemplate  =
`
[Unit]
Description=Service unit for docker service [[.Service]]
After=[[.Dependency]].service
BindsTo=[[.Dependency]].service
Conflicts=shutdown.target reboot.target halt.target

[Service]
Environment=APP='[[.Service]]'
TimeoutStartSec=0
TimeoutStopSec=30
Restart=always
RestartSec=10

WorkingDirectory=/opt/dockerapps/[[.AppName]]

ExecStartPre=-/usr/local/bin/docker-compose kill $APP
ExecStartPre=-/usr/local/bin/docker-compose rm $APP
ExecStartPre=-/usr/local/bin/docker-compose rm -f $APP
ExecStart=/usr/bin/envconsul -consul={{consul_addr}} -once -prefix=dc/{{env_prefix}} env /usr/local/bin/docker-compose up --force-recreate --no-deps $APP

ExecStop=/usr/local/bin/docker-compose stop $APP

NotifyAccess=all

[Install]
WantedBy=multi-user.target
`
