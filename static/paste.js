//Get the button:
mybutton = document.getElementById("myBtn");

// When the user scrolls down 20px from the top of the document, show the button
window.onscroll = function() {scrollFunction()};
function scrollFunction() {
  if (document.body.scrollTop > 20 || document.documentElement.scrollTop > 20) {
    mybutton.style.display = "block";
  } else {
    mybutton.style.display = "none";
  }
}

// When the user clicks on the button, scroll to the top of the document
function topFunction() {
  document.body.scrollTop = 0; // For Safari
  document.documentElement.scrollTop = 0; // For Chrome, Firefox, IE and Opera
}

function deletePaste() {
  console.log("Button clcik")
  var url = "/api/paste/8QlOY";
  var request = new XMLHttpRequest();
  request.open('DELETE', url, true);
  request.onload = function() { // request successful
   alert("paste deleted")
  };
  request.onerror = function () {
    console.log("Request failed")

    // request failed
  };
  request.send(null)
}