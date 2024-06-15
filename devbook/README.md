# Executable


## The MySQL

```bash
docker compose up
```

## The Dockerfile
```bash
docker build . -t devbook

docker run -p 6000:6000 devbook
```

## Rebuild the database

If you make any changes in the init.sql script, you might should delete the volume also.

So when you ```docker compose up```, be aware to use the ```-v``` flag to delete volumes or manually delete with ```docker volume rm name```

Ex:

```bash
docker compose down -v
docker compose up -d
```
also this is not working now lmao