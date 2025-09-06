#!/bin/bash

# ==========================================================
# Script Install Golang Versi Terbaru (Linux)
# ==========================================================

# Versi terbaru (cek di https://go.dev/dl/)
GO_VERSION="1.23.1"
GO_TAR="go${GO_VERSION}.linux-amd64.tar.gz"
GO_URL="https://go.dev/dl/${GO_TAR}"

echo "=== Download Golang ${GO_VERSION} ==="
wget -q --show-progress $GO_URL

echo "=== Hapus instalasi lama (jika ada) ==="
sudo rm -rf /usr/local/go

echo "=== Extract ke /usr/local ==="
sudo tar -C /usr/local -xzf $GO_TAR

echo "=== Konfigurasi PATH permanent ==="
if ! grep -q "/usr/local/go/bin" ~/.bashrc; then
    echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
fi

if ! grep -q "\$GOPATH/bin" ~/.bashrc; then
    echo 'export GOPATH=$HOME/go' >> ~/.bashrc
    echo 'export PATH=$PATH:$GOPATH/bin' >> ~/.bashrc
fi

echo "=== Bersihkan file tar ==="
rm $GO_TAR

# =============================================
# Aktifkan PATH untuk session sekarang juga
# =============================================
export PATH=$PATH:/usr/local/go/bin
export GOPATH=$HOME/go
export PATH=$PATH:$GOPATH/bin

echo "=== Cek versi Go ==="
go version
