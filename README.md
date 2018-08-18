# tailbridge

[![Yash Approved](https://img.shields.io/badge/yash-approved-brightgreen.svg)]()

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

Once the server is up and running, and assuming you have configured a linux machine with IP `192.168.1.33` with a user having proper access to the machine in your `config.yml`, one can navigate to the following url to tail one of the system logs file - 

```
http://localhost:9191#192.168.1.33,/var/log/messages
```

### Configuration
Here is a short definition of the various options that go in `config.yml` -

`listen` - Listening port for the server

`groups` - You can organize your similar machines in logical groups like web servers, databases, etc. A group can have certain settings which are common to the machines in that group

>`user` - Use this user to SSH on the machines in this group

>`port` - SSH server port on the remote machine

>`machines` - List of machines in this group

>`directories` - List of directories you want to allow the access to. Each element is a wildcard directory pattern.
 
