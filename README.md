# simple-lb
Simple load balancer written in Go


### Build From Source

#### Build UI
    cd ui_src
    yarn build

#### Build Executable
    go build main.go

#### Append Static Files
    rice append -i simplelb/ui --exec main.exe
