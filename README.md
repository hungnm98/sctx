# 🧠 sctx — Simple Shell Context Switcher

> Seamlessly switch between different development environments, companies, or projects — without messing up your global shell setup.

---

## ✨ Why sctx?

As a backend/system developer, you often work with **multiple environments** or **companies**, for example:

- `project1-dev`, `project1-staging`, `project1-prod`
- `project2-dev`, `personal`, `freelance`, ...
- each with its own `AWS_PROFILE`, `KUBECONFIG`, `GIT_SSH_COMMAND`, etc.

Instead of manually exporting variables or sourcing different scripts every time,  
👉 **`sctx`** lets you switch contexts instantly, in an isolated subshell — **safe, simple, and reproducible**.

Each profile has its own configuration, prompt color, and environment variables.  
When you `exit`, you’re back to your global shell, clean and untouched.

---

## 🚀 Features

- 🧩 **Per-profile environment isolation**
- 🎨 **Colored shell prompt** showing current profile
- 🧠 **Auto-load last used context**
- 🛠 **Editable profiles** with `$EDITOR`
- 🪄 **Cross-shell support** (Zsh & Bash)
- ⚡ **Zero dependencies** (pure Go, single binary)

---

## 🧰 Installation

### 📦 Download from GitHub Release

Download the latest prebuilt binary for your platform:

👉 [https://github.com/hungnm98/sctx/releases](https://github.com/hungnm98/sctx/releases)

Example for macOS (Apple Silicon):

```bash
curl -fsSL https://github.com/hungnm98/sctx/releases/download/v0.3.0/install.sh | bash

# Reload your shell
source ~/.zshrc

# Verify installation
sctx --version
```

For Bash users, replace `~/.zshrc` with `~/.bashrc`.

---

This allows your shell to automatically apply the last active profile on startup.

---

## 💻 Usage

### Create a new profile
```bash
sctx create cty1-dev
```

### Edit your profile
```bash
sctx edit cty1-dev
```

Each profile is just a shell script.  
You can export variables, paths, aliases, etc.

Example:
```bash
# ~/.sctx/profiles/cty1-dev
export AWS_PROFILE=company1-dev
export KUBECONFIG=~/.kube/cty1-dev.yaml
export GIT_SSH_COMMAND="ssh -i ~/.ssh/cty1_dev"

# optional startup message
echo "[cty1-dev] Environment loaded"
```

### List all profiles
```bash
sctx ls
```

### Switch to a profile
```bash
sctx use cty1-dev
```
→ Opens a new subshell with red `[cty1-dev]` prefix in your prompt.  
Type `exit` to return to your main shell.

### Interactive switcher
```bash
sctx ctx
```
→ Select profiles using arrow keys (like `kubie ctx`).

### Set default profile
```bash
sctx default cty1-dev
```

### Unset current profile
```bash
sctx unset
```

---

## 🧩 Example Workflow

| Action | Command | Result |
|--------|----------|--------|
| Enter company 1 dev env | `sctx use cty1-dev` | `[cty1-dev] user@macbook %` |
| Switch to staging | `sctx use cty1-staging` | `[cty1-staging] user@macbook %` |
| Work on freelance project | `sctx use freelance` | `[freelance] user@macbook %` |
| Return to normal shell | `exit` | back to global env |

Each context has its own `$PATH`, `$AWS_PROFILE`, `$KUBECONFIG`, `$SSH_AUTH_SOCK`, etc.  
No accidental cross-contamination between environments.

---

## 🧑‍💻 Editor Support

`$EDITOR` is used if set.  
Otherwise, `sctx` will ask you once and save your choice in `~/.sctx/config.yaml`.

---

## 🏗 Directory Structure

```
~/.sctx/
├── bin/               # binary location
├── config.yaml        # user config (editor, preferences)
├── current            # name of the active profile
└── profiles/
    ├── project1-dev
    ├── project2-staging
    └── freelance
```

---

## 🧾 License

MIT © [Hung Nguyen](https://github.com/hungnm98)

---

## 💬 Inspiration

Inspired by [`kubie`](https://github.com/sbstp/kubie), [`direnv`](https://github.com/direnv/direnv`), and [`asdf`](https://github.com/asdf-vm/asdf`) —  
but designed to be **simpler**, **shell-agnostic**, and focused on **per-company / per-project** development environments.
