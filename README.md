### RabbitMQ Client IT Tests

I needed to verify some amqp instances and the existing client tools were taking me longer to be useful then just rolling what I needed here. fork, rip, PR anything that is helpful.

It  adds and deletes queues on a remote RabbitMQ instance using amqp wire proto (in this case 5672/tcp) for messaging.

### Build

    go get github.com/nerdalert/rabbitmq-client-testing

or

    git clone https://github.com/nerdalert/rabbitmq-client-it.git
    cd rabbitmq-client-it
    go build 

### Run

after `go build`, you will have a binary in the root dir you can run with:

    ./rabbitmq-client-it \
    -s fubijar.com \
    -u username \
    -p topsecpassword \
    -v vhostname \
    -q anyname \
    -t 3

### Help / Usage

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
