package main

import (
	"context"

	"os"
	"os/signal"

	"time"

	"net"
	"net/http"

	"fmt"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func hello(c echo.Context) error {
	ip, _ := getMyInterfaceAddr()
	return c.String(http.StatusOK, fmt.Sprintf("Hello World from server with IP=%s! Now is %s",
		ip.String(),
		time.Now().String()))
}

func getMyInterfaceAddr() (net.IP, error) {
    ifaces, err := net.Interfaces()
    if err != nil {
        return nil, err
    }
    addresses := []net.IP{}
    for _, iface := range ifaces {

        if iface.Flags&net.FlagUp == 0 {
            continue // interface down
        }
        if iface.Flags&net.FlagLoopback != 0 {
            continue // loopback interface
        }
        addrs, err := iface.Addrs()
        if err != nil {
            continue
        }

        for _, addr := range addrs {
            var ip net.IP
            switch v := addr.(type) {
            case *net.IPNet:
                ip = v.IP
            case *net.IPAddr:
                ip = v.IP
            }
            if ip == nil || ip.IsLoopback() {
                continue
            }
            ip = ip.To4()
            if ip == nil {
                continue // not an ipv4 address
            }
            addresses = append(addresses, ip)
        }
    }
    if len(addresses) == 0 {
        return nil, fmt.Errorf("no address Found, net.InterfaceAddrs: %v", addresses)
    }
    //only need first
    return addresses[0], nil
}

func main() {
	// Setup
	e := echo.New()
	e.Logger.SetLevel(log.INFO)

	// Routes
	e.GET("/", hello)

	// Start server
	go func() {
		if err := e.Start(":8080"); err != nil {
			e.Logger.Info("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 10 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}
}
