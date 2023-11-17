let base, limite;
base = window.prompt("Ingrese la base de la tabla: ",0);
limite = window.prompt("Ingrese el límite de la tabla: ",0);

for (let i = 0; i < limite; i++) {
    console.log(i * base);
}