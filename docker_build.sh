#/bin/bash

docker build -t mehanizm/valera .
docker push mehanizm/valera

# docker run --name=valera -v /Users/mike_berezin/go/src/github.com/mehanizm/valera/configs:/configs mehanizm/valera