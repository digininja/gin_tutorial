// https://gilfink.medium.com/quick-tip-long-polling-using-jquery-fd26799921d

var timeout;
const MAX_CALLBACKS = 10;
var callbackCount = 0;

function doCallback(uuid) {
	callbackCount++;
	console.log ("Callback count: " + callbackCount);

	if (callbackCount > MAX_CALLBACKS) {
		alert ("Hit max callbacks, aborting");
		return;
	}

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
				if (data.status == "processing") {
					timeout = setTimeout(function() {doCallback(uuid);}, 1000);
				} else {
					if (data.status == "complete") {
						alert("success: " + data.message);
					} else {
						alert ("Something odd happened");
					}
					clearTimeout(timeout);
				}
			}
        },
        error: function(XMLHttpRequest, textStatus, errorThrown) {
            var obj = JSON.parse(XMLHttpRequest.responseText);
            alert(obj.errors[0].field);

            clearTimeout(timeout);
        }
    });
}
function submitURL() {
    var urlData = {
        "URL": $("#url").val()
    };

    console.log("Submitting: " + JSON.stringify(urlData))
	callbackCount = 0;

    $.ajax({
        type: "POST",
        url: "/submitURL",
        data: JSON.stringify(urlData),
        contentType: "application/json; charset=utf-8",
        dataType: "json",
        cache: false,
        success: function(data) {
            console.log("URL submitted, UUID: " + data.ID);
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
