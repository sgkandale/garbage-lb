# garbage-lb
Garbage Load Balancer


### Build From Source

#### Build UI
    cd ui_src
    yarn build

#### Build Executable
    go build main.go

#### Append Static Files
    rice append -i garbagelb/adminServer --exec main.exe


### Usage
Sample config.yaml file is available in the repo.  
Edit it to your values and start the load balancer.  
Use flag --config to specify the directory where the config.yaml file is located.   


### Config

#### Supported Listener Protocols
- http

#### Supported Listener Rules
- path
- header
- cookie
- source_ip
- source_port
- referrer/referer
- method
- host

#### Supported Backend Cluster Protocols
- http

#### Supported Cluster Routing Policies
- round_robin
- random
- least_connections
