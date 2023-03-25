#!/bin/bash
source secrets.prod

geoserver="bazel-bin/geoserver/tarball.tar"
bazelisk build geoserver --define prod=true

scp $geoserver $sshCred:geoserver/geoserver.tar
# scp $geoserver $sshCred:docker-compose.geoserver.yml 

ssh $sshCred << EOF
    docker stop geo
    docker container rm geo
    docker image rm -f geoserver
    docker load -i geoserver/geoserver.tar
    echo "Loaded in the tar \n"
    docker run -itd --network=garden-land -expose=6001:6001 --name geo geoserver
    echo "done! \n"
EOF