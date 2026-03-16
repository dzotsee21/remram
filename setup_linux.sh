#!/bin/bash

RULE_FILE="/etc/udev/rules.d/99-remram-uinput.rules"
RULE_CONTENT='KERNEL=="uinput", GROUP="input", MODE="0660", OPTIONS+="static_node=uinput"'

echo "--- RemRam Linux Environment Setup ---"

echo "[*] Creating udev rule at $RULE_FILE..."
echo "$RULE_CONTENT" | sudo tee $RULE_FILE > /dev/null

if ! getent group input > /dev/null; then
    echo "[*] Creating 'input' group..."
    sudo groupadd input
fi

echo "[*] Adding user $USER to the 'input' group..."
sudo usermod -aG input $USER

echo "[*] Reloading udev rules..."
sudo udevadm control --reload-rules
sudo udevadm trigger

echo "--- Setup Complete ---"
echo "IMPORTANT: You MUST log out and log back in for group changes to take effect."
