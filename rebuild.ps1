docker build -t capreturns:dev -q .
waitfor /T 10
docker run --name capreturns -it -v C:\Users\hildebrandt\Documents\code\Go\Capture-Returns\:/workdir -u dev capreturns:dev bash