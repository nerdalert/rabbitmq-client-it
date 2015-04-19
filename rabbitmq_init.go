// Integration tests for amqp, run using:
// go build
// ./rabbitmq-client-testing  -s tiger.cloudamqp.com -u grumpycat -p topsec -v vhostname -q randomqueue -t 3
package main

import (
	"fmt"
	log "github.com/Sirupsen/logrus"
	"github.com/jessevdk/go-flags"
	"os"
)

func main() {

	var opts struct {
		SrvAddr   string `short:"s" long:"server" description:"target ip or domain name of the amqp host"`
		User      string `short:"u" long:"user" description:"username of the target amqp host"`
		Passwd    string `short:"p" long:"pass" description:"password of the target amqp host"`
		Vhost     string `short:"v" long:"vhost" description:"name of the amqp vhost"`
		Debug     bool   `short:"d" long:"debug" description:"print credentials and remote host"`
		QueueName string `short:"q" long:"queue" description:"test queue name to use"`
		TestCase  int    `short:"t" long:"test" description:"test to run. 1) Add a queue. 2) Delete a queue. 3) Both."`
		Help      bool   `short:"h" long:"help" description:"show amqp tester help"`
	}
	_, err := flags.Parse(&opts)
	if err != nil {
		showUsage()
		os.Exit(1)
	}
	if opts.Help == true {
		showUsage()
		os.Exit(1)
	}
	if opts.TestCase > 3 {
		showUsage()
		log.Fatal("Please choose a test case number listed in --help or leave blank to run all tests")
		os.Exit(1)
	}
	if opts.SrvAddr == "" || opts.User == "" || opts.Passwd == "" || opts.Vhost == "" {
		showUsage()
		log.Fatal("Required fields missing or null")
		os.Exit(1)
	}
	if opts.Debug == true {
		log.SetLevel(log.DebugLevel)
		log.Debug("Logging level is set to : ", log.GetLevel())
	}
	amqp := new(AmqpHost)
	amqp.SetAmqpAddr(opts.SrvAddr)
	amqp.SetUid(opts.User)
	amqp.SetPwd(opts.Passwd)
	amqp.SetVhost(opts.Vhost)
	amqp.SetQueue(opts.QueueName)
	log.Debugf("Testing server [ %s ] with the following paramters [ %s ]", amqp.GetAmqpAddr(), amqp.String())
	switch opts.TestCase {

	case 1:
		log.Info("Waiting for transaction to complete. To exit press CTRL+C abort")
		log.Infof("Adding queue [ %s ] to target amqp service at [ %s ] ", amqp.GetQueue(), amqp.GetAmqpAddr())
		AddQueue(*amqp)
		log.Infof("Transaction successful, queue [ %s ] was added.", amqp.GetQueue())

	case 2:
		log.Info("Waiting for transaction to complete. To exit press CTRL+C abort")
		log.Infof("Deleting queue [ %s ] to target amqp service at [ %s ]", amqp.GetQueue(), amqp.GetAmqpAddr())
		DeleteQueue(*amqp)
		log.Infof("Transaction successful, queue [ %s ] was deleted.", amqp.GetQueue())

	case 3:
		log.Info("Waiting for transaction to complete. To exit press CTRL+C abort")
		log.Infof("Adding and Deleting queue [ %s ] to target amqp service at [ %s ]", amqp.GetQueue(), amqp.GetAmqpAddr())
		AddQueue(*amqp)
		log.Infof("Transaction successful, queue [ %s ] was added.", amqp.GetQueue())
		defer DeleteQueue(*amqp)
		log.Infof("Transaction successful, queue [ %s ] was deleted.", amqp.GetQueue())

	default:
		log.Info("Waiting for transaction to complete. To exit press CTRL+C abort")
		log.Warn("No test specified, defaulting to running all integration tests against the target amqp service at [ %s ]")
		log.Infof("Adding and Deleting queue [ %s ] to target amqp service at [ %s ]", amqp.GetQueue(), amqp.GetAmqpAddr())
		AddQueue(*amqp)
		log.Infof("Transaction successful, queue [ %s ] was added.", amqp.GetQueue())
		defer DeleteQueue(*amqp)
		log.Infof("Transaction successful, queue [ %s ] was deleted.", amqp.GetQueue())
	}
}

type toString interface {
	String() string
}

func (a *AmqpHost) String() string {
	s := fmt.Sprintf("Username: [ %s ] \n"+
		"AMQP Server Addr: [ %s ] \n"+
		"Username: [ %s ] \n"+
		"Passwd: [ %s ] \n"+
		"VHost: [ %s ] \n",
		a.GetUid(), a.GetAmqpAddr(), a.GetUid(), a.GetPwd(), a.GetVhost())
	return s
}

func showUsage() {
	var usage string
	usage = `
Usage:
  main [OPTIONS] [collector IP address] [collector port number]

Application Options:
  -s, --server= (required) target ip or domain name of the amqp service
  -u, --user= (required) username of the target amqp host
  -p, --pass (required) password of the target amqp host
  -v, --vhost (required)  name of the amqp vhost
  -t, --test (default is BOTH) warning: Don't use a production queue. The delete will not descriminate.
  		1) Add a queue.
  		2) Delete a queue.
  		3) Both.
  -dbg, --debug (optional) print debug information (warning: prints username/passwd to stdout)
  -

Example:

    ./rabbitmq-client-it -h amqp.foo.org -u grumpy -p cat -v vhostname -q testqueue -t 3
    - with debug logs to stdout
    ./rabbitmq-client-it -h amqp.foo.org -u grumpy -p cat -v vhostname -q testqueue -t 3 -d

Help Options:
  -h, --help    Show this help message
  `
	fmt.Print(usage)
}
