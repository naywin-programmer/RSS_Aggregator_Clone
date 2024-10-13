conn = new WebSocket("ws://" + document.location.host + "/ws-devtools");
conn.onclose = function (evt) {
    console.log("Connection Closed")
    setTimeout(function () {
        location.reload();
    }, 1000);
};