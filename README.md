<p align="center">
  <a href="" rel="noopener">
 <img width=200px height=200px  alt="net cat logo"></a>
</p>

<h3 align="center">NET CAT SERVER<br>
---
<p align="center"> This project consists on recreating the NetCat in a Server-Client Architecture that can run in a server mode on a specified port listening for incoming connections, and it can be used in client mode, trying to connect to a specified port and transmitting information to the server.
    <br> 
</p>

## 📝 Table of Contents

- [About](#about)
- [Usage](#usage)
- [Description](https://learn.zone01dakar.sn/git/root/public/src/branch/master/subjects/groupie-tracker/search-bar)
- [Authors](#authors)
- [Implementation details](#built_using)

## 🧐 About <a name = "about"></a>



    NetCat, nc system command, is a command-line utility that reads and writes data across network connections using TCP or UDP. It is used for anything involving TCP, UDP, or UNIX-domain sockets, it is able to open TCP connections, send UDP packages, listen on arbitrary TCP and UDP ports and many more.

    To see more information about NetCat inspect the manual man nc.


###  Description

Instructions
    

   Your project must work in a similar way that the original NetCat works, in other words, you must create a group chat. The project must have the following features :

    TCP connection between server and multiple clients (relation of 1 to many).
    A name requirement to the client.
    Control connections quantity.
    Clients must be able to send messages to the chat.
    Do not broadcast EMPTY messages from a client.
    Messages sent, must be identified by the time that was sent and the user name of who sent the message, example : [2020-01-20 15:48:41][client.name]:[client.message]
    If a Client joins the chat, all the previous messages sent to the chat must be uploaded to the new Client.
    If a Client connects to the server, the rest of the Clients must be informed by the server that the Client joined the group.
    If a Client exits the chat, the rest of the Clients must be informed by the server that the Client left.
    All Clients must receive the messages sent by other Clients.
    If a Client leaves the chat, the rest of the Clients must not disconnect.
    If there is no port specified, then set as default the port 8989. Otherwise, program must respond with usage message: [USAGE]: ./TCPChat $port

    

    Your project must be written in Go

    Start TCP server, listen and accept connections

    Your project must have Go-routines

    Your project must have channels or Mutexes

    Maximum 10 connections

    The code must respect the good practices

    It is recommended to have test files for unit testing both the server connection and the client.

    You have to be able to handle the errors from server side and client side


## ✍️ Authors <a name = "authors"></a>

Author:  ZONE 01 COPYRIGHT

ymadike
daiba
ndiba

## 🎈 Usage <a name="usage"></a>

To run this project you to do :
- copy that command to start the server
```
go run . and find an specific port or use the defult port 8989
```
After that we have a pssibilité to create client ant try to discusse whenbthe the server is started

## ⛏️ Implementation details <a name = "built_using"></a>

This project will help you learn about :

    Manipulation of structures.
    Net-Cat
    TCP/UDP
        TCP/UDP connection
        TCP/UDP socket
    Go concurrency
        Channels
        Goroutines
    Mutexes
    IP and ports


that all for this project 