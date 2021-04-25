# osquery exporter

`osquery_exporter` is an project written in go that exposes the `osqueryi` style
CLI interface over HTTP for querying [`osquery`](https://osquery.io).

## Interacting with the exporter
The daemon runs an http server over the default port 5000 which accepts a single
query over the POST body payload and replies with a JSON list of JSON objects.

```bash=
curl http://localhost:5000/query -H "Content-type: application/json" \
    -d '{"query": "SELECT * FROM osquery_info"}'
```
The above request returns
```json
[
  {
    "build_distro": "10.12",
    "build_platform": "darwin",
    "config_hash": "9ba4d23815e349372b139b9e5e4c030802a55f73",
    "config_valid": "1",
    "extensions": "active",
    "instance_id": "ec5aec49-e7b7-4156-8d7f-c8a764ea47a5",
    "pid": "30941",
    "platform_mask": "21",
    "start_time": "1615470145",
    "uuid": "6DEEC135-7E1A-5D4D-8674-4023D3200969",
    "version": "4.2.0",
    "watcher": "30940"
  }
]
```

## Building the project
The project is written in and go and can be built using GNU `make`.
```bash
$ make build
$ ./osquery_exporter --version
$ ./osquery_exporter --socket /path/to/osquery.sock --addr "0.0.0.0:5000"
```

This requires osqueryd running already with the socket file pointing to the same
socket file. The daemon is an extension that registers with the osqueryd process
explicitly and is not managed by osquery using its extension autoload feature.

## Building a container with osquery and the extension running
A container can be built out of the extension that is running osqueryd as well
as the extension. To build the container,
```bash
$ make docker
```

To run the container
```bash
$ docker run --rm -it -p 5000:5000 osquery_exporter
```
After the container is running, the regular curl request should work.
