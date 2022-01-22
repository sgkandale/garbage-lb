# garbage-lb
Garbage load balancer


### Build From Source

#### Build UI
    cd ui_src
    yarn build

#### Build Executable
    go build main.go

#### Append Static Files
    rice append -i garbage/ui --exec main.exe
