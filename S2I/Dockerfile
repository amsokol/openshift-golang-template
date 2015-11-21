
# sti-labs-golang1
FROM docker.io/openshift/base-centos7

# Put the maintainer name in the image metadata
MAINTAINER Alexander Sokolovsky <amsokol@gmail.com>

# Rename the builder environment variable to inform users about application you provide them
ENV GOLANG_BUILDER_VERSION 1.0

# Set labels used in OpenShift to describe the builder image
LABEL io.k8s.description="Builder for golang web projects" \
      io.k8s.display-name="Golang builder 1.0.0" \
      io.openshift.expose-services="8080:http" \
      io.openshift.tags="builder,golang"

ENV GOLANG_VERSION 1.5.1
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA1 46eecd290d8803887dec718c691cc243f2175fe0

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
	&& echo "$GOLANG_DOWNLOAD_SHA1  golang.tar.gz" | sha1sum -c - \
	&& tar -C /usr/local -xzf golang.tar.gz \
	&& rm golang.tar.gz

ENV PATH /usr/local/go/bin:$PATH

# Copy the S2I scripts to /usr/local/sti, since openshift/base-centos7 image sets io.openshift.s2i.scripts-url label that way, or update that label
COPY ./.sti/bin/ $STI_SCRIPTS_PATH

# Drop the root user and make the content of /opt/openshift owned by user 1001
RUN chown -R 1001:1001 /opt/app-root

# This default user is created in the openshift/base-centos7 image
USER 1001

ENV GOPATH /opt/app-root

# Set the default port for applications built using this image
EXPOSE 8080

# Set the default CMD for the image
CMD $STI_SCRIPTS_PATH/usage
