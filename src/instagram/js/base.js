$(document).ready(function() {
    console.log("Bienvenido a Instagram");
    var formularioRegistro = $("#formularioRegistro > form > input:last-child");
    console.log(formularioRegistro);

    $(formularioRegistro).click(function() {
        var palabra = $("#nombre").val();
        var usuario = $("#username").val();
        var electronico = $("#correo").val();
        var contrasena = $("#contrasena").val();
        console.log(nombre, username, correo, contrasena);
        
        
        var envio = {
            nombre: palabra,
            usuario: usuario,
            email: electronico,
            password: contrasena
        };
        console.log(envio);
        $.post({
            url:"/insert",
            data: JSON.stringify(envio),
            success: function(data, status, jqXHR) {
                console.log(data);
            },      
            dataType: "json"

        }).done(function(data) {
            console.log("Petición realizada");

        }).fail(function(data) {
            console.log("Petición fallida");
   
        }).always(function(data){
            console.log("Petición completa");
        });
    });
});

