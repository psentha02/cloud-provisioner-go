# 🌩️ Cloud Resource Group Manager — Go + Azure

A beginner-friendly cloud automation CLI tool built with **Go** that uses the Azure SDK to manage Resource Groups from your terminal.

This project is part of my journey into cloud development and DevOps using Go, designed for clean code, modular design, and real-world cloud integration.

---

## 🚀 Features

- 🔑 Authenticate to Azure using a Service Principal JSON file.
- 📦 Create Azure Resource Groups.
- 📋 List all Azure Resource Groups in your subscription.
- 🗑️ Delete Azure Resource Groups.
- 💡 Written with clean, extendable Go structure — perfect for future cloud automation.

---

## 🧑‍💻 Prerequisites

- [Go](https://go.dev/doc/install) installed.
- [Azure CLI](https://learn.microsoft.com/en-us/cli/azure/install-azure-cli) installed and logged in:


```bash
az login
```

- Azure Service Principal credentials JSON file generated:

```bash
az ad sp create-for-rbac --sdk-auth > azureAuth.json
```

Save the `azureAuth.json` securely in your system.

---

## ⚙️ Environment Setup

1. Move your `azureAuth.json` file to a safe place, for example:
```
$HOME/.azure/azureAuth.json
```

2. Set an environment variable to point to this file:

### Mac/Linux:
```bash
export AZURE_AUTH_LOCATION=$HOME/.azure/azureAuth.json
```

### Windows (Command Prompt):
```cmd
set AZURE_AUTH_LOCATION=C:\Users\<YourUser>\.azure\azureAuth.json
```

⚠️ **Note:**  
You must run this export command each time you open a terminal,  
or add it to your shell profile (`.bashrc`, `.zshrc`, `PowerShell profile`) for automatic loading.

---

## 💡 Usage

### Create a Resource Group:
```bash
go run ./cmd/main.go create-rg <resourceGroupName> <location>
```
Example:
```bash
go run ./cmd/main.go create-rg mygroup eastus
```

---

### List All Resource Groups:
```bash
go run ./cmd/main.go list-rg
```

---

### Delete a Resource Group:
```bash
go run ./cmd/main.go delete-rg <resourceGroupName>
```
Example:
```bash
go run ./cmd/main.go delete-rg mygroup
```

---

## 📂 Project Structure

```
cloud-provisioner-go/
├── cmd/
│   └── main.go          # Command-line entry point for the tool
├── internal/
│   └── azure/
│       └── auth.go      # Azure API logic: authentication, create, list, delete
├── go.mod               # Go module definition
├── go.sum               # Dependency checksums
└── README.md            # Project guide
```

---

## 🧪 To-Do List

- [ ] Write unit tests for Azure logic.
- [ ] Add support for creating and managing Virtual Machines.
- [ ] Add logging and error handling improvements.
- [ ] Package this as a binary executable for multiple platforms.

---

## 📌 Why This Project?

This project is part of my journey to learn **Go** and apply it to **Cloud Automation**  
as part of real-world DevOps workflows. It’s built for:

- 💡 Clean and modular code.
- 🧰 Practical Azure SDK usage.
- 📈 Scalable for future cloud tooling.
