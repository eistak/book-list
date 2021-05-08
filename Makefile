NAME=book-list
VERSION=0.1

build:
    docker build -t $(NAME):$(VERSION) .

restart: stop start

start:
    docker run -itd \
        -p xxx:xxx \
        -v xxx:xxx \
        --name $(NAME) \
        $(NAME):$(VERSION) bash
