var timeout; // this is a variable the timout will be assigned to
var init = 1;
var timestamp = null;

function waitForMsg() {
    var urlData = {
        "URL": $("#url").val()
    };

    console.log("Submitting: " + JSON.stringify(urlData))

    $.ajax({
        type: "POST",
        url: "/robin",
        data: JSON.stringify(urlData),
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        cache: false,
        success: function(data) {
            clearTimeout(timeout); // here we clear the timeout
            alert("success: " + data.ID);
            /*
            		var json = eval('('+data+ ')');
            		var str = json['out'].replace(/\n/g, "<br>");
            		tag.html(str).dialog({
            			title: 'Ping',
            			modal: false,
            			width: 480,
            			height: 550,
            			close: function() {
            				alert('close');
            				clearTimeout(timeout); // here we clear the timeout
            			}
            		}).dialog('open');
            		timestamp = json["tsp"];
            		init = 0;
            		timeout = setTimeout("waitForMsg()", 1000); // here we assign the timout
            */
        },
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            var obj = JSON.parse(XMLHttpRequest.responseText);
            alert(obj.errors[0].field);

            clearTimeout(timeout); // here we clear the timeout
            //		 setTimeout("waitForMsg()", 15000);
        }
    });
}
$(function() {
    $("#submit").on('click', function() {
        waitForMsg();
    });
});