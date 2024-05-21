FROM golang:1.22.3

WORKDIR /workspace

ENTRYPOINT ["/workspace/script/build.sh"]