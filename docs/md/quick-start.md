# Quick-Start guide 🚀

This document is intended as a quick guide on how to use FrostBite. For a more detailed guide, please refer to the [extended documentation](extended.md).

## Terminology 📖

Here is a brief explanation of terminology used in this document.

- **Encryption** - The process of converting data into a form that cannot be read by anyone except the intended recipient. Regard it as "Locking" the data or a file.
- **Decryption** - The process of converting data back into a form that can be read by anyone. Regard it as "Unlocking" the data or a file.
- **FrostBite** - The client program that is used to encrypt and decrypt files.
- **Admin** - Complementary program to FrostBite that is used to manage the encryption keys (simpler than it sounds).

## The two approaches ☯️

### Two programs

FrostBite consists of 2 parts. The **Admin** and the **FrostBite**, these are two separate programs that work together to provide a simpler and more flexible approach to file encryption.
**Admin** is used to manage the encryption keys, including generation and decryption. **FrostBite** is used to encrypt and decrypt (lock and unlock) the files on the user's computer.

### Embedded vs. local key 💻

There are two ways to use FrostBite, one enables you to embed the encryption key into the binary file itself *(advanced)*, and the other one is going to be explained in this file. Refer to the [extended documentation](extended.md) for more information on the embedded approach.

## Getting started 🚀

### Downloading the software 📥

If you are reading this, you probably want to download the software without compiling it from source code. The download links are available on in the repository's resources section on GitHub.

Upon downloading the archive for your operating system, extract it to a folder of your choice. You can delete the archive after extracting it.

### Generating the encryption key 🔑

Before you can start encrypting files, you need to generate a key pair. This is done using the **Admin** program. To generate a key pair, run the **Admin** and it will automatically generate a key pair named `private.key` and `public.key` in the `keys` folder.

⚠️ **IMPORTANT:** Do not share the `private.key` file with anyone and do not lose it. If you lose it, you will not be able to decrypt the files you encrypted with it. It is advised that you do not store the `private.key` file on the same computer as the files you want to encrypt.

It is strongly advised that you store the `private.key` file on a USB flash drive or a cloud storage service, and have a backup of it in case you lose it.

### Encrypting a single folder 📁

FrostFire by default encrypts the folder it is located in and all of its subfolders. To encrypt a single folder, you need to move the **FrostBite** program to the folder you want to encrypt and run it together with it's `public.key` file. The **FrostBite** program will encrypt the folder and all of its subfolders once you run it.

### Encrypting the entire computer 🖥️

I warmly advise you to read the [extended documentation](extended.md) before you attempt to encrypt the entire computer to understand all limitations.

To encrypt the entire computer, you need create a file named `THIS MAY LOCK MY DATA PERMANENTLY` in the same folder as the **FrostBite** program. Don't worry, the program will not actually destroy your computer. It is just a precaution to prevent accidental encryption of the entire computer.

Once you have created the file, you can run the **FrostBite** program and it will encrypt the entire computer. This process can take a while, depending on the size of your computer.

## Decrypting files 🔓

After running the **FrostBite** program, you will notice that the files have been encrypted, and each folder has a file named `encrypted.key` in it, they are all identical by the way, there are located in every folder just as a precaution.

### Getting the `decrypted.key` file 🗝️

You will use the `encrypted.key` and `private.key` to generate `decrypted.key` which will be used to decrypt the files. This might sound complicated, but it is actually very simple, just follow the steps below.

Remember the `keys` folder that you created earlier using the **Admin** program? You are going to use it now. Move the `encrypted.key` and `private.key` (if it's not already there for some reason) files to the `keys` folder. Then, run the **Admin** program and it will generate the `decrypted.key` file in the `keys` folder. Once done, you have successfully generated the `decrypted.key` file.

### Decrypting the files (Single folder and whole system) 📂

To decrypt the files, you need to move the `decrypted.key` file to the folder you want to decrypt. Then, run the **FrostBite** program and it will decrypt the folder and all of its subfolders.

For decrypting the entire computer the process is pretty much the same. You still need to move the `decrypted.key` and **FrostBite** to the same folder (does not matter where on the system) but this time include the `THIS MAY LOCK MY DATA PERMANENTLY` file. Then run the **FrostBite** program and it will decrypt the entire computer.
