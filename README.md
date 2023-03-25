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
`bazel-out/k8-fastbuild/bin/geoserver/tarball.tar`
to prod as tarball.tar

3. run `docker load < tarball.tar`
4. `docker run geoserver:latest`