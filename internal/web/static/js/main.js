function request(method, path, callback) {
    let xhttp = new XMLHttpRequest();
    xhttp.onreadystatechange = function() {
        if (xhttp.readyState != XMLHttpRequest.DONE) return;
        switch (xhttp.status) {
            case 200:
                callback(xhttp.responseText)
                break;
            default:
                console.warn(`Failed to call server: code ${xhttp.status} ${xhttp.statusText}`)
        }
    }
    xhttp.open(method, `http://${self.location.host}${path}`, true);
    xhttp.send();
}

function registerUpdateStatusPoll() {
    setTimeout(() => {
        request("GET", "/api/v1/status", function(r) {
            console.log(r)
            registerUpdateStatusPoll();
        });
    }, 1000);
}

registerUpdateStatusPoll();
