YACC_SRC = ./parser/parser.go.y
YACC_DIST = ./parser/parser.go
TEST_FILES = $(wildcard ./**/*_test.go)

yacc: $(YACC_SRC)
	goyacc -o $(YACC_DIST) $(YACC_SRC)

test: $(YACC_DIST) $(TEST_FILES)
	go test ./...

run:
	go run main.go

clean:
	rm $(YACC_DIST)
