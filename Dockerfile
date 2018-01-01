# iron/go:dev is the alpine image with the go tools added
FROM iron/go:dev

# Set an env var that matches github repo name
ENV SRC_DIR=${HOME}/gocode/src/crypto-historical-data/

# Add the source code:
ADD . $SRC_DIR
RUN rm ${SRC_DIR}/update_ticker.go

# Build it:
RUN cd $SRC_DIR;\
	go get ./;\
    go build -o api;
    
ENTRYPOINT ["/gocode/src/crypto-historical-data/api"]
