# autopull

Checks for new Docker images and executes a given command if a new image is available.

## Usage

```bash
./autopull /path/to/configuration.json
```

## Configuration

Create a text file with an array of JSON Dicts:

```json
[
  {
    "Image": "debian",
    "Tag": "10",
    "Actions": ["systemctl restart foo.service"]
  },
  {
    "Image": "my-custom-image",
    "Tag": "latest",
    "Actions": ["docker stop foo",
                "docker rm foo",
                "docker run --name foo my-custom-image:latest"
                ]
  }
]
```

Log levels can be set with an environment variable `LOGLEVEL`.

Valid log levels are:

* info
* error
* warn
* debug

If not set or another level is specified, the application will default to info.

## Deployment

Drop the binary onto a host, along with an appropriate configuration file, and install a crontab.

Example (check every 5 minutes):

```cron
*/5 * * * * /usr/local/bin/autopull /etc/autopull_conf.json
```

## Notes

If the specified image is not present on the Docker host, no action will be taken.
