document.addEventListener("DOMContentLoaded", function() {
    const apiUrl = "http://localhost:5050/mostrarClientes";

    fetch(apiUrl)
        .then(response => {
            if (!response.ok) {
                throw new Error(`Error de red - ${response.status}`);
            }
            return response.json();
        })
        .then(data => {
            mostrarClientes(data);
        })
        .catch(error => {
            console.error("Error al recuperar los datos de clientes:", error);
            mostrarResultadoError();
        });
});

function mostrarClientes(clientes) {
    const clientesListaDiv = document.getElementById("clientes-lista");
    clientesListaDiv.innerHTML = "<h2>Lista de Clientes:</h2>";
    
    const listaUl = document.createElement("ul");

    clientes.forEach(cliente => {
        const itemLi = document.createElement("li");
        itemLi.innerHTML = `
            <strong>Nombre:</strong> ${cliente.nombre}<br>
            <strong>Identificación:</strong> ${cliente.identificacion}<br>
            <strong>Teléfono:</strong> ${cliente.telefono}<br>
            <strong>Correo:</strong> ${cliente.correo}<br>
        `;
        listaUl.appendChild(itemLi);
    });

    clientesListaDiv.appendChild(listaUl);
}

function mostrarResultadoError() {
    const clientesListaDiv = document.getElementById("clientes-lista");
    clientesListaDiv.innerHTML = "<h2>Error al obtener datos de clientes</h2>";
}
