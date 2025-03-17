function changeMessage(){
    const messageElement = document.getElementById("message");
    const currentMessage = messageElement.textContent;

    const alternateMessages = [
        "Gin es mejor que Revel",
        "Edicion limitada",
        "chromium es un monopolio",
        "Error de capa 8"
    ];

    let newMessage;

    do {
        newMessage = alternateMessages[Math.floor(Math.random() * alternateMessages.length)];
    } while (newMessage === currentMessage);

    messageElement.textContent = newMessage;
}