FROM golang:1.12

# I am in /build/test
#COPY ../../ /app toto robim preto, aby som dostal cely projekt a spustil go get, ale mam forbidden path outside the build context
WORKDIR /app
RUN ["go", "get", "github.com/cespare/reflex"]
#RUN ["go", "get", "./..."] toto chcem spravit tak ze si sem nakopcim

CMD [ "reflex" ,"-c" ,"build/test/reflex.conf", "--decoration=fancy" ]
