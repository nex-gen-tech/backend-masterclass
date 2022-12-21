# Docker Cheat Sheet

## Here is a list of 50 commonly used commands and shortcuts in Docker:

1. **`docker run`**: Run a container

- **`docker run -it`**: Run a container in interactive mode
- **`docker run -d`**: Run a container in detached mode
- **`docker run -p`**: Run a container and map a port

2. **`docker ps`**: List running containers

- **`docker ps -a`**: List all containers
- **`docker ps -q`**: List container IDs
- **`docker ps -l`**: List the last container
- **`docker ps -n`**: List the last n containers

3. **`docker stop`**: Stop a running container

- **`docker stop $(docker ps -a -q)`**: Stop all running containers
- **`docker stop $(docker ps -q)`**: Stop all running containers (shorter)
- **`docker stop $(docker ps -q -f ancestor=<image>)`**: Stop all running containers from an image
- **`docker stop $(docker ps -q -f name=<name>)`**: Stop all running containers with a name

4. **`docker rm`**: Remove a container

- **`docker rm $(docker ps -a -q)`**: Remove all containers
- **`docker rm $(docker ps -a -q -f status=exited)`**: Remove all exited containers
- **`docker rm $(docker ps -a -q -f status=created)`**: Remove all created containers
- **`docker rm $(docker ps -a -q -f status=exited -f status=created)`**: Remove all exited and created containers

5. **`docker images`**: List images

- **`docker images -a`**: List all images (default hides intermediate images)
- **`docker images -q`**: List image IDs
- **`docker images -f dangling=true`**: List dangling images
- **`docker images -f "dangling=false"`**: List non-dangling images
- **`docker images -f "dangling=false" -q`**: List IDs of non-dangling images

6. **`docker rmi`**: Remove an image

- **`docker rmi $(docker images -q)`**: Remove all images
- **`docker rmi $(docker images -q -f dangling=true)`**: Remove all dangling images
- **`docker rmi $(docker images -q -f dangling=false)`**: Remove all non-dangling images

7. **`docker build`**: Build an image from a Dockerfile

- **`docker build -t <name> .`**: Build an image from a Dockerfile in the current directory and tag it
- **`docker build -t <name> -f <file> .`**: Build an image from a specific Dockerfile and tag it
- **`docker build -t <name> --no-cache .`**: Build an image without using cache
- **`docker build -t <name> --rm .`**: Build an image and remove intermediate containers
- **`docker build -t <name> --force-rm .`**: Build an image and always remove intermediate containers
- **`docker build -t <name> --pull .`**: Build an image and always attempt to pull a newer version of the base image
- **`docker build -t <name> --no-cache --pull .`**: Build an image without using cache and always attempt to pull a newer version of the base image

8. **`docker pull`**: Pull an image or a repository from a registry

- **`docker pull <name>`**: Pull an image or a repository from a registry
- **`docker pull <name>:<tag>`**: Pull a specific tag of an image or a repository from a registry
- **`docker pull <name>@<digest>`**: Pull a specific digest of an image or a repository from a registry

9. **`docker push`**: Push an image or a repository to a registry

- **`docker push <name>`**: Push an image or a repository to a registry
- **`docker push <name>:<tag>`**: Push a specific tag of an image or a repository to a registry
- **`docker push <name>@<digest>`**: Push a specific digest of an image or a repository to a registry

10. **`docker login`**: Log in to a registry

- **`docker login`**: Log in interactively
- **`docker login -u <username> -p <password>`**: Log in non-interactively
- **`docker login -u <username> -p <password> <registry>`**: Log in non-interactively to a specific registry

11. **`docker logout`**: Log out from a registry

- **`docker logout`**: Log out from the default registry
- **`docker logout <registry>`**: Log out from a specific registry

12. **`docker search`**: Search the Docker Hub for images

- **`docker search <term>`**: Search the Docker Hub for images
- **`docker search <term> --filter is-official=true`**: Search the Docker Hub for official images
- **`docker search <term> --filter is-automated=true`**: Search the Docker Hub for automated builds

13. **`docker tag`**: Create a tag TARGET_IMAGE that refers to SOURCE_IMAGE

- **`docker tag <source> <target>`**: Create a tag target that refers to source
- **`docker tag <source> <registry>/<target>`**: Create a tag registry/target that refers to source
- **`docker tag <source> <registry>/<target>:<tag>`**: Create a tag registry/target:tag that refers to source

14. **`docker inspect`**: Return low-level information on Docker objects

- **`docker inspect <name>`**: Return low-level information on a container, image or task
- **`docker inspect <name> --format '{{ .State.Running }}'`**: Return the running status of a container
- **`docker inspect <name> --format '{{ .NetworkSettings.IPAddress }}'`**: Return the IP address of a container

15. **`docker history`**: Show the history of an image

- **`docker history <name>`**: Show the history of an image
- **`docker history <name> --no-trunc`**: Show the history of an image without truncating the output

16. **`docker commit`**: Create a new image from a container's changes

- **`docker commit <container> <name>`**: Create a new image from a container's changes
- **`docker commit <container> <registry>/<name>`**: Create a new image from a container's changes and tag it
- **`docker commit <container> <registry>/<name>:<tag>`**: Create a new image from a container's changes, tag it and set the tag

17. **`docker save`**: Save one or more images to a tar archive (streamed to STDOUT by default)

- **`docker save <name> > <file>`**: Save an image to a tar archive
- **`docker save <name> -o <file>`**: Save an image to a specific file
- **`docker save <name> | gzip > <file>.gz`**: Save an image to a gzipped tar archive

18. **`docker load`**: Load an image from a tar archive or STDIN

- **`docker load < <file>`**: Load an image from a tar archive
- **`docker load -i <file>`**: Load an image from a specific file
- **`gunzip -c <file>.gz | docker load`**: Load an image from a gzipped tar archive

19. **`docker cp`**: Copy files/folders between a container and the local filesystem

- **`docker cp <container>:<path> <path>`**: Copy files/folders from a container's filesystem to the local filesystem
- **`docker cp <path> <container>:<path>`**: Copy files/folders from the local filesystem to a container's filesystem

20. **`docker diff`**: Inspect changes to files or directories on a container's filesystem

- **`docker diff <name>`**: Inspect changes to files or directories on a container's filesystem
- **`docker diff <name> --no-trunc`**: Inspect changes to files or directories on a container's filesystem without truncating the output

21. **`docker export`**: Export a container's filesystem as a tar archive

- **`docker export <name> > <file>`**: Export a container's filesystem as a tar archive
- **`docker export <name> -o <file>`**: Export a container's filesystem as a specific file
- **`docker export <name> | gzip > <file>.gz`**: Export a container's filesystem as a gzipped tar archive

22. **`docker import`**: Create a new filesystem image from the contents of a tarball

- **`docker import <file> <name>`**: Create a new filesystem image from the contents of a tarball
- **`docker import <file> <registry>/<name>`**: Create a new filesystem image from the contents of a tarball and tag it
- **`docker import <file> <registry>/<name>:<tag>`**: Create a new filesystem image from the contents of a tarball, tag it and set the tag

23. **`docker volume create`**: Create a volume

- **`docker volume create <name>`**: Create a volume
- **`docker volume create --driver <driver> <name>`**: Create a volume using a specific driver
- **`docker volume create --driver <driver> --opt <key>=<value> <name>`**: Create a volume using a specific driver and set driver specific options

24. **`docker volume inspect`**: Display detailed information on one or more volumes

- **`docker volume inspect <name>`**: Display detailed information on a volume
- **`docker volume inspect <name> --format '{{ .Mountpoint }}'`**: Display the mountpoint of a volume
- **`docker volume inspect <name> --format '{{ .Options }}'`**: Display the options of a volume

25. **`docker volume ls`**: List volumes

- **`docker volume ls`**: List volumes
- **`docker volume ls -q`**: List volumes in quiet mode
- **`docker volume ls -f dangling=true`**: List dangling volumes

26. **`docker volume rm`**: Remove one or more volumes

- **`docker volume rm <name>`**: Remove a volume
- **`docker volume rm $(docker volume ls -q)`**: Remove all volumes

27. **`docker volume prune`**: Remove all unused local volumes

- **`docker volume prune`**: Remove all unused local volumes
- **`docker volume prune -f`**: Remove all unused local volumes without confirmation

28. **`docker network create`**: Create a network

- **`docker network create <name>`**: Create a network
- **`docker network create --driver <driver> <name>`**: Create a network using a specific driver

29. **`docker network connect`**: Connect a container to a network

- **`docker network connect <network> <container>`**: Connect a container to a network
- **`docker network connect <network> <container> --alias <alias>`**: Connect a container to a network and set an alias

30. **`docker network disconnect`**: Disconnect a container from a network

- **`docker network disconnect <network> <container>`**: Disconnect a container from a network
- **`docker network disconnect <network> <container> --force`**: Disconnect a container from a network without confirmation

31. **`docker network inspect`**: Display detailed information on one or more networks

- **`docker network inspect <name>`**: Display detailed information on a network
- **`docker network inspect <name> --format '{{ .IPAM.Config }}'`**: Display the IPAM configuration of a network

32. **`docker network ls`**: List networks

- **`docker network ls`**: List networks
- **`docker network ls -q`**: List networks in quiet mode

33. **`docker network rm`**: Remove one or more networks

- **`docker network rm <name>`**: Remove a network
- **`docker network rm $(docker network ls -q)`**: Remove all networks

34. **`docker network prune`**: Remove all unused networks

- **`docker network prune`**: Remove all unused networks
- **`docker network prune -f`**: Remove all unused networks without confirmation

35. **`docker network disconnect`**: Disconnect a container from a network

- **`docker network disconnect <network> <container>`**: Disconnect a container from a network
- **`docker network disconnect <network> <container> --force`**: Disconnect a container from a network without confirmation

36. **`docker system df`**: Show docker disk usage

- **`docker system df`**: Show docker disk usage
- **`docker system df -v`**: Show docker disk usage and verbose output
- **`docker system df -v --format '{{ .Type }}'`**: Show docker disk usage and filter by type

37. **`docker system events`**: Get real time events from the server

- **`docker system events`**: Get real time events from the server
- **`docker system events --since 2018-01-01T00:00:00`**: Get real time events from the server since a specific date

38. **`docker system info`**: Display system-wide information

- **`docker system info`**: Display system-wide information
- **`docker system info --format '{{ .Name }}'`**: Display system-wide information and filter by name

39. **`docker system prune`**: Remove unused data

- **`docker system prune`**: Remove unused data
- **`docker system prune -a`**: Remove all unused data
- **`docker system prune -f`**: Remove unused data without confirmation

40. **`docker system stats`**: Display a live stream of container(s) resource usage statistics

- **`docker system stats`**: Display a live stream of container(s) resource usage statistics
- **`docker system stats --no-stream`**: Display a live stream of container(s) resource usage statistics without streaming

41. **`docker system version`**: Show the Docker version information

- **`docker system version`**: Show the Docker version information
- **`docker system version --format '{{ .Server.Version }}'`**: Show the Docker version information and filter by server version

42. **`docker top`**: Display the running processes of a container

- **`docker top <name>`**: Display the running processes of a container
- **`docker top <name> <ps args>`**: Display the running processes of a container and pass arguments to ps

43. **`docker version`**: Show the Docker version information

- **`docker version`**: Show the Docker version information
- **`docker version --format '{{ .Server.Version }}'`**: Show the Docker version information and filter by server version

44. **`docker wait`**: Block until one or more containers stop, then print their exit codes

- **`docker wait <name>`**: Block until one or more containers stop, then print their exit codes
- **`docker wait $(docker ps -q)`**: Block until one or more containers stop, then print their exit codes
