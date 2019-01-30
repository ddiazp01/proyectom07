$(document).ready (function () {

    $("#txt-texto").keyup(function(event){
        if (event.keyCode === 13){
            $("#btn-enviar").click();
        }
    });
    
    $("#btn-enviar").click(function(){
        var texto =  $("#txt-texto").val();
        var fecha = new Date();
        var peticion= {
            
            Palabra: texto,
            Fecha: fecha.getDate() + '/' + fecha.getMonth()+1 + '/'+ fecha.getFullYear()
        }


        $.ajax({
            type: "POST",
            url: "http://localhost:8080/ejercicio",
            data: JSON.stringify(peticion),
            dataType: "JSON",
            contentType:"application/json",
        })
        .done(function(data) {
            $("#respuesta p").html(data.Palabra);
        })
        .fail(function(data) {
            console.log( "La solicitud no se ha completado correctamente." );
        })
        .always(function(data){
            console.log("complete")
        })
    });
});
