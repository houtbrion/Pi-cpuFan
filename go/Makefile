SRC=cpufan.go
CFG_FILE=cpufan.cfg
SERVICE_FILE=cpufan.service
BINARY=cpufan
MOD_FILE=go.mod
SUM_FILE=go.sum
GO=go
RM=/usr/bin/rm
BASE_DIR=/usr/local
BIN_DIR=${BASE_DIR}/bin
CONF_DIR=${BASE_DIR}/etc
SYSTEMD_DIR=/etc/systemd/system
CP=/usr/bin/cp
SYSTEMCTL=/usr/bin/systemctl

all : ${BINARY}

${MOD_FILE} : ${SRC}
	${GO} mod init main

${SUM_FILE} : ${MOD_FILE}
	${GO} get periph.io/x/conn/v3/gpio
	${GO} get periph.io/x/host/v3

${BINARY} : ${SUM_FILE}
	${GO} build ${SRC}


clean:
	${RM} ${BINARY} ${SUM_FILE} ${MOD_FILE}

install : ${BINARY}
	${CP} ${BINARY} ${BIN_DIR}
	${CP} ${CFG_FILE} ${CONF_DIR}
	${CP} ${SERVICE_FILE} ${SYSTEMD_DIR}
	${SYSTEMCTL} daemon-reload
	${SYSTEMCTL} enable ${SERVICE_FILE}
