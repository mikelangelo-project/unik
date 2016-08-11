FROM ubuntu:14.04

# Install prerequisites
RUN apt-get update -y
RUN apt-get install -y qemu wget curl build-essential


# Copy files for docker
COPY docker_files/bin /bin
COPY docker_files/root /root

# Copy files for osv
COPY osv_files /osv_files

# Create mount point (important: must be after copying files for OSv)
RUN mkdir /project_directory
VOLUME /project_directory


# Pull basic capstan images and packages
# RUN /bin/capstan pull mike/osv-loader  # using precompiled binary until remote repo contains patched image
# RUN /bin/capstan package pull eu.mikelangelo-project.osv.bootstrap
# RUN /bin/capstan package pull eu.mikelangelo-project.app.node-4.4.5


# Compose boot image and copy it to /project_directory (unik expects it there)
ENTRYPOINT cd /project_directory && \
	cp /osv_files/* . -R && \
	sed -e '/\/\/CALL_NODE_MAIN_HERE/ {' -e 'r app.js' -e 'd' -e '}' -i node-wrapper-udp.js && \
	/bin/capstan package compose --run "/node /node-wrapper-udp.js" unik/node-image && \
	#qemu-img convert -f qcow2 -O qcow2 -o compat=0.10 /root/.capstan/repository/unik/node-image/node-image.qemu /project_directory/boot.qcow2
	cp /root/.capstan/repository/unik/node-image/node-image.qemu /project_directory/boot.qcow2

# Build this container with
# docker build -t projectunik/compilers-osv-nodejs:v0.0 .

# Run this container with
# docker run -ti --volume="/home/miha/PLAYGROUND/docker/Example-HelloWorldOSv:/project_directory" projectunik/compilers-osv-nodejs:v0.0

