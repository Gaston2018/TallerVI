$(document).ready(function(){

    $("#puto").click(function(){
        $.get("../public/js/Resume.txt", function(dato, status, xhr){
            console.log("Mensaje: "+ dato+ "\nStatus: " + status + " "+xhr.status);
        });
        return false;
    });
});