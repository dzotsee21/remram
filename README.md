<img width="330" height="82" alt="image" src="https://github.com/user-attachments/assets/1386e20d-3af9-4d84-ae30-3f6a7a13da2b" />


# Remram - Remote Device Controller
Remram is a lightweight Go utility that turns another device into a remote trackpad and keyboard for your computer.

It is mainly designed for phone → PC control, but it also works for PC → PC control.

Remram supports Linux (X11 & Wayland) & Windows.

## Requirements:
Remram assumes you have Golang, MSYS2 and GCC installed on your hosting device.

# Linux
## Installation:
```bash
> chmod +x setup.sh
> ./setup_linux.sh
```
Replace `DEVICE_IP` with the local IP address of your phone. (keep in mind that your `DEVICE_IP` may change)
<br />
```bash
> sudo ufw allow from DEVICE_IP to any port 8080
```

## How to run

start the server:
```bash
> go run .
```

# Windows
## Installation:
Open Powershell, Replace `DEVICE_IP` with the local IP address of your phone. (keep in mind that your `DEVICE_IP` may change)
<br />
```bash
> New-NetFirewallRule -DisplayName "Allow Remram Access" `
    -Direction Inbound `
    -Action Allow `
    -Protocol TCP `
    -LocalPort 8080 `
    -RemoteAddress 'DEVICE_IP'
```

## How to run

start the server:
```bash
> go build -o remram.exe
> remram.exe
```

## Todo
- smoother scrolling
- [x] customizable settings
- less installation steps
- [x] support for other platforms (Windows and MacOS primarily)

## Credits
Remram relies on the [uinput](https://github.com/bendahl/uinput) and [RobotGo](https://github.com/go-vgo/robotgo) packages to emulate keyboard and mouse input.
