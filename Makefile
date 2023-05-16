TARGET=interpreter
DEBUG_TARGET=interpreter_debug

${TARGET}:
	go build -o ${TARGET}

${DEBUG_TARGET}:
	go build -o ${DEBUG_TARGET} -gcflags "-l -N"

.PHONY : run debug clean debug_lexer

run: ${TARGET}
	./${TARGET} -i

run_file: ${TARGET}
	./${TARGET} -f=${FILE}

debug: ${DEBUG_TARGET}
	dlv attach $(shell ps | fzf | awk '{print $$1}')

debug_lexer: ${TARGET}
	./${TARGET} -dl -e="${EXPR}"

clean:
	-rm ${TARGET} ${DEBUG_TARGET}
