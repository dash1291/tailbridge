var socket = io.connect();
var streamOut = document.querySelector('#stream-output');
var statusElement = document.querySelector('#status-text');
var filterInput = document.querySelector('#regexInput');
var filterBtn = document.querySelector('#filterBtn');
var filterRegex = new RegExp('', 'g');
var firstStream = true;
var machine = {};

var filterLogs = function(logs) {
  var filterLogs = [];
  for (var i = 0; i < logs.length; i++) {
    if (filterRegex.test(logs[i]))
      filterLogs.push(logs[i]);
  }
  return filterLogs;
};

var updateFilter = function() {
  filterRegex = new RegExp(filterInput.value, 'g');
  console.log('Update regex filter');
}
filterBtn.addEventListener('click', updateFilter, false);

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
    socket.emit('init', hash);
  }
});

socket.on('stream', function(data) {
  if (firstStream) {
    statusElement.innerHTML = "Tailing the file <b>" + machine.file + "</b> on <b>" + machine.ip + "</b>...";
    firstStream = false;
  }

  var data = filterLogs(data.split('\n')).join('<br>');
  var atBottom = (streamOut.scrollHeight - streamOut.scrollTop) === streamOut.offsetHeight;
	streamOut.innerHTML += data + '<br>';
  if (atBottom) {
    streamOut.scrollTop = streamOut.scrollHeight - streamOut.offsetHeight
  }
});

socket.on('denied', function(data) {
  statusElement.innerHTML = "Tailbridge cannot access the requested file.";
});

socket.on('invalid_ip', function(data) {
  statusElement.innerHTML = "Requested IP not found in the config.";
});
