var socket = io.connect();
var streamOut = document.querySelector('#stream-output');

socket.on('connect', function() {
  if (window.location.hash) {
    socket.emit('init', window.location.hash.slice(1));
  }
});

socket.on('stream', function(data) {
  var data = data.replace('\n', '<br>');
	streamOut.innerHTML += data;
})

