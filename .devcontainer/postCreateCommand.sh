#!/bin/bash

# Install Node.js dependencies. Used primarily for frontend.
npm i

# Install Air for hot rebuilding.
# See: https://github.com/air-verse/air
go install github.com/air-verse/air@latest
