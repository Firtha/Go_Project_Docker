const mongoose = require('mongoose');
const bodyParser = require('body-parser');
const app = require('express')();


// connect to Mongo daemon
mongoose
  .connect(
    'mongodb://mongo:27017/example',
    { useNewUrlParser: true }
  )
  .then(() => console.log('MongoDB Connected'))
  .catch(err => console.log(err));


// DB schema
const IdentitySchema = new mongoose.Schema({
  id : {
    type: Number,
    required: true
  },
  publickeyeth : {
    type: String,
  },
  Firstname : {
    type:String,
  },
  Lastname : {
    type:String,
  },
  Email : {
    type:String,
  },
  DateOfBirth: {
    type:Date,
  }
});

Ident = mongoose.model('item', IdentitySchema);
app.set('view engine', 'ejs');
app.use(bodyParser.urlencoded({ extended: false }));

app.get("/", (req, res) => {
  Ident.find((err, items) => {
    if (err) throw err;
    res.render("index", { items });
  });
});




const port = 8080;
app.listen(port, () => console.log('Server running...'));
