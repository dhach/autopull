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
    "Action": "systemctl restart foo.service"
  },
  {
    "Image": "hello-world",
    "Tag": "latest",
    "Action": "touch /tmp/foo"
  }
]
```

## Deployment

Drop the binary onto a host, along with an appropriate configuration file, and install a crontab.

Example (check every 5 minutes):

```cron
*/5 * * * * /usr/local/bin/autopull /etc/autopull_conf.json
```

## Notes

If the specified image is not present on the Docker host, no action will be taken.

Also note that, as of now, you cannot chain commands together, using semicolons or ampersands.
