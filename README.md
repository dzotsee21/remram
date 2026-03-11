<img width="330" height="82" alt="image" src="https://github.com/user-attachments/assets/1386e20d-3af9-4d84-ae30-3f6a7a13da2b" />


# Remram - Remote Device Controller
Remram is a lightweight Go utility that turns another device into a remote trackpad and keyboard for your computer.

It is mainly designed for phone → PC control, but it also works for PC → PC control.

Remram currently supports Linux (X11 & Wayland)

## Installation:
```bash
> chmod +x setup.sh
> ./setup.sh
```
Replace `DEVICE_IP` with the local IP address of your phone. (keep in mind that your DEVICE_IP may change)
<br />
```bash
> sudo ufw allow from DEVICE_IP to any port 8080
```

## How to run

start the server:
```bash
> go run .
```

## Todo
- smoother scrolling
- customizable settings
- less installation steps
- support for other platforms (Windows and MacOS primarily)

## Credits
Remram relies on the [uinput](https://github.com/bendahl/uinput) package to emulate keyboard and mouse input.
