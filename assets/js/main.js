var botao = document.getElementById("botao");
var progressContainer = document.getElementById("progress-container");
var progressBar = document.getElementById("progress-bar");
var infoDiv = document.querySelector(".info");

botao.onclick = function() {
    // Mostra o container da barra de progresso e esconde as informações
    progressContainer.classList.remove("hidden");
    infoDiv.classList.add("hidden");

    // Reseta a barra de progresso
    progressBar.style.width = "0%";
    let width = 0;

    // Simula o progresso de "scaneamento" usando setInterval
    var scanning = setInterval(function() {
        if (width >= 100) {
            clearInterval(scanning); // Para o scaneamento quando atinge 100%
            progressContainer.classList.add("hidden"); // Esconde a barra de progresso
            infoDiv.classList.remove("hidden"); // Mostra as informações
            botao.classList.add("hidden"); // Remove o botão após o scaneamento
        } else {
            width += 10; // Incrementa a largura da barra de progresso
            progressBar.style.width = width + "%";
        }
    }, 300); // Tempo para cada incremento da barra (0.3s)
};
