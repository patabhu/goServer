to-do crud apis.
----------------------------------------------------
To start the server manually follow the steps below.
1. cd goServer in go path.
2. source vars.debug.env for environmental variables for http host port and db configs.
3. in goServer run main.go
4. check-out endpoints in controller in todo.go
---------------------------------------------------------------------------------------
Running 2 different Docker containers follow below instructions:-

For postgres container
1. Run cmd "sudo docker pull postgres:10" for downloading docker postgres image.

2. Run your postgres image in container with cmd "sudo docker run -it --network=bridge -p 2345:5432 -e POSTGRES_HOST_AUTH_METHOD=trust --name postgres -d postgres:10.12".

3. postgres images will be running in container name as postgres.

For Go server container
1. Run cmd "sudo docker build -t goServer ." from in your goServer dir.

2. Run your Go server image in container with cmd "sudo docker run -it --rm -e HTTP_HOST=0.0.0.0:10000 -e DB_HOST=172.17.0.2 -e DB_PORT=5432 -e DB_USER=postgres -e DB_NAME=postgres -e DB_TYPE=postgres -p 10000:10000 --network=bridge --name goServer -d goServer".

3. Go server images will be running in container name as goServer.
