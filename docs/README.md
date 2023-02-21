# FrostBite

![IceDragon, Lee Kent on ArtStation](https://i.pinimg.com/originals/ee/0f/c0/ee0fc09c9df5f84c37c4d21a07a3b603.jpg)

If you want instructions on how to use the software, and aren't interested in how it works go straight to the [How to use it?](#how-to-use-it) section.

- [FrostBite](#frostbite)
  - [What is FrostBite?](#what-is-frostbite)
    - [So what is special with Frostbite? 🥶](#so-what-is-special-with-frostbite-)
      - [Cross-platform and cross-architecture support](#cross-platform-and-cross-architecture-support)
      - [Performance](#performance)
      - [Public key cryptography](#public-key-cryptography)
  - [How it works?](#how-it-works)
    - [Overview 📚](#overview-)
    - [Symmetric vs Asymmetric encryption](#symmetric-vs-asymmetric-encryption)
      - [Symmetric](#symmetric)
      - [Asymmetric](#asymmetric)
    - [Hybrid encryption](#hybrid-encryption)
  - [How to use it?](#how-to-use-it)
    - [IMPORTANT ⚠️](#important-️)
    - [Building from source vs using prebuilt binaries](#building-from-source-vs-using-prebuilt-binaries)
    - [Creating public and private keys 🔑](#creating-public-and-private-keys-)
    - [Locking files 🔒](#locking-files-)
      - [Locking one folder/directory](#locking-one-folderdirectory)
      - [Locking the whole system](#locking-the-whole-system)
    - [Unlocking files 🔓](#unlocking-files-)
      - [Getting the `decrypt.key`](#getting-the-decryptkey)
      - [Using the `decrypt.key`](#using-the-decryptkey)
    - [Limitations](#limitations)
  - [Use cases; Why have I made it? 🤔](#use-cases-why-have-i-made-it-)

## What is FrostBite?

Frostbite is a file encryption software targeting activists, journalists, and anyone who needs their data protected in transit.

### So what is special with Frostbite? 🥶

#### Cross-platform and cross-architecture support

Being written in Go, FrostBite can easily be compiled and run on Linux as well as Windows systems. It supports x64 as well as x86 for older systems, on Linux it even supports ARM for both 64 and 32 bit systems.

Besides the binary itself, it automatically detects if it's run on Linux or Windows and adjusts the directory scanning accordingly.

#### Performance

Besides Go being in the C neighborhood when it comes to performance, there a few more tricks that FrostBite uses to gain an edge.\
While scanning the system, each disk is designated to a separate Go worker which for simplicity can be thought of as a separate CPU thread. That means that as long as your system has more CPU threads (Usually 2x more than CPU cores) than disks, scanning will take as much time as the slowest disk read time.

Encryption and decryption use the same approach and delegates each file to a separate worker. Number of workers is limited to avoid freezing the system and hogging all the resources, but on HDD systems the bottleneck will probably be the disk anyways.

#### Public key cryptography

Unlike regular encryption software with which you can lock and unlock your files using the same password/key, frostbite generates a new key on each run. The "new key for each lock" makes it impossible to prepare the unlock key in advance and can delegate the Locking and Unlocking to separate parties.

## How it works?

### Overview 📚

This project consists of 2 separate programs that work as one system. Those are the **Admin** and **FrostBite** itself.

🤵‍♂️ **Admin** is used to generate the keys and should be kept on a separate system together with the Private Key.

🥶 **FrostBite** is the part that encrypts (locks) the files, the same program is used to unlock the files after acquiring the AES key (password). The unlocking part will be discussed in detail.  
[🔝Go to top](#frostbite)

### Symmetric vs Asymmetric encryption

To understand Hybrid cryptography which this project is based on, we first need to refresh our knowledge on Symmetric and Asymmetric cryptography and remind ourselves why they are amazing.

#### Symmetric

![Symmetric graph](png/symmetric.png)

**Symmetric encryption** is the most common and wide-spread kind of encryption. It uses only one secret key to cipher and decipher the data. The secret key can be provided by a user and treated as a standard password, or it can be randomly generated.

It's **advantage** is it's simplicity of having only one static key and the ability to encrypt large amounts of data, making it perfect for file encryption.

It's **disadvantages** is that there is no safe way to transport the key. If you want to send an encrypted message, you would need to send the key along with it so that the receiver can read it. Sending the key defeats the purpose of encrypting it in the first place.

#### Asymmetric

![Asymmetric](png/asymmetric.png)

Asymmetric encryption is also known as public key cryptography, which is a relatively new method, compared to symmetric encryption. Asymmetric encryption uses two keys to encrypt a plain text. Secret keys are exchanged over the Internet or a large network. It ensures that malicious persons do not misuse the keys. It is important to note that anyone with a secret key can decrypt the message and this is why asymmetric encryption uses two related keys to boosting security. A public key is made freely available to anyone who might want to send you a message. The second private key is kept a secret so that you can only know.

Asymmetric encryption takes a different approach. Instead of having a single key for encrypting and decrypting, we generate a key pair consisting of a Private key, and a Public key. Public key is used only for encrypting the message, and the private key is used only for decrypting the message.

**How is asymmetric encryption used?**
The public key is given to anyone who wishes to send us a message, and upon receiving it we are free to decrypt it using our private key. The private key should NEVER ever be made available to the public

A message that is encrypted using a public key can only be decrypted using a private key, while also, a message encrypted using a private key can be decrypted using a public key. Security of the public key is not required because it is publicly available and can be passed over the internet. Asymmetric key has a far better power in ensuring the security of information transmitted during communication.

[🔝Go to top](#frostbite)

### Hybrid encryption

![Hybrid](png/hybrid.png)

Using both asymmetric and symmetric encryption methods. A secret key is generated and the data are encrypted using the newly generated key (symmetric method). The data are sent to the recipient along with the key via the public key method (asymmetric method). Recipients use their private key to decrypt the secret key, which is then used to decrypt the message. See cryptography.

---
[🔝Go to top](#frostbite)

## How to use it?

### IMPORTANT ⚠️

Keep private key on a different system and NEVER copy/generate it on the system you wish to lock.

### Building from source vs using prebuilt binaries

### Creating public and private keys 🔑

### Locking files 🔒

#### Locking one folder/directory

#### Locking the whole system

### Unlocking files 🔓

#### Getting the `decrypt.key`

#### Using the `decrypt.key`

---

### Limitations

| Limitation | Description | Reason |
| ---------- | ---------- | ---------- |
| File size | Limited file size to 1GB | I determined that most files larger than 1GB are usually .iso, VM, temp or similar useless files that take too much time to process. There is a plan to add an option which file size should be excluded before running the encryption as I do understand that some video files, archives and large documents will be excluded, but this might put unneeded strain on older systems. |
| Whole system scan Windows 🪟 | Can only scan `C:/Users` on C: disk | Safety feature to avoid encrypting vital system files. Will be fixed in the future, but not bricking the system is a priority |
| Whole system scan Linux 🐧 | Can only scan `/home` on main disk | Same as Windows, but will probably be mediated earlier |

## Use cases; Why have I made it? 🤔
