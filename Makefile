TARGET=interpreter

${TARGET}:
	go build -o ${TARGET} -gcflags "-l -N" .

.PHONY : run debug clean

run: ${TARGET}
	./${TARGET} -i

debug: ${TARGET}
	dlv attach $(shell ps | fzf | awk '{print $$1}')

clean:
	rm ${TARGET}
