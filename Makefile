TARGET=interpreter
DEBUG_TARGET=interpreter_debug

${TARGET}:
	go build -o ${TARGET}

${DEBUG_TARGET}:
	go build -o ${DEBUG_TARGET} -gcflags "-l -N"

.PHONY : run debug clean

run: ${TARGET}
	./${TARGET} -i

debug: ${DEBUG_TARGET}
	dlv attach $(shell ps | fzf | awk '{print $$1}')

clean:
	-rm ${TARGET} ${DEBUG_TARGET}
