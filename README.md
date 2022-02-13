# garbage-lb
Garbage Load Balancer


### Build From Source

#### Build UI
    cd ui_src
    yarn build

#### Build Executable
    cd cmd/garbagelb
    go build main.go
    
or

    go install ./cmd/garbagelb


#### Append Static Files
    rice append -i garbagelb/adminServer --exec main.exe


### Usage
Sample config.yaml file is available in the repo.  
Edit it to your values and start the load balancer.  
Use flag --config to specify the directory where the config.yaml file is located.   


### Config

#### Supported Listener Protocols
- http
- tcp

#### Supported Listener Rules
- http
    - path
    - header
    - cookie
    - source_ip
    - source_port
    - referrer/referer
    - method
    - host

- tcp
    - tcp_check
    - source_ip
    - source_port

#### Rule Comparison Methods
- equals
- not_equals
- contains
- not_contains
- starts_with
- ends_with
- matches
- not_matches
- greater_than
- less_than
- greater_than_or_equal
- less_than_or_equal
- not_required

#### Supported Backend Cluster Protocols
- http / https
- tcp (tcp, tcp4, tcp6)

#### Supported Cluster Routing Policies
- round_robin
- random
- least_connections
