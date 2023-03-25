Sorry.
Bazel.

#To start
`bazelisk run server`

#Deps
bazelisk

# BUILD IMAGE LOCALY
`docker load < bazel-out/k8-fastbuild/bin/geoserver/tarball.tar`
different on k8 is something different on mac. Won't work on mac probably.

# GO TO PROD!
1. `bazelisk build image --define prod=true`

2. copy
`bazel-bin/geoserver/tarball.tar`
to prod as tarball.tar

3. run `docker load < tarball.tar`
4. `docker run geoserver:latest`

# Test locally
http://127.0.0.1:6001/?service=WFS&version=1.1.0&request=GetFeature&typename=community&outputFormat=application/json&srsname=EPSG:3857&bbox=-1630949.8645041138,6199729.394767177,40614977.800791524,20354068.98252296,EPSG:3857