package tailbridge

import (
    "io/ioutil"

    "gopkg.in/yaml.v2"
)

type Config struct {
    Listen int
    Groups map[string]Group
}

type Group struct {
    User string
    Port int
    Machines []string
}

// Machine-group map
var machines map[string]string
var config Config

func ReadConfig(config_path string) Config {
    yamlFile, err := ioutil.ReadFile(config_path)

    if err != nil {
       panic(err)
    }

    err = yaml.Unmarshal(yamlFile, &config)

    if err != nil {
        panic(err)
    }

    BuildMachinesIndex(config.Groups)
    return config
}

func BuildMachinesIndex(groups map[string]Group) {
    machines = make(map[string]string)

    for group_name := range groups {
        for _, ip := range groups[group_name].Machines {
            machines[ip] = group_name
        }
    }
}

func GetMachineParams(ip string) (user string, port int, succes bool) {
    group_name, ok := machines[ip]

    if ok {
        port := config.Groups[group_name].Port
        user := config.Groups[group_name].User
        return user, port, ok
    } else {
        return "", 0, ok
    }
}
