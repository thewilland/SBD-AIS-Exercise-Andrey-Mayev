# todo note commands

## Initialising
    docker swarm init

## Creating secrets
    docker secret create postgres_user docker/postgres_user_secret
    docker secret create postgres_password docker/postgres_password_secret
    docker secret create s3_user docker/s3_user_secret
    docker secret create s3_password docker/s3_password_secret

## Deploying stack
    docker stack deploy -c docker-compose.yml stuff

## Cheking that everything runs
    docker service ls

## Removing stack
    docker stack rm stuff

## Leaving the swarm
    docker swarm leave --force
