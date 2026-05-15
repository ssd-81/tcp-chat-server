## session 1
### what do I need to establish before getting started? 
To begin this, I need a simple thing working; two nodes (I don't know how I will simulate that) to talk to each other. 

### where do I begin? 
- what can I do to connect two computers? 
- that is a big question? what is something I can test right now? 
- how do I simulate two computers on the same system 
    while researching this topic, I came across a topic called sockets which can be used to do this. 
- I don't have any idea what sockets are and how to use them, but I guess this is a starting point. 
- **AF_INET** sockets and **AF_LOCAL** sockets seem like something which could be used for the chat application for now. 
- sockets are likely the endpoints used for communication between two applications over a network. 

### how to setup two nodes on the same system? 
- the socket trajectory seems like promising for now. 
- I will try to understand this and build on top of that. 
- while trying to research sockets, I came across something called **WebSockets** and the author said that they are a good choice for real time systems. I am building a chat application; what is more real time than a chat. So, websockets are likely the answer. 
- to connect two sockets; one socket needs to be a server and one needs to be client. 
- let's start coding a bit. 


### should I run client and server in separate files? 
- I think to first test out the main idea, I should just get them working state, from there we will try to optimize it further. 
- Current task is simple: get client to send data to server and server confirming back. 
- I am not sure if the flow that I have designed is working as expected. I need a way to check if the model I have of flow is correct or not? 


## session 2
### code fixing session
- I have some code for now, but right now I am clearly unaware of what is the flow and what each component is doing. I will be spending some time on figuring out assumptions and what needs to be done. 
- one thing I understand pretty clearly; listen function creates a server and dial function is used by a client to connect to a server. so I likely need to setup the server first. 
 