FROM --platform=${TARGETPLATFORM} alpine:3.18

ARG TARGETPLATFORM
ARG CMD_NAME
ENV COMMAND_NAME=${CMD_NAME}

COPY ${TARGETPLATFORM}/${CMD_NAME} /usr/local/bin/

CMD ["/bin/sh", "-c", "${COMMAND_NAME}"]
