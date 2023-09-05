FROM registry.access.redhat.com/ubi8/ubi-minimal:8.5

RUN microdnf install shadow-utils
RUN useradd -u 10000 manager

WORKDIR /opt/bin/

COPY controller-manager /usr/local/bin/controller-manager
COPY node-agent /usr/local/bin/node-agent

USER manager

ENTRYPOINT ["controller-manager"]
