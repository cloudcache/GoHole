package main

import (
    "log"
    "flag"

    "GoHole/config"
    "GoHole/dnsserver"
    "GoHole/dnscache"
)


func main(){

    // Command line options
    port := flag.String("p", "", "Set DNS server port")
    cfgFile := flag.String("c", "./config.json", "Config file")

    // option to start the DNS server
    startDNS := flag.Bool("s", false, "Start DNS server")
    
    // Add domain to blacklist by command line
    // example: gohole -ad google.com -ip4 0.0.0.0 -ip6 "::1"
    domain := flag.String("ad", "", "Domain")
    ipv4 := flag.String("ip4", "", "IPv4 Address for the domain")
    ipv6 := flag.String("ip6", "", "IPv6 Address for the domain")

    // Flush Cache&Blacklist DB (RedisDB)
    // example: gohole -fcache
    flushCache := flag.Bool("fcache", false, "Domain")

    
    flag.Parse()

    log.Printf("Loading config..")
    config.CreateInstance(*cfgFile)
    if *port != ""{
        config.GetInstance().DNSPort = *port
    }


    if *domain != "" && *ipv4 != "" && *ipv6 != ""{
        err := dnscache.AddDomainIPv4(*domain, *ipv4, 0)
        if err != nil{
            log.Printf("Error: %s", err)
        }
        err = dnscache.AddDomainIPv6(*domain, *ipv6, 0)
        if err != nil{
            log.Printf("Error: %s", err)
        }
    }

    if *flushCache{
        err := dnscache.Flush()
        if err != nil{
            log.Printf("Error: %s", err)
        }else{
            log.Printf("Cache flushed!")
        }
    }

    if *startDNS{
        dnsserver.ListenAndServe()
    }

}