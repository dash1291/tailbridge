var socket = io.connect();
var streamOut = document.querySelector('#stream-output');
var statusElement = document.querySelector('#status')
var firstStream = true;
var machine = {};

socket.on('connect', function() {
  var hash = window.location.hash;
  if (hash) {
    firstStream = true;
    hash = hash.slice(1);
    var hashParts = hash.split(',')
    machine.ip = hashParts[0];
    machine.file = hashParts[1];

    statusElement.innerHTML = "Connecting to tailbridge to fetch logs from <b>" + machine.ip + "</b>...";
    console.log(status.innerHTML);
    console.log('yo');
    socket.emit('init', hash);
  }
});

socket.on('stream', function(data) {
  if (firstStream) {
    statusElement.innerHTML = "Tailing the file <b>" + machine.file + "</b> on <b>" + machine.ip + "</b>...";
    firstStream = false;
  }

  var data = data.replace('\n', '<br>');
  var atBottom = (streamOut.scrollHeight - streamOut.scrollTop) === streamOut.offsetHeight;
	streamOut.innerHTML += data + '<br>';
  if (atBottom) {
    streamOut.scrollTop = streamOut.scrollHeight - streamOut.offsetHeight
  }
})

