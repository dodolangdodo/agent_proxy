#!/bin/bash
# Install Agent Proxy V2 as a systemd service
# Usage: sudo ./install-service.sh

set -e

INSTALL_DIR="/opt/agent-proxy-v2"
SERVICE_NAME="agent-proxy"
SOURCE_DIR="$(cd "$(dirname "$0")" && pwd)"
RUN_USER="${SUDO_USER:-$(whoami)}"

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

log_info() {
    echo -e "${GREEN}[INFO]${NC} $1"
}

log_warn() {
    echo -e "${YELLOW}[WARN]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check if running as root
if [ "$EUID" -ne 0 ]; then
    log_error "Please run as root or with sudo: sudo $0"
    exit 1
fi

# Detect Go (check SUDO_USER's home first since we run as root)
GO_CMD=""
SUDO_USER_HOME="$(eval echo ~"$RUN_USER")"

if [ -x "$SUDO_USER_HOME/go-local/go/bin/go" ]; then
    GO_CMD="$SUDO_USER_HOME/go-local/go/bin/go"
    export PATH="$SUDO_USER_HOME/go-local/go/bin:$PATH"
    export GOTOOLCHAIN=local
elif [ -x "$HOME/go-local/go/bin/go" ]; then
    GO_CMD="$HOME/go-local/go/bin/go"
    export PATH="$HOME/go-local/go/bin:$PATH"
    export GOTOOLCHAIN=local
elif [ -x "/usr/local/go/bin/go" ]; then
    GO_CMD="/usr/local/go/bin/go"
    export PATH="/usr/local/go/bin:$PATH"
elif command -v go &> /dev/null; then
    GO_CMD="go"
fi

# Build binary if Go is available
if [ -n "$GO_CMD" ]; then
    log_info "Building proxy binary with Go ($($GO_CMD version))..."
    cd "$SOURCE_DIR"
    $GO_CMD build -o proxy ./cmd/proxy/
    log_info "Build successful."
else
    log_warn "Go not found. Using existing binary if available."
fi

# Check if binary exists
if [ ! -f "$SOURCE_DIR/proxy" ]; then
    log_error "Proxy binary not found at $SOURCE_DIR/proxy"
    log_info "Install Go and build: cd $SOURCE_DIR && go build -o proxy ./cmd/proxy/"
    exit 1
fi

# Stop existing service BEFORE copying binary (avoids "Text file busy")
if systemctl is-active --quiet "$SERVICE_NAME" 2>/dev/null; then
    log_warn "Stopping existing $SERVICE_NAME service..."
    systemctl stop "$SERVICE_NAME"
fi

# Create install directory
log_info "Creating install directory: $INSTALL_DIR"
mkdir -p "$INSTALL_DIR"

# Copy binary
log_info "Copying proxy binary..."
cp "$SOURCE_DIR/proxy" "$INSTALL_DIR/"
chmod +x "$INSTALL_DIR/proxy"

# Copy state file if it exists
if [ -f "$SOURCE_DIR/proxy_state.json" ]; then
    log_info "Copying existing state file..."
    cp "$SOURCE_DIR/proxy_state.json" "$INSTALL_DIR/"
fi

# Copy web UI if it exists
if [ -d "$SOURCE_DIR/web/dist" ]; then
    log_info "Copying web UI..."
    mkdir -p "$INSTALL_DIR/web"
    cp -r "$SOURCE_DIR/web/dist" "$INSTALL_DIR/web/"
fi

# Set ownership
log_info "Setting ownership to $RUN_USER:$RUN_USER..."
chown -R "$RUN_USER:$RUN_USER" "$INSTALL_DIR"

# Stop existing service if running
if systemctl is-active --quiet "$SERVICE_NAME" 2>/dev/null; then
    log_warn "Stopping existing $SERVICE_NAME service..."
    systemctl stop "$SERVICE_NAME"
fi

# Disable old service if it exists
if systemctl is-enabled --quiet "$SERVICE_NAME" 2>/dev/null; then
    log_warn "Disabling old $SERVICE_NAME service..."
    systemctl disable "$SERVICE_NAME"
fi

# Write systemd service file
log_info "Installing systemd service..."
cat > "/etc/systemd/system/${SERVICE_NAME}.service" << EOF
[Unit]
Description=Agent Proxy V2 - AI Gateway
After=network-online.target
Wants=network-online.target

[Service]
Type=simple
User=${RUN_USER}
Group=${RUN_USER}
WorkingDirectory=/opt/agent-proxy-v2
ExecStart=/opt/agent-proxy-v2/proxy
Restart=always
RestartSec=5

Environment="PROXY_LISTEN_ADDR=:9999"
Environment="PROXY_ADMIN_KEY=admin-change-me"
Environment="PROXY_LOG_LEVEL=info"

[Install]
WantedBy=multi-user.target
EOF

# Reload systemd
log_info "Reloading systemd..."
systemctl daemon-reload

# Enable and start service
log_info "Enabling $SERVICE_NAME service..."
systemctl enable "$SERVICE_NAME"

log_info "Starting $SERVICE_NAME service..."
systemctl start "$SERVICE_NAME"

# Wait a moment for service to start
sleep 2

# Check status
if systemctl is-active --quiet "$SERVICE_NAME"; then
    log_info "Service is running!"
    systemctl status "$SERVICE_NAME" --no-pager
else
    log_error "Service failed to start. Check logs:"
    journalctl -u "$SERVICE_NAME" --no-pager -n 20
    exit 1
fi

log_info "Installation complete!"
log_info "Proxy running on: http://localhost:9999"
log_info "Admin API: http://localhost:9999/api/health"
log_info "Web UI: http://localhost:9999/ui/"
log_info ""
log_info "To check status: sudo systemctl status $SERVICE_NAME"
log_info "To view logs: sudo journalctl -u $SERVICE_NAME -f"
log_info "To restart: sudo systemctl restart $SERVICE_NAME"
