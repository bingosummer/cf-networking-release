package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"lib/marshal"
	"log"
	"net/http"
	"netman-agent/config"
	"netman-agent/handlers"
	"netman-agent/policy_client"
	"netman-agent/rule_updater"
	"netman-agent/store"
	"os"
	"time"

	"github.com/coreos/go-iptables/iptables"
	"github.com/pivotal-golang/lager"
	"github.com/tedsuo/ifrit"
	"github.com/tedsuo/ifrit/grouper"
	"github.com/tedsuo/ifrit/http_server"
	"github.com/tedsuo/ifrit/sigmon"
	"github.com/tedsuo/rata"
)

func main() {
	conf := &config.Config{}

	configFilePath := flag.String("config-file", "", "path to config file")
	flag.Parse()

	logger := lager.NewLogger("netman-agent")
	logger.RegisterSink(lager.NewWriterSink(os.Stdout, lager.DEBUG))

	configBytes, err := ioutil.ReadFile(*configFilePath)
	if err != nil {
		log.Fatal("error reading config")
	}

	pollInterval := time.Duration(conf.PollInterval) * time.Second
	if pollInterval == 0 {
		pollInterval = time.Second
	}

	err = json.Unmarshal(configBytes, conf)
	if err != nil {
		log.Fatal("error unmarshalling config")
	}

	store := store.New()

	cniResultHandler := &handlers.CNIResult{
		Logger:      logger,
		StoreWriter: store,
	}

	routes := rata.Routes{
		{Name: "add", Method: "POST", Path: "/cni_result"},
		{Name: "del", Method: "DELETE", Path: "/cni_result"},
	}

	rataHandlers := rata.Handlers{
		"add": cniResultHandler,
		"del": cniResultHandler,
	}

	router, err := rata.NewRouter(routes, rataHandlers)
	if err != nil {
		log.Fatalf("unable to create rata Router: %s", err) // not tested
	}

	policyClient := policy_client.New(
		logger.Session("policy-client"),
		http.DefaultClient,
		conf.PolicyServerURL,
		marshal.UnmarshalFunc(json.Unmarshal),
	)

	ipt, err := iptables.New()
	if err != nil {
		log.Fatal(err)
	}

	ruleUpdater, err := rule_updater.New(
		logger.Session("rules-updater"),
		store,
		policyClient,
		ipt,
		conf.VNI,
	)
	if err != nil {
		log.Fatal(err)
	}

	policyPoller := ifrit.RunFunc(func(signals <-chan os.Signal, ready chan<- struct{}) error {
		close(ready)
		for {
			select {
			case <-signals:
				return nil
			case <-time.After(pollInterval):
				err = ruleUpdater.Update()
				if err != nil {
					return err
				}
			}
		}
	})

	httpServer := http_server.New(fmt.Sprintf("%s:%d", conf.ListenHost, conf.ListenPort), router)

	members := grouper.Members{
		{"http_server", httpServer},
		{"policy_poller", policyPoller},
	}

	monitor := ifrit.Invoke(sigmon.New(grouper.NewOrdered(os.Interrupt, members)))
	err = <-monitor.Wait()
	if err != nil {
		log.Fatalf("daemon terminated: %s", err)
	}
}