# Free SSL on Kubernetes with cert-manager

SSL certificates is crucial component ensuring secure communication between server and client. Today we will secure out Kubernetes cluster Ingress traffic with a free SSL certificate issued from **Let's Encrypt** using cert-manager.

## What is cert-manager?

Cert-manager is a native Kubernetes certificate management controller. It can help with issuing certificates from a variety of sources, such as Let’s Encrypt, HashiCorp Vault etc.

<br/>

It will ensure certificates are valid and up to date, and attempt to renew certificates at a configured time before expiry.

<br/>

