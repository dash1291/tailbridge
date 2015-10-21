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

func ReadConfig(config_path string) Config {
    yamlFile, err := ioutil.ReadFile(config_path)

    if err != nil {
       panic(err)
    }

    var config Config
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
