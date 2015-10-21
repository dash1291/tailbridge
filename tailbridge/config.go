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

    return config
}
