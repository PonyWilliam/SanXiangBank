#!/bin/bash
set -ex
make
./Seckill --config=config.conf
