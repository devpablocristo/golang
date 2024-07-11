pragma solidity ^0.4.17;

//guarda un mensaje que alguien puede leer
contract Inbox {
    
    // crea un funcion llamada message, por eso getMessage es redundante
    string public message;
    
    function Inbox(string initialMessage) public {
        message = initialMessage;
    }
    
    function setMessage(string newMessage) public {
        message = newMessage;
    }
    
    /*
    redundante 
    function getMessage() public view returns (string) {
        return message;
    }
    */
}