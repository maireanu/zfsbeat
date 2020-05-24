# Zfsbeat

Welcome to Zfsbeat.

Ensure that this folder is at the following location:
`${GOPATH}/src/github.com/maireanu/zfsbeat`

## Getting Started with Zfsbeat

### Requirements

* [Golang](https://golang.org/dl/) 1.7

### Install and run the zfs beat

Download the latest release
https://github.com/maireanu/zfsbeat/releases

Get the config zfsbeat.yml file example from the repo and adjust according to your needs

```
zfsbeat:
  # Defines how often an event is sent to the output
  period: 1s
  # Defines the information needed from the beat
  source_zpool: true
  source_filesystem: true
  source_snapshot: true

name: "Zfsbeat"
seccomp.enabled: false 

# Defines the output of the beat
output.logstash:
  # Array of hosts to connect to.
  hosts: ["localhost:5044"]  
```

### Run the beat

The beat needs to have enought permision in order to use the zfs
you can give permisions by using zfs delegate ( if the zfs version is > 0.7.0 )  or by running it under the root user
https://docs.oracle.com/cd/E23823_01/html/819-5461/gbchv.html#scrolltoc

```
./zfsbeat -c zfsbeat.yml &

# In order to run it on debug mode use
./zfsbeat -c zfsbeat.yml -e -d "*"
```

### Init Project
To get running with Zfsbeat and also install the
dependencies, run the following command:

```
make setup
```

It will create a clean git history for each major step. Note that you can always rewrite the history if you wish before pushing your changes.

To push Zfsbeat in the git repository, run the following commands:

```
git remote set-url origin https://github.com/maireanu/zfsbeat
git push origin master
```

For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).

### Build

To build the binary for Zfsbeat run the command below. This will generate a binary
in the same directory with the name zfsbeat.

```
make
```


### Run

To run Zfsbeat with debugging output enabled, run:

```
./zfsbeat -c zfsbeat.yml -e -d "*"
```


### Test

To test Zfsbeat, run the following command:

```
make testsuite
```

alternatively:
```
make unit-tests
make system-tests
make integration-tests
make coverage-report
```

The test coverage is reported in the folder `./build/coverage/`

### Update

Each beat has a template for the mapping in elasticsearch and a documentation for the fields
which is automatically generated based on `fields.yml` by running the following command.

```
make update
```


### Cleanup

To clean  Zfsbeat source code, run the following command:

```
make fmt
```

To clean up the build directory and generated artifacts, run:

```
make clean
```


### Clone

To clone Zfsbeat from the git repository, run the following commands:

```
mkdir -p ${GOPATH}/src/github.com/maireanu/zfsbeat
git clone https://github.com/maireanu/zfsbeat ${GOPATH}/src/github.com/maireanu/zfsbeat
```


For further development, check out the [beat developer guide](https://www.elastic.co/guide/en/beats/libbeat/current/new-beat.html).


## Packaging

The beat frameworks provides tools to crosscompile and package your beat for different platforms. This requires [docker](https://www.docker.com/) and vendoring as described above. To build packages of your beat, run the following command:

```
make release
```

This will fetch and create all images required for the build process. The whole process to finish can take several minutes.
