#!/bin/sh
# -*- tab-width:4;indent-tabs-mode:nil -*-
# ex: ts=4 sw=4 et

make && make rel && cd rel/emqttd && ./bin/emqttd console
