# What this is

This is just a small project to practice docker,docker-compose, and K8s
# Accessing Applications (Docker)

- Heimdall: localhost:80
- Grafana: localhost:3000
- Portainer: https://localhost:9443
- Home Assistant: http://localhost:8123
- Plex: http://localhost:32400

# Accessing Applications (K8s)

- Heimdall: http://0.0.0.0:30777
- Grafana: http://0.0.0.0:30776

# Raspberry Pi Installation

Note: I am using a Raspberry Pi 3B with ARMV7 cpu architecture (32 bit)

- Download Raspberry Pi OS: https://www.raspberrypi.com/software/operating-systems/
	- Download the version for your architecture (in my case) I downloaded the one for ARMv7: 2022-09-22-raspios-bullseye-armhf-lite
		- Download the lite version if you don't want a desktop environment
- Insert SD card into computer and SD card reader
- Use balenaEtcher to flash the Raspberry Pi OS image to the SD card
- Connect the SD card into the Pi
- Connect the Pi to power
- Go through set up prompts for new account
- Configure Wifi
	- Use command ```sudo raspi-config``` to connect to wifi
	- ```reboot```
- SSH into Pi
	- Use sudo raspi-config to enable SSH
	- Find IP Address of Pi: ```hostname -I```
	- From host machine, ```ssh <ip-address>```
	- If you get the following warning: Host key for <raspberry-pi-ip> has changed and you have requested strict checking
		- ```ssh-keygen -R <ip-of-raspberry-pi)```
	- From host machine: ssh raspberry-pi-username@raspberry-pi-ip-address
- sudo apt-get update && sudo apt-get upgrade
- Install docker
	- ```curl -fsSL https://get.docker.com -o get-docker.sh```
	- ```sudo sh get-docker.sh```
	- ```docker --version```
	- ```docker ps```
		- If you get a permission denied error
			- Add our user to the docker group (Automatically, only root users or those with administrative privileges can run the containers. If you are logged out as the root, you can utilize the sudo prefix.  To execute docker commands and avoid typing the sudo each time, you can also add on-root users to the docker group)
				- ```sudo usermod -a -G docker $(whoami)```
				- ```sudo reboot```
				- ```docker ps```
- Install docker-compose (use lscpu to get architecture type and match that with the correct release binary on the releases page)
	- https://github.com/docker/compose#docker-compose-v2
	- ```wget -P .docker/cli-plugins/ https://github.com/docker/compose/releases/download/v2.16.0/docker-compose-linux-armv7```
	- ```mv /home/logan/.docker/cli-plugins/docker-compose-linux-armv7 /home/logan/.docker/cli-plugins/docker-compose```
	- ```chmod +x /home/logan/.docker/cli-plugins/docker-compose```
	- Add folder to path:
		- ```vi ~/.bashrc```
		- Add the following line somewhere in the file ```export PATH="~/.docker/cli-plugins:$PATH"```
		- ```source ~/.bashrc```
		- ```docker-compose --version to verify installation```
- Transfer folder to Pi: ```scp -r ~/Desktop/host logan@192.168.254.63:~```
- ```docker-compose up```
- You can now view applications using the IP address of the Pi (as long as it is connected to your home network) (http://192.168.254.63/)
- Now we have to set up the Loki Docker plugin so that all Docker container logs get sent to Loki
	- ```docker plugin install grafana/loki-docker-driver:latest --alias loki --grant-all-permissions```
	- The recent Loki plugin Docker images do not work on 32 bit architecture, so we have to use version 2.4.0. (https://github.com/grafana/loki/issues/5388#issuecomment-1304561794)
		- Use the following image as a workaround: grafana/loki:2.4.0
	- The loki Docker plugin is currently only packaged for amd64 architecture. To get this to work on ARM architecture, we have to cross compile it. This is explained here (https://github.com/grafana/loki/issues/974#issuecomment-766390258) and here (https://github.com/grafana/loki/issues/974#issuecomment-897165660)
- Now we have to configure the Docker daemon to send all logs to Loki
- ```sudo apt install vim```
- ```vim /etc/docker/daemon.json```

```
{
	"log-driver": "loki",
    	"log-opts": {
        	"loki-url": "http://localhost:3100/loki/api/v1/push",
        	"loki-batch-size": "400"
    }

}
```

- ```udo systemctl restart docker```
- ```docker-compose up --force-recreate``` (Docker plugin will only work on newly created containers)

# Running on Mac

https://grafana.com/docs/loki/latest/clients/docker-driver/configuration/


docker plugin install grafana/loki-docker-driver:latest --alias loki --grant-all-permissions

docker plugin disable eba6e5024973 (if you want to disable)

Vim /Users/logan/.docker/daemon.json

```
{
	"builder": { "gc": { "defaultKeepStorage": "20GB", "enabled": true } },
	"experimental": false,
	"features": { "buildkit": true },
	"log-driver": "loki",
    	"log-opts": {
        	"loki-url": "http://localhost:3100/loki/api/v1/push",
        	"loki-batch-size": "400"
    }

}
```

- After changing daemon.json, restart the Docker daemon for the changes to take effect. All newly created containers from that host will then send logs to Loki via the driver
	- ```killall Docker && open /Applications/Docker.app```
	- Recreate all containers: ```docker-compose up --force-recreate```


https://grafana.com/docs/loki/latest/clients/promtail/scraping/

docker-compose up

# Kubernetes

## Mac

- Install Kind: go install sigs.k8s.io/kind@v0.17.0
- Set up cluster
- On Mac, we need to do some extra configuration due to how Docker works. That is why the kind.config.yaml file is needed. Read me here: https://kind.sigs.k8s.io/docs/user/known-issues/#docker-desktop-for-macos-and-windows

vim kind.config.yaml

```
# Save to 'kind.config.yaml'
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
nodes:
- role: control-plane
  extraPortMappings:
  - containerPort: 30777
    hostPort: 30777
    listenAddress: "0.0.0.0"
```

```kind create cluster --config=kind.config.yaml```

kubectl apply -f kubernetes



If you want to delete the cluster and start over: ```kind delete cluster```


```
kind create cluster --config=kind.config.yaml
kubectl apply -f ../Downloads/portainer\(1\).yaml

```





Delete:  ~/go/bin/kind and ~/go/pkg/mod/sigs.k8s.io
go install sigs.k8s.io/kind@v0.17.0
kind delete cluster
kind create cluster
