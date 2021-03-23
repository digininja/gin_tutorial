  var timeout; // this is a variable the timout will be assigned to
  var init = 1;
  var timestamp = null;
  var tag = $("<div></div>");
  function waitForMsg() {
    $.ajax({
      type: "GET",
      url: "/robin",
      cache: false,
      success: function(data) {
            clearTimeout(timeout); // here we clear the timeout
			alert ("success");
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
        alert("error: "+textStatus + " "+ errorThrown);
		//clearTimeout(timeout); // here we clear the timeout
         setTimeout("waitForMsg()", 15000);
      }
    });
  }
  $(function() {
    $("#submit").on('click', function() {
      waitForMsg();
    });
  });
