# LogParsingNewRelic

sudo apt update

sudo apt install git

git clone https://github.com/alonsoaz/LogParsingNewRelic.git

cd LogParsingNewRelic

sudo apt install curl
curl -L https://toolbelt.treasuredata.com/sh/install-debian-buster-td-agent3.sh | sh

sudo td-agent-gem install fluent-plugin-newrelic

sudo rm -f /etc/td-agent/td-agent.conf

sudo mv td-agent.conf /etc/td-agent/

sudo apt install wget

wget https://dl.google.com/go/go1.13.linux-amd64.tar.gz

#To Verify the downloaded file

sha256sum go1.13.linux-amd64.tar.gz

sudo tar -C /usr/local -xzf go1.13.linux-amd64.tar.gz

export PATH=$PATH:/usr/local/go/bin

source ~/.profile

go version

sudo td-agent-gem install fluent-plugin-grok-parser

sudo systemctl restart td-agent.service

go run scanner.go

