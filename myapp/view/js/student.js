function addStudent(){
 var data = {
  stdid : parseInt(document.getElementById("sid").value),
  fname : document.getElementById("fname").value,
  lname : document.getElementById("lname").value,
  email : document.getElementById("email").value
 }

 fetch('/student', {
  method: "POST",
  body: JSON.stringify(data),
  headers: {"Content-type": "application/json; charset=UTF-8"}
 })
 .then(response1 => {
    var sid = data.stdid;
    if (response1.ok) {
        fetch('/student/' + sid)
        .then(response2 => response2.text())
        .then(data => showStudent(data))
    }
 })
}

function showStudent(data) {
 const student = JSON.parse(data)
 // Find a <table> element with id="myTable":
 var table = document.getElementById("myTable");
 // Create an empty <tr> element and add to the last position of the table:
 var row = table.insertRow(table.length);
 // Insert new cells (<td> elements) at the 1st and 2nd position of the
"new" <tr> element:
 var td=[]
 for(i=0; i<table.rows[0].cells.length; i++){
 td[i] = row.insertCell(i);
 }
 // Add student detail to the new cells:
 td[0].innerHTML = student.stdid;
 td[1].innerHTML = student.fname;
 td[2].innerHTML = student.lname;
 td[3].innerHTML = student.email;
 td[4].innerHTML = '<input type="button" onclick="deleteStudent(this)"
value="delete" id="button-1">';
 td[5].innerHTML = '<input type="button" onclick="updateStudent(this)"
value="edit" id="button-2">';
}