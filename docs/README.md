# FrostBite

![IceDragon, Lee Kent on ArtStation](https://i.pinimg.com/originals/ee/0f/c0/ee0fc09c9df5f84c37c4d21a07a3b603.jpg)

If you want instructions on how to use the software, and aren't interested in how it works go straight to the [How to use it?](#how-to-use-it) section.

- [FrostBite](#frostbite)
  - [How it works?](#how-it-works)
    - [Overview 📚](#overview-)
    - [Symmetric vs Asymmetric encryption](#symmetric-vs-asymmetric-encryption)
      - [Symmetric](#symmetric)
      - [Asymmetric](#asymmetric)
    - [Hybrid encryption](#hybrid-encryption)
  - [How to use it?](#how-to-use-it)
    - [IMPORTANT ⚠️](#important-️)
    - [Creating public and private keys 🔑](#creating-public-and-private-keys-)
    - [Locking files 🔒](#locking-files-)
      - [Locking one folder/directory](#locking-one-folderdirectory)
      - [Locking the whole system](#locking-the-whole-system)
    - [Unlocking files 🔓](#unlocking-files-)
      - [Getting the `decrypt.key`](#getting-the-decryptkey)
      - [Using the `decrypt.key`](#using-the-decryptkey)
  - [Usecases; Why have I made it? 🤔](#usecases-why-have-i-made-it-)

## How it works?

### Overview 📚

This project consists of 2 separate programs that work as one system. Those are the **Admin** and **FrostBite** itself.

**Admin** is used to generate the keys and should be kept on a separate system together with the Private Key.

**FrostBite** is the part that encrypts (locks) the files, the same program is used to unlock the files after acquiring the AES key (password). The unlocking part will be discussed in detail.

### Symmetric vs Asymmetric encryption

To understand Hybrid cryptography which this project is based on, we first need to refresh our knowledge on Symmetric and Asymmetric cryptography and remind ourselves why they are amazing.

#### Symmetric

![Symmetric graph](png/symmetric.png)

This is the simplest kind of encryption that involves only one secret key to cipher and decipher information. Symmetric encryption is an old and best-known technique. It uses a secret key that can either be a number, a word or a string of random letters. It is a blended with the plain text of a message to change the content in a particular way. The sender and the recipient should know the secret key that is used to encrypt and decrypt all the messages.

The main disadvantage of the symmetric key encryption is that all parties involved have to exchange the key used to encrypt the data before they can decrypt it.

#### Asymmetric

![Asymmetric](png/asymmetric.png)

Asymmetric encryption is also known as public key cryptography, which is a relatively new method, compared to symmetric encryption. Asymmetric encryption uses two keys to encrypt a plain text. Secret keys are exchanged over the Internet or a large network. It ensures that malicious persons do not misuse the keys. It is important to note that anyone with a secret key can decrypt the message and this is why asymmetric encryption uses two related keys to boosting security. A public key is made freely available to anyone who might want to send you a message. The second private key is kept a secret so that you can only know.

A message that is encrypted using a public key can only be decrypted using a private key, while also, a message encrypted using a private key can be decrypted using a public key. Security of the public key is not required because it is publicly available and can be passed over the internet. Asymmetric key has a far better power in ensuring the security of information transmitted during communication.

### Hybrid encryption

---

## How to use it?

### IMPORTANT ⚠️

Keep private key on a different system and NEVER copy/generate it on the system you wish to lock.

[🔝Go to top](#frostbite)

### Creating public and private keys 🔑

### Locking files 🔒

#### Locking one folder/directory

#### Locking the whole system

### Unlocking files 🔓

#### Getting the `decrypt.key`

#### Using the `decrypt.key`

---

## Usecases; Why have I made it? 🤔
