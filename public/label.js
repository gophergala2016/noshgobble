$(document).ready(function() {
  $('#label').nutritionLabel({
    width: '300'
  });
  $("#process").submit(function(event) {
    event.preventDefault();
      $.getJSON("/process", {ingredients: $("#ingredients").val()}).done(function(data){
        console.log(data);
      }).error(function(err) {
        console.log(err);
      });
      return false;
  });
});
