


# godo-cli

godo-cli was heavily inspired by the tugboat.

[![Build Status](https://drone.io/github.com/masayukioguni/godo-cli/status.png)](https://drone.io/github.com/masayukioguni/godo-cli/latest)
[![Build Status](https://travis-ci.org/masayukioguni/godo-cli.svg?branch=master)](https://travis-ci.org/masayukioguni/godo-cli)

DigitalOcean API v2 command line tool for interacting with your [DigitalOcean](https://www.digitalocean.com/) droplets.

## References


## Installation

    $ go get -u github.com/masayukioguni/godo-cli

-u switch updates your previous install, stay current!

## Configuration

### Authorize

Run the configuration utility, `godo-cli authorize`. You can grab your keys
[here](https://cloud.digitalocean.com/settings/applications).

    $ godo-cli authorize
    Enter your API key:foo
    
    Authentication with DigitalOcean was successful!

### Configuration

    $ godo-cli config set -region=nyc3
    successful!

    $ godo-cli config get 
    Defaults
    image: 9801954
    Size: 512mb
    Region: nyc3
    Keys: xxxxxx


## Usage


### Retrieve a list of your droplets

    $ godo-cli droplets
    test (ip: xxx.xxx.xxx.xxx, status: active, region :nyc1, id: 3395705)
    test1(ip: xxx.xxx.xxx.xxx, status: active, region :nyc1, id: 3395706)

### Create a droplet

    $ godo-cli create -name=test 
    Queueing creation of droplet 'test1' ...done

### Droplet power commands

You can either power on, off or cycle a droplet

    $ godo-cli power -id=3395702 -mode=on # power on droplet id 
    $ godo-cli power -name=foo -mode=cycle # power cycle (off to on) droplet by name 


### Take a snapshot of your droplet

    $ godo-cli snapshot -id=3395702 -snapshot=fizz # create a snapshot called fizz for a droplet by id
    $ godo-cli snapshot -name=foo -snapshot=buzz # create a snapshot called buzz for a droplet by name 

N.B. Your droplet needs to be powered off.

### Destroy a droplet

    $ godo-cli destroy -id=3402715
    Queuing destroy for 3402715 ...done

### List Available Images

You can list images that you have created.

list images provided by DigitalOcean as well.

    $ godo-cli images list
    Maintenance Mode (id: 11732785, distro: Debian) 
    633.1.0 (stable) (id: 11420434, distro: CoreOS) coreos-stable
    647.0.0 (beta) (id: 11434448, distro: CoreOS) coreos-beta
    668.2.0 (alpha) (id: 11657005, distro: CoreOS) coreos-alpha
    ...

### List Available Sizes

    $ godo-cli sizes
    slug:512mb memory:   512mb vcpus: 1 disk: 20gb
    slug:  1gb memory:  1024mb vcpus: 1 disk: 30gb
    slug:  2gb memory:  2048mb vcpus: 2 disk: 40gb
    slug:  4gb memory:  4096mb vcpus: 2 disk: 60gb
    ...

### List Available Regions

    $ godo-cli regions
    Regions:
    slug: nyc1 name: New York 1
    slug: ams1 name: Amsterdam 1
    slug: sfo1 name: San Francisco 1
    slug: nyc2 name: New York 2
    slug: ams2 name: Amsterdam 2
    …

### List SSH Keys

    $ godo-cli keys
    id:xxxxxx name:masayukixxxxx@xxxxx
    ...

## Help

If you're curious about command flags for a specific command, you can
ask godo-cli about it.

    $ godo-cli --help create

For a complete overview of all of the available commands, run:

    $ godo-cli  help

## Reporting Bugs

Yes, please!

You can create a new issue [here](https://github.com/masayukioguni/godo-cli/issues/new). Thank you!

## Contributing

1. Fork it!
2. Create your feature branch: `git checkout -b my-new-feature`
3. Commit your changes: `git commit -am 'Add some feature'`
4. Push to the branch: `git push origin my-new-feature`
5. Submit a pull request :D

## History

+ 0.0.3 Added snapshot droplet command
+ 0.0.2 Added power droplet command
+ 0.0.1 first release


## License
MIT
