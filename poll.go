package main

import (
    "time"

    "github.com/censhin/eve-feeds/config"
    ec "github.com/censhin/eve-client"
)

var conf *config.Config = config.GetConfig()
var client *ec.Client = ec.NewClient(conf.BaseUrl, conf.KeyId, conf.VCode)

// This function receives an EVE client function it will call to poll the
// XML API for the set time passed until the set condition is met.
func LongPoll(function func(), pollTime time.Duration, condition bool) {
    for {
        function() // TODO: determine how to effectively pass a function
        if condition { break } // TODO: determine how condition will be defined
        time.Sleep(pollTime)
    }
}

// This function initializes the poller by executing a list of goroutines
// that poll the EVE API for data.
func InitPoller() {}

