$(document).ready(function() {

    ActualizarHistorial();
    $("#idioma").keyup(function(event) {
        if (event.keyCode === 13) {
            $("#btn_enviar").click();
        }
    });
    $("#btn_enviar").click(function() {
        var idioma = $("#txt_idioma").val();
        var traduccion = $("#txt_traductor").val();
        var texto = $("#txt_texto").val();
        console.log(idioma, traduccion, texto);
        
        var envio = {
            nombre: idioma
        };

        $.post({
            url:"/envioIdioma",
            data: JSON.stringify(envio),
            success: function(data, status, jqXHR) {
                console.log(data);
                $("#txt_idioma").val('')
            },
            dataType: "json"

        }).done(function(data) {
            console.log("Petición realizada");
            ActualizarHistorial();
        
        }).fail(function(data) {
            console.log("Petición fallida");
        
        }).always(function(data){
            console.log("Petición completa");
        });
    });
});

function ActualizarHistorial() {   
    $.ajax({
        url: "/listadoIdioma",
        method: "POST",
        dataType: "json",
        contentType: "application/json",
        success: function(data) {
            if(data != null)
                console.log(data.length + " objetos obtenidos");
            Historial_Idioma(data);
        },
        error: function(data) {
            console.log(data);
        }
    });
}

function Historial_Idioma(array) {
    var select = $("#idiomas select");
    select.children().remove();
    if(array != null && array.length > 0) {

        for(var x = 0; x < array.length; x++) {
            select.append(
                "<option value='>" + array[x].ID + 
                "'>" + array[x].Nombre + 
                "</option>");
        }
    } else {
        select.append('<tr><td colspan="3">No hay registros de hoy</td></tr>');
        
    }

}