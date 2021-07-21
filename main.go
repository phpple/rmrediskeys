package main

import (
    "flag"
    "fmt"
    "gopkg.in/yaml.v2"
    "io/ioutil"
    "log"
    "os"
    "strconv"
    "strings"
)
import "gopkg.in/redis.v4"

func main() {
    configFile := *flag.String("config", "", "config file")

    var pattern string
    if len(os.Args) > 1 {
        pattern = os.Args[1]
    } else {
        help()
        os.Exit(1)
    }

    if pattern == "*" || pattern == "go.mod" {
        log.Fatalln("* is not allowed to use for the security risk.")
        os.Exit(2)
    }

    fmt.Println("pattern:" + pattern)
    if configFile == "" {
        pwd, _ := os.Getwd()
        configFile = pwd + "/config.yml"
    }
    redisConfig := getConfig(configFile)

    client := createClient(redisConfig)
    keys := client.Keys(pattern).Val()

    if len(keys) == 0 {
        fmt.Println("Found nothing")
        return
    }

    fmt.Println("Found ", len(keys),"keys:", keys)
    fmt.Println("Do you want delete all of them: y,n?")

    var yesOrNo string
    fmt.Scanln(&yesOrNo)
    yesOrNo = strings.ToLower(strings.TrimSpace(yesOrNo))
    if yesOrNo == "y" {
        client.Del(keys...)
        log.Println("Keys deleted:", keys)
    }
}

func help() {
    fmt.Println("usage ./rmrediskeys somepattern.*")
}

// ge redis client
func createClient(config RedisConfig) *redis.Client {
    client := redis.NewClient(&redis.Options{
        Addr:     config.Host + ":" + strconv.Itoa(config.Port),
        Password: config.Password,
        DB:       config.Database,
    })

    // check if connected successfully with cient.Ping()
    _, err := client.Ping().Result()
    if err != nil {
        fmt.Println(err)
        os.Exit(2)
    }

    return client
}

func getConfig(configFile string) RedisConfig {
    _, err := os.Stat(configFile)
    log.Println(err)
    var conf RedisConfig

    if err != nil && os.IsNotExist(err) {
        log.Println("config file not found, use default configuration")

        conf = RedisConfig{}
        conf.Host = "localhost"
        conf.Port = 6379
        log.Println(conf)
        return conf
    }

    redisConfig, err := ioutil.ReadFile(configFile)

    if err != nil {
        log.Fatalln("redisConfig file read error:", err)
    }
    err = yaml.Unmarshal(redisConfig, &redisConfig)
    if err != nil {
        log.Fatalln("yaml parse failed:", err)
    }

    if conf.Host == "" {
        conf.Host = "localhost"
    }
    if conf.Port == 0 {
        conf.Port = 6379
    }
    return conf
}
