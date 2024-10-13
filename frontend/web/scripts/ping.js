function ping() {
	// fetch("http://ivanlogvynenko.ddns.net:8080/ping")
	fetch("http://localhost:8080/ping")
		.then(resp => {
			if(!resp.ok) {
				console.log("Request failed")
				document.getElementById("ping-status").innerHTML = "Error";
			}
			return resp.json()
		})
		.then(data => {
			console.log("Ping successful");
			document.getElementById("ping-status").innerHTML = "Running...";	
        });
}

document.getElementById("btn-ping").addEventListener("click", ping);