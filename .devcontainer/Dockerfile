FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

# Install additional OS packages
RUN apt-get update && export DEBIAN_FRONTEND=noninteractive \
    && apt-get -y install --no-install-recommends \
    postgresql-client \
    && apt-get clean -y \
    && rm -rf /var/lib/apt/lists/*

COPY .env* /workspace/
WORKDIR /workspace

# [Optional] Uncomment if you want to install more global node packages
# RUN npm install -g pnpm typescript ts-node