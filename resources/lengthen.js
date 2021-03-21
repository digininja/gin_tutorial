// https://gilfink.medium.com/quick-tip-long-polling-using-jquery-fd26799921d

function submitURL() {
	alert ("x");
	longPoll();
	return false;
}

function longPoll(){
    $.ajax({ 
        url: "/robin",
        success: function(data){
            alert ("success");
        },
        error: function(err) {
            alert ("error");
        },
        type: "POST", 
        //dataType: "json", 
        complete: longPoll,
        timeout: 60000 // timeout every one minute
    });
}

// https://www.shift8web.ca/2015/05/ajax-polling-to-your-restful-api/

function updateMessages () {
  $.getScript("/conversations.js")
  setTimeout(updateMessages, 10000);
}
