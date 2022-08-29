## DevSpace Visual Studio Code Example

For a full workthrough, please take a look at the [guide in the DevSpace docs](https://devspace.sh/docs/ide-integration/visual-studio-code).

## Prerequisites

The following components need to be installed before you can use DevSpace with Visual Studio Code:
- [DevSpace](https://devspace.sh/docs/getting-started/installation)
- A Kubernetes cluster either locally (e.g. Docker Desktop, Rancher Desktop, minikube etc.) or in a Cloud Environment (e.g. GKE, AKS, EKS etc.) and a **valid** Kubernetes context configured locally.
- [Visual Studio Code](https://code.visualstudio.com/)
- [Remote - SSH Extension for Visual Studio Code](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-ssh)
- [Visual Studio Code - Command Line Interface](https://code.visualstudio.com/docs/editor/command-line)

## TL;DR

Run the following commands in a terminal
```
# Clone the example project
git clone https://github.com/loft-sh/devspace-vscode-example.git

# Switch to the folder
cd devspace-vscode-example

# Open Visual Studio Code in a Container
devspace dev -n my-namespace
```

