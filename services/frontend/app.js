const express = require('express')
const app = express()
const port = 3000

app.use(express.static(__dirname + '/html'));

app.get('*', function(req, res){
  res.sendfile(__dirname + '/html/index.html');
});

app.listen(port, () => {
    console.log(`Example app listening on port ${port}`)
})