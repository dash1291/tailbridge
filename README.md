# tailbridge

*(Work in progress. Stay tuned)*

This is supposed to be a very simple tool which does a very simple thing - lets you view logs on a remote machine without having to SSH into it.

## Setup and Usage

### Go dependencies  
You would need to install a couple of dependencies before you can run or build the project.

```
go get github.com/googollee/go-socket.io gopkg.in/yaml.v2
```

### Running the server
To run the server you need to have a `config.yml` present in the current directory, which contains all the necessary configuration settings. There is a template `config.yml.sample` which you can copy as `config.yml` and tweak the latter according to your needs.

After the configuration is done, the following command will start the server on the default port 9191 -
```
go run main.go
```
