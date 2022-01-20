# garbage-lb
Garbage load balancer written in Go


### Build From Source

#### Build UI
    cd ui_src
    yarn build

#### Build Executable
    go build main.go

#### Append Static Files
    rice append -i garbage/ui --exec main.exe
