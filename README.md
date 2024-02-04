# What this is

This is just a small project to practice docker,docker-compose, and K8s

Kind version used: `kind v0.17.0 go1.18 darwin/amd64`

# TODO

- Add https://hub.docker.com/r/hkotel/mealie to both docker-compose and K8s
- Add WatchTower to both docker-compose and K8s: https://github.com/containrrr/watchtower
- Create Diagram for Kubernetes (show pods and how they interact with each other)
- Add Kafka And KHQ: https://github.com/tchiotludo/akhq
- Create Go app that runs as Cron Job in K8s and sends data to kafka every 15 seconds (discount scraper)
- Add https://github.com/qarmin/czkawka to Docker and K8s
- Add https://github.com/louislam/uptime-kuma to Docker and K8s
- Add Redis to Docker and K8s
- Add Prometheus Kafka exporter: https://prometheus.io/docs/instrumenting/exporters/
- Add https://github.com/wger-project/wger to Docker and K8s
- Alert manager for prometheus
	- https://devopscube.com/alert-manager-kubernetes-guide/
	- https://grafana.com/blog/2020/02/25/step-by-step-guide-to-setting-up-prometheus-alertmanager-with-slack-pagerduty-and-gmail/
	- https://prometheus.io/docs/alerting/latest/alertmanager/
	- https://github.com/prometheus/alertmanager
- Create simple golang app for Kubernetes that writes random data to Kafka
- Fix Plex (https://www.debontonline.com/2021/01/part-14-deploy-plexserver-yaml-with.html)
- Redix (https://www.youtube.com/watch?v=JmCn7k0PlV4)
- https://www.reddit.com/r/kubernetes/comments/18oxqe4/are_there_any_better_kubernetes_dashboard_options/
	- https://www.reddit.com/r/kubernetes/comments/18oxqe4/are_there_any_better_kubernetes_dashboard_options/
	- https://github.com/kubernetes/dashboard
	- https://k9scli.io/
- Python container that gets pushed to Dockerhub and gets pulled and run as cronjob in K8s
- https://github.com/gethomepage/homepage , https://gethomepage.dev/latest/installation/k8s/
- https://github.com/ArchiveBox/ArchiveBox (add to docker and k8s)
- https://opentelemetry.io/
- https://github.com/louislam/uptime-kuma
	- https://github.com/kdubb1337/examples/tree/main/uptime-kuma/kubernetes
- https://github.com/rancher/rancher
- Tie images to specific version instead of using :latest
- https://github.com/kubernetes-sigs/metrics-server
- Reverse proxy
- Jelly Fin: https://www.debontonline.com/2021/11/kubernetes-part-16-deploy-jellyfin.html
- https://docs.celeryq.dev/en/stable/getting-started/first-steps-with-celery.html

# Accessing Applications (Docker)

- Heimdall: localhost:80
- Grafana: localhost:3000
- Portainer: https://localhost:9443
- Home Assistant: http://localhost:8123
- Plex: http://localhost:32400
- TODO: Kafka
- TODO: AKHQ
- TODO: Add Prometheus URL (should already be available)

# Accessing Applications (K8s)

- Heimdall: http://0.0.0.0:30777
- Grafana: http://0.0.0.0:30776
- Portainer: http://0.0.0.0:30775
- Home Assistant: http://0.0.0.0:30774
- Prometheus: http://0.0.0.0:30773
- Plex: http://0.0.0.0:30772
	- First, go to https://www.plex.tv/claim/ to get the claim and put it in the YAML file
	- Once pod is running, go here: http://0.0.0.0:30772/
	- The setup wizard should come up
	- If you have any trouble, first check if the plexserver container is running: `kubectl get pods <pod-name> -o jsonpath='{.spec.containers[*].name}' -n plexserver` and check the logs of the pod
	- Note: view this link to see how Plex reads media: https://support.plex.tv/articles/naming-and-organizing-your-movie-media-files/
- Kafka: Read below to learn how to interact with Kafka
- AKHQ:  http://0.0.0.0:30771
- Homepage: http://0.0.0.0:30770
- tautulli: http://0.0.0.0:30769/
	- Plex Hostname: plex-service.plexserver.svc.cluster.local
	- Plex Port: 32400
- python project: http://0.0.0.0:30768/

- Setting context
	- cat ~/.kube/config
	- kubectl config set-context main_context --namespace=kubernetes-project --cluster=kind-kind --user=kind-kind
	- kubectl config use-context main_context
	- kubectl config current-context
	- kubectl config unset contexts.main_context

## Kafka

First, enter the Kafka Broker pod `kubectl exec -it <pod_name> -n kubernetes-project -- /bin/bash`

- `kafka-topics.sh --create --topic test-topic --partitions 3 --replication-factor 1 --bootstrap-server kafka-service:9092`
- `kafka-topics.sh --list --bootstrap-server kafka-service:9092`
- `kafka-console-producer.sh --broker-list kafka-service:9092 --topic test-topic`
- `kafka-console-consumer.sh --bootstrap-server kafka-service:9092 --topic test-topic --from-beginning`


- Fining config inside pod
	- kubectl exec -it kafka-broker-5d4c874c5f-xj5mp -n kubernetes-project  -- /bin/bash
	- find . -name server*
	- cat ./opt/kafka_2.13-2.8.1/config/server.properties | grep "advertised"

# Raspberry Pi Installation

# Domain Names

https://raspberrypi.stackexchange.com/questions/37920/how-do-i-set-up-networking-wifi-static-ip-address-on-raspbian-raspberry-pi-os/74428#74428

1. Go through Raspberry PI OS installation
2. Set up SSH and Wifi
3. Get IP address of Pi: hostname -I
4. SSH into PI: `sudo ssh logan@192.168.254.63`
5. Install Pi Hole onto Pi: `curl -sSL https://install.pi-hole.net | bash`
	1. It will warn you that it you need to have a static IP address for the PI. This is not necessary. You just have to be aware that restarting the PI could cause a new IP address to be assigned to it and cause failures on your network, depending on what devices are connecting to the PI
	2. Interface: wlan0
6. Access at localhost:80/admin if on the PI or 192.168.254.63:80/admin if you are connecting to PiHole from another device on the network

Now, we want to be able to access the PI at a domain name instead of using the IP address. To do this, we need to add an A record

Local DNS > DNS Records

Domain: logan.homelab	
IP: 192.168.254.63	

Now, configure the DNS server on your device to use the PI as a DNS Server. On your device, your DNS server is probably the IP address of your router. Change this to be the IP address of the Pi. If you have 2 DNS severs listed, it probably will still go through your router

Now you can access Pi Hole by http://logan.homelab/admin

Install Pi Hole directly on Pi

- Transfer folder to Pi: ```scp -r /Users/logan/docker-project logan@192.168.254.63:~```

sudo kubectl apply -f docker-project/kubernetes/

sudo kubectl get pods -A

## K3s

https://k3s.io/

The install.sh script provides a convenient way to download K3s and add a service to systemd or openrc.

To install k3s as a service, run

curl -sfL https://get.k3s.io | sh -

if you see this error: [INFO]  Failed to find memory cgroup, you may need to add "cgroup_memory=1 cgroup_enable=memory" to your linux cmdline (/boot/cmdline.txt on a Raspberry Pi) (vim /etc/cmdline.txt > cgroup_memory=1 cgroup_enable=memory)

go to this link: https://github.com/k3s-io/k3s/issues/2067#issuecomment-664048424, restart the pi (sudo reboot) and install again

Run systemctl status k3s.service to double check that it install correctly

sudo vim ./etc/lighttpd/lighttpd.conf (Change port to 81)


systemctl | grep "lighttpd"

sudo systemctl restart lighttpd.service

sudo systemctl status lighttpd.service

 Now it should work again:

http://192.168.254.63:81/admin/login.php
http://logan.homelab:81/admin/



/usr/local/bin/k3s-uninstall.sh

## Docker Installation

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

- ```sudo systemctl restart docker```
- ```docker-compose up --force-recreate``` (Docker plugin will only work on newly created containers)

## K3s Installation


https://k3s.io/

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

Apply all YAML files: `kubectl apply -f kubernetes`



If you want to delete the cluster and start over: ```kind delete cluster```


```
kind create cluster --config=kind.config.yaml
kubectl apply -f kubernetes/
```

- To delete Kind from system:
	- Delete:  ~/go/bin/kind and ~/go/pkg/mod/sigs.k8s.io
- To reinstall:
	- go install sigs.k8s.io/kind@v0.17.0
	- kind delete cluster
	- kind create cluster

## Raspberry Pi

We are going to use K3s for this: https://k3s.io/


## Architecture

Promtail is an agent which ships the contents of local logs to a private Grafana Loki instance or Grafana Cloud. It is usually deployed to every machine that has applications needed to be monitored

Zookeeper is used to manage and coordinate Kafka brokers and maintain metadata about Kafka topics, partitions, and consumer groups. Zookeeper tracks the status of Kafka cluster nodes and helps to maintain data consistency among the brokers. Kafka relies on Zookeeper to maintain the configuration information, and without it, Kafka would not be able to function properly.

## Troubleshooting

### Kubernetes

You will have to run `kubectl apply -f kubernetes` twice, due to the order in which the K8s objects are created.

It does take several minutes for all the containers to get into a running state. Sometimes it takes around 8 minutes.

If any pods are stuck in a 'Creating' state, deleting the pod so that it recreates will probably fix it (For example, a PVC might not have been created in time for the Pod).



- kubectl get pods -A | awk '{print $4}'  - only print out status column
