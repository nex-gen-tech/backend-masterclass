# Docker COmpose Cheat Sheet

## Here is a list of 50 commonly used commands and shortcuts in Docker Compose:

1. **`docker-compose up`**: Create and start containers

- **`docker-compose up -d`**: Create and start containers in the background
- **`docker-compose up --build`**: Rebuild images before starting containers
- **`docker-compose up --force-recreate`**: Recreate containers even if their configuration and image haven't changed

2. **`docker-compose start`**: Start existing containers

- **`docker-compose start -d`**: Start containers in the background
- **`docker-compose start --build`**: Rebuild images before starting containers
- **`docker-compose start --force-recreate`**: Recreate containers even if their configuration and image haven't changed

3. **`docker-compose stop`**: Stop running containers without removing them

- **`docker-compose stop -t 10`**: Stop containers and wait for 10 seconds for them to stop
- **`docker-compose stop -t 0`**: Stop containers immediately
- **`docker-compose stop --timeout 10`**: Stop containers and wait for 10 seconds for them to stop

4. **`docker-compose down`**: Stop and remove containers, networks, images, and volumes

- **`docker-compose down -v`**: Stop and remove containers, networks, images, and volumes
- **`docker-compose down --remove-orphans`**: Stop and remove containers for services not defined in the Compose file
- **`docker-compose down --rmi all`**: Stop and remove containers, networks, images, and volumes

5. **`docker-compose ps`**: List containers

- **`docker-compose ps -a`**: List all containers (default shows just running)
- **`docker-compose ps -q`**: List container names only
- **`docker-compose ps -q -f status=exited`**: List containers by status

6. **`docker-compose logs`**: View output from containers

- **`docker-compose logs -f`**: Follow log output
- **`docker-compose logs -f -t`**: Follow log output and show timestamps
- **`docker-compose logs -f -t --tail 10`**: Follow log output and show timestamps for the last 10 lines

7. **`docker-compose exec`**: Run a command in a running container

- **`docker-compose exec -T`**: Run a command in a running container without allocating a TTY
- **`docker-compose exec -d`**: Run a command in a running container in the background
- **`docker-compose exec -u`**: Run a command in a running container as a specific user

8. **`docker-compose run`**: Run a one-off command

- **`docker-compose run -T`**: Run a one-off command without allocating a TTY
- **`docker-compose run -d`**: Run a one-off command in the background
- **`docker-compose run -u`**: Run a one-off command as a specific user

9. **`docker-compose build`**: Build or rebuild services

- **`docker-compose build --no-cache`**: Build or rebuild services, without using cache
- **`docker-compose build --pull`**: Build or rebuild services, always attempting to pull a newer version of the image
- **`docker-compose build --force-rm`**: Build or rebuild services, removing intermediate containers

10. **`docker-compose pull`**: Pull service images

- **`docker-compose pull --ignore-pull-failures`**: Pull service images, ignoring if some service image can't be pulled
- **`docker-compose pull --parallel`**: Pull service images in parallel

11. **`docker-compose push`**: Push service images

- **`docker-compose push --ignore-push-failures`**: Push service images, ignoring if some service image can't be pushed
- **`docker-compose push --parallel`**: Push service images in parallel
- **`docker-compose push --ignore-push-failures --parallel`**: Push service images in parallel, ignoring if some service image can't be pushed

12. **`docker-compose images`**: List images used by the created containers

- **`docker-compose images -q`**: List images used by the created containers, only displaying IDs
- **`docker-compose images -q --filter dangling=true`**: List images used by the created containers, only displaying IDs, and only for dangling images
- **`docker-compose images -q --filter dangling=false`**: List images used by the created containers, only displaying IDs, and only for non-dangling images

13. **`docker-compose port`**: Print the public port for a port binding

- **`docker-compose port --protocol=tcp`**: Print the public port for a port binding, only for TCP
- **`docker-compose port --protocol=udp`**: Print the public port for a port binding, only for UDP
- **`docker-compose port --protocol=tcp --index=1`**: Print the public port for a port binding, only for TCP, and only for the first port binding

14. **`docker-compose top`**: Display the running processes

- **`docker-compose top --pid=PID`**: Display the running processes, only for the container with the specified PID
- **`docker-compose top --pid=PID --services`**: Display the running processes, only for the container with the specified PID, and only for the service name

15. **`docker-compose config`**: Validate and view the Compose file

- **`docker-compose config --quiet`**: Validate and view the Compose file, only displaying errors
- **`docker-compose config --services`**: Validate and view the Compose file, only displaying service names

16. **`docker-compose events`**: Receive real time events from containers

- **`docker-compose events --json`**: Receive real time events from containers, only displaying JSON events
- **`docker-compose events --filter event=stop`**: Receive real time events from containers, only displaying events of type stop

17. **`docker-compose pause`**: Pause services

- **`docker-compose pause SERVICE`**: Pause services, only for the specified service
- **`docker-compose pause --all`**: Pause services, for all services

18. **`docker-compose unpause`**: Unpause services

- **`docker-compose unpause SERVICE`**: Unpause services, only for the specified service
- **`docker-compose unpause --all`**: Unpause services, for all services

19. **`docker-compose kill`**: Kill containers

- **`docker-compose kill SERVICE`**: Kill containers, only for the specified service
- **`docker-compose kill --signal=SIGKILL`**: Kill containers, sending the specified signal

20. **`docker-compose restart`**: Restart services

- **`docker-compose restart SERVICE`**: Restart services, only for the specified service
- **`docker-compose restart --timeout 10`**: Restart services, waiting for 10 seconds for them to stop before killing them

21. **`docker-compose rm`**: Remove stopped containers

- **`docker-compose rm -f`**: Remove stopped containers, including non-stopped containers
- **`docker-compose rm -v`**: Remove stopped containers, including anonymous volumes attached to containers
