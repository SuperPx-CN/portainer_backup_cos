## Install
Can be installed through Docker containers
```shell
docker run -d \
  -e BACKUP_INTERVAL=15m \
  -e BACKUP_LIMIT=4321 \
  -e PORTAINER_URL=http://127.0.0.1:9000 \
  --name portainer_backup_cos \
  superpx/portainer_backup_cos

```

## Environments

| Key             | Required | Default               | description                      |
|-----------------|----------|-----------------------|----------------------------------|
| BACKUP_INTERVAL | false    | 10m                   | The interval between each backup |
| BACKUP_LIMIT    | false    | 4321                  | Save up to 'Limit' backup files  |
| PORTAINER_URL   | false    | http://127.0.0.1:9000 | Portainer url                    |


