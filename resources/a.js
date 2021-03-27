var timeout; // this is a variable the timout will be assigned to
var init = 1;
var timestamp = null;

function doCallback(uuid) {
    var urlData = {
        "UUID": uuid
    };

    console.log("Submitting: " + JSON.stringify(urlData))

    $.ajax({
        type: "POST",
        url: "/callback",
        data: JSON.stringify(urlData),
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        cache: false,
        success: function(data) {
			if (data.status == "error") {
				alert (data.message);
			} else {
				if (data.status == "not ready yet") {
					timeout = setTimeout(function() {doCallback(uuid);}, 1000); // here we assign the timout
				} else {
					if (data.status == "ready") {
						alert("success: " + data.message);
					} else {
						alert ("Something odd happened");
					}
					clearTimeout(timeout); // here we clear the timeout
				}
			}
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
function submitURL() {
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
			doCallback(data.ID);
        },
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            var obj = JSON.parse(XMLHttpRequest.responseText);
			alert ("URL not accepted");
            alert(obj.errors[0].field);
        }
    });
}

$(function() {
    $("#submit").on('click', function() {
        submitURL();
    });
});
