# AGoTTHT
**A**nother **Go**, **T**ailwind, **T**empl, **H**TMX **T**emplate

---

## Quickstart Guide

Follow these steps to get your development environment up and running.

### 1. Install Go 
* **Download & Learn:** [go.dev/learn](https://go.dev/learn/)
initialize your Go module:
```bash
go mod init <directory>
```

### 2. Install Templ
Install the Templ CLI tool
```bash
go install [github.com/a-h/templ](https://github.com/a-h/templ)
```
*(Note: to add Templ to your project dependencies, also run `go get github.com/a-h/templ`)*

### 3. Install Tailwind CSS
You can install Tailwind CSS either via NPM or by using the standalone executable. The makefile by defualt is set to use the standalone executable

#### Option A: Via NPM
Initialize NPM and install the Tailwind CLI:

```bash
npm init -y
npm install -D tailwindcss @tailwindcss/cli
```

Run the CLI to compile your CSS:

```bash
npx @tailwindcss/cli -i ./assets/css/input.css -o ./static/css/styles.css
```

#### Option B: Standalone Executable
If you prefer to work without Node.js, you can download the standalone executable.
* **Download:** [Tailwind CSS Releases on GitHub](https://github.com/tailwindlabs/tailwindcss/releases)

After downloading the correct executable for your operating system, rename it and move it to your binaries folder (Linux/macOS example), this can and should be skipped if already done before:

```bash
sudo mv tailwindcss /usr/local/bin/
```

### 4. Live Reloading with Air
To get live reloading for your Go server, you can use [Air](https://github.com/air-verse/air).

Install Air globally:

```bash
go install [github.com/air-verse/air@latest](https://github.com/air-verse/air@latest)
```

Initialize Air in your project directory (this creates a `.air.toml` config file):

```bash
air init
```

Start your development server with live reloading:

```bash
air
```

Assuming all went well for you the command:
```bash
make dev
```
should spin your localhost to life on :8080