#!/bin/bash

# Function to pause and wait for user input
function pause() {
    read -n1 -rsp $'Press any key to continue...\n'
}
echo "go run ./cmd/ https://pbs.twimg.com/media/EMtmPFLWkAA8CIS.jpg"
pause
go run ./cmd/ https://pbs.twimg.com/media/EMtmPFLWkAA8CIS.jpg
pause
clear

echo "go run ./cmd/ https://reboot01.com"
pause
go run ./cmd/ https://reboot01.com
pause
clear

echo "go run ./cmd/ https://golang.org/dl/go1.16.3.linux-amd64.tar.gz"
pause
go run ./cmd/ https://golang.org/dl/go1.16.3.linux-amd64.tar.gz
pause
clear

echo "go run ./cmd/ http://ipv4.download.thinkbroadband.com/100MB.zip"
pause
go run ./cmd/ t http://ipv4.download.thinkbroadband.com/100MB.zip
pause
clear

echo "go run ./cmd/ -O=test_20MB.zip http://ipv4.download.thinkbroadband.com/20MB.zip"
pause
go run ./cmd/ -O=test_20MB.zip http://ipv4.download.thinkbroadband.com/20MB.zip
pause
clear

echo "go run ./cmd/ -O=test_20MB.zip -P=~/Downloads/ http://ipv4.download.thinkbroadband.com/20MB.zip"
pause
go run ./cmd/ -O=test_20MB.zip -P=~/Downloads/ http://ipv4.download.thinkbroadband.com/20MB.zip
pause
clear

echo "go run ./cmd/ --rate-limit=300k http://ipv4.download.thinkbroadband.com/20MB.zip"
pause
go run ./cmd/ --rate-limit=300k http://ipv4.download.thinkbroadband.com/20MB.zip
pause
clear

echo "go run ./cmd/ --rate-limit=700k http://ipv4.download.thinkbroadband.com/20MB.zip"
pause
go run ./cmd/ --rate-limit=700k http://ipv4.download.thinkbroadband.com/20MB.zip
pause
clear

echo "go run ./cmd/ --rate-limit=2M http://ipv4.download.thinkbroadband.com/20MB.zip"
pause
go run ./cmd/ --rate-limit=2M http://ipv4.download.thinkbroadband.com/20MB.zip
pause
clear

echo "go run ./cmd/ -i=downloads.txt"
pause
go run ./cmd/ -i=downloads.txt
pause
clear

echo "go run ./cmd/ -B http://ipv4.download.thinkbroadband.com/20MB.zip"
pause
go run ./cmd/ -B http://ipv4.download.thinkbroadband.com/20MB.zip
pause
clear

echo "go run ./cmd/ --mirror --convert-links http://corndog.io/"
pause
go run ./cmd/ --mirror --convert-links http://corndog.io/
pause
clear

echo "go run ./cmd/ --mirror https://oct82.com/"
pause
go run ./cmd/ --mirror https://oct82.com/
pause
clear

echo "go run ./cmd/ --mirror --reject=gif https://oct82.com/"
pause
go run ./cmd/ --mirror --reject=gif https://oct82.com/
pause
clear

echo "go run ./cmd/ --mirror https://trypap.com/"
pause
go run ./cmd/ --mirror https://trypap.com/
pause
clear

echo "go run ./cmd/ --mirror -X=/img https://trypap.com/"
pause
go run ./cmd/ --mirror -X=/img https://trypap.com/
pause
clear

echo "go run ./cmd/ --mirror https://theuselessweb.com/"
pause
go run ./cmd/ --mirror https://theuselessweb.com/
pause
clear

echo "go run ./cmd/ --mirror <https://link_of_your_choice.com>"
pause
