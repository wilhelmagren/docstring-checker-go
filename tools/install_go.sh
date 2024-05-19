#!/usr/bin/env bash

wget https://go.dev/dl/go1.22.3.linux-amd64.tar.gz;
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.22.3.linux-amd64.tar.gz;
rm go1.22.3.linux-amd64.tar.gz;

printf "Now you need to run 'export PATH=$PATH:/usr/local/go/bin'";
exit 0;

